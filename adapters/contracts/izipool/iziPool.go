// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izipool

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

// IziPoolMetaData contains all meta data concerning the IziPool contract.
var IziPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"liquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenOwedY\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPrice_96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"currentPoint\",\"type\":\"int24\"},{\"internalType\":\"uint16\",\"name\":\"observationCurrentIndex\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"observationNextQueueLen\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityX\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IziPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IziPoolMetaData.ABI instead.
var IziPoolABI = IziPoolMetaData.ABI

// IziPool is an auto generated Go binding around an Ethereum contract.
type IziPool struct {
	IziPoolCaller     // Read-only binding to the contract
	IziPoolTransactor // Write-only binding to the contract
	IziPoolFilterer   // Log filterer for contract events
}

// IziPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IziPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IziPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IziPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IziPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IziPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IziPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IziPoolSession struct {
	Contract     *IziPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IziPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IziPoolCallerSession struct {
	Contract *IziPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IziPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IziPoolTransactorSession struct {
	Contract     *IziPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IziPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IziPoolRaw struct {
	Contract *IziPool // Generic contract binding to access the raw methods on
}

// IziPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IziPoolCallerRaw struct {
	Contract *IziPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IziPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IziPoolTransactorRaw struct {
	Contract *IziPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIziPool creates a new instance of IziPool, bound to a specific deployed contract.
func NewIziPool(address common.Address, backend bind.ContractBackend) (*IziPool, error) {
	contract, err := bindIziPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IziPool{IziPoolCaller: IziPoolCaller{contract: contract}, IziPoolTransactor: IziPoolTransactor{contract: contract}, IziPoolFilterer: IziPoolFilterer{contract: contract}}, nil
}

// NewIziPoolCaller creates a new read-only instance of IziPool, bound to a specific deployed contract.
func NewIziPoolCaller(address common.Address, caller bind.ContractCaller) (*IziPoolCaller, error) {
	contract, err := bindIziPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IziPoolCaller{contract: contract}, nil
}

// NewIziPoolTransactor creates a new write-only instance of IziPool, bound to a specific deployed contract.
func NewIziPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IziPoolTransactor, error) {
	contract, err := bindIziPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IziPoolTransactor{contract: contract}, nil
}

// NewIziPoolFilterer creates a new log filterer instance of IziPool, bound to a specific deployed contract.
func NewIziPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IziPoolFilterer, error) {
	contract, err := bindIziPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IziPoolFilterer{contract: contract}, nil
}

// bindIziPool binds a generic wrapper to an already deployed contract.
func bindIziPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IziPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IziPool *IziPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IziPool.Contract.IziPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IziPool *IziPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IziPool.Contract.IziPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IziPool *IziPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IziPool.Contract.IziPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IziPool *IziPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IziPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IziPool *IziPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IziPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IziPool *IziPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IziPool.Contract.contract.Transact(opts, method, params...)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IziPool *IziPoolCaller) Liquidity(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	var out []interface{}
	err := _IziPool.contract.Call(opts, &out, "liquidity", key)

	outstruct := new(struct {
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		TokenOwedX       *big.Int
		TokenOwedY       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedX = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokenOwedY = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IziPool *IziPoolSession) Liquidity(key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _IziPool.Contract.Liquidity(&_IziPool.CallOpts, key)
}

// Liquidity is a free data retrieval call binding the contract method 0xb0f59257.
//
// Solidity: function liquidity(bytes32 key) view returns(uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 tokenOwedX, uint256 tokenOwedY)
func (_IziPool *IziPoolCallerSession) Liquidity(key [32]byte) (struct {
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	TokenOwedX       *big.Int
	TokenOwedY       *big.Int
}, error) {
	return _IziPool.Contract.Liquidity(&_IziPool.CallOpts, key)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IziPool *IziPoolCaller) State(opts *bind.CallOpts) (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	var out []interface{}
	err := _IziPool.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SqrtPrice96             *big.Int
		CurrentPoint            *big.Int
		ObservationCurrentIndex uint16
		ObservationQueueLen     uint16
		ObservationNextQueueLen uint16
		Locked                  bool
		Liquidity               *big.Int
		LiquidityX              *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPrice96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationCurrentIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationQueueLen = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationNextQueueLen = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Locked = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Liquidity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityX = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IziPool *IziPoolSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _IziPool.Contract.State(&_IziPool.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint160 sqrtPrice_96, int24 currentPoint, uint16 observationCurrentIndex, uint16 observationQueueLen, uint16 observationNextQueueLen, bool locked, uint128 liquidity, uint128 liquidityX)
func (_IziPool *IziPoolCallerSession) State() (struct {
	SqrtPrice96             *big.Int
	CurrentPoint            *big.Int
	ObservationCurrentIndex uint16
	ObservationQueueLen     uint16
	ObservationNextQueueLen uint16
	Locked                  bool
	Liquidity               *big.Int
	LiquidityX              *big.Int
}, error) {
	return _IziPool.Contract.State(&_IziPool.CallOpts)
}
