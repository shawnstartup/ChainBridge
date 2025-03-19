package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ChainSafe/ChainBridge/bindings/ERC20Handler"
	"github.com/ChainSafe/ChainBridge/chains"
	"github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

type QueryProposalResponse struct {
	*ProposalRecord      `json:"proposalRecord"`
	*BridgeProposal      `json:"bridgeProposal"`
	*BridgeVaultProposal `json:"bridgeVaultProposal"`
	Error                string `json:"error,omitempty"`
}

type ProposalRecord struct {
	TokenAddress                   string   `json:"tokenAddress"`
	LenDestinationRecipientAddress uint8    `json:"lenDestinationRecipientAddress"`
	DestinationChainID             uint8    `json:"destinationChainID"`
	ResourceID                     string   `json:"resourceID"`
	DestinationRecipientAddress    string   `json:"destinationRecipientAddress"`
	Depositer                      string   `json:"depositer"`
	Amount                         *big.Int `json:"amount"`
}

type BridgeProposal struct {
	ResourceID    string           `json:"resourceID"`
	DataHash      string           `json:"dataHash"`
	YesVotes      []common.Address `json:"yesVotes"`
	NoVotes       []common.Address `json:"noVotes"`
	Status        string           `json:"status"`
	ProposedBlock *big.Int         `json:"proposedBlock"`
}

type BridgeVaultProposal struct {
	ResourceID    string `json:"resourceID"`
	DataHash      string `json:"dataHash"`
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        string
	ProposedBlock *big.Int
	TxIdHash      string `json:"txIdHash"`
	TxKey         string `json:"txKey"`
}

func queryProposalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := &QueryProposalResponse{
		Error: "",
	}

	srcChainIdStr := r.FormValue("srcChainId")
	if srcChainIdStr == "" {
		http.Error(w, "srcChainId is required", http.StatusBadRequest)

		return
	}
	srcChainId, err := strconv.Atoi(srcChainIdStr)
	if err != nil {
		http.Error(w, "Invalid srcChainId", http.StatusBadRequest)

		return
	}
	dstChainIdStr := r.FormValue("dstChainId")
	if dstChainIdStr == "" {
		http.Error(w, "dstChainId is required", http.StatusBadRequest)

		return
	}
	dstChainId, err := strconv.Atoi(dstChainIdStr)
	if err != nil {
		http.Error(w, "Invalid dstChainId", http.StatusBadRequest)

		return
	}

	nonceStr := r.FormValue("nonce")
	if nonceStr == "" {
		http.Error(w, "nonce is required", http.StatusBadRequest)

		return
	}
	nonce, err := strconv.Atoi(nonceStr)
	if err != nil {
		http.Error(w, "Invalid nonce", http.StatusBadRequest)

		return
	}
	var depositRecord ERC20Handler.ERC20HandlerDepositRecord
	if srcChain, ok := handlerRegistry[msg.ChainId(srcChainId)]; ok {
		depositRecord, err = srcChain.QueryErc20DepositRecord(msg.ChainId(dstChainId), msg.Nonce(nonce))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Failed to serve queryProposalHandler")

			return
		}

		depositRecordData := &ProposalRecord{}
		depositRecordData.Depositer = depositRecord.Depositer.String()
		depositRecordData.Amount = depositRecord.Amount
		depositRecordData.LenDestinationRecipientAddress = depositRecord.LenDestinationRecipientAddress
		depositRecordData.DestinationRecipientAddress = hex.EncodeToString(depositRecord.DestinationRecipientAddress)
		depositRecordData.DestinationChainID = depositRecord.DestinationChainID
		//slice := make([]byte, 32)
		//copy(slice, depositRecord.ResourceID[:])
		depositRecordData.ResourceID = hex.EncodeToString(depositRecord.ResourceID[:])
		depositRecordData.TokenAddress = depositRecord.TokenAddress.String()

		response.ProposalRecord = depositRecordData
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if dstChain, ok := handlerRegistry[msg.ChainId(dstChainId)]; ok {
		proposal, getProposalErr := dstChain.GetProposal(depositRecord, uint8(srcChainId), msg.Nonce(nonce))
		if getProposalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Failed to serve GetProposal")

			return
		}

		proposalData := &BridgeProposal{}
		proposalData.ProposedBlock = proposal.ProposedBlock
		proposalData.DataHash = hex.EncodeToString(proposal.DataHash[:])
		proposalData.ResourceID = hex.EncodeToString(proposal.ResourceID[:])

		//     enum ProposalStatus {
		//        Inactive,
		//        Active,
		//        Passed,
		//        Executed,
		//        Cancelled
		//    }
		proposalData.Status = chains.ProposalStatus(proposal.Status).String()

		proposalData.NoVotes = proposal.NoVotes
		proposalData.YesVotes = proposal.YesVotes

		response.BridgeProposal = proposalData

		vaultProposal, getVaultProposalErr := dstChain.GetVaultProposal(depositRecord, uint8(srcChainId), msg.Nonce(nonce))
		if getVaultProposalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Failed to serve GetVaultProposal")

			return
		}

		vaultProposalData := &BridgeVaultProposal{}
		vaultProposalData.ProposedBlock = vaultProposal.ProposedBlock
		vaultProposalData.DataHash = hex.EncodeToString(vaultProposal.DataHash[:])

		//     enum VaultProposalStatus {
		//        Inactive,
		//        Active,
		//        Passed,
		//        Executed,
		//        Cancelled
		//    }
		vaultProposalData.Status = chains.ProposalStatus(vaultProposal.Status).String()

		vaultProposalData.ResourceID = hex.EncodeToString(vaultProposal.ResourceID[:])
		vaultProposalData.YesVotes = vaultProposal.YesVotes
		vaultProposalData.NoVotes = vaultProposal.NoVotes
		vaultProposalData.TxKey = vaultProposal.TxKey
		vaultProposalData.TxIdHash = hex.EncodeToString(vaultProposal.TxIdHash[:])

		response.BridgeVaultProposal = vaultProposalData
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	respErr := json.NewEncoder(w).Encode(response)
	if respErr != nil {
		log.Error("Failed to serve queryProposalHandler")
	}
}

func executeProposalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var srcChainId int
	var dstChainId int
	var nonce int

	switch r.Method {
	case http.MethodPost:
		contentType := r.Header.Get("Content-Type")

		if r.Method == http.MethodPost && strings.Contains(contentType, "application/json") {
			var params struct {
				SrcChainId int `json:"srcChainId"`
				DstChainId int `json:"dstChainId"`
				Nonce      int `json:"nonce"`
			}
			if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)

				return
			}
			srcChainId = params.SrcChainId
			dstChainId = params.DstChainId
			nonce = params.Nonce
		} else {
			http.Error(w, "Invalid Content-Type", http.StatusBadRequest)

			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	response := &QueryProposalResponse{
		Error: "",
	}

	var srcChain chains.ChainHandler
	var depositRecord ERC20Handler.ERC20HandlerDepositRecord
	var ok bool
	if srcChain, ok = handlerRegistry[msg.ChainId(srcChainId)]; ok {
		var queryDepositRecordErr error
		depositRecord, queryDepositRecordErr = srcChain.QueryErc20DepositRecord(msg.ChainId(dstChainId), msg.Nonce(nonce))
		if queryDepositRecordErr != nil {
			log.Error("Failed to serve queryProposalHandler", "err", queryDepositRecordErr)
			response.Error = "Failed to QueryErc20DepositRecord"
		} else {
			depositRecordData := &ProposalRecord{}
			depositRecordData.Depositer = depositRecord.Depositer.String()
			depositRecordData.Amount = depositRecord.Amount
			depositRecordData.LenDestinationRecipientAddress = depositRecord.LenDestinationRecipientAddress
			depositRecordData.DestinationRecipientAddress = hex.EncodeToString(depositRecord.DestinationRecipientAddress)
			depositRecordData.DestinationChainID = depositRecord.DestinationChainID
			slice := make([]byte, 32)
			copy(slice, depositRecord.ResourceID[:])
			depositRecordData.ResourceID = hex.EncodeToString(slice)
			depositRecordData.TokenAddress = depositRecord.TokenAddress.String()

			response.ProposalRecord = depositRecordData
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if dstChain, ok := handlerRegistry[msg.ChainId(dstChainId)]; ok {
		proposal, getProposalErr := dstChain.GetProposal(depositRecord, uint8(srcChainId), msg.Nonce(nonce))
		if getProposalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Failed to serve GetProposal")

			return
		}

		proposalData := &BridgeProposal{}
		proposalData.ProposedBlock = proposal.ProposedBlock
		proposalData.DataHash = hex.EncodeToString(proposal.DataHash[:])
		proposalData.ResourceID = hex.EncodeToString(proposal.ResourceID[:])
		proposalData.Status = chains.ProposalStatus(proposal.Status).String()
		proposalData.NoVotes = proposal.NoVotes
		proposalData.YesVotes = proposal.YesVotes

		response.BridgeProposal = proposalData

		vaultProposal, getVaultProposalErr := dstChain.GetVaultProposal(depositRecord, uint8(srcChainId), msg.Nonce(nonce))
		if getVaultProposalErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Failed to serve GetVaultProposal")

			return
		}

		vaultProposalData := &BridgeVaultProposal{}
		vaultProposalData.ProposedBlock = vaultProposal.ProposedBlock
		vaultProposalData.DataHash = hex.EncodeToString(vaultProposal.DataHash[:])
		vaultProposalData.Status = chains.ProposalStatus(vaultProposal.Status).String()
		vaultProposalData.ResourceID = hex.EncodeToString(vaultProposal.ResourceID[:])
		vaultProposalData.YesVotes = vaultProposal.YesVotes
		vaultProposalData.NoVotes = vaultProposal.NoVotes
		vaultProposalData.TxKey = vaultProposal.TxKey
		vaultProposalData.TxIdHash = hex.EncodeToString(vaultProposal.TxIdHash[:])

		response.BridgeVaultProposal = vaultProposalData
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if response.BridgeProposal.Status != chains.Passed.String() {
		response.Error = "status of proposal is not passed"
	} else {
		handleDepositRecordErr := srcChain.HandleErc20DepositedRecord(depositRecord, msg.Nonce(nonce))
		if handleDepositRecordErr != nil {
			log.Error("Failed to HandleErc20DepositedRecord", "err", handleDepositRecordErr)
			response.Error = "Failed to HandleErc20DepositedRecord"
		}
	}

	w.WriteHeader(http.StatusOK)
	respErr := json.NewEncoder(w).Encode(response)
	if respErr != nil {
		log.Error("Failed to serve queryProposalHandler")
	}
}
