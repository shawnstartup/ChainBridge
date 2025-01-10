// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ChainSafe/ChainBridge/vault"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"

	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
)

// Number of blocks to wait for an finalization event
const ExecuteBlockWatchLimit = 2000

// Time between retrying a failed tx
const TxRetryInterval = time.Second * 2

// Maximum number of tx retries before exiting
const TxRetryLimit = 10

var ErrNonceTooLow = errors.New("nonce too low")
var ErrTxUnderpriced = errors.New("replacement transaction underpriced")
var ErrFatalTx = errors.New("submission of transaction failed")
var ErrFatalQuery = errors.New("query of chain state failed")

// vaultProposalIsComplete returns true if the proposal state is Executed
func (w *writer) vaultProposalIsComplete(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetVaultProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check vault proposal existence", "err", err)
		return false
	}
	return prop.Status == VaultExecutedStatus
}

// vaultProposalIsActive returns true if the proposal state is Active
func (w *writer) vaultProposalIsActive(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetVaultProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check vault proposal existence", "err", err)
		return false
	}
	return prop.Status == VaultActiveStatus
}

// vaultProposalIsActive returns true if the proposal state is Passed
func (w *writer) vaultProposalIsPassed(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetVaultProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check vault proposal existence", "err", err)
		return false
	}
	return prop.Status == VaultPassedStatus
}

// proposalIsComplete returns true if the proposal state is either Passed, Transferred or Cancelled
func (w *writer) proposalIsComplete(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus || prop.Status == TransferredStatus || prop.Status == CancelledStatus
}

// proposalIsComplete returns true if the proposal state is Transferred or Cancelled
func (w *writer) proposalIsFinalized(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == TransferredStatus || prop.Status == CancelledStatus // Transferred (3)
}

func (w *writer) proposalIsPassed(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	prop, err := w.bridgeContract.GetProposal(w.conn.CallOpts(), uint8(srcId), uint64(nonce), dataHash)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}
	return prop.Status == PassedStatus
}

// hasVoted checks if this relayer has already voted
func (w *writer) hasVoted(srcId msg.ChainId, nonce msg.Nonce, dataHash [32]byte) bool {
	hasVoted, err := w.bridgeContract.HasVotedOnProposal(w.conn.CallOpts(), utils.IDAndNonce(srcId, nonce), dataHash, w.conn.Opts().From)
	if err != nil {
		w.log.Error("Failed to check proposal existence", "err", err)
		return false
	}

	return hasVoted
}

func (w *writer) shouldVote(m msg.Message, dataHash [32]byte) bool {
	// Check if proposal has passed and skip if Passed or Transferred
	if w.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("Proposal complete, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	// Check if relayer has previously voted
	if w.hasVoted(m.Source, m.DepositNonce, dataHash) {
		w.log.Info("Relayer has already voted, not voting", "src", m.Source, "nonce", m.DepositNonce)
		return false
	}

	return true
}

// createErc20Proposal creates an Erc20 proposal.
// Returns true if the proposal is successfully created or is complete
func (w *writer) createErc20Proposal(m msg.Message) bool {
	w.log.Info("Creating erc20 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := ConstructErc20ProposalData(m.Payload[0].([]byte), m.Payload[1].([]byte))
	dataHash := utils.Hash(append(w.cfg.erc20HandlerContract.Bytes(), data...))

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.log.Trace("GetVaultProposal", "src", m.Source, "nonce", m.DepositNonce)
			vaultProp, err := w.bridgeContract.GetVaultProposal(w.conn.CallOpts(), uint8(m.Source), uint64(m.DepositNonce), dataHash)
			if err != nil {
				w.log.Error("Failed to check vault proposal existence", "err", err)
				return false
			}
			if vaultProp.Status == VaultInactiveStatus {
				w.createVaultProposal(m, dataHash)
				// watch for execution event
				// Capture latest block so when know where to watch from
				latestBlock, err := w.conn.LatestBlock()
				if err != nil {
					w.log.Error("Unable to fetch latest block", "err", err)
					return false
				}
				go w.watchThenExecute(m, data, dataHash, latestBlock)
				return true
			} else if vaultProp.Status == VaultExecutedStatus {
				w.executeProposal(m, data, dataHash)
				return true
			} else {
				w.log.Trace("Ignoring event", "src", m.Source, "nonce", m.DepositNonce, "vaultStatus", vaultProp.Status)
				return true
			}
		} else {
			return false
		}
	}

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.LatestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, dataHash, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

// createErc721Proposal creates an Erc721 proposal.
// Returns true if the proposal is succesfully created or is complete
func (w *writer) createErc721Proposal(m msg.Message) bool {
	w.log.Info("Creating erc721 proposal", "src", m.Source, "nonce", m.DepositNonce)

	data := ConstructErc721ProposalData(m.Payload[0].([]byte), m.Payload[1].([]byte), m.Payload[2].([]byte))
	dataHash := utils.Hash(append(w.cfg.erc721HandlerContract.Bytes(), data...))

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.executeProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	// Capture latest block so we know where to watch from
	latestBlock, err := w.conn.LatestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, dataHash, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

// createGenericDepositProposal creates a generic proposal
// returns true if the proposal is complete or is succesfully created
func (w *writer) createGenericDepositProposal(m msg.Message) bool {
	w.log.Info("Creating generic proposal", "src", m.Source, "nonce", m.DepositNonce)

	metadata := m.Payload[0].([]byte)
	data := ConstructGenericProposalData(metadata)
	toHash := append(w.cfg.genericHandlerContract.Bytes(), data...)
	dataHash := utils.Hash(toHash)

	if !w.shouldVote(m, dataHash) {
		if w.proposalIsPassed(m.Source, m.DepositNonce, dataHash) {
			// We should not vote for this proposal but it is ready to be executed
			w.executeProposal(m, data, dataHash)
			return true
		} else {
			return false
		}
	}

	// Capture latest block so when know where to watch from
	latestBlock, err := w.conn.LatestBlock()
	if err != nil {
		w.log.Error("Unable to fetch latest block", "err", err)
		return false
	}

	// watch for execution event
	go w.watchThenExecute(m, data, dataHash, latestBlock)

	w.voteProposal(m, dataHash)

	return true
}

// watchThenExecute watches for the latest block and executes once the matching finalized event is found
func (w *writer) watchThenExecute(m msg.Message, data []byte, dataHash [32]byte, latestBlock *big.Int) {
	w.log.Info("Watching for finalization event", "src", m.Source, "nonce", m.DepositNonce)
	vaultTxKey := ""
	vaultTxCompleted := false

	var startBlock = big.NewInt(latestBlock.Int64())
	var currentBlock = big.NewInt(latestBlock.Int64())
	// watching for the latest block, querying and matching the finalized event will be retried up to ExecuteBlockWatchLimit times
	for currentBlock.Int64() < startBlock.Int64()+ExecuteBlockWatchLimit {
		select {
		case <-w.stop:
			return
		default:
			// watch for the lastest block, retry up to BlockRetryLimit times
			for waitRetrys := 0; waitRetrys < BlockRetryLimit; waitRetrys++ {
				err := w.conn.WaitForBlock(currentBlock, w.cfg.blockConfirmations)
				if err != nil {
					w.log.Error("Waiting for block failed", "err", err)
					// Exit if retries exceeded
					if waitRetrys+1 == BlockRetryLimit {
						w.log.Error("Waiting for block retries exceeded, shutting down")
						w.sysErr <- ErrFatalQuery
						return
					}
				} else {
					break
				}
			}

			endBlock := new(big.Int).Add(currentBlock, big.NewInt(w.cfg.blockConfirmations.Int64()-1))

			// query for ProposalEvent logs
			query := buildQuery(w.cfg.bridgeContract, utils.ProposalEvent, currentBlock, endBlock)
			evts, err := w.conn.Client().FilterLogs(context.Background(), query)
			if err != nil {
				w.log.Error("Failed to fetch logs", "err", err)
			}
			// execute the proposal once we find the matching finalized event
			for _, evt := range evts {
				sourceId := evt.Topics[1].Big().Uint64()
				depositNonce := evt.Topics[2].Big().Uint64()
				status := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					utils.IsFinalized(uint8(status)) {

					vaultProp, err := w.bridgeContract.GetVaultProposal(w.conn.CallOpts(), uint8(sourceId), depositNonce, dataHash)
					if err != nil {
						w.log.Error("Failed to check vault proposal existence", "err", err)
					} else {
						if vaultProp.Status == VaultInactiveStatus {
							w.createVaultProposal(m, dataHash)
							w.log.Info("createVaultProposal", "src", sourceId, "nonce", depositNonce)
						} else {
							w.log.Trace("Ignoring event", "src", sourceId, "nonce", depositNonce, "vaultStatus", vaultProp.Status)
						}
					}
				} else {
					w.log.Trace("Ignoring event", "src", sourceId, "nonce", depositNonce)
				}
			}
			w.log.Trace("No finalization event found in current block", "block", currentBlock, "src", m.Source, "nonce", m.DepositNonce)

			// query for VaultProposalEvent logs
			query = buildQuery(w.cfg.bridgeContract, utils.VaultProposalEvent, currentBlock, endBlock)
			evts, err = w.conn.Client().FilterLogs(context.Background(), query)
			if err != nil {
				w.log.Error("Failed to fetch vault logs", "err", err)
			}

			// execute the proposal once we find the matching finalized event
			for _, evt := range evts {
				sourceId := evt.Topics[1].Big().Uint64()
				depositNonce := evt.Topics[2].Big().Uint64()
				status := evt.Topics[3].Big().Uint64()

				if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					utils.IsExecuted(uint8(status)) {
					w.executeProposal(m, data, dataHash)
					return
				} else if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					utils.IsActive(uint8(status)) {
					w.executeVaultProposal(m, dataHash)
				} else if m.Source == msg.ChainId(sourceId) &&
					m.DepositNonce.Big().Uint64() == depositNonce &&
					utils.IsFinalized(uint8(status)) {
					vaultProposalEvent, err := w.bridgeContract.ParseVaultProposalEvent(evt)
					if err != nil {
						w.log.Error("Failed to ParseVaultProposalEvent", "err", err)
					} else {
						vaultTxKey = vaultProposalEvent.TxKey
					}
				} else {
					w.log.Trace("Ignoring event", "src", sourceId, "nonce", depositNonce)
				}
			}
			// retrieve vault transaction
			if vaultTxKey != "" && !vaultTxCompleted {
				txStatus, txSubStatus, err := w.vault.RetrieveTransaction(vaultTxKey, "")
				if err != nil {
					w.log.Error("Unable to retrieve vault transaction", "err", err, "txKey", vaultTxKey, "txStatus", txStatus, "txSubStatus", txSubStatus)
				}
				if txStatus == vault.TxStatusCompleted {
					// completeVaultProposal
					w.completeVaultProposal(m, dataHash)
					vaultTxCompleted = true
				}
			}

			w.log.Trace("No passed vault event found in current block", "block", currentBlock, "src", m.Source, "nonce", m.DepositNonce)

			currentBlock = new(big.Int).Add(currentBlock, w.cfg.blockConfirmations)

		}
	}
	log.Warn("Block watch limit exceeded, skipping execution", "source", m.Source, "dest", m.Destination, "nonce", m.DepositNonce)
}

// voteProposal submits a vote proposal
// a vote proposal will try to be submitted up to the TxRetryLimit times
func (w *writer) voteProposal(m msg.Message, dataHash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}
			// These store the gas limit and price before a transaction is sent for logging in case of a failure
			// This declaration is necessary as tx will be nil in the case of an error when sending VoteProposal()
			// We must also declare variables instead of using w.conn.Opts() directly as the opts are currently locked
			// here but for all the logging after line 272 the w.conn.Opts() is unlocked and could be changed by another process
			gasLimit := w.conn.Opts().GasLimit
			gasPrice := w.conn.Opts().GasPrice

			tx, err := w.bridgeContract.VoteProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				m.ResourceId,
				dataHash,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal vote", "tx", tx.Hash(), "src", m.Source, "depositNonce", m.DepositNonce, "gasPrice", tx.GasPrice().String())
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Voting failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "gasLimit", gasLimit, "gasPrice", gasPrice, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.proposalIsComplete(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("Proposal voting complete on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Vote transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

// executeProposal executes the proposal
func (w *writer) executeProposal(m msg.Message, data []byte, dataHash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update nonce", "err", err)
				return
			}
			// These store the gas limit and price before a transaction is sent for logging in case of a failure
			// This is necessary as tx will be nil in the case of an error when sending VoteProposal()
			gasLimit := w.conn.Opts().GasLimit
			gasPrice := w.conn.Opts().GasPrice

			tx, err := w.bridgeContract.ExecuteProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				data,
				m.ResourceId,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Submitted proposal execution", "tx", tx.Hash(), "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce, "gasPrice", tx.GasPrice().String())
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Error("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed, proposal may already be complete", "gasLimit", gasLimit, "gasPrice", gasPrice, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for execution, tx will fail if we aren't the first to execute,
			// but there is no need to retry
			if w.proposalIsFinalized(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("Proposal finalized on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Submission of Execute transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

func (w *writer) createVaultProposal(m msg.Message, dataHash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}
			// These store the gas limit and price before a transaction is sent for logging in case of a failure
			// This declaration is necessary as tx will be nil in the case of an error when sending VoteProposal()
			// We must also declare variables instead of using w.conn.Opts() directly as the opts are currently locked
			// here but for all the logging after line 272 the w.conn.Opts() is unlocked and could be changed by another process
			gasLimit := w.conn.Opts().GasLimit
			gasPrice := w.conn.Opts().GasPrice

			//  make Raw with below cutomer fields for audit by cosigner callback service
			txId := w.vault.MakeCustomerRefId(uint8(m.Source), uint8(m.Destination), uint64(m.DepositNonce))
			txIdHash := utils.Hash(append([]byte(txId)))

			tx, err := w.bridgeContract.CreateVaultProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				dataHash,
				txIdHash,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Create vaultProposal", "tx", tx.Hash(), "txId", txId, "txIdHash", hex.EncodeToString(txIdHash[:]), "resourceId", m.ResourceId.Hex(), "src", m.Source, "destination", m.Destination, "depositNonce", m.DepositNonce, "gasPrice", tx.GasPrice().String())
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "gasLimit", gasLimit, "gasPrice", gasPrice, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.vaultProposalIsActive(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("VaultProposal is active on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Creation of VaultProposal transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

func (w *writer) executeVaultProposal(m msg.Message, dataHash [32]byte) {
	// call vault api to transfer tokens from vault wallet
	addr, err := w.bridgeContract.ResourceIDToHandlerAddress(&bind.CallOpts{From: w.conn.Keypair().CommonAddress()}, m.ResourceId)
	if err != nil {
		w.log.Error("failed to get handler from resource ID", "resourceId", m.ResourceId, "err", err)
		return
	}

	//  make Raw with below cutomer fields for audit by cosigner callback service
	txId := w.vault.MakeCustomerRefId(uint8(m.Source), uint8(m.Destination), uint64(m.DepositNonce))
	w.log.Info("Vault sendTransaction", "destinationChainId", m.Destination, "destinationAddress", addr.String())
	amount := big.NewInt(0).SetBytes(m.Payload[0].([]byte))
	//recipientBytes := m.Payload[1]

	//customerRefIdWithTimestamp := fmt.Sprintf("%s-timestamp-%d", txId, time.Now().Unix())
	customerRefIdWithTimestamp := fmt.Sprintf("%s-timestamp-%d", txId, 0)
	txKey, err := w.vault.SendVaultTransaction(uint8(m.Destination), "0x"+m.ResourceId.Hex(), addr.String(), customerRefIdWithTimestamp, amount, false)
	if err != nil {
		w.log.Error("Vault sendTransaction", "destinationChainId", m.Destination, "destinationAddress", addr.String(), "amount", amount.String(), "err", err)
		return
		// TODO handle err
	}
	w.log.Info("Vault sendTransaction", "destinationChainId", m.Destination, "destinationAddress", addr.String(), "customerRefId", customerRefIdWithTimestamp, "txKey", txKey, "amount", amount.String())

	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}

			// These store the gas limit and price before a transaction is sent for logging in case of a failure
			// This declaration is necessary as tx will be nil in the case of an error when sending VoteProposal()
			// We must also declare variables instead of using w.conn.Opts() directly as the opts are currently locked
			// here but for all the logging after line 272 the w.conn.Opts() is unlocked and could be changed by another process
			gasLimit := w.conn.Opts().GasLimit
			gasPrice := w.conn.Opts().GasPrice

			tx, err := w.bridgeContract.PassVaultProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				dataHash,
				txKey,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Create vaultProposal", "tx", tx.Hash(), "resourceId", m.ResourceId.Hex(), "src", m.Source, "destination", m.Destination, "depositNonce", m.DepositNonce, "gasPrice", tx.GasPrice().String())
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("Execution failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "gasLimit", gasLimit, "gasPrice", gasPrice, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.vaultProposalIsPassed(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("VaultProposal is passed on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("Execution of VaultProposal transaction failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}

func (w *writer) completeVaultProposal(m msg.Message, dataHash [32]byte) {
	for i := 0; i < TxRetryLimit; i++ {
		select {
		case <-w.stop:
			return
		default:
			err := w.conn.LockAndUpdateOpts()
			if err != nil {
				w.log.Error("Failed to update tx opts", "err", err)
				continue
			}
			// These store the gas limit and price before a transaction is sent for logging in case of a failure
			// This declaration is necessary as tx will be nil in the case of an error when sending VoteProposal()
			// We must also declare variables instead of using w.conn.Opts() directly as the opts are currently locked
			// here but for all the logging after line 272 the w.conn.Opts() is unlocked and could be changed by another process
			gasLimit := w.conn.Opts().GasLimit
			gasPrice := w.conn.Opts().GasPrice

			tx, err := w.bridgeContract.ExecuteVaultProposal(
				w.conn.Opts(),
				uint8(m.Source),
				uint64(m.DepositNonce),
				dataHash,
			)
			w.conn.UnlockOpts()

			if err == nil {
				w.log.Info("Completed vaultProposal", "tx", tx.Hash(), "src", m.Source, "depositNonce", m.DepositNonce, "gasPrice", tx.GasPrice().String())
				if w.metrics != nil {
					w.metrics.VotesSubmitted.Inc()
				}
				return
			} else if err.Error() == ErrNonceTooLow.Error() || err.Error() == ErrTxUnderpriced.Error() {
				w.log.Debug("Nonce too low, will retry")
				time.Sleep(TxRetryInterval)
			} else {
				w.log.Warn("completeVaultProposal failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce, "gasLimit", gasLimit, "gasPrice", gasPrice, "err", err)
				time.Sleep(TxRetryInterval)
			}

			// Verify proposal is still open for voting, otherwise no need to retry
			if w.vaultProposalIsComplete(m.Source, m.DepositNonce, dataHash) {
				w.log.Info("VaultProposal is completed on chain", "src", m.Source, "dst", m.Destination, "nonce", m.DepositNonce)
				return
			}
		}
	}
	w.log.Error("completeVaultProposal failed", "source", m.Source, "dest", m.Destination, "depositNonce", m.DepositNonce)
	w.sysErr <- ErrFatalTx
}
