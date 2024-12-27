// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bridge

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// BridgeProposalRecord is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposalRecord struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}

// BridgeVaultProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeVaultProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
	TxId          string
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"VaultProposalEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"_txProposals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_vaultProposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txId\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminAddRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminRemoveRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"adminSetVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"getTxProposalRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBridge.ProposalRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getVaultProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txId\",\"type\":\"string\"}],\"internalType\":\"structBridge.VaultProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"createVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"passVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"executeVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620065b7380380620065b78339818101604052810190620000379190620004bc565b60008060006101000a81548160ff02191690831515021790555084600260006101000a81548160ff021916908360ff160217905550826003819055508160068190555080600781905550620000966000801b336200013460201b60201c565b620000c0604051620000a89062000601565b60405180910390206000801b6200014a60201b60201c565b60005b8451811015620001285762000108604051620000df9062000601565b6040518091039020868381518110620000f457fe5b60200260200101516200016960201b60201c565b6004600081548092919060010191905055508080600101915050620000c3565b50505050505062000746565b620001468282620001f860201b60201c565b5050565b8060016000848152602001908152602001600020600201819055505050565b620001a06001600084815260200190815260200160002060020154620001946200029c60201b60201c565b620002a460201b60201c565b620001e2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001d99062000618565b60405180910390fd5b620001f48282620001f860201b60201c565b5050565b620002278160016000858152602001908152602001600020600001620002dd60201b620036cc1790919060201c565b1562000298576200023d6200029c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6000620002d582600160008681526020019081526020016000206000016200031560201b620035d11790919060201c565b905092915050565b60006200030d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200034d60201b60201c565b905092915050565b600062000345836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003c760201b60201c565b905092915050565b6000620003618383620003c760201b60201c565b620003bc578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620003c1565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620003fb81620006f8565b92915050565b600082601f8301126200041357600080fd5b81516200042a620004248262000668565b6200063a565b915081818352602084019350602081019050838560208402820111156200045057600080fd5b60005b83811015620004845781620004698882620003ea565b84526020840193506020830192505060018101905062000453565b5050505092915050565b6000815190506200049f8162000712565b92915050565b600081519050620004b6816200072c565b92915050565b600080600080600060a08688031215620004d557600080fd5b6000620004e588828901620004a5565b955050602086015167ffffffffffffffff8111156200050357600080fd5b620005118882890162000401565b945050604062000524888289016200048e565b935050606062000537888289016200048e565b92505060806200054a888289016200048e565b9150509295509295909350565b600062000566602f8362000691565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000620005ce600c83620006a2565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b60006200060e82620005bf565b9150819050919050565b60006020820190508181036000830152620006338162000557565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200065e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200068057600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b6000620006ba82620006c1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6200070381620006ad565b81146200070f57600080fd5b50565b6200071d81620006e1565b81146200072957600080fd5b50565b6200073781620006eb565b81146200074357600080fd5b50565b615e6180620007566000396000f3fe6080604052600436106102885760003560e01c806380ae1c281161015a578063b4b04f32116100c1578063cdb0f73a1161007a578063cdb0f73a14610a31578063ce6cfa0f14610a5a578063d547741f14610a97578063d7a9cd7914610ac0578063e8437ee714610aeb578063ffaac0eb14610b1457610288565b8063b4b04f3214610909578063beab71311461094a578063c5b37c2214610975578063c5ec8970146109a0578063ca15c873146109cb578063cb10f21514610a0857610288565b806391d148541161011357806391d14854146107e5578063926d7d7f146108225780639d5773e01461084d5780639d82dd6314610878578063a217fddf146108a1578063a9cf69fa146108cc57610288565b806380ae1c28146106c257806384db809f146106d957806385717b42146107165780638c0c2631146107565780639010d07c1461077f57806391c404ac146107bc57610288565b80634603ae38116101fe5780635c975abb116101b75780635c975abb146105b45780635e1fab0f146105df5780636580a71a14610608578063780cf004146106315780637febe63f1461065a578063802aabe81461069757610288565b80634603ae381461047f57806347b8f8a4146104a85780634b0b919d146104d15780634e0560051461050e5780635059871914610537578063541d55481461057757610288565b806320857acc1161025057806320857acc1461034d578063248a9ca31461038a5780632f2ff15d146103c757806336568abe146103f05780633ee7094a146104195780634454b20d1461045657610288565b806305e2ca171461028d578063110b7e0b146102a957806317f03ce5146102d257806318b257ed146102fb5780631ff013f114610324575b600080fd5b6102a760048036038101906102a291906142cd565b610b2b565b005b3480156102b557600080fd5b506102d060048036038101906102cb9190613eb7565b610d88565b005b3480156102de57600080fd5b506102f960048036038101906102f49190614339565b610e05565b005b34801561030757600080fd5b50610322600480360381019061031d91906143eb565b610fb2565b005b34801561033057600080fd5b5061034b60048036038101906103469190614388565b61123d565b005b34801561035957600080fd5b50610374600480360381019061036f9190614173565b6119be565b604051610381919061595d565b60405180910390f35b34801561039657600080fd5b506103b160048036038101906103ac9190614091565b611a59565b6040516103be919061535e565b60405180910390f35b3480156103d357600080fd5b506103ee60048036038101906103e991906140ba565b611a79565b005b3480156103fc57600080fd5b50610417600480360381019061041291906140ba565b611aed565b005b34801561042557600080fd5b50610440600480360381019061043b91906141dd565b611b70565b60405161044d91906155b9565b60405180910390f35b34801561046257600080fd5b5061047d60048036038101906104789190614466565b611c2d565b005b34801561048b57600080fd5b506104a660048036038101906104a1919061401c565b611f6e565b005b3480156104b457600080fd5b506104cf60048036038101906104ca9190614339565b612014565b005b3480156104dd57600080fd5b506104f860048036038101906104f391906142a4565b612251565b60405161050591906159d7565b60405180910390f35b34801561051a57600080fd5b50610535600480360381019061053091906141b4565b612278565b005b34801561054357600080fd5b5061055e60048036038101906105599190614219565b6122b7565b60405161056e9493929190615410565b60405180910390f35b34801561058357600080fd5b5061059e60048036038101906105999190613e65565b612301565b6040516105ab9190615343565b60405180910390f35b3480156105c057600080fd5b506105c9612327565b6040516105d69190615343565b60405180910390f35b3480156105eb57600080fd5b5061060660048036038101906106019190613e65565b61233d565b005b34801561061457600080fd5b5061062f600480360381019061062a9190614339565b612362565b005b34801561063d57600080fd5b5061065860048036038101906106539190613ef3565b61262a565b005b34801561066657600080fd5b50610681600480360381019061067c9190614255565b6126ad565b60405161068e9190615343565b60405180910390f35b3480156106a357600080fd5b506106ac6126e9565b6040516106b991906159bc565b60405180910390f35b3480156106ce57600080fd5b506106d76126ef565b005b3480156106e557600080fd5b5061070060048036038101906106fb9190614091565b612701565b60405161070d91906152d6565b60405180910390f35b34801561072257600080fd5b5061073d60048036038101906107389190614132565b612734565b60405161074d9493929190615a0d565b60405180910390f35b34801561076257600080fd5b5061077d60048036038101906107789190613eb7565b61279b565b005b34801561078b57600080fd5b506107a660048036038101906107a191906140f6565b612818565b6040516107b391906152d6565b60405180910390f35b3480156107c857600080fd5b506107e360048036038101906107de91906141b4565b61284a565b005b3480156107f157600080fd5b5061080c600480360381019061080791906140ba565b6128a1565b6040516108199190615343565b60405180910390f35b34801561082e57600080fd5b506108376128d3565b604051610844919061535e565b60405180910390f35b34801561085957600080fd5b506108626128ea565b60405161086f91906159bc565b60405180910390f35b34801561088457600080fd5b5061089f600480360381019061089a9190613e65565b6128f0565b005b3480156108ad57600080fd5b506108b66129ca565b6040516108c3919061535e565b60405180910390f35b3480156108d857600080fd5b506108f360048036038101906108ee9190614339565b6129d1565b6040516109009190615978565b60405180910390f35b34801561091557600080fd5b50610930600480360381019061092b9190614219565b612bb2565b604051610941959493929190615455565b60405180910390f35b34801561095657600080fd5b5061095f612c9a565b60405161096c91906159f2565b60405180910390f35b34801561098157600080fd5b5061098a612cad565b60405161099791906159bc565b60405180910390f35b3480156109ac57600080fd5b506109b5612cb3565b6040516109c291906159bc565b60405180910390f35b3480156109d757600080fd5b506109f260048036038101906109ed9190614091565b612cb9565b6040516109ff91906159bc565b60405180910390f35b348015610a1457600080fd5b50610a2f6004803603810190610a2a9190613f56565b612ce0565b005b348015610a3d57600080fd5b50610a586004803603810190610a539190613e65565b612db2565b005b348015610a6657600080fd5b50610a816004803603810190610a7c9190614339565b612e8c565b604051610a8e919061599a565b60405180910390f35b348015610aa357600080fd5b50610abe6004803603810190610ab991906140ba565b61310f565b005b348015610acc57600080fd5b50610ad5613183565b604051610ae291906159bc565b60405180910390f35b348015610af757600080fd5b50610b126004803603810190610b0d9190613fa5565b613189565b005b348015610b2057600080fd5b50610b29613261565b005b610b33613273565b6006543414610b77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6e9061573d565b60405180910390fd5b60006009600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c169061579d565b60405180910390fd5b6000600860008760ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590508383600a60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008960ff1660ff1681526020019081526020016000209190610cc8929190613975565b5060008290508073ffffffffffffffffffffffffffffffffffffffff166338995da9878985338a8a6040518763ffffffff1660e01b8152600401610d119695949392919061555d565b600060405180830381600087803b158015610d2b57600080fd5b505af1158015610d3f573d6000803e3d6000fd5b505050508167ffffffffffffffff16868860ff167fdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed860405160405180910390a450505050505050565b610d906132c4565b60008290508073ffffffffffffffffffffffffffffffffffffffff16636817031b836040518263ffffffff1660e01b8152600401610dce91906152d6565b600060405180830381600087803b158015610de857600080fd5b505af1158015610dfc573d6000803e3d6000fd5b50505050505050565b610e0d613312565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600480811115610e7a57fe5b8160040160009054906101000a900460ff166004811115610e9757fe5b1415610ed8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ecf9061577d565b60405180910390fd5b600754610ee9438360050154613384565b11610f29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f20906157dd565b60405180910390fd5b60048160040160006101000a81548160ff02191690836004811115610f4a57fe5b0217905550600480811115610f5b57fe5b8467ffffffffffffffff168660ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041784600001548560010154604051610fa39291906153e7565b60405180910390a45050505050565b610fba613312565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002090506002600481111561102857fe5b8160040160009054906101000a900460ff16600481111561104557fe5b14611085576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107c9061583d565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008681526020019081526020016000209050600160048111156110d057fe5b8160040160009054906101000a900460ff1660048111156110ed57fe5b1415801561112357506002600481111561110357fe5b8160040160009054906101000a900460ff16600481111561112057fe5b14155b801561115757506003600481111561113757fe5b8160040160009054906101000a900460ff16600481111561115457fe5b14155b611196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118d906156fd565b60405180910390fd5b60018160040160006101000a81548160ff021916908360048111156111b757fe5b0217905550838160060190805190602001906111d49291906139f5565b50600160048111156111e257fe5b8667ffffffffffffffff168860ff167fd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b5856000015486600101548960405161122c939291906154af565b60405180910390a450505050505050565b6112456133ce565b61124d613273565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611351576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113489061591d565b60405180910390fd5b60018160040160009054906101000a900460ff16600481111561137057fe5b11156113b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a89061585d565b60405180910390fd5b600c60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611476576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161146d9061569d565b60405180910390fd5b60008160040160009054906101000a900460ff16600481111561149557fe5b14156116b557600560008154600101919050819055506040518060c0016040528085815260200184815260200160016040519080825280602002602001820160405280156114f25781602001602082028036833780820191505090505b50815260200160006040519080825280602002602001820160405280156115285781602001602082028036833780820191505090505b5081526020016001600481111561153b57fe5b815260200143815250600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906115ae929190613a75565b5060608201518160030190805190602001906115cb929190613a75565b5060808201518160040160006101000a81548160ff021916908360048111156115f057fe5b021790555060a08201518160050155905050338160020160008154811061161357fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600481111561166857fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516116a89291906153e7565b60405180910390a46117f8565b6007546116c6438360050154613384565b111561174b5760048160040160006101000a81548160ff021916908360048111156116ed57fe5b02179055506004808111156116fe57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417878760405161173e9291906153e7565b60405180910390a46117f7565b80600101548314611791576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611788906158fd565b60405180910390fd5b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b60048081111561180457fe5b8160040160009054906101000a900460ff16600481111561182157fe5b146119b6576001600c60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff1660048111156118d357fe5b8567ffffffffffffffff168760ff167f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b957434364087604051611911919061535e565b60405180910390a460016003541115806119345750600354816002018054905010155b156119b55760028160040160006101000a81548160ff0219169083600481111561195a57fe5b02179055506002600481111561196c57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516119ac9291906153e7565b60405180910390a45b5b505050505050565b6119c6613aff565b600e826040516119d69190615293565b90815260200160405180910390206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b600060016000838152602001908152602001600020600201549050919050565b611aa06001600084815260200190815260200160002060020154611a9b61342c565b6128a1565b611adf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ad69061567d565b60405180910390fd5b611ae98282613434565b5050565b611af561342c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b62576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b599061593d565b60405180910390fd5b611b6c82826134c8565b5050565b600a602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c255780601f10611bfa57610100808354040283529160200191611c25565b820191906000526020600020905b815481529060010190602001808311611c0857829003601f168201915b505050505081565b611c356133ce565b611c3d613273565b60006009600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008660ff1660088767ffffffffffffffff1668ffffffffffffffffff16901b1790506000828686604051602001611caf93929190615269565b6040516020818303038152906040528051906020012090506000600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020905060006004811115611d1257fe5b8160040160009054906101000a900460ff166004811115611d2f57fe5b1415611d70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d679061561d565b60405180910390fd5b60026004811115611d7d57fe5b8160040160009054906101000a900460ff166004811115611d9a57fe5b14611dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dd1906155fd565b60405180910390fd5b80600101548214611e20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e179061575d565b60405180910390fd5b60038160040160006101000a81548160ff02191690836004811115611e4157fe5b02179055506000600960008360000154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663e248cff283600001548a8a6040518463ffffffff1660e01b8152600401611ec39392919061552b565b600060405180830381600087803b158015611edd57600080fd5b505af1158015611ef1573d6000803e3d6000fd5b505050508160040160009054906101000a900460ff166004811115611f1257fe5b8967ffffffffffffffff168b60ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041785600001548660010154604051611f5a9291906153e7565b60405180910390a450505050505050505050565b611f766132c4565b60008090505b8484905081101561200d57848482818110611f9357fe5b9050602002016020810190611fa89190613e8e565b73ffffffffffffffffffffffffffffffffffffffff166108fc848484818110611fcd57fe5b905060200201359081150290604051600060405180830381858888f19350505050158015611fff573d6000803e3d6000fd5b508080600101915050611f7c565b5050505050565b61201c613312565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002090506002600481111561208a57fe5b8160040160009054906101000a900460ff1660048111156120a757fe5b146120e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120de9061583d565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002090506001600481111561213257fe5b8160040160009054906101000a900460ff16600481111561214f57fe5b148061218257506002600481111561216357fe5b8160040160009054906101000a900460ff16600481111561218057fe5b145b6121c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121b89061571d565b60405180910390fd5b60038160040160006101000a81548160ff021916908360048111156121e257fe5b0217905550600360048111156121f457fe5b8567ffffffffffffffff168760ff167fd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b58560000154866001015486600601604051612241939291906154ed565b60405180910390a4505050505050565b60086020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6122806132c4565b80600381905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a250565b600b602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff16908060050154905084565b6000612320604051612312906152c1565b6040518091039020836128a1565b9050919050565b60008060009054906101000a900460ff16905090565b6123456132c4565b6123526000801b82611a79565b61235f6000801b33611aed565b50565b61236a613312565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600260048111156123d857fe5b8160040160009054906101000a900460ff1660048111156123f557fe5b14612435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161242c9061583d565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002090506002600481111561248057fe5b8160040160009054906101000a900460ff16600481111561249d57fe5b141580156124d35750600360048111156124b357fe5b8160040160009054906101000a900460ff1660048111156124d057fe5b14155b612512576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612509906158bd565b60405180910390fd5b60028160040160006101000a81548160ff0219169083600481111561253357fe5b02179055506000600e8260060160405161254d91906152aa565b908152602001604051809103902090508260000154816001018190555082600101548160020181905550858160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550868160000160006101000a81548160ff021916908360ff160217905550600260048111156125cc57fe5b8667ffffffffffffffff168860ff167fd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b58660000154876001015487600601604051612619939291906154ed565b60405180910390a450505050505050565b6126326132c4565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d9caed128585856040518463ffffffff1660e01b81526004016126749392919061530c565b600060405180830381600087803b15801561268e57600080fd5b505af11580156126a2573d6000803e3d6000fd5b505050505050505050565b600c602052826000526040600020602052816000526040600020602052806000526040600020600092509250509054906101000a900460ff1681565b60045481565b6126f76132c4565b6126ff61355c565b565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600e818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900460ff16908060000160019054906101000a900467ffffffffffffffff16908060010154908060020154905084565b6127a36132c4565b60008290508073ffffffffffffffffffffffffffffffffffffffff166307b7ed99836040518263ffffffff1660e01b81526004016127e191906152d6565b600060405180830381600087803b1580156127fb57600080fd5b505af115801561280f573d6000803e3d6000fd5b50505050505050565b600061284282600160008681526020019081526020016000206000016135b790919063ffffffff16565b905092915050565b6128526132c4565b806006541415612897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161288e906158dd565b60405180910390fd5b8060068190555050565b60006128cb82600160008681526020019081526020016000206000016135d190919063ffffffff16565b905092915050565b6040516128df906152c1565b604051809103902081565b60055481565b6128f86132c4565b612915604051612907906152c1565b6040518091039020826128a1565b612954576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161294b906156dd565b60405180910390fd5b612971604051612963906152c1565b60405180910390208261310f565b8073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a26004600081548092919060019003919050555050565b6000801b81565b6129d9613b3a565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600b60008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060c0016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015612ada57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612a90575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015612b6857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612b1e575b505050505081526020016004820160009054906101000a900460ff166004811115612b8f57fe5b6004811115612b9a57fe5b81526020016005820154815250509150509392505050565b600d602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff1690806005015490806006018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612c905780601f10612c6557610100808354040283529160200191612c90565b820191906000526020600020905b815481529060010190602001808311612c7357829003601f168201915b5050505050905085565b600260009054906101000a900460ff1681565b60065481565b60075481565b6000612cd960016000848152602001908152602001600020600001613601565b9050919050565b612ce86132c4565b826009600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508073ffffffffffffffffffffffffffffffffffffffff1663b8fa373684846040518363ffffffff1660e01b8152600401612d7a929190615379565b600060405180830381600087803b158015612d9457600080fd5b505af1158015612da8573d6000803e3d6000fd5b5050505050505050565b612dba6132c4565b612dd7604051612dc9906152c1565b6040518091039020826128a1565b15612e17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e0e9061581d565b60405180910390fd5b612e34604051612e26906152c1565b604051809103902082611a79565b8073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a260046000815480929190600101919050555050565b612e94613b81565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600d60008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060e0016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015612f9557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612f4b575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561302357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612fd9575b505050505081526020016004820160009054906101000a900460ff16600481111561304a57fe5b600481111561305557fe5b815260200160058201548152602001600682018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156130fc5780601f106130d1576101008083540402835291602001916130fc565b820191906000526020600020905b8154815290600101906020018083116130df57829003601f168201915b5050505050815250509150509392505050565b613136600160008481526020019081526020016000206002015461313161342c565b6128a1565b613175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161316c906157bd565b60405180910390fd5b61317f82826134c8565b5050565b60035481565b6131916132c4565b846009600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008590508073ffffffffffffffffffffffffffffffffffffffff1663bba8185a868686866040518563ffffffff1660e01b815260040161322794939291906153a2565b600060405180830381600087803b15801561324157600080fd5b505af1158015613255573d6000803e3d6000fd5b50505050505050505050565b6132696132c4565b613271613616565b565b6000809054906101000a900460ff16156132c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132b9906157fd565b60405180910390fd5b565b6132d16000801b336128a1565b613310576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133079061589d565b60405180910390fd5b565b61331f6000801b336128a1565b806133435750613342604051613334906152c1565b6040518091039020336128a1565b5b613382576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133799061565d565b60405180910390fd5b565b60006133c683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613671565b905092915050565b6133eb6040516133dd906152c1565b6040518091039020336128a1565b61342a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134219061587d565b60405180910390fd5b565b600033905090565b61345c81600160008581526020019081526020016000206000016136cc90919063ffffffff16565b156134c45761346961342c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6134f081600160008581526020019081526020016000206000016136fc90919063ffffffff16565b15613558576134fd61342c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b613564613273565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336040516135ad91906152f1565b60405180910390a1565b60006135c6836000018361372c565b60001c905092915050565b60006135f9836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613799565b905092915050565b600061360f826000016137bc565b9050919050565b61361e6137cd565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa3360405161366791906152f1565b60405180910390a1565b60008383111582906136b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136b091906155db565b60405180910390fd5b5060008385039050809150509392505050565b60006136f4836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61381d565b905092915050565b6000613724836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61388d565b905092915050565b600081836000018054905011613777576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161376e9061563d565b60405180910390fd5b82600001828154811061378657fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b6000809054906101000a900460ff1661381b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613812906156bd565b60405180910390fd5b565b60006138298383613799565b613882578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613887565b600090505b92915050565b6000808360010160008481526020019081526020016000205490506000811461396957600060018203905060006001866000018054905003905060008660000182815481106138d857fe5b90600052602060002001549050808760000184815481106138f557fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061392d57fe5b6001900381819060005260206000200160009055905586600101600087815260200190815260200160002060009055600194505050505061396f565b60009150505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106139b657803560ff19168380011785556139e4565b828001600101855582156139e4579182015b828111156139e35782358255916020019190600101906139c8565b5b5090506139f19190613bcf565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613a3657805160ff1916838001178555613a64565b82800160010185558215613a64579182015b82811115613a63578251825591602001919060010190613a48565b5b509050613a719190613bcf565b5090565b828054828255906000526020600020908101928215613aee579160200282015b82811115613aed5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190613a95565b5b509050613afb9190613bf4565b5090565b6040518060800160405280600060ff168152602001600067ffffffffffffffff16815260200160008019168152602001600080191681525090565b6040518060c001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613b7457fe5b8152602001600081525090565b6040518060e001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613bbb57fe5b815260200160008152602001606081525090565b613bf191905b80821115613bed576000816000905550600101613bd5565b5090565b90565b613c3491905b80821115613c3057600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101613bfa565b5090565b90565b600081359050613c4681615d73565b92915050565b600081359050613c5b81615d8a565b92915050565b60008083601f840112613c7357600080fd5b8235905067ffffffffffffffff811115613c8c57600080fd5b602083019150836020820283011115613ca457600080fd5b9250929050565b60008083601f840112613cbd57600080fd5b8235905067ffffffffffffffff811115613cd657600080fd5b602083019150836020820283011115613cee57600080fd5b9250929050565b600081359050613d0481615da1565b92915050565b600081359050613d1981615db8565b92915050565b60008083601f840112613d3157600080fd5b8235905067ffffffffffffffff811115613d4a57600080fd5b602083019150836001820283011115613d6257600080fd5b9250929050565b600082601f830112613d7a57600080fd5b8135613d8d613d8882615a7f565b615a52565b91508082526020830160208301858383011115613da957600080fd5b613db4838284615cd5565b50505092915050565b600082601f830112613dce57600080fd5b8135613de1613ddc82615aab565b615a52565b91508082526020830160208301858383011115613dfd57600080fd5b613e08838284615cd5565b50505092915050565b600081359050613e2081615dcf565b92915050565b600081359050613e3581615de6565b92915050565b600081359050613e4a81615dfd565b92915050565b600081359050613e5f81615e14565b92915050565b600060208284031215613e7757600080fd5b6000613e8584828501613c37565b91505092915050565b600060208284031215613ea057600080fd5b6000613eae84828501613c4c565b91505092915050565b60008060408385031215613eca57600080fd5b6000613ed885828601613c37565b9250506020613ee985828601613c37565b9150509250929050565b60008060008060808587031215613f0957600080fd5b6000613f1787828801613c37565b9450506020613f2887828801613c37565b9350506040613f3987828801613c37565b9250506060613f4a87828801613e11565b91505092959194509250565b600080600060608486031215613f6b57600080fd5b6000613f7986828701613c37565b9350506020613f8a86828701613cf5565b9250506040613f9b86828701613c37565b9150509250925092565b600080600080600060a08688031215613fbd57600080fd5b6000613fcb88828901613c37565b9550506020613fdc88828901613cf5565b9450506040613fed88828901613c37565b9350506060613ffe88828901613d0a565b925050608061400f88828901613d0a565b9150509295509295909350565b6000806000806040858703121561403257600080fd5b600085013567ffffffffffffffff81111561404c57600080fd5b61405887828801613c61565b9450945050602085013567ffffffffffffffff81111561407757600080fd5b61408387828801613cab565b925092505092959194509250565b6000602082840312156140a357600080fd5b60006140b184828501613cf5565b91505092915050565b600080604083850312156140cd57600080fd5b60006140db85828601613cf5565b92505060206140ec85828601613c37565b9150509250929050565b6000806040838503121561410957600080fd5b600061411785828601613cf5565b925050602061412885828601613e11565b9150509250929050565b60006020828403121561414457600080fd5b600082013567ffffffffffffffff81111561415e57600080fd5b61416a84828501613d69565b91505092915050565b60006020828403121561418557600080fd5b600082013567ffffffffffffffff81111561419f57600080fd5b6141ab84828501613dbd565b91505092915050565b6000602082840312156141c657600080fd5b60006141d484828501613e11565b91505092915050565b600080604083850312156141f057600080fd5b60006141fe85828601613e26565b925050602061420f85828601613e50565b9150509250929050565b6000806040838503121561422c57600080fd5b600061423a85828601613e3b565b925050602061424b85828601613cf5565b9150509250929050565b60008060006060848603121561426a57600080fd5b600061427886828701613e3b565b935050602061428986828701613cf5565b925050604061429a86828701613c37565b9150509250925092565b6000602082840312156142b657600080fd5b60006142c484828501613e50565b91505092915050565b600080600080606085870312156142e357600080fd5b60006142f187828801613e50565b945050602061430287828801613cf5565b935050604085013567ffffffffffffffff81111561431f57600080fd5b61432b87828801613d1f565b925092505092959194509250565b60008060006060848603121561434e57600080fd5b600061435c86828701613e50565b935050602061436d86828701613e26565b925050604061437e86828701613cf5565b9150509250925092565b6000806000806080858703121561439e57600080fd5b60006143ac87828801613e50565b94505060206143bd87828801613e26565b93505060406143ce87828801613cf5565b92505060606143df87828801613cf5565b91505092959194509250565b6000806000806080858703121561440157600080fd5b600061440f87828801613e50565b945050602061442087828801613e26565b935050604061443187828801613cf5565b925050606085013567ffffffffffffffff81111561444e57600080fd5b61445a87828801613dbd565b91505092959194509250565b60008060008060006080868803121561447e57600080fd5b600061448c88828901613e50565b955050602061449d88828901613e26565b945050604086013567ffffffffffffffff8111156144ba57600080fd5b6144c688828901613d1f565b935093505060606144d988828901613cf5565b9150509295509295909350565b60006144f2838361450d565b60208301905092915050565b61450781615c7b565b82525050565b61451681615b8f565b82525050565b61452581615b8f565b82525050565b61453c61453782615b8f565b615d17565b82525050565b600061454d82615afc565b6145578185615b35565b935061456283615ad7565b8060005b8381101561459357815161457a88826144e6565b975061458583615b28565b925050600181019050614566565b5085935050505092915050565b6145a981615bb3565b82525050565b6145b881615bbf565b82525050565b6145c781615bbf565b82525050565b6145d681615bc9565b82525050565b60006145e88385615b46565b93506145f5838584615cd5565b6145fe83615d3b565b840190509392505050565b60006146158385615b57565b9350614622838584615cd5565b82840190509392505050565b600061463982615b07565b6146438185615b46565b9350614653818560208601615ce4565b61465c81615d3b565b840191505092915050565b61467081615c8d565b82525050565b61467f81615c8d565b82525050565b61468e81615c9f565b82525050565b61469d81615c9f565b82525050565b60006146ae82615b1d565b6146b88185615b73565b93506146c8818560208601615ce4565b6146d181615d3b565b840191505092915050565b60006146e782615b1d565b6146f18185615b84565b9350614701818560208601615ce4565b80840191505092915050565b600061471882615b12565b6147228185615b62565b9350614732818560208601615ce4565b61473b81615d3b565b840191505092915050565b600061475182615b12565b61475b8185615b73565b935061476b818560208601615ce4565b61477481615d3b565b840191505092915050565b60008154600181166000811461479c57600181146147c257614806565b607f60028304166147ad8187615b73565b955060ff198316865260208601935050614806565b600282046147d08187615b73565b95506147db85615ae7565b60005b828110156147fd578154818901526001820191506020810190506147de565b80880195505050505b505092915050565b60008154600181166000811461482b576001811461485057614894565b607f600283041661483c8187615b84565b955060ff1983168652808601935050614894565b6002820461485e8187615b84565b955061486985615ae7565b60005b8281101561488b5781548189015260018201915060208101905061486c565b82880195505050505b505092915050565b60006148a9601c83615b73565b91507f70726f706f73616c20616c7265616479207472616e73666572726564000000006000830152602082019050919050565b60006148e9601683615b73565b91507f70726f706f73616c206973206e6f7420616374697665000000000000000000006000830152602082019050919050565b6000614929602283615b73565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f64730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061498f601e83615b73565b91507f73656e646572206973206e6f742072656c61796572206f722061646d696e00006000830152602082019050919050565b60006149cf602f83615b73565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000614a35601583615b73565b91507f72656c6179657220616c726561647920766f74656400000000000000000000006000830152602082019050919050565b6000614a75601483615b73565b91507f5061757361626c653a206e6f74207061757365640000000000000000000000006000830152602082019050919050565b6000614ab5601f83615b73565b91507f6164647220646f65736e277420686176652072656c6179657220726f6c6521006000830152602082019050919050565b6000614af5601c83615b73565b91507f5661756c7450726f706f73616c20616c726561647920616374697665000000006000830152602082019050919050565b6000614b35601b83615b73565b91507f5661756c7450726f706f73616c206973206e6f742061637469766500000000006000830152602082019050919050565b6000614b75601683615b73565b91507f496e636f72726563742066656520737570706c696564000000000000000000006000830152602082019050919050565b6000614bb5601b83615b73565b91507f6461746120646f65736e2774206d6174636820646174616861736800000000006000830152602082019050919050565b6000614bf5601a83615b73565b91507f50726f706f73616c20616c72656164792063616e63656c6c65640000000000006000830152602082019050919050565b6000614c35602083615b73565b91507f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726000830152602082019050919050565b6000614c75603083615b73565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f207265766f6b65000000000000000000000000000000006020830152604082019050919050565b6000614cdb602083615b73565b91507f50726f706f73616c206e6f7420617420657870697279207468726573686f6c646000830152602082019050919050565b6000614d1b601083615b73565b91507f5061757361626c653a20706175736564000000000000000000000000000000006000830152602082019050919050565b6000614d5b601e83615b73565b91507f6164647220616c7265616479206861732072656c6179657220726f6c652100006000830152602082019050919050565b6000614d9b601683615b73565b91507f70726f706f73616c206973206e6f7420706173736564000000000000000000006000830152602082019050919050565b6000614ddb602a83615b73565b91507f70726f706f73616c20616c7265616479207061737365642f657865637574656460008301527f2f63616e63656c6c6564000000000000000000000000000000000000000000006020830152604082019050919050565b6000614e41602083615b73565b91507f73656e64657220646f65736e277420686176652072656c6179657220726f6c656000830152602082019050919050565b6000614e81601e83615b73565b91507f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006000830152602082019050919050565b6000614ec1602883615b73565b91507f5661756c7450726f706f73616c20616c726561647920706173736564206f722060008301527f65786563757465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614f27601f83615b73565b91507f43757272656e742066656520697320657175616c20746f206e657720666565006000830152602082019050919050565b6000614f67600c83615b84565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b6000614fa7601183615b73565b91507f6461746168617368206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000614fe7601983615b73565b91507f6e6f2068616e646c657220666f72207265736f757263654944000000000000006000830152602082019050919050565b6000615027602f83615b73565b91507f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008301527f20726f6c657320666f722073656c6600000000000000000000000000000000006020830152604082019050919050565b608082016000820151615096600085018261524b565b5060208201516150a9602085018261522d565b5060408201516150bc60408501826145af565b5060608201516150cf60608501826145af565b50505050565b600060c0830160008301516150ed60008601826145af565b50602083015161510060208601826145af565b50604083015184820360408601526151188282614542565b915050606083015184820360608601526151328282614542565b91505060808301516151476080860182614667565b5060a083015161515a60a086018261520f565b508091505092915050565b600060e08301600083015161517d60008601826145af565b50602083015161519060208601826145af565b50604083015184820360408601526151a88282614542565b915050606083015184820360608601526151c28282614542565b91505060808301516151d76080860182614685565b5060a08301516151ea60a086018261520f565b5060c083015184820360c0860152615202828261470d565b9150508091505092915050565b61521881615c3b565b82525050565b61522781615c3b565b82525050565b61523681615c45565b82525050565b61524581615c45565b82525050565b61525481615c6e565b82525050565b61526381615c6e565b82525050565b6000615275828661452b565b601482019150615286828486614609565b9150819050949350505050565b600061529f82846146dc565b915081905092915050565b60006152b6828461480e565b915081905092915050565b60006152cc82614f5a565b9150819050919050565b60006020820190506152eb600083018461451c565b92915050565b600060208201905061530660008301846144fe565b92915050565b6000606082019050615321600083018661451c565b61532e602083018561451c565b61533b604083018461521e565b949350505050565b600060208201905061535860008301846145a0565b92915050565b600060208201905061537360008301846145be565b92915050565b600060408201905061538e60008301856145be565b61539b602083018461451c565b9392505050565b60006080820190506153b760008301876145be565b6153c4602083018661451c565b6153d160408301856145cd565b6153de60608301846145cd565b95945050505050565b60006040820190506153fc60008301856145be565b61540960208301846145be565b9392505050565b600060808201905061542560008301876145be565b61543260208301866145be565b61543f6040830185614676565b61544c606083018461521e565b95945050505050565b600060a08201905061546a60008301886145be565b61547760208301876145be565b6154846040830186614694565b615491606083018561521e565b81810360808301526154a38184614746565b90509695505050505050565b60006060820190506154c460008301866145be565b6154d160208301856145be565b81810360408301526154e381846146a3565b9050949350505050565b600060608201905061550260008301866145be565b61550f60208301856145be565b8181036040830152615521818461477f565b9050949350505050565b600060408201905061554060008301866145be565b81810360208301526155538184866145dc565b9050949350505050565b600060a08201905061557260008301896145be565b61557f602083018861525a565b61558c604083018761523c565b61559960608301866144fe565b81810360808301526155ac8184866145dc565b9050979650505050505050565b600060208201905081810360008301526155d3818461462e565b905092915050565b600060208201905081810360008301526155f581846146a3565b905092915050565b600060208201905081810360008301526156168161489c565b9050919050565b60006020820190508181036000830152615636816148dc565b9050919050565b600060208201905081810360008301526156568161491c565b9050919050565b6000602082019050818103600083015261567681614982565b9050919050565b60006020820190508181036000830152615696816149c2565b9050919050565b600060208201905081810360008301526156b681614a28565b9050919050565b600060208201905081810360008301526156d681614a68565b9050919050565b600060208201905081810360008301526156f681614aa8565b9050919050565b6000602082019050818103600083015261571681614ae8565b9050919050565b6000602082019050818103600083015261573681614b28565b9050919050565b6000602082019050818103600083015261575681614b68565b9050919050565b6000602082019050818103600083015261577681614ba8565b9050919050565b6000602082019050818103600083015261579681614be8565b9050919050565b600060208201905081810360008301526157b681614c28565b9050919050565b600060208201905081810360008301526157d681614c68565b9050919050565b600060208201905081810360008301526157f681614cce565b9050919050565b6000602082019050818103600083015261581681614d0e565b9050919050565b6000602082019050818103600083015261583681614d4e565b9050919050565b6000602082019050818103600083015261585681614d8e565b9050919050565b6000602082019050818103600083015261587681614dce565b9050919050565b6000602082019050818103600083015261589681614e34565b9050919050565b600060208201905081810360008301526158b681614e74565b9050919050565b600060208201905081810360008301526158d681614eb4565b9050919050565b600060208201905081810360008301526158f681614f1a565b9050919050565b6000602082019050818103600083015261591681614f9a565b9050919050565b6000602082019050818103600083015261593681614fda565b9050919050565b600060208201905081810360008301526159568161501a565b9050919050565b60006080820190506159726000830184615080565b92915050565b6000602082019050818103600083015261599281846150d5565b905092915050565b600060208201905081810360008301526159b48184615165565b905092915050565b60006020820190506159d1600083018461521e565b92915050565b60006020820190506159ec600083018461523c565b92915050565b6000602082019050615a07600083018461525a565b92915050565b6000608082019050615a22600083018761525a565b615a2f602083018661523c565b615a3c60408301856145be565b615a4960608301846145be565b95945050505050565b6000604051905081810181811067ffffffffffffffff82111715615a7557600080fd5b8060405250919050565b600067ffffffffffffffff821115615a9657600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115615ac257600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000615b9a82615c1b565b9050919050565b6000615bac82615c1b565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050615c0382615d59565b919050565b6000819050615c1682615d66565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600068ffffffffffffffffff82169050919050565b600060ff82169050919050565b6000615c8682615cb1565b9050919050565b6000615c9882615bf5565b9050919050565b6000615caa82615c08565b9050919050565b6000615cbc82615cc3565b9050919050565b6000615cce82615c1b565b9050919050565b82818337600083830152505050565b60005b83811015615d02578082015181840152602081019050615ce7565b83811115615d11576000848401525b50505050565b6000615d2282615d29565b9050919050565b6000615d3482615d4c565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b60058110615d6357fe5b50565b60058110615d7057fe5b50565b615d7c81615b8f565b8114615d8757600080fd5b50565b615d9381615ba1565b8114615d9e57600080fd5b50565b615daa81615bbf565b8114615db557600080fd5b50565b615dc181615bc9565b8114615dcc57600080fd5b50565b615dd881615c3b565b8114615de357600080fd5b50565b615def81615c45565b8114615dfa57600080fd5b50565b615e0681615c59565b8114615e1157600080fd5b50565b615e1d81615c6e565b8114615e2857600080fd5b5056fea264697066735822122088f9f47af612d3d87e1b15b1d3e5cd2d68c5a04e120119222b509b98b72ba4bc64736f6c63430006040033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, chainID uint8, initialRelayers []common.Address, initialRelayerThreshold *big.Int, fee *big.Int, expiry *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, chainID, initialRelayers, initialRelayerThreshold, fee, expiry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeCaller) ChainID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeSession) ChainID() (uint8, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeCallerSession) ChainID() (uint8, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCaller) DepositCounts(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCallerSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 uint8) ([]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCallerSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCallerSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCallerSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCallerSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_proposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ProposedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCallerSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCaller) TotalRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalRelayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// TxProposals is a free data retrieval call binding the contract method 0x85717b42.
//
// Solidity: function _txProposals(string ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeCaller) TxProposals(opts *bind.CallOpts, arg0 string) (struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_txProposals", arg0)

	outstruct := new(struct {
		OriginChainID uint8
		DepositNonce  uint64
		ResourceID    [32]byte
		DataHash      [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginChainID = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.DepositNonce = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ResourceID = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TxProposals is a free data retrieval call binding the contract method 0x85717b42.
//
// Solidity: function _txProposals(string ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeSession) TxProposals(arg0 string) (struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}, error) {
	return _Bridge.Contract.TxProposals(&_Bridge.CallOpts, arg0)
}

// TxProposals is a free data retrieval call binding the contract method 0x85717b42.
//
// Solidity: function _txProposals(string ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeCallerSession) TxProposals(arg0 string) (struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}, error) {
	return _Bridge.Contract.TxProposals(&_Bridge.CallOpts, arg0)
}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, string _txId)
func (_Bridge *BridgeCaller) VaultProposals(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxId          string
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_vaultProposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
		TxId          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ProposedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TxId = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, string _txId)
func (_Bridge *BridgeSession) VaultProposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxId          string
}, error) {
	return _Bridge.Contract.VaultProposals(&_Bridge.CallOpts, arg0, arg1)
}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, string _txId)
func (_Bridge *BridgeCallerSession) VaultProposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxId          string
}, error) {
	return _Bridge.Contract.VaultProposals(&_Bridge.CallOpts, arg0, arg1)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCaller) GetProposal(opts *bind.CallOpts, originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeSession) GetProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCallerSession) GetProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetTxProposalRecord is a free data retrieval call binding the contract method 0x20857acc.
//
// Solidity: function getTxProposalRecord(string txId) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeCaller) GetTxProposalRecord(opts *bind.CallOpts, txId string) (BridgeProposalRecord, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getTxProposalRecord", txId)

	if err != nil {
		return *new(BridgeProposalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposalRecord)).(*BridgeProposalRecord)

	return out0, err

}

// GetTxProposalRecord is a free data retrieval call binding the contract method 0x20857acc.
//
// Solidity: function getTxProposalRecord(string txId) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeSession) GetTxProposalRecord(txId string) (BridgeProposalRecord, error) {
	return _Bridge.Contract.GetTxProposalRecord(&_Bridge.CallOpts, txId)
}

// GetTxProposalRecord is a free data retrieval call binding the contract method 0x20857acc.
//
// Solidity: function getTxProposalRecord(string txId) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeCallerSession) GetTxProposalRecord(txId string) (BridgeProposalRecord, error) {
	return _Bridge.Contract.GetTxProposalRecord(&_Bridge.CallOpts, txId)
}

// GetVaultProposal is a free data retrieval call binding the contract method 0xce6cfa0f.
//
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,string))
func (_Bridge *BridgeCaller) GetVaultProposal(opts *bind.CallOpts, originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeVaultProposal, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getVaultProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeVaultProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeVaultProposal)).(*BridgeVaultProposal)

	return out0, err

}

// GetVaultProposal is a free data retrieval call binding the contract method 0xce6cfa0f.
//
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,string))
func (_Bridge *BridgeSession) GetVaultProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeVaultProposal, error) {
	return _Bridge.Contract.GetVaultProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetVaultProposal is a free data retrieval call binding the contract method 0xce6cfa0f.
//
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,string))
func (_Bridge *BridgeCallerSession) GetVaultProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeVaultProposal, error) {
	return _Bridge.Contract.GetVaultProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCaller) IsRelayer(opts *bind.CallOpts, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isRelayer", relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCallerSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminAddRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminAddRelayer", relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminRemoveRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminRemoveRelayer", relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactor) AdminSetGenericResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetGenericResource", handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactorSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetVault is a paid mutator transaction binding the contract method 0x110b7e0b.
//
// Solidity: function adminSetVault(address handlerAddress, address vaultAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetVault(opts *bind.TransactOpts, handlerAddress common.Address, vaultAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetVault", handlerAddress, vaultAddress)
}

// AdminSetVault is a paid mutator transaction binding the contract method 0x110b7e0b.
//
// Solidity: function adminSetVault(address handlerAddress, address vaultAddress) returns()
func (_Bridge *BridgeSession) AdminSetVault(handlerAddress common.Address, vaultAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetVault(&_Bridge.TransactOpts, handlerAddress, vaultAddress)
}

// AdminSetVault is a paid mutator transaction binding the contract method 0x110b7e0b.
//
// Solidity: function adminSetVault(address handlerAddress, address vaultAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetVault(handlerAddress common.Address, vaultAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetVault(&_Bridge.TransactOpts, handlerAddress, vaultAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) CancelProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "cancelProposal", chainID, depositNonce, dataHash)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) CancelProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x17f03ce5.
//
// Solidity: function cancelProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) CancelProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x18b257ed.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txId) returns()
func (_Bridge *BridgeTransactor) CreateVaultProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte, txId string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "createVaultProposal", chainID, depositNonce, dataHash, txId)
}

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x18b257ed.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txId) returns()
func (_Bridge *BridgeSession) CreateVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txId string) (*types.Transaction, error) {
	return _Bridge.Contract.CreateVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txId)
}

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x18b257ed.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txId) returns()
func (_Bridge *BridgeTransactorSession) CreateVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txId string) (*types.Transaction, error) {
	return _Bridge.Contract.CreateVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txId)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeSession) Deposit(destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) Deposit(destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactor) ExecuteProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeProposal", chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeSession) ExecuteProposal(chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactorSession) ExecuteProposal(chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// ExecuteVaultProposal is a paid mutator transaction binding the contract method 0x47b8f8a4.
//
// Solidity: function executeVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) ExecuteVaultProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeVaultProposal", chainID, depositNonce, dataHash)
}

// ExecuteVaultProposal is a paid mutator transaction binding the contract method 0x47b8f8a4.
//
// Solidity: function executeVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) ExecuteVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// ExecuteVaultProposal is a paid mutator transaction binding the contract method 0x47b8f8a4.
//
// Solidity: function executeVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) ExecuteVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// PassVaultProposal is a paid mutator transaction binding the contract method 0x6580a71a.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) PassVaultProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "passVaultProposal", chainID, depositNonce, dataHash)
}

// PassVaultProposal is a paid mutator transaction binding the contract method 0x6580a71a.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) PassVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.PassVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// PassVaultProposal is a paid mutator transaction binding the contract method 0x6580a71a.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) PassVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.PassVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactor) TransferFunds(opts *bind.TransactOpts, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferFunds", addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactorSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) VoteProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteProposal", chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) VoteProposal(chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) VoteProposal(chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	DepositNonce       uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (*BridgeDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Bridge contract.
type BridgeProposalEventIterator struct {
	Event *BridgeProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalEvent represents a ProposalEvent event raised by the Bridge contract.
type BridgeProposalEvent struct {
	OriginChainID uint8
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalEventIterator{contract: _Bridge.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *BridgeProposalEvent, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalEvent)
				if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalEvent is a log parse operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) ParseProposalEvent(log types.Log) (*BridgeProposalEvent, error) {
	event := new(BridgeProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Bridge contract.
type BridgeProposalVoteIterator struct {
	Event *BridgeProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalVote represents a ProposalVote event raised by the Bridge contract.
type BridgeProposalVote struct {
	OriginChainID uint8
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) FilterProposalVote(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalVoteIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVoteIterator{contract: _Bridge.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeProposalVote, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVote is a log parse operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) ParseProposalVote(log types.Log) (*BridgeProposalVote, error) {
	event := new(BridgeProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the Bridge contract.
type BridgeRelayerAddedIterator struct {
	Event *BridgeRelayerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerAdded represents a RelayerAdded event raised by the Bridge contract.
type BridgeRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerAdded(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerAddedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerAddedIterator{contract: _Bridge.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *BridgeRelayerAdded, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerAdded)
				if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerAdded(log types.Log) (*BridgeRelayerAdded, error) {
	event := new(BridgeRelayerAdded)
	if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the Bridge contract.
type BridgeRelayerRemovedIterator struct {
	Event *BridgeRelayerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerRemoved represents a RelayerRemoved event raised by the Bridge contract.
type BridgeRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerRemoved(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerRemovedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerRemovedIterator{contract: _Bridge.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *BridgeRelayerRemoved, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerRemoved)
				if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerRemoved(log types.Log) (*BridgeRelayerRemoved, error) {
	event := new(BridgeRelayerRemoved)
	if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Bridge contract.
type BridgeRelayerThresholdChangedIterator struct {
	Event *BridgeRelayerThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Bridge contract.
type BridgeRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*BridgeRelayerThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerThresholdChangedIterator{contract: _Bridge.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeRelayerThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) ParseRelayerThresholdChanged(log types.Log) (*BridgeRelayerThresholdChanged, error) {
	event := new(BridgeRelayerThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeVaultProposalEventIterator is returned from FilterVaultProposalEvent and is used to iterate over the raw logs and unpacked data for VaultProposalEvent events raised by the Bridge contract.
type BridgeVaultProposalEventIterator struct {
	Event *BridgeVaultProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeVaultProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeVaultProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeVaultProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeVaultProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeVaultProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeVaultProposalEvent represents a VaultProposalEvent event raised by the Bridge contract.
type BridgeVaultProposalEvent struct {
	OriginChainID uint8
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
	TxId          string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVaultProposalEvent is a free log retrieval operation binding the contract event 0xd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b5.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, string txId)
func (_Bridge *BridgeFilterer) FilterVaultProposalEvent(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeVaultProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "VaultProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeVaultProposalEventIterator{contract: _Bridge.contract, event: "VaultProposalEvent", logs: logs, sub: sub}, nil
}

// WatchVaultProposalEvent is a free log subscription operation binding the contract event 0xd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b5.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, string txId)
func (_Bridge *BridgeFilterer) WatchVaultProposalEvent(opts *bind.WatchOpts, sink chan<- *BridgeVaultProposalEvent, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "VaultProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeVaultProposalEvent)
				if err := _Bridge.contract.UnpackLog(event, "VaultProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultProposalEvent is a log parse operation binding the contract event 0xd952684087b1cb4dd542500339d62a5e317e9421c682b0aa2f4ca6593041e5b5.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, string txId)
func (_Bridge *BridgeFilterer) ParseVaultProposalEvent(log types.Log) (*BridgeVaultProposalEvent, error) {
	event := new(BridgeVaultProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "VaultProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
