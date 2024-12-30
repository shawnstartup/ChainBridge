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
	TxIdHash      [32]byte
	TxKey         string
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txIdHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"VaultProposalEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_txProposals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_vaultProposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txIdHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_txKey\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminAddRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminRemoveRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"adminSetVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txIdHash\",\"type\":\"bytes32\"}],\"name\":\"getTxProposalRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBridge.ProposalRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getVaultProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.VaultProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txIdHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_txKey\",\"type\":\"string\"}],\"internalType\":\"structBridge.VaultProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txIdHash\",\"type\":\"bytes32\"}],\"name\":\"createVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"txKey\",\"type\":\"string\"}],\"name\":\"passVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"executeVaultProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200646f3803806200646f8339818101604052810190620000379190620004bc565b60008060006101000a81548160ff02191690831515021790555084600260006101000a81548160ff021916908360ff160217905550826003819055508160068190555080600781905550620000966000801b336200013460201b60201c565b620000c0604051620000a89062000601565b60405180910390206000801b6200014a60201b60201c565b60005b8451811015620001285762000108604051620000df9062000601565b6040518091039020868381518110620000f457fe5b60200260200101516200016960201b60201c565b6004600081548092919060010191905055508080600101915050620000c3565b50505050505062000746565b620001468282620001f860201b60201c565b5050565b8060016000848152602001908152602001600020600201819055505050565b620001a06001600084815260200190815260200160002060020154620001946200029c60201b60201c565b620002a460201b60201c565b620001e2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001d99062000618565b60405180910390fd5b620001f48282620001f860201b60201c565b5050565b620002278160016000858152602001908152602001600020600001620002dd60201b620036c61790919060201c565b1562000298576200023d6200029c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6000620002d582600160008681526020019081526020016000206000016200031560201b620035cb1790919060201c565b905092915050565b60006200030d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200034d60201b60201c565b905092915050565b600062000345836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003c760201b60201c565b905092915050565b6000620003618383620003c760201b60201c565b620003bc578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620003c1565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620003fb81620006f8565b92915050565b600082601f8301126200041357600080fd5b81516200042a620004248262000668565b6200063a565b915081818352602084019350602081019050838560208402820111156200045057600080fd5b60005b83811015620004845781620004698882620003ea565b84526020840193506020830192505060018101905062000453565b5050505092915050565b6000815190506200049f8162000712565b92915050565b600081519050620004b6816200072c565b92915050565b600080600080600060a08688031215620004d557600080fd5b6000620004e588828901620004a5565b955050602086015167ffffffffffffffff8111156200050357600080fd5b620005118882890162000401565b945050604062000524888289016200048e565b935050606062000537888289016200048e565b92505060806200054a888289016200048e565b9150509295509295909350565b600062000566602f8362000691565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000620005ce600c83620006a2565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b60006200060e82620005bf565b9150819050919050565b60006020820190508181036000830152620006338162000557565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200065e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200068057600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b6000620006ba82620006c1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6200070381620006ad565b81146200070f57600080fd5b50565b6200071d81620006e1565b81146200072957600080fd5b50565b6200073781620006eb565b81146200074357600080fd5b50565b615d1980620007566000396000f3fe6080604052600436106102885760003560e01c80638c0c26311161015a578063beab7131116100c1578063ce6cfa0f1161007a578063ce6cfa0f14610a1e578063d547741f14610a5b578063d7a9cd7914610a84578063e8437ee714610aaf578063e87b285014610ad8578063ffaac0eb14610b1557610288565b8063beab71311461090e578063c5b37c2214610939578063c5ec897014610964578063ca15c8731461098f578063cb10f215146109cc578063cdb0f73a146109f557610288565b80639d5773e0116101135780639d5773e0146107e75780639d82dd6314610812578063a217fddf1461083b578063a9cf69fa14610866578063ae527b6e146108a3578063b4b04f32146108cc57610288565b80638c0c2631146106b05780639010d07c146106d957806391c404ac1461071657806391d148541461073f578063926d7d7f1461077c5780639b267216146107a757610288565b806347b8f8a4116101fe5780635e1fab0f116101b75780635e1fab0f146105a2578063780cf004146105cb5780637febe63f146105f4578063802aabe81461063157806380ae1c281461065c57806384db809f1461067357610288565b806347b8f8a41461046b5780634b0b919d146104945780634e056005146104d157806350598719146104fa578063541d55481461053a5780635c975abb1461057757610288565b80632f2ff15d116102505780632f2ff15d1461036157806336568abe1461038a5780633754cc4c146103b35780633ee7094a146103dc5780634454b20d146104195780634603ae381461044257610288565b806305e2ca171461028d578063110b7e0b146102a957806317f03ce5146102d25780631ff013f1146102fb578063248a9ca314610324575b600080fd5b6102a760048036038101906102a291906141fc565b610b2c565b005b3480156102b557600080fd5b506102d060048036038101906102cb9190613e68565b610d89565b005b3480156102de57600080fd5b506102f960048036038101906102f49190614268565b610e06565b005b34801561030757600080fd5b50610322600480360381019061031d91906142b7565b610fb3565b005b34801561033057600080fd5b5061034b60048036038101906103469190614042565b611734565b60405161035891906151ce565b60405180910390f35b34801561036d57600080fd5b506103886004803603810190610383919061406b565b611754565b005b34801561039657600080fd5b506103b160048036038101906103ac919061406b565b6117c8565b005b3480156103bf57600080fd5b506103da60048036038101906103d591906142b7565b61184b565b005b3480156103e857600080fd5b5061040360048036038101906103fe919061410c565b611ac6565b604051610410919061549d565b60405180910390f35b34801561042557600080fd5b50610440600480360381019061043b9190614395565b611b83565b005b34801561044e57600080fd5b5061046960048036038101906104649190613fcd565b611ec4565b005b34801561047757600080fd5b50610492600480360381019061048d9190614268565b611f6a565b005b3480156104a057600080fd5b506104bb60048036038101906104b691906141d3565b6121ad565b6040516104c891906158bb565b60405180910390f35b3480156104dd57600080fd5b506104f860048036038101906104f391906140e3565b6121d4565b005b34801561050657600080fd5b50610521600480360381019061051c9190614148565b612213565b6040516105319493929190615362565b60405180910390f35b34801561054657600080fd5b50610561600480360381019061055c9190613e16565b61225d565b60405161056e91906151b3565b60405180910390f35b34801561058357600080fd5b5061058c612283565b60405161059991906151b3565b60405180910390f35b3480156105ae57600080fd5b506105c960048036038101906105c49190613e16565b612299565b005b3480156105d757600080fd5b506105f260048036038101906105ed9190613ea4565b6122be565b005b34801561060057600080fd5b5061061b60048036038101906106169190614184565b612341565b60405161062891906151b3565b60405180910390f35b34801561063d57600080fd5b5061064661237d565b60405161065391906158a0565b60405180910390f35b34801561066857600080fd5b50610671612383565b005b34801561067f57600080fd5b5061069a60048036038101906106959190614042565b612395565b6040516106a79190615146565b60405180910390f35b3480156106bc57600080fd5b506106d760048036038101906106d29190613e68565b6123c8565b005b3480156106e557600080fd5b5061070060048036038101906106fb91906140a7565b612445565b60405161070d9190615146565b60405180910390f35b34801561072257600080fd5b5061073d600480360381019061073891906140e3565b612477565b005b34801561074b57600080fd5b506107666004803603810190610761919061406b565b6124ce565b60405161077391906151b3565b60405180910390f35b34801561078857600080fd5b50610791612500565b60405161079e91906151ce565b60405180910390f35b3480156107b357600080fd5b506107ce60048036038101906107c99190614042565b612517565b6040516107de94939291906158f1565b60405180910390f35b3480156107f357600080fd5b506107fc612568565b60405161080991906158a0565b60405180910390f35b34801561081e57600080fd5b5061083960048036038101906108349190613e16565b61256e565b005b34801561084757600080fd5b50610850612648565b60405161085d91906151ce565b60405180910390f35b34801561087257600080fd5b5061088d60048036038101906108889190614268565b61264f565b60405161089a919061585c565b60405180910390f35b3480156108af57600080fd5b506108ca60048036038101906108c5919061431a565b612830565b005b3480156108d857600080fd5b506108f360048036038101906108ee9190614148565b612b0b565b604051610905969594939291906153a7565b60405180910390f35b34801561091a57600080fd5b50610923612bf9565b60405161093091906158d6565b60405180910390f35b34801561094557600080fd5b5061094e612c0c565b60405161095b91906158a0565b60405180910390f35b34801561097057600080fd5b50610979612c12565b60405161098691906158a0565b60405180910390f35b34801561099b57600080fd5b506109b660048036038101906109b19190614042565b612c18565b6040516109c391906158a0565b60405180910390f35b3480156109d857600080fd5b506109f360048036038101906109ee9190613f07565b612c3f565b005b348015610a0157600080fd5b50610a1c6004803603810190610a179190613e16565b612d11565b005b348015610a2a57600080fd5b50610a456004803603810190610a409190614268565b612deb565b604051610a52919061587e565b60405180910390f35b348015610a6757600080fd5b50610a826004803603810190610a7d919061406b565b613079565b005b348015610a9057600080fd5b50610a996130ed565b604051610aa691906158a0565b60405180910390f35b348015610abb57600080fd5b50610ad66004803603810190610ad19190613f56565b6130f3565b005b348015610ae457600080fd5b50610aff6004803603810190610afa9190614042565b6131cb565b604051610b0c9190615841565b60405180910390f35b348015610b2157600080fd5b50610b2a61325b565b005b610b3461326d565b6006543414610b78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6f90615621565b60405180910390fd5b60006009600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1790615681565b60405180910390fd5b6000600860008760ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590508383600a60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008960ff1660ff1681526020019081526020016000209190610cc992919061396f565b5060008290508073ffffffffffffffffffffffffffffffffffffffff166338995da9878985338a8a6040518763ffffffff1660e01b8152600401610d1296959493929190615441565b600060405180830381600087803b158015610d2c57600080fd5b505af1158015610d40573d6000803e3d6000fd5b505050508167ffffffffffffffff16868860ff167fdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed860405160405180910390a450505050505050565b610d916132be565b60008290508073ffffffffffffffffffffffffffffffffffffffff16636817031b836040518263ffffffff1660e01b8152600401610dcf9190615146565b600060405180830381600087803b158015610de957600080fd5b505af1158015610dfd573d6000803e3d6000fd5b50505050505050565b610e0e61330c565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600480811115610e7b57fe5b8160040160009054906101000a900460ff166004811115610e9857fe5b1415610ed9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed090615661565b60405180910390fd5b600754610eea43836005015461337e565b11610f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f21906156c1565b60405180910390fd5b60048160040160006101000a81548160ff02191690836004811115610f4b57fe5b0217905550600480811115610f5c57fe5b8467ffffffffffffffff168660ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041784600001548560010154604051610fa4929190615257565b60405180910390a45050505050565b610fbb6133c8565b610fc361326d565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156110c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110be90615801565b60405180910390fd5b60018160040160009054906101000a900460ff1660048111156110e657fe5b1115611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111e90615741565b60405180910390fd5b600c60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156111ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e390615581565b60405180910390fd5b60008160040160009054906101000a900460ff16600481111561120b57fe5b141561142b57600560008154600101919050819055506040518060c0016040528085815260200184815260200160016040519080825280602002602001820160405280156112685781602001602082028036833780820191505090505b508152602001600060405190808252806020026020018201604052801561129e5781602001602082028036833780820191505090505b508152602001600160048111156112b157fe5b815260200143815250600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906113249291906139ef565b5060608201518160030190805190602001906113419291906139ef565b5060808201518160040160006101000a81548160ff0219169083600481111561136657fe5b021790555060a08201518160050155905050338160020160008154811061138957fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160048111156113de57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417878760405161141e929190615257565b60405180910390a461156e565b60075461143c43836005015461337e565b11156114c15760048160040160006101000a81548160ff0219169083600481111561146357fe5b021790555060048081111561147457fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516114b4929190615257565b60405180910390a461156d565b80600101548314611507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114fe906157e1565b60405180910390fd5b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b60048081111561157a57fe5b8160040160009054906101000a900460ff16600481111561159757fe5b1461172c576001600c60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff16600481111561164957fe5b8567ffffffffffffffff168760ff167f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b95743436408760405161168791906151ce565b60405180910390a460016003541115806116aa5750600354816002018054905010155b1561172b5760028160040160006101000a81548160ff021916908360048111156116d057fe5b0217905550600260048111156116e257fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab6504178787604051611722929190615257565b60405180910390a45b5b505050505050565b600060016000838152602001908152602001600020600201549050919050565b61177b6001600084815260200190815260200160002060020154611776613426565b6124ce565b6117ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117b190615561565b60405180910390fd5b6117c4828261342e565b5050565b6117d0613426565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461183d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183490615821565b60405180910390fd5b61184782826134c2565b5050565b61185361330c565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050600260048111156118c157fe5b8160040160009054906101000a900460ff1660048111156118de57fe5b1461191e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191590615721565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002090506001600481111561196957fe5b8160040160009054906101000a900460ff16600481111561198657fe5b141580156119bc57506002600481111561199c57fe5b8160040160009054906101000a900460ff1660048111156119b957fe5b14155b80156119f05750600360048111156119d057fe5b8160040160009054906101000a900460ff1660048111156119ed57fe5b14155b611a2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a26906155e1565b60405180910390fd5b60018160040160006101000a81548160ff02191690836004811115611a5057fe5b021790555083816006018190555060016004811115611a6b57fe5b8667ffffffffffffffff168860ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e42998560000154866001015489604051611ab593929190615318565b60405180910390a450505050505050565b600a602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611b7b5780601f10611b5057610100808354040283529160200191611b7b565b820191906000526020600020905b815481529060010190602001808311611b5e57829003601f168201915b505050505081565b611b8b6133c8565b611b9361326d565b60006009600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008660ff1660088767ffffffffffffffff1668ffffffffffffffffff16901b1790506000828686604051602001611c0593929190615107565b6040516020818303038152906040528051906020012090506000600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020905060006004811115611c6857fe5b8160040160009054906101000a900460ff166004811115611c8557fe5b1415611cc6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611cbd90615501565b60405180910390fd5b60026004811115611cd357fe5b8160040160009054906101000a900460ff166004811115611cf057fe5b14611d30576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d27906154e1565b60405180910390fd5b80600101548214611d76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d6d90615641565b60405180910390fd5b60038160040160006101000a81548160ff02191690836004811115611d9757fe5b02179055506000600960008360000154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663e248cff283600001548a8a6040518463ffffffff1660e01b8152600401611e199392919061540f565b600060405180830381600087803b158015611e3357600080fd5b505af1158015611e47573d6000803e3d6000fd5b505050508160040160009054906101000a900460ff166004811115611e6857fe5b8967ffffffffffffffff168b60ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041785600001548660010154604051611eb0929190615257565b60405180910390a450505050505050505050565b611ecc6132be565b60008090505b84849050811015611f6357848482818110611ee957fe5b9050602002016020810190611efe9190613e3f565b73ffffffffffffffffffffffffffffffffffffffff166108fc848484818110611f2357fe5b905060200201359081150290604051600060405180830381858888f19350505050158015611f55573d6000803e3d6000fd5b508080600101915050611ed2565b5050505050565b611f7261330c565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000848152602001908152602001600020905060026004811115611fe057fe5b8160040160009054906101000a900460ff166004811115611ffd57fe5b1461203d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161203490615721565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002090506001600481111561208857fe5b8160040160009054906101000a900460ff1660048111156120a557fe5b14806120d85750600260048111156120b957fe5b8160040160009054906101000a900460ff1660048111156120d657fe5b145b612117576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161210e90615601565b60405180910390fd5b60038160040160006101000a81548160ff0219169083600481111561213857fe5b02179055506003600481111561214a57fe5b8567ffffffffffffffff168760ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e42998560000154866001015486600601548760070160405161219d94939291906152cc565b60405180910390a4505050505050565b60086020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b6121dc6132be565b80600381905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a250565b600b602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff16908060050154905084565b600061227c60405161226e90615131565b6040518091039020836124ce565b9050919050565b60008060009054906101000a900460ff16905090565b6122a16132be565b6122ae6000801b82611754565b6122bb6000801b336117c8565b50565b6122c66132be565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d9caed128585856040518463ffffffff1660e01b81526004016123089392919061517c565b600060405180830381600087803b15801561232257600080fd5b505af1158015612336573d6000803e3d6000fd5b505050505050505050565b600c602052826000526040600020602052816000526040600020602052806000526040600020600092509250509054906101000a900460ff1681565b60045481565b61238b6132be565b612393613556565b565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6123d06132be565b60008290508073ffffffffffffffffffffffffffffffffffffffff166307b7ed99836040518263ffffffff1660e01b815260040161240e9190615146565b600060405180830381600087803b15801561242857600080fd5b505af115801561243c573d6000803e3d6000fd5b50505050505050565b600061246f82600160008681526020019081526020016000206000016135b190919063ffffffff16565b905092915050565b61247f6132be565b8060065414156124c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124bb906157c1565b60405180910390fd5b8060068190555050565b60006124f882600160008681526020019081526020016000206000016135cb90919063ffffffff16565b905092915050565b60405161250c90615131565b604051809103902081565b600e6020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900467ffffffffffffffff16908060010154908060020154905084565b60055481565b6125766132be565b61259360405161258590615131565b6040518091039020826124ce565b6125d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125c9906155c1565b60405180910390fd5b6125ef6040516125e190615131565b604051809103902082613079565b8073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a26004600081548092919060019003919050555050565b6000801b81565b612657613a79565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600b60008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561275857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161270e575b50505050508152602001600382018054806020026020016040519081016040528092919081815260200182805480156127e657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161279c575b505050505081526020016004820160009054906101000a900460ff16600481111561280d57fe5b600481111561281857fe5b81526020016005820154815250509150509392505050565b61283861330c565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050600260048111156128a657fe5b8160040160009054906101000a900460ff1660048111156128c357fe5b14612903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128fa90615721565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002090506002600481111561294e57fe5b8160040160009054906101000a900460ff16600481111561296b57fe5b141580156129a157506003600481111561298157fe5b8160040160009054906101000a900460ff16600481111561299e57fe5b14155b6129e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129d7906157a1565b60405180910390fd5b60028160040160006101000a81548160ff02191690836004811115612a0157fe5b021790555083816007019080519060200190612a1e929190613ac0565b506000600e60008360060154815260200190815260200160002090508260000154816001018190555082600101548160020181905550868160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550878160000160006101000a81548160ff021916908360ff16021790555060026004811115612aa957fe5b8767ffffffffffffffff168960ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e42998660000154876001015487600601548b604051612af99493929190615280565b60405180910390a45050505050505050565b600d602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff1690806005015490806006015490806007018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612bef5780601f10612bc457610100808354040283529160200191612bef565b820191906000526020600020905b815481529060010190602001808311612bd257829003601f168201915b5050505050905086565b600260009054906101000a900460ff1681565b60065481565b60075481565b6000612c38600160008481526020019081526020016000206000016135fb565b9050919050565b612c476132be565b826009600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508073ffffffffffffffffffffffffffffffffffffffff1663b8fa373684846040518363ffffffff1660e01b8152600401612cd99291906151e9565b600060405180830381600087803b158015612cf357600080fd5b505af1158015612d07573d6000803e3d6000fd5b5050505050505050565b612d196132be565b612d36604051612d2890615131565b6040518091039020826124ce565b15612d76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d6d90615701565b60405180910390fd5b612d93604051612d8590615131565b604051809103902082611754565b8073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a260046000815480929190600101919050555050565b612df3613b40565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600d60008268ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060405180610100016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015612ef557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612eab575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015612f8357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612f39575b505050505081526020016004820160009054906101000a900460ff166004811115612faa57fe5b6004811115612fb557fe5b81526020016005820154815260200160068201548152602001600782018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156130665780601f1061303b57610100808354040283529160200191613066565b820191906000526020600020905b81548152906001019060200180831161304957829003601f168201915b5050505050815250509150509392505050565b6130a0600160008481526020019081526020016000206002015461309b613426565b6124ce565b6130df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130d6906156a1565b60405180910390fd5b6130e982826134c2565b5050565b60035481565b6130fb6132be565b846009600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008590508073ffffffffffffffffffffffffffffffffffffffff1663bba8185a868686866040518563ffffffff1660e01b81526004016131919493929190615212565b600060405180830381600087803b1580156131ab57600080fd5b505af11580156131bf573d6000803e3d6000fd5b50505050505050505050565b6131d3613b99565b600e60008381526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b6132636132be565b61326b613610565b565b6000809054906101000a900460ff16156132bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132b3906156e1565b60405180910390fd5b565b6132cb6000801b336124ce565b61330a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161330190615781565b60405180910390fd5b565b6133196000801b336124ce565b8061333d575061333c60405161332e90615131565b6040518091039020336124ce565b5b61337c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161337390615541565b60405180910390fd5b565b60006133c083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f77000081525061366b565b905092915050565b6133e56040516133d790615131565b6040518091039020336124ce565b613424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161341b90615761565b60405180910390fd5b565b600033905090565b61345681600160008581526020019081526020016000206000016136c690919063ffffffff16565b156134be57613463613426565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6134ea81600160008581526020019081526020016000206000016136f690919063ffffffff16565b15613552576134f7613426565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b61355e61326d565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336040516135a79190615161565b60405180910390a1565b60006135c08360000183613726565b60001c905092915050565b60006135f3836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613793565b905092915050565b6000613609826000016137b6565b9050919050565b6136186137c7565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516136619190615161565b60405180910390a1565b60008383111582906136b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136aa91906154bf565b60405180910390fd5b5060008385039050809150509392505050565b60006136ee836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613817565b905092915050565b600061371e836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613887565b905092915050565b600081836000018054905011613771576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161376890615521565b60405180910390fd5b82600001828154811061378057fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b6000809054906101000a900460ff16613815576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161380c906155a1565b60405180910390fd5b565b60006138238383613793565b61387c578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613881565b600090505b92915050565b6000808360010160008481526020019081526020016000205490506000811461396357600060018203905060006001866000018054905003905060008660000182815481106138d257fe5b90600052602060002001549050808760000184815481106138ef57fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061392757fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050613969565b60009150505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106139b057803560ff19168380011785556139de565b828001600101855582156139de579182015b828111156139dd5782358255916020019190600101906139c2565b5b5090506139eb9190613bd4565b5090565b828054828255906000526020600020908101928215613a68579160200282015b82811115613a675782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190613a0f565b5b509050613a759190613bf9565b5090565b6040518060c001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613ab357fe5b8152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613b0157805160ff1916838001178555613b2f565b82800160010185558215613b2f579182015b82811115613b2e578251825591602001919060010190613b13565b5b509050613b3c9190613bd4565b5090565b6040518061010001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613b7b57fe5b81526020016000815260200160008019168152602001606081525090565b6040518060800160405280600060ff168152602001600067ffffffffffffffff16815260200160008019168152602001600080191681525090565b613bf691905b80821115613bf2576000816000905550600101613bda565b5090565b90565b613c3991905b80821115613c3557600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101613bff565b5090565b90565b600081359050613c4b81615c2b565b92915050565b600081359050613c6081615c42565b92915050565b60008083601f840112613c7857600080fd5b8235905067ffffffffffffffff811115613c9157600080fd5b602083019150836020820283011115613ca957600080fd5b9250929050565b60008083601f840112613cc257600080fd5b8235905067ffffffffffffffff811115613cdb57600080fd5b602083019150836020820283011115613cf357600080fd5b9250929050565b600081359050613d0981615c59565b92915050565b600081359050613d1e81615c70565b92915050565b60008083601f840112613d3657600080fd5b8235905067ffffffffffffffff811115613d4f57600080fd5b602083019150836001820283011115613d6757600080fd5b9250929050565b600082601f830112613d7f57600080fd5b8135613d92613d8d82615963565b615936565b91508082526020830160208301858383011115613dae57600080fd5b613db9838284615b8d565b50505092915050565b600081359050613dd181615c87565b92915050565b600081359050613de681615c9e565b92915050565b600081359050613dfb81615cb5565b92915050565b600081359050613e1081615ccc565b92915050565b600060208284031215613e2857600080fd5b6000613e3684828501613c3c565b91505092915050565b600060208284031215613e5157600080fd5b6000613e5f84828501613c51565b91505092915050565b60008060408385031215613e7b57600080fd5b6000613e8985828601613c3c565b9250506020613e9a85828601613c3c565b9150509250929050565b60008060008060808587031215613eba57600080fd5b6000613ec887828801613c3c565b9450506020613ed987828801613c3c565b9350506040613eea87828801613c3c565b9250506060613efb87828801613dc2565b91505092959194509250565b600080600060608486031215613f1c57600080fd5b6000613f2a86828701613c3c565b9350506020613f3b86828701613cfa565b9250506040613f4c86828701613c3c565b9150509250925092565b600080600080600060a08688031215613f6e57600080fd5b6000613f7c88828901613c3c565b9550506020613f8d88828901613cfa565b9450506040613f9e88828901613c3c565b9350506060613faf88828901613d0f565b9250506080613fc088828901613d0f565b9150509295509295909350565b60008060008060408587031215613fe357600080fd5b600085013567ffffffffffffffff811115613ffd57600080fd5b61400987828801613c66565b9450945050602085013567ffffffffffffffff81111561402857600080fd5b61403487828801613cb0565b925092505092959194509250565b60006020828403121561405457600080fd5b600061406284828501613cfa565b91505092915050565b6000806040838503121561407e57600080fd5b600061408c85828601613cfa565b925050602061409d85828601613c3c565b9150509250929050565b600080604083850312156140ba57600080fd5b60006140c885828601613cfa565b92505060206140d985828601613dc2565b9150509250929050565b6000602082840312156140f557600080fd5b600061410384828501613dc2565b91505092915050565b6000806040838503121561411f57600080fd5b600061412d85828601613dd7565b925050602061413e85828601613e01565b9150509250929050565b6000806040838503121561415b57600080fd5b600061416985828601613dec565b925050602061417a85828601613cfa565b9150509250929050565b60008060006060848603121561419957600080fd5b60006141a786828701613dec565b93505060206141b886828701613cfa565b92505060406141c986828701613c3c565b9150509250925092565b6000602082840312156141e557600080fd5b60006141f384828501613e01565b91505092915050565b6000806000806060858703121561421257600080fd5b600061422087828801613e01565b945050602061423187828801613cfa565b935050604085013567ffffffffffffffff81111561424e57600080fd5b61425a87828801613d24565b925092505092959194509250565b60008060006060848603121561427d57600080fd5b600061428b86828701613e01565b935050602061429c86828701613dd7565b92505060406142ad86828701613cfa565b9150509250925092565b600080600080608085870312156142cd57600080fd5b60006142db87828801613e01565b94505060206142ec87828801613dd7565b93505060406142fd87828801613cfa565b925050606061430e87828801613cfa565b91505092959194509250565b6000806000806080858703121561433057600080fd5b600061433e87828801613e01565b945050602061434f87828801613dd7565b935050604061436087828801613cfa565b925050606085013567ffffffffffffffff81111561437d57600080fd5b61438987828801613d6e565b91505092959194509250565b6000806000806000608086880312156143ad57600080fd5b60006143bb88828901613e01565b95505060206143cc88828901613dd7565b945050604086013567ffffffffffffffff8111156143e957600080fd5b6143f588828901613d24565b9350935050606061440888828901613cfa565b9150509295509295909350565b6000614421838361443c565b60208301905092915050565b61443681615b33565b82525050565b61444581615a47565b82525050565b61445481615a47565b82525050565b61446b61446682615a47565b615bcf565b82525050565b600061447c826159b4565b61448681856159ed565b93506144918361598f565b8060005b838110156144c25781516144a98882614415565b97506144b4836159e0565b925050600181019050614495565b5085935050505092915050565b6144d881615a6b565b82525050565b6144e781615a77565b82525050565b6144f681615a77565b82525050565b61450581615a81565b82525050565b600061451783856159fe565b9350614524838584615b8d565b61452d83615bf3565b840190509392505050565b60006145448385615a0f565b9350614551838584615b8d565b82840190509392505050565b6000614568826159bf565b61457281856159fe565b9350614582818560208601615b9c565b61458b81615bf3565b840191505092915050565b61459f81615b45565b82525050565b6145ae81615b45565b82525050565b6145bd81615b57565b82525050565b6145cc81615b57565b82525050565b60006145dd826159d5565b6145e78185615a2b565b93506145f7818560208601615b9c565b61460081615bf3565b840191505092915050565b6000614616826159ca565b6146208185615a1a565b9350614630818560208601615b9c565b61463981615bf3565b840191505092915050565b600061464f826159ca565b6146598185615a2b565b9350614669818560208601615b9c565b61467281615bf3565b840191505092915050565b60008154600181166000811461469a57600181146146c057614704565b607f60028304166146ab8187615a2b565b955060ff198316865260208601935050614704565b600282046146ce8187615a2b565b95506146d98561599f565b60005b828110156146fb578154818901526001820191506020810190506146dc565b80880195505050505b505092915050565b6000614719601c83615a2b565b91507f70726f706f73616c20616c7265616479207472616e73666572726564000000006000830152602082019050919050565b6000614759601683615a2b565b91507f70726f706f73616c206973206e6f7420616374697665000000000000000000006000830152602082019050919050565b6000614799602283615a2b565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f64730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006147ff601e83615a2b565b91507f73656e646572206973206e6f742072656c61796572206f722061646d696e00006000830152602082019050919050565b600061483f602f83615a2b565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b60006148a5601583615a2b565b91507f72656c6179657220616c726561647920766f74656400000000000000000000006000830152602082019050919050565b60006148e5601483615a2b565b91507f5061757361626c653a206e6f74207061757365640000000000000000000000006000830152602082019050919050565b6000614925601f83615a2b565b91507f6164647220646f65736e277420686176652072656c6179657220726f6c6521006000830152602082019050919050565b6000614965601c83615a2b565b91507f5661756c7450726f706f73616c20616c726561647920616374697665000000006000830152602082019050919050565b60006149a5601b83615a2b565b91507f5661756c7450726f706f73616c206973206e6f742061637469766500000000006000830152602082019050919050565b60006149e5601683615a2b565b91507f496e636f72726563742066656520737570706c696564000000000000000000006000830152602082019050919050565b6000614a25601b83615a2b565b91507f6461746120646f65736e2774206d6174636820646174616861736800000000006000830152602082019050919050565b6000614a65601a83615a2b565b91507f50726f706f73616c20616c72656164792063616e63656c6c65640000000000006000830152602082019050919050565b6000614aa5602083615a2b565b91507f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726000830152602082019050919050565b6000614ae5603083615a2b565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f207265766f6b65000000000000000000000000000000006020830152604082019050919050565b6000614b4b602083615a2b565b91507f50726f706f73616c206e6f7420617420657870697279207468726573686f6c646000830152602082019050919050565b6000614b8b601083615a2b565b91507f5061757361626c653a20706175736564000000000000000000000000000000006000830152602082019050919050565b6000614bcb601e83615a2b565b91507f6164647220616c7265616479206861732072656c6179657220726f6c652100006000830152602082019050919050565b6000614c0b601683615a2b565b91507f70726f706f73616c206973206e6f7420706173736564000000000000000000006000830152602082019050919050565b6000614c4b602a83615a2b565b91507f70726f706f73616c20616c7265616479207061737365642f657865637574656460008301527f2f63616e63656c6c6564000000000000000000000000000000000000000000006020830152604082019050919050565b6000614cb1600083615a2b565b9150600082019050919050565b6000614ccb602083615a2b565b91507f73656e64657220646f65736e277420686176652072656c6179657220726f6c656000830152602082019050919050565b6000614d0b601e83615a2b565b91507f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006000830152602082019050919050565b6000614d4b602883615a2b565b91507f5661756c7450726f706f73616c20616c726561647920706173736564206f722060008301527f65786563757465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614db1601f83615a2b565b91507f43757272656e742066656520697320657175616c20746f206e657720666565006000830152602082019050919050565b6000614df1600c83615a3c565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b6000614e31601183615a2b565b91507f6461746168617368206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000614e71601983615a2b565b91507f6e6f2068616e646c657220666f72207265736f757263654944000000000000006000830152602082019050919050565b6000614eb1602f83615a2b565b91507f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008301527f20726f6c657320666f722073656c6600000000000000000000000000000000006020830152604082019050919050565b608082016000820151614f2060008501826150e9565b506020820151614f3360208501826150cb565b506040820151614f4660408501826144de565b506060820151614f5960608501826144de565b50505050565b600060c083016000830151614f7760008601826144de565b506020830151614f8a60208601826144de565b5060408301518482036040860152614fa28282614471565b91505060608301518482036060860152614fbc8282614471565b9150506080830151614fd16080860182614596565b5060a0830151614fe460a08601826150ad565b508091505092915050565b60006101008301600083015161500860008601826144de565b50602083015161501b60208601826144de565b50604083015184820360408601526150338282614471565b9150506060830151848203606086015261504d8282614471565b915050608083015161506260808601826145b4565b5060a083015161507560a08601826150ad565b5060c083015161508860c08601826144de565b5060e083015184820360e08601526150a0828261460b565b9150508091505092915050565b6150b681615af3565b82525050565b6150c581615af3565b82525050565b6150d481615afd565b82525050565b6150e381615afd565b82525050565b6150f281615b26565b82525050565b61510181615b26565b82525050565b6000615113828661445a565b601482019150615124828486614538565b9150819050949350505050565b600061513c82614de4565b9150819050919050565b600060208201905061515b600083018461444b565b92915050565b6000602082019050615176600083018461442d565b92915050565b6000606082019050615191600083018661444b565b61519e602083018561444b565b6151ab60408301846150bc565b949350505050565b60006020820190506151c860008301846144cf565b92915050565b60006020820190506151e360008301846144ed565b92915050565b60006040820190506151fe60008301856144ed565b61520b602083018461444b565b9392505050565b600060808201905061522760008301876144ed565b615234602083018661444b565b61524160408301856144fc565b61524e60608301846144fc565b95945050505050565b600060408201905061526c60008301856144ed565b61527960208301846144ed565b9392505050565b600060808201905061529560008301876144ed565b6152a260208301866144ed565b6152af60408301856144ed565b81810360608301526152c181846145d2565b905095945050505050565b60006080820190506152e160008301876144ed565b6152ee60208301866144ed565b6152fb60408301856144ed565b818103606083015261530d818461467d565b905095945050505050565b600060808201905061532d60008301866144ed565b61533a60208301856144ed565b61534760408301846144ed565b818103606083015261535881614ca4565b9050949350505050565b600060808201905061537760008301876144ed565b61538460208301866144ed565b61539160408301856145a5565b61539e60608301846150bc565b95945050505050565b600060c0820190506153bc60008301896144ed565b6153c960208301886144ed565b6153d660408301876145c3565b6153e360608301866150bc565b6153f060808301856144ed565b81810360a08301526154028184614644565b9050979650505050505050565b600060408201905061542460008301866144ed565b818103602083015261543781848661450b565b9050949350505050565b600060a08201905061545660008301896144ed565b61546360208301886150f8565b61547060408301876150da565b61547d606083018661442d565b818103608083015261549081848661450b565b9050979650505050505050565b600060208201905081810360008301526154b7818461455d565b905092915050565b600060208201905081810360008301526154d981846145d2565b905092915050565b600060208201905081810360008301526154fa8161470c565b9050919050565b6000602082019050818103600083015261551a8161474c565b9050919050565b6000602082019050818103600083015261553a8161478c565b9050919050565b6000602082019050818103600083015261555a816147f2565b9050919050565b6000602082019050818103600083015261557a81614832565b9050919050565b6000602082019050818103600083015261559a81614898565b9050919050565b600060208201905081810360008301526155ba816148d8565b9050919050565b600060208201905081810360008301526155da81614918565b9050919050565b600060208201905081810360008301526155fa81614958565b9050919050565b6000602082019050818103600083015261561a81614998565b9050919050565b6000602082019050818103600083015261563a816149d8565b9050919050565b6000602082019050818103600083015261565a81614a18565b9050919050565b6000602082019050818103600083015261567a81614a58565b9050919050565b6000602082019050818103600083015261569a81614a98565b9050919050565b600060208201905081810360008301526156ba81614ad8565b9050919050565b600060208201905081810360008301526156da81614b3e565b9050919050565b600060208201905081810360008301526156fa81614b7e565b9050919050565b6000602082019050818103600083015261571a81614bbe565b9050919050565b6000602082019050818103600083015261573a81614bfe565b9050919050565b6000602082019050818103600083015261575a81614c3e565b9050919050565b6000602082019050818103600083015261577a81614cbe565b9050919050565b6000602082019050818103600083015261579a81614cfe565b9050919050565b600060208201905081810360008301526157ba81614d3e565b9050919050565b600060208201905081810360008301526157da81614da4565b9050919050565b600060208201905081810360008301526157fa81614e24565b9050919050565b6000602082019050818103600083015261581a81614e64565b9050919050565b6000602082019050818103600083015261583a81614ea4565b9050919050565b60006080820190506158566000830184614f0a565b92915050565b600060208201905081810360008301526158768184614f5f565b905092915050565b600060208201905081810360008301526158988184614fef565b905092915050565b60006020820190506158b560008301846150bc565b92915050565b60006020820190506158d060008301846150da565b92915050565b60006020820190506158eb60008301846150f8565b92915050565b600060808201905061590660008301876150f8565b61591360208301866150da565b61592060408301856144ed565b61592d60608301846144ed565b95945050505050565b6000604051905081810181811067ffffffffffffffff8211171561595957600080fd5b8060405250919050565b600067ffffffffffffffff82111561597a57600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000615a5282615ad3565b9050919050565b6000615a6482615ad3565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050615abb82615c11565b919050565b6000819050615ace82615c1e565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600068ffffffffffffffffff82169050919050565b600060ff82169050919050565b6000615b3e82615b69565b9050919050565b6000615b5082615aad565b9050919050565b6000615b6282615ac0565b9050919050565b6000615b7482615b7b565b9050919050565b6000615b8682615ad3565b9050919050565b82818337600083830152505050565b60005b83811015615bba578082015181840152602081019050615b9f565b83811115615bc9576000848401525b50505050565b6000615bda82615be1565b9050919050565b6000615bec82615c04565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b60058110615c1b57fe5b50565b60058110615c2857fe5b50565b615c3481615a47565b8114615c3f57600080fd5b50565b615c4b81615a59565b8114615c5657600080fd5b50565b615c6281615a77565b8114615c6d57600080fd5b50565b615c7981615a81565b8114615c8457600080fd5b50565b615c9081615af3565b8114615c9b57600080fd5b50565b615ca781615afd565b8114615cb257600080fd5b50565b615cbe81615b11565b8114615cc957600080fd5b50565b615cd581615b26565b8114615ce057600080fd5b5056fea26469706673582212209676c8c9684d2ae0c279a786c25e1b668d4c0b7690c61ee3ea5e7b00990a706b64736f6c63430006040033",
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

// TxProposals is a free data retrieval call binding the contract method 0x9b267216.
//
// Solidity: function _txProposals(bytes32 ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeCaller) TxProposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
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

// TxProposals is a free data retrieval call binding the contract method 0x9b267216.
//
// Solidity: function _txProposals(bytes32 ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeSession) TxProposals(arg0 [32]byte) (struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}, error) {
	return _Bridge.Contract.TxProposals(&_Bridge.CallOpts, arg0)
}

// TxProposals is a free data retrieval call binding the contract method 0x9b267216.
//
// Solidity: function _txProposals(bytes32 ) view returns(uint8 _originChainID, uint64 _depositNonce, bytes32 _resourceID, bytes32 _dataHash)
func (_Bridge *BridgeCallerSession) TxProposals(arg0 [32]byte) (struct {
	OriginChainID uint8
	DepositNonce  uint64
	ResourceID    [32]byte
	DataHash      [32]byte
}, error) {
	return _Bridge.Contract.TxProposals(&_Bridge.CallOpts, arg0)
}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, bytes32 _txIdHash, string _txKey)
func (_Bridge *BridgeCaller) VaultProposals(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxIdHash      [32]byte
	TxKey         string
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_vaultProposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
		TxIdHash      [32]byte
		TxKey         string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ProposedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TxIdHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.TxKey = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, bytes32 _txIdHash, string _txKey)
func (_Bridge *BridgeSession) VaultProposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxIdHash      [32]byte
	TxKey         string
}, error) {
	return _Bridge.Contract.VaultProposals(&_Bridge.CallOpts, arg0, arg1)
}

// VaultProposals is a free data retrieval call binding the contract method 0xb4b04f32.
//
// Solidity: function _vaultProposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock, bytes32 _txIdHash, string _txKey)
func (_Bridge *BridgeCallerSession) VaultProposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
	TxIdHash      [32]byte
	TxKey         string
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

// GetTxProposalRecord is a free data retrieval call binding the contract method 0xe87b2850.
//
// Solidity: function getTxProposalRecord(bytes32 txIdHash) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeCaller) GetTxProposalRecord(opts *bind.CallOpts, txIdHash [32]byte) (BridgeProposalRecord, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getTxProposalRecord", txIdHash)

	if err != nil {
		return *new(BridgeProposalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposalRecord)).(*BridgeProposalRecord)

	return out0, err

}

// GetTxProposalRecord is a free data retrieval call binding the contract method 0xe87b2850.
//
// Solidity: function getTxProposalRecord(bytes32 txIdHash) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeSession) GetTxProposalRecord(txIdHash [32]byte) (BridgeProposalRecord, error) {
	return _Bridge.Contract.GetTxProposalRecord(&_Bridge.CallOpts, txIdHash)
}

// GetTxProposalRecord is a free data retrieval call binding the contract method 0xe87b2850.
//
// Solidity: function getTxProposalRecord(bytes32 txIdHash) view returns((uint8,uint64,bytes32,bytes32))
func (_Bridge *BridgeCallerSession) GetTxProposalRecord(txIdHash [32]byte) (BridgeProposalRecord, error) {
	return _Bridge.Contract.GetTxProposalRecord(&_Bridge.CallOpts, txIdHash)
}

// GetVaultProposal is a free data retrieval call binding the contract method 0xce6cfa0f.
//
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,bytes32,string))
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
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,bytes32,string))
func (_Bridge *BridgeSession) GetVaultProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeVaultProposal, error) {
	return _Bridge.Contract.GetVaultProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetVaultProposal is a free data retrieval call binding the contract method 0xce6cfa0f.
//
// Solidity: function getVaultProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256,bytes32,string))
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

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x3754cc4c.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, bytes32 txIdHash) returns()
func (_Bridge *BridgeTransactor) CreateVaultProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte, txIdHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "createVaultProposal", chainID, depositNonce, dataHash, txIdHash)
}

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x3754cc4c.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, bytes32 txIdHash) returns()
func (_Bridge *BridgeSession) CreateVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txIdHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CreateVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txIdHash)
}

// CreateVaultProposal is a paid mutator transaction binding the contract method 0x3754cc4c.
//
// Solidity: function createVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, bytes32 txIdHash) returns()
func (_Bridge *BridgeTransactorSession) CreateVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txIdHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CreateVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txIdHash)
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

// PassVaultProposal is a paid mutator transaction binding the contract method 0xae527b6e.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txKey) returns()
func (_Bridge *BridgeTransactor) PassVaultProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, dataHash [32]byte, txKey string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "passVaultProposal", chainID, depositNonce, dataHash, txKey)
}

// PassVaultProposal is a paid mutator transaction binding the contract method 0xae527b6e.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txKey) returns()
func (_Bridge *BridgeSession) PassVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txKey string) (*types.Transaction, error) {
	return _Bridge.Contract.PassVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txKey)
}

// PassVaultProposal is a paid mutator transaction binding the contract method 0xae527b6e.
//
// Solidity: function passVaultProposal(uint8 chainID, uint64 depositNonce, bytes32 dataHash, string txKey) returns()
func (_Bridge *BridgeTransactorSession) PassVaultProposal(chainID uint8, depositNonce uint64, dataHash [32]byte, txKey string) (*types.Transaction, error) {
	return _Bridge.Contract.PassVaultProposal(&_Bridge.TransactOpts, chainID, depositNonce, dataHash, txKey)
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
	TxIdHash      [32]byte
	TxKey         string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVaultProposalEvent is a free log retrieval operation binding the contract event 0x227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e4299.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, bytes32 txIdHash, string txKey)
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

// WatchVaultProposalEvent is a free log subscription operation binding the contract event 0x227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e4299.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, bytes32 txIdHash, string txKey)
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

// ParseVaultProposalEvent is a log parse operation binding the contract event 0x227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e4299.
//
// Solidity: event VaultProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash, bytes32 txIdHash, string txKey)
func (_Bridge *BridgeFilterer) ParseVaultProposalEvent(log types.Log) (*BridgeVaultProposalEvent, error) {
	event := new(BridgeVaultProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "VaultProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
