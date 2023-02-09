// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Caller

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

// CallerMetaData contains all meta data concerning the Caller contract.
var CallerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Response\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"callSetX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"delegatecallSetX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506102f1806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630c55699c146100515780631750baac1461006c5780633ca065bb146100815780635197c7aa14610094575b600080fd5b61005a60005481565b60405190815260200160405180910390f35b61007f61007a366004610207565b61009c565b005b61007f61008f366004610207565b610175565b60005461005a565b600080836001600160a01b0316836040516024016100bc91815260200190565b60408051601f198184030181529181526020820180516001600160e01b031663200c6cd560e11b179052516100f19190610263565b600060405180830381855af49150503d806000811461012c576040519150601f19603f3d011682016040523d82523d6000602084013e610131565b606091505b50915091507f13848c3e38f8886f3f5d2ad9dff80d8092c2bbb8efd5b887a99c2c6cfc09ac2a828260405161016792919061027f565b60405180910390a150505050565b600080836001600160a01b03168360405160240161019591815260200190565b60408051601f198184030181529181526020820180516001600160e01b031663200c6cd560e11b179052516101ca9190610263565b6000604051808303816000865af19150503d806000811461012c576040519150601f19603f3d011682016040523d82523d6000602084013e610131565b6000806040838503121561021a57600080fd5b82356001600160a01b038116811461023157600080fd5b946020939093013593505050565b60005b8381101561025a578181015183820152602001610242565b50506000910152565b6000825161027581846020870161023f565b9190910192915050565b821515815260406020820152600082518060408401526102a681606085016020870161023f565b601f01601f191691909101606001939250505056fea26469706673582212205c3cb319e2b0b46aad6451df1ac8a68bd3e76b52f19e1a5b01f7b231eb711e4a64736f6c63430008110033",
}

// CallerABI is the input ABI used to generate the binding from.
// Deprecated: Use CallerMetaData.ABI instead.
var CallerABI = CallerMetaData.ABI

// CallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallerMetaData.Bin instead.
var CallerBin = CallerMetaData.Bin

// DeployCaller deploys a new Ethereum contract, binding an instance of Caller to it.
func DeployCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Caller, error) {
	parsed, err := CallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// Caller is an auto generated Go binding around an Ethereum contract.
type Caller struct {
	CallerCaller     // Read-only binding to the contract
	CallerTransactor // Write-only binding to the contract
	CallerFilterer   // Log filterer for contract events
}

// CallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallerSession struct {
	Contract     *Caller           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallerCallerSession struct {
	Contract *CallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallerTransactorSession struct {
	Contract     *CallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallerRaw struct {
	Contract *Caller // Generic contract binding to access the raw methods on
}

// CallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallerCallerRaw struct {
	Contract *CallerCaller // Generic read-only contract binding to access the raw methods on
}

// CallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallerTransactorRaw struct {
	Contract *CallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCaller creates a new instance of Caller, bound to a specific deployed contract.
func NewCaller(address common.Address, backend bind.ContractBackend) (*Caller, error) {
	contract, err := bindCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Caller{CallerCaller: CallerCaller{contract: contract}, CallerTransactor: CallerTransactor{contract: contract}, CallerFilterer: CallerFilterer{contract: contract}}, nil
}

// NewCallerCaller creates a new read-only instance of Caller, bound to a specific deployed contract.
func NewCallerCaller(address common.Address, caller bind.ContractCaller) (*CallerCaller, error) {
	contract, err := bindCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallerCaller{contract: contract}, nil
}

// NewCallerTransactor creates a new write-only instance of Caller, bound to a specific deployed contract.
func NewCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CallerTransactor, error) {
	contract, err := bindCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallerTransactor{contract: contract}, nil
}

// NewCallerFilterer creates a new log filterer instance of Caller, bound to a specific deployed contract.
func NewCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CallerFilterer, error) {
	contract, err := bindCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallerFilterer{contract: contract}, nil
}

// bindCaller binds a generic wrapper to an already deployed contract.
func bindCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.CallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.CallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Caller *CallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Caller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Caller *CallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Caller *CallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Caller.Contract.contract.Transact(opts, method, params...)
}

// GetX is a free data retrieval call binding the contract method 0x5197c7aa.
//
// Solidity: function getX() view returns(uint256)
func (_Caller *CallerCaller) GetX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "getX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetX is a free data retrieval call binding the contract method 0x5197c7aa.
//
// Solidity: function getX() view returns(uint256)
func (_Caller *CallerSession) GetX() (*big.Int, error) {
	return _Caller.Contract.GetX(&_Caller.CallOpts)
}

// GetX is a free data retrieval call binding the contract method 0x5197c7aa.
//
// Solidity: function getX() view returns(uint256)
func (_Caller *CallerCallerSession) GetX() (*big.Int, error) {
	return _Caller.Contract.GetX(&_Caller.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Caller *CallerCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Caller.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Caller *CallerSession) X() (*big.Int, error) {
	return _Caller.Contract.X(&_Caller.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Caller *CallerCallerSession) X() (*big.Int, error) {
	return _Caller.Contract.X(&_Caller.CallOpts)
}

// CallSetX is a paid mutator transaction binding the contract method 0x3ca065bb.
//
// Solidity: function callSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerTransactor) CallSetX(opts *bind.TransactOpts, _addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "callSetX", _addr, _x)
}

// CallSetX is a paid mutator transaction binding the contract method 0x3ca065bb.
//
// Solidity: function callSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerSession) CallSetX(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.CallSetX(&_Caller.TransactOpts, _addr, _x)
}

// CallSetX is a paid mutator transaction binding the contract method 0x3ca065bb.
//
// Solidity: function callSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerTransactorSession) CallSetX(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.CallSetX(&_Caller.TransactOpts, _addr, _x)
}

// DelegatecallSetX is a paid mutator transaction binding the contract method 0x1750baac.
//
// Solidity: function delegatecallSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerTransactor) DelegatecallSetX(opts *bind.TransactOpts, _addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.contract.Transact(opts, "delegatecallSetX", _addr, _x)
}

// DelegatecallSetX is a paid mutator transaction binding the contract method 0x1750baac.
//
// Solidity: function delegatecallSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerSession) DelegatecallSetX(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.DelegatecallSetX(&_Caller.TransactOpts, _addr, _x)
}

// DelegatecallSetX is a paid mutator transaction binding the contract method 0x1750baac.
//
// Solidity: function delegatecallSetX(address _addr, uint256 _x) returns()
func (_Caller *CallerTransactorSession) DelegatecallSetX(_addr common.Address, _x *big.Int) (*types.Transaction, error) {
	return _Caller.Contract.DelegatecallSetX(&_Caller.TransactOpts, _addr, _x)
}

// CallerResponseIterator is returned from FilterResponse and is used to iterate over the raw logs and unpacked data for Response events raised by the Caller contract.
type CallerResponseIterator struct {
	Event *CallerResponse // Event containing the contract specifics and raw log

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
func (it *CallerResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallerResponse)
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
		it.Event = new(CallerResponse)
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
func (it *CallerResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallerResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallerResponse represents a Response event raised by the Caller contract.
type CallerResponse struct {
	Success bool
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResponse is a free log retrieval operation binding the contract event 0x13848c3e38f8886f3f5d2ad9dff80d8092c2bbb8efd5b887a99c2c6cfc09ac2a.
//
// Solidity: event Response(bool success, bytes data)
func (_Caller *CallerFilterer) FilterResponse(opts *bind.FilterOpts) (*CallerResponseIterator, error) {

	logs, sub, err := _Caller.contract.FilterLogs(opts, "Response")
	if err != nil {
		return nil, err
	}
	return &CallerResponseIterator{contract: _Caller.contract, event: "Response", logs: logs, sub: sub}, nil
}

// WatchResponse is a free log subscription operation binding the contract event 0x13848c3e38f8886f3f5d2ad9dff80d8092c2bbb8efd5b887a99c2c6cfc09ac2a.
//
// Solidity: event Response(bool success, bytes data)
func (_Caller *CallerFilterer) WatchResponse(opts *bind.WatchOpts, sink chan<- *CallerResponse) (event.Subscription, error) {

	logs, sub, err := _Caller.contract.WatchLogs(opts, "Response")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallerResponse)
				if err := _Caller.contract.UnpackLog(event, "Response", log); err != nil {
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

// ParseResponse is a log parse operation binding the contract event 0x13848c3e38f8886f3f5d2ad9dff80d8092c2bbb8efd5b887a99c2c6cfc09ac2a.
//
// Solidity: event Response(bool success, bytes data)
func (_Caller *CallerFilterer) ParseResponse(log types.Log) (*CallerResponse, error) {
	event := new(CallerResponse)
	if err := _Caller.contract.UnpackLog(event, "Response", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
