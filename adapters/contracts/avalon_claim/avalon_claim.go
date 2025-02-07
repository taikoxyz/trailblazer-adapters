// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package avalon_claim

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

// AvalonClaimMetaData contains all meta data concerning the AvalonClaim contract.
var AvalonClaimMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"internalType\":\"uint256\",\"name\":\"avlAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internalType\":\"uint256\",\"name\":\"usdaAmount\",\"type\":\"uint256\",\"indexed\":false}],\"type\":\"event\",\"name\":\"Claimed\",\"anonymous\":false}]",
}

// AvalonClaimABI is the input ABI used to generate the binding from.
// Deprecated: Use AvalonClaimMetaData.ABI instead.
var AvalonClaimABI = AvalonClaimMetaData.ABI

// AvalonClaim is an auto generated Go binding around an Ethereum contract.
type AvalonClaim struct {
	AvalonClaimCaller     // Read-only binding to the contract
	AvalonClaimTransactor // Write-only binding to the contract
	AvalonClaimFilterer   // Log filterer for contract events
}

// AvalonClaimCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvalonClaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalonClaimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvalonClaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalonClaimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvalonClaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvalonClaimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvalonClaimSession struct {
	Contract     *AvalonClaim      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvalonClaimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvalonClaimCallerSession struct {
	Contract *AvalonClaimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AvalonClaimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvalonClaimTransactorSession struct {
	Contract     *AvalonClaimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AvalonClaimRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvalonClaimRaw struct {
	Contract *AvalonClaim // Generic contract binding to access the raw methods on
}

// AvalonClaimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvalonClaimCallerRaw struct {
	Contract *AvalonClaimCaller // Generic read-only contract binding to access the raw methods on
}

// AvalonClaimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvalonClaimTransactorRaw struct {
	Contract *AvalonClaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvalonClaim creates a new instance of AvalonClaim, bound to a specific deployed contract.
func NewAvalonClaim(address common.Address, backend bind.ContractBackend) (*AvalonClaim, error) {
	contract, err := bindAvalonClaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvalonClaim{AvalonClaimCaller: AvalonClaimCaller{contract: contract}, AvalonClaimTransactor: AvalonClaimTransactor{contract: contract}, AvalonClaimFilterer: AvalonClaimFilterer{contract: contract}}, nil
}

// NewAvalonClaimCaller creates a new read-only instance of AvalonClaim, bound to a specific deployed contract.
func NewAvalonClaimCaller(address common.Address, caller bind.ContractCaller) (*AvalonClaimCaller, error) {
	contract, err := bindAvalonClaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvalonClaimCaller{contract: contract}, nil
}

// NewAvalonClaimTransactor creates a new write-only instance of AvalonClaim, bound to a specific deployed contract.
func NewAvalonClaimTransactor(address common.Address, transactor bind.ContractTransactor) (*AvalonClaimTransactor, error) {
	contract, err := bindAvalonClaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvalonClaimTransactor{contract: contract}, nil
}

// NewAvalonClaimFilterer creates a new log filterer instance of AvalonClaim, bound to a specific deployed contract.
func NewAvalonClaimFilterer(address common.Address, filterer bind.ContractFilterer) (*AvalonClaimFilterer, error) {
	contract, err := bindAvalonClaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvalonClaimFilterer{contract: contract}, nil
}

// bindAvalonClaim binds a generic wrapper to an already deployed contract.
func bindAvalonClaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvalonClaimMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalonClaim *AvalonClaimRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalonClaim.Contract.AvalonClaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalonClaim *AvalonClaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalonClaim.Contract.AvalonClaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalonClaim *AvalonClaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalonClaim.Contract.AvalonClaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvalonClaim *AvalonClaimCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvalonClaim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvalonClaim *AvalonClaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvalonClaim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvalonClaim *AvalonClaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvalonClaim.Contract.contract.Transact(opts, method, params...)
}

// AvalonClaimClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the AvalonClaim contract.
type AvalonClaimClaimedIterator struct {
	Event *AvalonClaimClaimed // Event containing the contract specifics and raw log

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
func (it *AvalonClaimClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvalonClaimClaimed)
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
		it.Event = new(AvalonClaimClaimed)
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
func (it *AvalonClaimClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvalonClaimClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvalonClaimClaimed represents a Claimed event raised by the AvalonClaim contract.
type AvalonClaimClaimed struct {
	User       common.Address
	AvlAmount  *big.Int
	UsdaAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 avlAmount, uint256 usdaAmount)
func (_AvalonClaim *AvalonClaimFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*AvalonClaimClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AvalonClaim.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &AvalonClaimClaimedIterator{contract: _AvalonClaim.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 avlAmount, uint256 usdaAmount)
func (_AvalonClaim *AvalonClaimFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *AvalonClaimClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AvalonClaim.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvalonClaimClaimed)
				if err := _AvalonClaim.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x987d620f307ff6b94d58743cb7a7509f24071586a77759b77c2d4e29f75a2f9a.
//
// Solidity: event Claimed(address indexed user, uint256 avlAmount, uint256 usdaAmount)
func (_AvalonClaim *AvalonClaimFilterer) ParseClaimed(log types.Log) (*AvalonClaimClaimed, error) {
	event := new(AvalonClaimClaimed)
	if err := _AvalonClaim.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
