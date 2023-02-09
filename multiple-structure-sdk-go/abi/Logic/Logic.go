// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Logic

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
)

// LogicMetaData contains all meta data concerning the Logic contract.
var LogicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractData\",\"name\":\"_data\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"setX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"setXFromAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506101d9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634a2ad99f1461003b578063cb58703914610050575b600080fd5b61004e61004936600461015e565b610063565b005b61004e61005e36600461015e565b6100d3565b60405163200c6cd560e11b8152600481018290526001600160a01b03831690634018d9aa906024016020604051808303816000875af11580156100aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100ce919061018a565b505050565b60405163200c6cd560e11b81526004810182905282906001600160a01b03821690634018d9aa906024016020604051808303816000875af115801561011c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610140919061018a565b50505050565b6001600160a01b038116811461015b57600080fd5b50565b6000806040838503121561017157600080fd5b823561017c81610146565b946020939093013593505050565b60006020828403121561019c57600080fd5b505191905056fea26469706673582212205e0250759d0459b39005fdfc7804fbf1d5209155355eb0e8d125bf61a842095164736f6c63430008110033",
}

// LogicABI is the input ABI used to generate the binding from.
// Deprecated: Use LogicMetaData.ABI instead.
var LogicABI = LogicMetaData.ABI

// LogicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LogicMetaData.Bin instead.
var LogicBin = LogicMetaData.Bin

// DeployLogic deploys a new Ethereum contract, binding an instance of Logic to it.
func DeployLogic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Logic, error) {
	parsed, err := LogicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LogicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Logic{LogicCaller: LogicCaller{contract: contract}, LogicTransactor: LogicTransactor{contract: contract}, LogicFilterer: LogicFilterer{contract: contract}}, nil
}

// Logic is an auto generated Go binding around an Ethereum contract.
type Logic struct {
	LogicCaller     // Read-only binding to the contract
	LogicTransactor // Write-only binding to the contract
	LogicFilterer   // Log filterer for contract events
}

// LogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type LogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LogicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LogicSession struct {
	Contract     *Logic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LogicCallerSession struct {
	Contract *LogicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LogicTransactorSession struct {
	Contract     *LogicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type LogicRaw struct {
	Contract *Logic // Generic contract binding to access the raw methods on
}

// LogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LogicCallerRaw struct {
	Contract *LogicCaller // Generic read-only contract binding to access the raw methods on
}

// LogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LogicTransactorRaw struct {
	Contract *LogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogic creates a new instance of Logic, bound to a specific deployed contract.
func NewLogic(address common.Address, backend bind.ContractBackend) (*Logic, error) {
	contract, err := bindLogic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Logic{LogicCaller: LogicCaller{contract: contract}, LogicTransactor: LogicTransactor{contract: contract}, LogicFilterer: LogicFilterer{contract: contract}}, nil
}

// NewLogicCaller creates a new read-only instance of Logic, bound to a specific deployed contract.
func NewLogicCaller(address common.Address, caller bind.ContractCaller) (*LogicCaller, error) {
	contract, err := bindLogic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LogicCaller{contract: contract}, nil
}

// NewLogicTransactor creates a new write-only instance of Logic, bound to a specific deployed contract.
func NewLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*LogicTransactor, error) {
	contract, err := bindLogic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LogicTransactor{contract: contract}, nil
}

// NewLogicFilterer creates a new log filterer instance of Logic, bound to a specific deployed contract.
func NewLogicFilterer(address common.Address, filterer bind.ContractFilterer) (*LogicFilterer, error) {
	contract, err := bindLogic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LogicFilterer{contract: contract}, nil
}

// bindLogic binds a generic wrapper to an already deployed contract.
func bindLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logic *LogicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logic.Contract.LogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logic *LogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logic.Contract.LogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logic *LogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logic.Contract.LogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Logic *LogicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Logic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Logic *LogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Logic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Logic *LogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Logic.Contract.contract.Transact(opts, method, params...)
}

// SetX is a paid mutator transaction binding the contract method 0x4a2ad99f.
//
// Solidity: function setX(address _data, uint256 _x) returns()
func (_Logic *LogicTransactor) SetX(opts *bind.TransactOpts, _data common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.contract.Transact(opts, "setX", _data, _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4a2ad99f.
//
// Solidity: function setX(address _data, uint256 _x) returns()
func (_Logic *LogicSession) SetX(_data common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.Contract.SetX(&_Logic.TransactOpts, _data, _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4a2ad99f.
//
// Solidity: function setX(address _data, uint256 _x) returns()
func (_Logic *LogicTransactorSession) SetX(_data common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.Contract.SetX(&_Logic.TransactOpts, _data, _x)
}

// SetXFromAddress is a paid mutator transaction binding the contract method 0xcb587039.
//
// Solidity: function setXFromAddress(address _addr, uint256 _x) returns()
func (_Logic *LogicTransactor) SetXFromAddress(opts *bind.TransactOpts, _addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.contract.Transact(opts, "setXFromAddress", _addr, _x)
}

// SetXFromAddress is a paid mutator transaction binding the contract method 0xcb587039.
//
// Solidity: function setXFromAddress(address _addr, uint256 _x) returns()
func (_Logic *LogicSession) SetXFromAddress(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.Contract.SetXFromAddress(&_Logic.TransactOpts, _addr, _x)
}

// SetXFromAddress is a paid mutator transaction binding the contract method 0xcb587039.
//
// Solidity: function setXFromAddress(address _addr, uint256 _x) returns()
func (_Logic *LogicTransactorSession) SetXFromAddress(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Logic.Contract.SetXFromAddress(&_Logic.TransactOpts, _addr, _x)
}
