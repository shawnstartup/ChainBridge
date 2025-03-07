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
	Bin: "0x60806040523480156200001157600080fd5b50604051620063d6380380620063d68339818101604052810190620000379190620004bc565b60008060006101000a81548160ff02191690831515021790555084600260006101000a81548160ff021916908360ff160217905550826003819055508160068190555080600781905550620000966000801b336200013460201b60201c565b620000c0604051620000a89062000601565b60405180910390206000801b6200014a60201b60201c565b60005b8451811015620001285762000108604051620000df9062000601565b6040518091039020868381518110620000f457fe5b60200260200101516200016960201b60201c565b6004600081548092919060010191905055508080600101915050620000c3565b50505050505062000746565b620001468282620001f860201b60201c565b5050565b8060016000848152602001908152602001600020600201819055505050565b620001a06001600084815260200190815260200160002060020154620001946200029c60201b60201c565b620002a460201b60201c565b620001e2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001d99062000618565b60405180910390fd5b620001f48282620001f860201b60201c565b5050565b620002278160016000858152602001908152602001600020600001620002dd60201b6200368d1790919060201c565b1562000298576200023d6200029c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6000620002d582600160008681526020019081526020016000206000016200031560201b620035921790919060201c565b905092915050565b60006200030d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200034d60201b60201c565b905092915050565b600062000345836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620003c760201b60201c565b905092915050565b6000620003618383620003c760201b60201c565b620003bc578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620003c1565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620003fb81620006f8565b92915050565b600082601f8301126200041357600080fd5b81516200042a620004248262000668565b6200063a565b915081818352602084019350602081019050838560208402820111156200045057600080fd5b60005b83811015620004845781620004698882620003ea565b84526020840193506020830192505060018101905062000453565b5050505092915050565b6000815190506200049f8162000712565b92915050565b600081519050620004b6816200072c565b92915050565b600080600080600060a08688031215620004d557600080fd5b6000620004e588828901620004a5565b955050602086015167ffffffffffffffff8111156200050357600080fd5b620005118882890162000401565b945050604062000524888289016200048e565b935050606062000537888289016200048e565b92505060806200054a888289016200048e565b9150509295509295909350565b600062000566602f8362000691565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b6000620005ce600c83620006a2565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b60006200060e82620005bf565b9150819050919050565b60006020820190508181036000830152620006338162000557565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200065e57600080fd5b8060405250919050565b600067ffffffffffffffff8211156200068057600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b600081905092915050565b6000620006ba82620006c1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6200070381620006ad565b81146200070f57600080fd5b50565b6200071d81620006e1565b81146200072957600080fd5b50565b6200073781620006eb565b81146200074357600080fd5b50565b615c8080620007566000396000f3fe6080604052600436106102885760003560e01c80638c0c26311161015a578063beab7131116100c1578063ce6cfa0f1161007a578063ce6cfa0f14610a1e578063d547741f14610a5b578063d7a9cd7914610a84578063e8437ee714610aaf578063e87b285014610ad8578063ffaac0eb14610b1557610288565b8063beab71311461090e578063c5b37c2214610939578063c5ec897014610964578063ca15c8731461098f578063cb10f215146109cc578063cdb0f73a146109f557610288565b80639d5773e0116101135780639d5773e0146107e75780639d82dd6314610812578063a217fddf1461083b578063a9cf69fa14610866578063ae527b6e146108a3578063b4b04f32146108cc57610288565b80638c0c2631146106b05780639010d07c146106d957806391c404ac1461071657806391d148541461073f578063926d7d7f1461077c5780639b267216146107a757610288565b806347b8f8a4116101fe5780635e1fab0f116101b75780635e1fab0f146105a2578063780cf004146105cb5780637febe63f146105f4578063802aabe81461063157806380ae1c281461065c57806384db809f1461067357610288565b806347b8f8a41461046b5780634b0b919d146104945780634e056005146104d157806350598719146104fa578063541d55481461053a5780635c975abb1461057757610288565b80632f2ff15d116102505780632f2ff15d1461036157806336568abe1461038a5780633754cc4c146103b35780633ee7094a146103dc5780634454b20d146104195780634603ae381461044257610288565b806305e2ca171461028d578063110b7e0b146102a957806317f03ce5146102d25780631ff013f1146102fb578063248a9ca314610324575b600080fd5b6102a760048036038101906102a291906141c3565b610b2c565b005b3480156102b557600080fd5b506102d060048036038101906102cb9190613e2f565b610d89565b005b3480156102de57600080fd5b506102f960048036038101906102f4919061422f565b610e06565b005b34801561030757600080fd5b50610322600480360381019061031d919061427e565b610fb3565b005b34801561033057600080fd5b5061034b60048036038101906103469190614009565b611734565b6040516103589190615155565b60405180910390f35b34801561036d57600080fd5b5061038860048036038101906103839190614032565b611754565b005b34801561039657600080fd5b506103b160048036038101906103ac9190614032565b6117c8565b005b3480156103bf57600080fd5b506103da60048036038101906103d5919061427e565b61184b565b005b3480156103e857600080fd5b5061040360048036038101906103fe91906140d3565b611b44565b6040516104109190615424565b60405180910390f35b34801561042557600080fd5b50610440600480360381019061043b919061435c565b611c01565b005b34801561044e57600080fd5b5061046960048036038101906104649190613f94565b611f09565b005b34801561047757600080fd5b50610492600480360381019061048d919061422f565b611faf565b005b3480156104a057600080fd5b506104bb60048036038101906104b6919061419a565b6121f2565b6040516104c89190615822565b60405180910390f35b3480156104dd57600080fd5b506104f860048036038101906104f391906140aa565b612219565b005b34801561050657600080fd5b50610521600480360381019061051c919061410f565b612258565b60405161053194939291906152e9565b60405180910390f35b34801561054657600080fd5b50610561600480360381019061055c9190613ddd565b6122a2565b60405161056e919061513a565b60405180910390f35b34801561058357600080fd5b5061058c6122c8565b604051610599919061513a565b60405180910390f35b3480156105ae57600080fd5b506105c960048036038101906105c49190613ddd565b6122de565b005b3480156105d757600080fd5b506105f260048036038101906105ed9190613e6b565b612303565b005b34801561060057600080fd5b5061061b6004803603810190610616919061414b565b612386565b604051610628919061513a565b60405180910390f35b34801561063d57600080fd5b506106466123c2565b6040516106539190615807565b60405180910390f35b34801561066857600080fd5b506106716123c8565b005b34801561067f57600080fd5b5061069a60048036038101906106959190614009565b6123da565b6040516106a791906150cd565b60405180910390f35b3480156106bc57600080fd5b506106d760048036038101906106d29190613e2f565b61240d565b005b3480156106e557600080fd5b5061070060048036038101906106fb919061406e565b61248a565b60405161070d91906150cd565b60405180910390f35b34801561072257600080fd5b5061073d600480360381019061073891906140aa565b6124bc565b005b34801561074b57600080fd5b5061076660048036038101906107619190614032565b612513565b604051610773919061513a565b60405180910390f35b34801561078857600080fd5b50610791612545565b60405161079e9190615155565b60405180910390f35b3480156107b357600080fd5b506107ce60048036038101906107c99190614009565b61255c565b6040516107de9493929190615858565b60405180910390f35b3480156107f357600080fd5b506107fc6125ad565b6040516108099190615807565b60405180910390f35b34801561081e57600080fd5b5061083960048036038101906108349190613ddd565b6125b3565b005b34801561084757600080fd5b5061085061268d565b60405161085d9190615155565b60405180910390f35b34801561087257600080fd5b5061088d6004803603810190610888919061422f565b612694565b60405161089a91906157c3565b60405180910390f35b3480156108af57600080fd5b506108ca60048036038101906108c591906142e1565b612875565b005b3480156108d857600080fd5b506108f360048036038101906108ee919061410f565b612ad2565b6040516109059695949392919061532e565b60405180910390f35b34801561091a57600080fd5b50610923612bc0565b604051610930919061583d565b60405180910390f35b34801561094557600080fd5b5061094e612bd3565b60405161095b9190615807565b60405180910390f35b34801561097057600080fd5b50610979612bd9565b6040516109869190615807565b60405180910390f35b34801561099b57600080fd5b506109b660048036038101906109b19190614009565b612bdf565b6040516109c39190615807565b60405180910390f35b3480156109d857600080fd5b506109f360048036038101906109ee9190613ece565b612c06565b005b348015610a0157600080fd5b50610a1c6004803603810190610a179190613ddd565b612cd8565b005b348015610a2a57600080fd5b50610a456004803603810190610a40919061422f565b612db2565b604051610a5291906157e5565b60405180910390f35b348015610a6757600080fd5b50610a826004803603810190610a7d9190614032565b613040565b005b348015610a9057600080fd5b50610a996130b4565b604051610aa69190615807565b60405180910390f35b348015610abb57600080fd5b50610ad66004803603810190610ad19190613f1d565b6130ba565b005b348015610ae457600080fd5b50610aff6004803603810190610afa9190614009565b613192565b604051610b0c91906157a8565b60405180910390f35b348015610b2157600080fd5b50610b2a613222565b005b610b34613234565b6006543414610b78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6f90615588565b60405180910390fd5b60006009600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c17906155e8565b60405180910390fd5b6000600860008760ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590508383600a60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008960ff1660ff1681526020019081526020016000209190610cc9929190613936565b5060008290508073ffffffffffffffffffffffffffffffffffffffff166338995da9878985338a8a6040518763ffffffff1660e01b8152600401610d12969594939291906153c8565b600060405180830381600087803b158015610d2c57600080fd5b505af1158015610d40573d6000803e3d6000fd5b505050508167ffffffffffffffff16868860ff167fdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed860405160405180910390a450505050505050565b610d91613285565b60008290508073ffffffffffffffffffffffffffffffffffffffff16636817031b836040518263ffffffff1660e01b8152600401610dcf91906150cd565b600060405180830381600087803b158015610de957600080fd5b505af1158015610dfd573d6000803e3d6000fd5b50505050505050565b610e0e6132d3565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600480811115610e7b57fe5b8160040160009054906101000a900460ff166004811115610e9857fe5b1415610ed9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed0906155c8565b60405180910390fd5b600754610eea438360050154613345565b11610f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2190615628565b60405180910390fd5b60048160040160006101000a81548160ff02191690836004811115610f4b57fe5b0217905550600480811115610f5c57fe5b8467ffffffffffffffff168660ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041784600001548560010154604051610fa49291906151de565b60405180910390a45050505050565b610fbb61338f565b610fc3613234565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff166009600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156110c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110be90615768565b60405180910390fd5b60018160040160009054906101000a900460ff1660048111156110e657fe5b1115611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111e906156a8565b60405180910390fd5b600c60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156111ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e3906154e8565b60405180910390fd5b60008160040160009054906101000a900460ff16600481111561120b57fe5b141561142b57600560008154600101919050819055506040518060c0016040528085815260200184815260200160016040519080825280602002602001820160405280156112685781602001602082028036833780820191505090505b508152602001600060405190808252806020026020018201604052801561129e5781602001602082028036833780820191505090505b508152602001600160048111156112b157fe5b815260200143815250600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906113249291906139b6565b5060608201518160030190805190602001906113419291906139b6565b5060808201518160040160006101000a81548160ff0219169083600481111561136657fe5b021790555060a08201518160050155905050338160020160008154811061138957fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160048111156113de57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417878760405161141e9291906151de565b60405180910390a461156e565b60075461143c438360050154613345565b11156114c15760048160040160006101000a81548160ff0219169083600481111561146357fe5b021790555060048081111561147457fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516114b49291906151de565b60405180910390a461156d565b80600101548314611507576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114fe90615748565b60405180910390fd5b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b60048081111561157a57fe5b8160040160009054906101000a900460ff16600481111561159757fe5b1461172c576001600c60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff16600481111561164957fe5b8567ffffffffffffffff168760ff167f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640876040516116879190615155565b60405180910390a460016003541115806116aa5750600354816002018054905010155b1561172b5760028160040160006101000a81548160ff021916908360048111156116d057fe5b0217905550600260048111156116e257fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516117229291906151de565b60405180910390a45b5b505050505050565b600060016000838152602001908152602001600020600201549050919050565b61177b60016000848152602001908152602001600020600201546117766133ed565b612513565b6117ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117b1906154c8565b60405180910390fd5b6117c482826133f5565b5050565b6117d06133ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461183d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183490615788565b60405180910390fd5b6118478282613489565b5050565b6118536132d3565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050600260048111156118c157fe5b8160040160009054906101000a900460ff1660048111156118de57fe5b1461191e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161191590615688565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002090506001600481111561196957fe5b8160040160009054906101000a900460ff16600481111561198657fe5b141580156119bc57506002600481111561199c57fe5b8160040160009054906101000a900460ff1660048111156119b957fe5b14155b80156119f05750600360048111156119d057fe5b8160040160009054906101000a900460ff1660048111156119ed57fe5b14155b611a2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a2690615548565b60405180910390fd5b60018160040160006101000a81548160ff02191690836004811115611a5057fe5b02179055508381600601819055506000600e60008360060154815260200190815260200160002090508260000154816001018190555082600101548160020181905550868160000160016101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550878160000160006101000a81548160ff021916908360ff16021790555060016004811115611ae857fe5b8767ffffffffffffffff168960ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e4299866000015487600101548a604051611b329392919061529f565b60405180910390a45050505050505050565b600a602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611bf95780601f10611bce57610100808354040283529160200191611bf9565b820191906000526020600020905b815481529060010190602001808311611bdc57829003601f168201915b505050505081565b611c0961338f565b611c11613234565b60006009600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008660ff1660088767ffffffffffffffff1668ffffffffffffffffff16901b1790506000828686604051602001611c839392919061508e565b6040516020818303038152906040528051906020012090506000600b60008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020905060006004811115611ce657fe5b8160040160009054906101000a900460ff166004811115611d0357fe5b1415611d44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d3b90615468565b60405180910390fd5b80600101548214611d8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d81906155a8565b60405180910390fd5b60026004811115611d9757fe5b8160040160009054906101000a900460ff166004811115611db457fe5b1415611efe5760038160040160006101000a81548160ff02191690836004811115611ddb57fe5b02179055506000600960008360000154815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663e248cff283600001548a8a6040518463ffffffff1660e01b8152600401611e5d93929190615396565b600060405180830381600087803b158015611e7757600080fd5b505af1158015611e8b573d6000803e3d6000fd5b505050508160040160009054906101000a900460ff166004811115611eac57fe5b8967ffffffffffffffff168b60ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041785600001548660010154604051611ef49291906151de565b60405180910390a4505b505050505050505050565b611f11613285565b60008090505b84849050811015611fa857848482818110611f2e57fe5b9050602002016020810190611f439190613e06565b73ffffffffffffffffffffffffffffffffffffffff166108fc848484818110611f6857fe5b905060200201359081150290604051600060405180830381858888f19350505050158015611f9a573d6000803e3d6000fd5b508080600101915050611f17565b5050505050565b611fb76132d3565b60008360ff1660088467ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002090506002600481111561202557fe5b8160040160009054906101000a900460ff16600481111561204257fe5b14612082576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161207990615688565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050600160048111156120cd57fe5b8160040160009054906101000a900460ff1660048111156120ea57fe5b148061211d5750600260048111156120fe57fe5b8160040160009054906101000a900460ff16600481111561211b57fe5b145b61215c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161215390615568565b60405180910390fd5b60038160040160006101000a81548160ff0219169083600481111561217d57fe5b02179055506003600481111561218f57fe5b8567ffffffffffffffff168760ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e4299856000015486600101548660060154876007016040516121e29493929190615253565b60405180910390a4505050505050565b60086020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b612221613285565b80600381905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a250565b600b602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff16908060050154905084565b60006122c16040516122b3906150b8565b604051809103902083612513565b9050919050565b60008060009054906101000a900460ff16905090565b6122e6613285565b6122f36000801b82611754565b6123006000801b336117c8565b50565b61230b613285565b60008490508073ffffffffffffffffffffffffffffffffffffffff1663d9caed128585856040518463ffffffff1660e01b815260040161234d93929190615103565b600060405180830381600087803b15801561236757600080fd5b505af115801561237b573d6000803e3d6000fd5b505050505050505050565b600c602052826000526040600020602052816000526040600020602052806000526040600020600092509250509054906101000a900460ff1681565b60045481565b6123d0613285565b6123d861351d565b565b60096020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612415613285565b60008290508073ffffffffffffffffffffffffffffffffffffffff166307b7ed99836040518263ffffffff1660e01b815260040161245391906150cd565b600060405180830381600087803b15801561246d57600080fd5b505af1158015612481573d6000803e3d6000fd5b50505050505050565b60006124b4826001600086815260200190815260200160002060000161357890919063ffffffff16565b905092915050565b6124c4613285565b806006541415612509576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161250090615728565b60405180910390fd5b8060068190555050565b600061253d826001600086815260200190815260200160002060000161359290919063ffffffff16565b905092915050565b604051612551906150b8565b604051809103902081565b600e6020528060005260406000206000915090508060000160009054906101000a900460ff16908060000160019054906101000a900467ffffffffffffffff16908060010154908060020154905084565b60055481565b6125bb613285565b6125d86040516125ca906150b8565b604051809103902082612513565b612617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161260e90615528565b60405180910390fd5b612634604051612626906150b8565b604051809103902082613040565b8073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a26004600081548092919060019003919050555050565b6000801b81565b61269c613a40565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600b60008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060c001604052908160008201548152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561279d57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612753575b505050505081526020016003820180548060200260200160405190810160405280929190818152602001828054801561282b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116127e1575b505050505081526020016004820160009054906101000a900460ff16600481111561285257fe5b600481111561285d57fe5b81526020016005820154815250509150509392505050565b61287d6132d3565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000600b60008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008581526020019081526020016000209050600260048111156128eb57fe5b8160040160009054906101000a900460ff16600481111561290857fe5b14612948576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161293f90615688565b60405180910390fd5b6000600d60008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002090506002600481111561299357fe5b8160040160009054906101000a900460ff1660048111156129b057fe5b141580156129e65750600360048111156129c657fe5b8160040160009054906101000a900460ff1660048111156129e357fe5b14155b612a25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a1c90615708565b60405180910390fd5b60028160040160006101000a81548160ff02191690836004811115612a4657fe5b021790555083816007019080519060200190612a63929190613a87565b5060026004811115612a7157fe5b8667ffffffffffffffff168860ff167f227e817446962d42b6f4495d22603cbd95531d9f1d5737550db8bc02597e42998560000154866001015486600601548a604051612ac19493929190615207565b60405180910390a450505050505050565b600d602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff1690806005015490806006015490806007018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612bb65780601f10612b8b57610100808354040283529160200191612bb6565b820191906000526020600020905b815481529060010190602001808311612b9957829003601f168201915b5050505050905086565b600260009054906101000a900460ff1681565b60065481565b60075481565b6000612bff600160008481526020019081526020016000206000016135c2565b9050919050565b612c0e613285565b826009600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508073ffffffffffffffffffffffffffffffffffffffff1663b8fa373684846040518363ffffffff1660e01b8152600401612ca0929190615170565b600060405180830381600087803b158015612cba57600080fd5b505af1158015612cce573d6000803e3d6000fd5b5050505050505050565b612ce0613285565b612cfd604051612cef906150b8565b604051809103902082612513565b15612d3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d3490615668565b60405180910390fd5b612d5a604051612d4c906150b8565b604051809103902082611754565b8073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a260046000815480929190600101919050555050565b612dba613b07565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050600d60008268ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060405180610100016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015612ebc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612e72575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015612f4a57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612f00575b505050505081526020016004820160009054906101000a900460ff166004811115612f7157fe5b6004811115612f7c57fe5b81526020016005820154815260200160068201548152602001600782018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561302d5780601f106130025761010080835404028352916020019161302d565b820191906000526020600020905b81548152906001019060200180831161301057829003601f168201915b5050505050815250509150509392505050565b61306760016000848152602001908152602001600020600201546130626133ed565b612513565b6130a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161309d90615608565b60405180910390fd5b6130b08282613489565b5050565b60035481565b6130c2613285565b846009600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008590508073ffffffffffffffffffffffffffffffffffffffff1663bba8185a868686866040518563ffffffff1660e01b81526004016131589493929190615199565b600060405180830381600087803b15801561317257600080fd5b505af1158015613186573d6000803e3d6000fd5b50505050505050505050565b61319a613b60565b600e60008381526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b61322a613285565b6132326135d7565b565b6000809054906101000a900460ff1615613283576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161327a90615648565b60405180910390fd5b565b6132926000801b33612513565b6132d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132c8906156e8565b60405180910390fd5b565b6132e06000801b33612513565b8061330457506133036040516132f5906150b8565b604051809103902033612513565b5b613343576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161333a906154a8565b60405180910390fd5b565b600061338783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613632565b905092915050565b6133ac60405161339e906150b8565b604051809103902033612513565b6133eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133e2906156c8565b60405180910390fd5b565b600033905090565b61341d816001600085815260200190815260200160002060000161368d90919063ffffffff16565b156134855761342a6133ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6134b181600160008581526020019081526020016000206000016136bd90919063ffffffff16565b15613519576134be6133ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b613525613234565b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583360405161356e91906150e8565b60405180910390a1565b600061358783600001836136ed565b60001c905092915050565b60006135ba836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61375a565b905092915050565b60006135d08260000161377d565b9050919050565b6135df61378e565b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa3360405161362891906150e8565b60405180910390a1565b600083831115829061367a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136719190615446565b60405180910390fd5b5060008385039050809150509392505050565b60006136b5836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6137de565b905092915050565b60006136e5836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61384e565b905092915050565b600081836000018054905011613738576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161372f90615488565b60405180910390fd5b82600001828154811061374757fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081600001805490509050919050565b6000809054906101000a900460ff166137dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137d390615508565b60405180910390fd5b565b60006137ea838361375a565b613843578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613848565b600090505b92915050565b6000808360010160008481526020019081526020016000205490506000811461392a576000600182039050600060018660000180549050039050600086600001828154811061389957fe5b90600052602060002001549050808760000184815481106138b657fe5b90600052602060002001819055506001830187600101600083815260200190815260200160002081905550866000018054806138ee57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050613930565b60009150505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061397757803560ff19168380011785556139a5565b828001600101855582156139a5579182015b828111156139a4578235825591602001919060010190613989565b5b5090506139b29190613b9b565b5090565b828054828255906000526020600020908101928215613a2f579160200282015b82811115613a2e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906139d6565b5b509050613a3c9190613bc0565b5090565b6040518060c001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613a7a57fe5b8152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10613ac857805160ff1916838001178555613af6565b82800160010185558215613af6579182015b82811115613af5578251825591602001919060010190613ada565b5b509050613b039190613b9b565b5090565b6040518061010001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613b4257fe5b81526020016000815260200160008019168152602001606081525090565b6040518060800160405280600060ff168152602001600067ffffffffffffffff16815260200160008019168152602001600080191681525090565b613bbd91905b80821115613bb9576000816000905550600101613ba1565b5090565b90565b613c0091905b80821115613bfc57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101613bc6565b5090565b90565b600081359050613c1281615b92565b92915050565b600081359050613c2781615ba9565b92915050565b60008083601f840112613c3f57600080fd5b8235905067ffffffffffffffff811115613c5857600080fd5b602083019150836020820283011115613c7057600080fd5b9250929050565b60008083601f840112613c8957600080fd5b8235905067ffffffffffffffff811115613ca257600080fd5b602083019150836020820283011115613cba57600080fd5b9250929050565b600081359050613cd081615bc0565b92915050565b600081359050613ce581615bd7565b92915050565b60008083601f840112613cfd57600080fd5b8235905067ffffffffffffffff811115613d1657600080fd5b602083019150836001820283011115613d2e57600080fd5b9250929050565b600082601f830112613d4657600080fd5b8135613d59613d54826158ca565b61589d565b91508082526020830160208301858383011115613d7557600080fd5b613d80838284615af4565b50505092915050565b600081359050613d9881615bee565b92915050565b600081359050613dad81615c05565b92915050565b600081359050613dc281615c1c565b92915050565b600081359050613dd781615c33565b92915050565b600060208284031215613def57600080fd5b6000613dfd84828501613c03565b91505092915050565b600060208284031215613e1857600080fd5b6000613e2684828501613c18565b91505092915050565b60008060408385031215613e4257600080fd5b6000613e5085828601613c03565b9250506020613e6185828601613c03565b9150509250929050565b60008060008060808587031215613e8157600080fd5b6000613e8f87828801613c03565b9450506020613ea087828801613c03565b9350506040613eb187828801613c03565b9250506060613ec287828801613d89565b91505092959194509250565b600080600060608486031215613ee357600080fd5b6000613ef186828701613c03565b9350506020613f0286828701613cc1565b9250506040613f1386828701613c03565b9150509250925092565b600080600080600060a08688031215613f3557600080fd5b6000613f4388828901613c03565b9550506020613f5488828901613cc1565b9450506040613f6588828901613c03565b9350506060613f7688828901613cd6565b9250506080613f8788828901613cd6565b9150509295509295909350565b60008060008060408587031215613faa57600080fd5b600085013567ffffffffffffffff811115613fc457600080fd5b613fd087828801613c2d565b9450945050602085013567ffffffffffffffff811115613fef57600080fd5b613ffb87828801613c77565b925092505092959194509250565b60006020828403121561401b57600080fd5b600061402984828501613cc1565b91505092915050565b6000806040838503121561404557600080fd5b600061405385828601613cc1565b925050602061406485828601613c03565b9150509250929050565b6000806040838503121561408157600080fd5b600061408f85828601613cc1565b92505060206140a085828601613d89565b9150509250929050565b6000602082840312156140bc57600080fd5b60006140ca84828501613d89565b91505092915050565b600080604083850312156140e657600080fd5b60006140f485828601613d9e565b925050602061410585828601613dc8565b9150509250929050565b6000806040838503121561412257600080fd5b600061413085828601613db3565b925050602061414185828601613cc1565b9150509250929050565b60008060006060848603121561416057600080fd5b600061416e86828701613db3565b935050602061417f86828701613cc1565b925050604061419086828701613c03565b9150509250925092565b6000602082840312156141ac57600080fd5b60006141ba84828501613dc8565b91505092915050565b600080600080606085870312156141d957600080fd5b60006141e787828801613dc8565b94505060206141f887828801613cc1565b935050604085013567ffffffffffffffff81111561421557600080fd5b61422187828801613ceb565b925092505092959194509250565b60008060006060848603121561424457600080fd5b600061425286828701613dc8565b935050602061426386828701613d9e565b925050604061427486828701613cc1565b9150509250925092565b6000806000806080858703121561429457600080fd5b60006142a287828801613dc8565b94505060206142b387828801613d9e565b93505060406142c487828801613cc1565b92505060606142d587828801613cc1565b91505092959194509250565b600080600080608085870312156142f757600080fd5b600061430587828801613dc8565b945050602061431687828801613d9e565b935050604061432787828801613cc1565b925050606085013567ffffffffffffffff81111561434457600080fd5b61435087828801613d35565b91505092959194509250565b60008060008060006080868803121561437457600080fd5b600061438288828901613dc8565b955050602061439388828901613d9e565b945050604086013567ffffffffffffffff8111156143b057600080fd5b6143bc88828901613ceb565b935093505060606143cf88828901613cc1565b9150509295509295909350565b60006143e88383614403565b60208301905092915050565b6143fd81615a9a565b82525050565b61440c816159ae565b82525050565b61441b816159ae565b82525050565b61443261442d826159ae565b615b36565b82525050565b60006144438261591b565b61444d8185615954565b9350614458836158f6565b8060005b8381101561448957815161447088826143dc565b975061447b83615947565b92505060018101905061445c565b5085935050505092915050565b61449f816159d2565b82525050565b6144ae816159de565b82525050565b6144bd816159de565b82525050565b6144cc816159e8565b82525050565b60006144de8385615965565b93506144eb838584615af4565b6144f483615b5a565b840190509392505050565b600061450b8385615976565b9350614518838584615af4565b82840190509392505050565b600061452f82615926565b6145398185615965565b9350614549818560208601615b03565b61455281615b5a565b840191505092915050565b61456681615aac565b82525050565b61457581615aac565b82525050565b61458481615abe565b82525050565b61459381615abe565b82525050565b60006145a48261593c565b6145ae8185615992565b93506145be818560208601615b03565b6145c781615b5a565b840191505092915050565b60006145dd82615931565b6145e78185615981565b93506145f7818560208601615b03565b61460081615b5a565b840191505092915050565b600061461682615931565b6146208185615992565b9350614630818560208601615b03565b61463981615b5a565b840191505092915050565b6000815460018116600081146146615760018114614687576146cb565b607f60028304166146728187615992565b955060ff1983168652602086019350506146cb565b600282046146958187615992565b95506146a085615906565b60005b828110156146c2578154818901526001820191506020810190506146a3565b80880195505050505b505092915050565b60006146e0601683615992565b91507f70726f706f73616c206973206e6f7420616374697665000000000000000000006000830152602082019050919050565b6000614720602283615992565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f64730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614786601e83615992565b91507f73656e646572206973206e6f742072656c61796572206f722061646d696e00006000830152602082019050919050565b60006147c6602f83615992565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b600061482c601583615992565b91507f72656c6179657220616c726561647920766f74656400000000000000000000006000830152602082019050919050565b600061486c601483615992565b91507f5061757361626c653a206e6f74207061757365640000000000000000000000006000830152602082019050919050565b60006148ac601f83615992565b91507f6164647220646f65736e277420686176652072656c6179657220726f6c6521006000830152602082019050919050565b60006148ec601c83615992565b91507f5661756c7450726f706f73616c20616c726561647920616374697665000000006000830152602082019050919050565b600061492c601b83615992565b91507f5661756c7450726f706f73616c206973206e6f742061637469766500000000006000830152602082019050919050565b600061496c601683615992565b91507f496e636f72726563742066656520737570706c696564000000000000000000006000830152602082019050919050565b60006149ac601b83615992565b91507f6461746120646f65736e2774206d6174636820646174616861736800000000006000830152602082019050919050565b60006149ec601a83615992565b91507f50726f706f73616c20616c72656164792063616e63656c6c65640000000000006000830152602082019050919050565b6000614a2c602083615992565b91507f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726000830152602082019050919050565b6000614a6c603083615992565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f207265766f6b65000000000000000000000000000000006020830152604082019050919050565b6000614ad2602083615992565b91507f50726f706f73616c206e6f7420617420657870697279207468726573686f6c646000830152602082019050919050565b6000614b12601083615992565b91507f5061757361626c653a20706175736564000000000000000000000000000000006000830152602082019050919050565b6000614b52601e83615992565b91507f6164647220616c7265616479206861732072656c6179657220726f6c652100006000830152602082019050919050565b6000614b92601683615992565b91507f70726f706f73616c206973206e6f7420706173736564000000000000000000006000830152602082019050919050565b6000614bd2602a83615992565b91507f70726f706f73616c20616c7265616479207061737365642f657865637574656460008301527f2f63616e63656c6c6564000000000000000000000000000000000000000000006020830152604082019050919050565b6000614c38600083615992565b9150600082019050919050565b6000614c52602083615992565b91507f73656e64657220646f65736e277420686176652072656c6179657220726f6c656000830152602082019050919050565b6000614c92601e83615992565b91507f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006000830152602082019050919050565b6000614cd2602883615992565b91507f5661756c7450726f706f73616c20616c726561647920706173736564206f722060008301527f65786563757465640000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614d38601f83615992565b91507f43757272656e742066656520697320657175616c20746f206e657720666565006000830152602082019050919050565b6000614d78600c836159a3565b91507f52454c415945525f524f4c4500000000000000000000000000000000000000006000830152600c82019050919050565b6000614db8601183615992565b91507f6461746168617368206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000614df8601983615992565b91507f6e6f2068616e646c657220666f72207265736f757263654944000000000000006000830152602082019050919050565b6000614e38602f83615992565b91507f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008301527f20726f6c657320666f722073656c6600000000000000000000000000000000006020830152604082019050919050565b608082016000820151614ea76000850182615070565b506020820151614eba6020850182615052565b506040820151614ecd60408501826144a5565b506060820151614ee060608501826144a5565b50505050565b600060c083016000830151614efe60008601826144a5565b506020830151614f1160208601826144a5565b5060408301518482036040860152614f298282614438565b91505060608301518482036060860152614f438282614438565b9150506080830151614f58608086018261455d565b5060a0830151614f6b60a0860182615034565b508091505092915050565b600061010083016000830151614f8f60008601826144a5565b506020830151614fa260208601826144a5565b5060408301518482036040860152614fba8282614438565b91505060608301518482036060860152614fd48282614438565b9150506080830151614fe9608086018261457b565b5060a0830151614ffc60a0860182615034565b5060c083015161500f60c08601826144a5565b5060e083015184820360e086015261502782826145d2565b9150508091505092915050565b61503d81615a5a565b82525050565b61504c81615a5a565b82525050565b61505b81615a64565b82525050565b61506a81615a64565b82525050565b61507981615a8d565b82525050565b61508881615a8d565b82525050565b600061509a8286614421565b6014820191506150ab8284866144ff565b9150819050949350505050565b60006150c382614d6b565b9150819050919050565b60006020820190506150e26000830184614412565b92915050565b60006020820190506150fd60008301846143f4565b92915050565b60006060820190506151186000830186614412565b6151256020830185614412565b6151326040830184615043565b949350505050565b600060208201905061514f6000830184614496565b92915050565b600060208201905061516a60008301846144b4565b92915050565b600060408201905061518560008301856144b4565b6151926020830184614412565b9392505050565b60006080820190506151ae60008301876144b4565b6151bb6020830186614412565b6151c860408301856144c3565b6151d560608301846144c3565b95945050505050565b60006040820190506151f360008301856144b4565b61520060208301846144b4565b9392505050565b600060808201905061521c60008301876144b4565b61522960208301866144b4565b61523660408301856144b4565b81810360608301526152488184614599565b905095945050505050565b600060808201905061526860008301876144b4565b61527560208301866144b4565b61528260408301856144b4565b81810360608301526152948184614644565b905095945050505050565b60006080820190506152b460008301866144b4565b6152c160208301856144b4565b6152ce60408301846144b4565b81810360608301526152df81614c2b565b9050949350505050565b60006080820190506152fe60008301876144b4565b61530b60208301866144b4565b615318604083018561456c565b6153256060830184615043565b95945050505050565b600060c08201905061534360008301896144b4565b61535060208301886144b4565b61535d604083018761458a565b61536a6060830186615043565b61537760808301856144b4565b81810360a0830152615389818461460b565b9050979650505050505050565b60006040820190506153ab60008301866144b4565b81810360208301526153be8184866144d2565b9050949350505050565b600060a0820190506153dd60008301896144b4565b6153ea602083018861507f565b6153f76040830187615061565b61540460608301866143f4565b81810360808301526154178184866144d2565b9050979650505050505050565b6000602082019050818103600083015261543e8184614524565b905092915050565b600060208201905081810360008301526154608184614599565b905092915050565b60006020820190508181036000830152615481816146d3565b9050919050565b600060208201905081810360008301526154a181614713565b9050919050565b600060208201905081810360008301526154c181614779565b9050919050565b600060208201905081810360008301526154e1816147b9565b9050919050565b600060208201905081810360008301526155018161481f565b9050919050565b600060208201905081810360008301526155218161485f565b9050919050565b600060208201905081810360008301526155418161489f565b9050919050565b60006020820190508181036000830152615561816148df565b9050919050565b600060208201905081810360008301526155818161491f565b9050919050565b600060208201905081810360008301526155a18161495f565b9050919050565b600060208201905081810360008301526155c18161499f565b9050919050565b600060208201905081810360008301526155e1816149df565b9050919050565b6000602082019050818103600083015261560181614a1f565b9050919050565b6000602082019050818103600083015261562181614a5f565b9050919050565b6000602082019050818103600083015261564181614ac5565b9050919050565b6000602082019050818103600083015261566181614b05565b9050919050565b6000602082019050818103600083015261568181614b45565b9050919050565b600060208201905081810360008301526156a181614b85565b9050919050565b600060208201905081810360008301526156c181614bc5565b9050919050565b600060208201905081810360008301526156e181614c45565b9050919050565b6000602082019050818103600083015261570181614c85565b9050919050565b6000602082019050818103600083015261572181614cc5565b9050919050565b6000602082019050818103600083015261574181614d2b565b9050919050565b6000602082019050818103600083015261576181614dab565b9050919050565b6000602082019050818103600083015261578181614deb565b9050919050565b600060208201905081810360008301526157a181614e2b565b9050919050565b60006080820190506157bd6000830184614e91565b92915050565b600060208201905081810360008301526157dd8184614ee6565b905092915050565b600060208201905081810360008301526157ff8184614f76565b905092915050565b600060208201905061581c6000830184615043565b92915050565b60006020820190506158376000830184615061565b92915050565b6000602082019050615852600083018461507f565b92915050565b600060808201905061586d600083018761507f565b61587a6020830186615061565b61588760408301856144b4565b61589460608301846144b4565b95945050505050565b6000604051905081810181811067ffffffffffffffff821117156158c057600080fd5b8060405250919050565b600067ffffffffffffffff8211156158e157600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006159b982615a3a565b9050919050565b60006159cb82615a3a565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050615a2282615b78565b919050565b6000819050615a3582615b85565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600068ffffffffffffffffff82169050919050565b600060ff82169050919050565b6000615aa582615ad0565b9050919050565b6000615ab782615a14565b9050919050565b6000615ac982615a27565b9050919050565b6000615adb82615ae2565b9050919050565b6000615aed82615a3a565b9050919050565b82818337600083830152505050565b60005b83811015615b21578082015181840152602081019050615b06565b83811115615b30576000848401525b50505050565b6000615b4182615b48565b9050919050565b6000615b5382615b6b565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b60058110615b8257fe5b50565b60058110615b8f57fe5b50565b615b9b816159ae565b8114615ba657600080fd5b50565b615bb2816159c0565b8114615bbd57600080fd5b50565b615bc9816159de565b8114615bd457600080fd5b50565b615be0816159e8565b8114615beb57600080fd5b50565b615bf781615a5a565b8114615c0257600080fd5b50565b615c0e81615a64565b8114615c1957600080fd5b50565b615c2581615a78565b8114615c3057600080fd5b50565b615c3c81615a8d565b8114615c4757600080fd5b5056fea264697066735822122044358a7b7212cef875010d470a3d201534f9f4c563d0d1bfbeec59569d6543d764736f6c63430006040033",
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
