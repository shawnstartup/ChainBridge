// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HandlerHelpers

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

// HandlerHelpersMetaData contains all meta data concerning the HandlerHelpers contract.
var HandlerHelpersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610857806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637f79bea8116100665780637f79bea81461023f578063b8fa37361461029b578063c3b65816146102e9578063c8ba6c8714610333578063d9caed121461038b5761009e565b806307b7ed99146100a35780630a6d55d8146100e7578063318c136e146101555780636817031b1461019f5780636a70d081146101e3575b600080fd5b6100e5600480360360208110156100b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103f9565b005b610113600480360360208110156100fd57600080fd5b810190808035906020019092919050505061040d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61015d610440565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101e1600480360360208110156101b557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610465565b005b610225600480360360208110156101f957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104b1565b604051808215151515815260200191505060405180910390f35b6102816004803603602081101561025557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104d1565b604051808215151515815260200191505060405180910390f35b6102e7600480360360408110156102b157600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104f1565b005b6102f1610507565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103756004803603602081101561034957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061052d565b6040518082815260200191505060405180910390f35b6103f7600480360360608110156103a157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610545565b005b61040161054a565b61040a8161060e565b50565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61046d61054a565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60056020528060005260406000206000915054906101000a900460ff1681565b60046020528060005260406000206000915054906101000a900460ff1681565b6104f961054a565b610503828261070b565b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090505481565b505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f73656e646572206d7573742062652062726964676520636f6e7472616374000081525060200191505060405180910390fd5b565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166106b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806107fe6024913960400191505060405180910390fd5b6001600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b806002600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505056fe70726f766964656420636f6e7472616374206973206e6f742077686974656c6973746564a2646970667358221220766be0a20aeb8b96d4847db8e3d221578f650f31e573bc29b15a458e8bb0990864736f6c63430006040033",
}

// HandlerHelpersABI is the input ABI used to generate the binding from.
// Deprecated: Use HandlerHelpersMetaData.ABI instead.
var HandlerHelpersABI = HandlerHelpersMetaData.ABI

// HandlerHelpersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HandlerHelpersMetaData.Bin instead.
var HandlerHelpersBin = HandlerHelpersMetaData.Bin

// DeployHandlerHelpers deploys a new Ethereum contract, binding an instance of HandlerHelpers to it.
func DeployHandlerHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HandlerHelpers, error) {
	parsed, err := HandlerHelpersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HandlerHelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// HandlerHelpers is an auto generated Go binding around an Ethereum contract.
type HandlerHelpers struct {
	HandlerHelpersCaller     // Read-only binding to the contract
	HandlerHelpersTransactor // Write-only binding to the contract
	HandlerHelpersFilterer   // Log filterer for contract events
}

// HandlerHelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HandlerHelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HandlerHelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HandlerHelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HandlerHelpersSession struct {
	Contract     *HandlerHelpers   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HandlerHelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HandlerHelpersCallerSession struct {
	Contract *HandlerHelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HandlerHelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HandlerHelpersTransactorSession struct {
	Contract     *HandlerHelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HandlerHelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HandlerHelpersRaw struct {
	Contract *HandlerHelpers // Generic contract binding to access the raw methods on
}

// HandlerHelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HandlerHelpersCallerRaw struct {
	Contract *HandlerHelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HandlerHelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HandlerHelpersTransactorRaw struct {
	Contract *HandlerHelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHandlerHelpers creates a new instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpers(address common.Address, backend bind.ContractBackend) (*HandlerHelpers, error) {
	contract, err := bindHandlerHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpers{HandlerHelpersCaller: HandlerHelpersCaller{contract: contract}, HandlerHelpersTransactor: HandlerHelpersTransactor{contract: contract}, HandlerHelpersFilterer: HandlerHelpersFilterer{contract: contract}}, nil
}

// NewHandlerHelpersCaller creates a new read-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersCaller(address common.Address, caller bind.ContractCaller) (*HandlerHelpersCaller, error) {
	contract, err := bindHandlerHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersCaller{contract: contract}, nil
}

// NewHandlerHelpersTransactor creates a new write-only instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HandlerHelpersTransactor, error) {
	contract, err := bindHandlerHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersTransactor{contract: contract}, nil
}

// NewHandlerHelpersFilterer creates a new log filterer instance of HandlerHelpers, bound to a specific deployed contract.
func NewHandlerHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HandlerHelpersFilterer, error) {
	contract, err := bindHandlerHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HandlerHelpersFilterer{contract: contract}, nil
}

// bindHandlerHelpers binds a generic wrapper to an already deployed contract.
func bindHandlerHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HandlerHelpersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.HandlerHelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.HandlerHelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HandlerHelpers *HandlerHelpersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HandlerHelpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HandlerHelpers *HandlerHelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) BridgeAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.BridgeAddress(&_HandlerHelpers.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.BurnList(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_HandlerHelpers *HandlerHelpersCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _HandlerHelpers.Contract.ContractWhitelist(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _HandlerHelpers.Contract.ResourceIDToTokenContractAddress(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_HandlerHelpers *HandlerHelpersCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _HandlerHelpers.Contract.TokenContractAddressToResourceID(&_HandlerHelpers.CallOpts, arg0)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HandlerHelpers.contract.Call(opts, &out, "_vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersSession) VaultAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.VaultAddress(&_HandlerHelpers.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0xc3b65816.
//
// Solidity: function _vaultAddress() view returns(address)
func (_HandlerHelpers *HandlerHelpersCallerSession) VaultAddress() (common.Address, error) {
	return _HandlerHelpers.Contract.VaultAddress(&_HandlerHelpers.CallOpts)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetBurnable(&_HandlerHelpers.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetResource(&_HandlerHelpers.TransactOpts, resourceID, contractAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) SetVault(opts *bind.TransactOpts, vaultAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "setVault", vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_HandlerHelpers *HandlerHelpersSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetVault(&_HandlerHelpers.TransactOpts, vaultAddress)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address vaultAddress) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) SetVault(vaultAddress common.Address) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.SetVault(&_HandlerHelpers.TransactOpts, vaultAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_HandlerHelpers *HandlerHelpersTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _HandlerHelpers.contract.Transact(opts, "withdraw", tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_HandlerHelpers *HandlerHelpersSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.Withdraw(&_HandlerHelpers.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_HandlerHelpers *HandlerHelpersTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _HandlerHelpers.Contract.Withdraw(&_HandlerHelpers.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}
