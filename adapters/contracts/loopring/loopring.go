// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package loopring

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

// LoopringMetaData contains all meta data concerning the Loopring contract.
var LoopringMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_exchange\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LoopringABI is the input ABI used to generate the binding from.
// Deprecated: Use LoopringMetaData.ABI instead.
var LoopringABI = LoopringMetaData.ABI

// Loopring is an auto generated Go binding around an Ethereum contract.
type Loopring struct {
	LoopringCaller     // Read-only binding to the contract
	LoopringTransactor // Write-only binding to the contract
	LoopringFilterer   // Log filterer for contract events
}

// LoopringCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoopringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoopringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoopringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoopringSession struct {
	Contract     *Loopring         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoopringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoopringCallerSession struct {
	Contract *LoopringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LoopringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoopringTransactorSession struct {
	Contract     *LoopringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LoopringRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoopringRaw struct {
	Contract *Loopring // Generic contract binding to access the raw methods on
}

// LoopringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoopringCallerRaw struct {
	Contract *LoopringCaller // Generic read-only contract binding to access the raw methods on
}

// LoopringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoopringTransactorRaw struct {
	Contract *LoopringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoopring creates a new instance of Loopring, bound to a specific deployed contract.
func NewLoopring(address common.Address, backend bind.ContractBackend) (*Loopring, error) {
	contract, err := bindLoopring(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Loopring{LoopringCaller: LoopringCaller{contract: contract}, LoopringTransactor: LoopringTransactor{contract: contract}, LoopringFilterer: LoopringFilterer{contract: contract}}, nil
}

// NewLoopringCaller creates a new read-only instance of Loopring, bound to a specific deployed contract.
func NewLoopringCaller(address common.Address, caller bind.ContractCaller) (*LoopringCaller, error) {
	contract, err := bindLoopring(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoopringCaller{contract: contract}, nil
}

// NewLoopringTransactor creates a new write-only instance of Loopring, bound to a specific deployed contract.
func NewLoopringTransactor(address common.Address, transactor bind.ContractTransactor) (*LoopringTransactor, error) {
	contract, err := bindLoopring(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoopringTransactor{contract: contract}, nil
}

// NewLoopringFilterer creates a new log filterer instance of Loopring, bound to a specific deployed contract.
func NewLoopringFilterer(address common.Address, filterer bind.ContractFilterer) (*LoopringFilterer, error) {
	contract, err := bindLoopring(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoopringFilterer{contract: contract}, nil
}

// bindLoopring binds a generic wrapper to an already deployed contract.
func bindLoopring(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoopringMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loopring *LoopringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loopring.Contract.LoopringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loopring *LoopringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loopring.Contract.LoopringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loopring *LoopringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loopring.Contract.LoopringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loopring *LoopringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loopring.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loopring *LoopringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loopring.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loopring *LoopringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loopring.Contract.contract.Transact(opts, method, params...)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Loopring *LoopringCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loopring.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Loopring *LoopringSession) Exchange() (common.Address, error) {
	return _Loopring.Contract.Exchange(&_Loopring.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_Loopring *LoopringCallerSession) Exchange() (common.Address, error) {
	return _Loopring.Contract.Exchange(&_Loopring.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8742da3a.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, uint256 duration, bytes extraData) payable returns()
func (_Loopring *LoopringTransactor) Deposit(opts *bind.TransactOpts, from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, duration *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Loopring.contract.Transact(opts, "deposit", from, to, tokenAddress, amount, duration, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0x8742da3a.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, uint256 duration, bytes extraData) payable returns()
func (_Loopring *LoopringSession) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, duration *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Loopring.Contract.Deposit(&_Loopring.TransactOpts, from, to, tokenAddress, amount, duration, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0x8742da3a.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, uint256 duration, bytes extraData) payable returns()
func (_Loopring *LoopringTransactorSession) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, duration *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Loopring.Contract.Deposit(&_Loopring.TransactOpts, from, to, tokenAddress, amount, duration, extraData)
}

// LoopringDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Loopring contract.
type LoopringDepositedIterator struct {
	Event *LoopringDeposited // Event containing the contract specifics and raw log

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
func (it *LoopringDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringDeposited)
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
		it.Event = new(LoopringDeposited)
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
func (it *LoopringDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringDeposited represents a Deposited event raised by the Loopring contract.
type LoopringDeposited struct {
	From     common.Address
	To       common.Address
	Token    common.Address
	Amount   *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xe3683a3633d7c4dea824891208cc5661856f1c9441073cd1a0b999e189c8bb41.
//
// Solidity: event Deposited(address from, address to, address token, uint96 amount, uint256 duration)
func (_Loopring *LoopringFilterer) FilterDeposited(opts *bind.FilterOpts) (*LoopringDepositedIterator, error) {

	logs, sub, err := _Loopring.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &LoopringDepositedIterator{contract: _Loopring.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xe3683a3633d7c4dea824891208cc5661856f1c9441073cd1a0b999e189c8bb41.
//
// Solidity: event Deposited(address from, address to, address token, uint96 amount, uint256 duration)
func (_Loopring *LoopringFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *LoopringDeposited) (event.Subscription, error) {

	logs, sub, err := _Loopring.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringDeposited)
				if err := _Loopring.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xe3683a3633d7c4dea824891208cc5661856f1c9441073cd1a0b999e189c8bb41.
//
// Solidity: event Deposited(address from, address to, address token, uint96 amount, uint256 duration)
func (_Loopring *LoopringFilterer) ParseDeposited(log types.Log) (*LoopringDeposited, error) {
	event := new(LoopringDeposited)
	if err := _Loopring.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
