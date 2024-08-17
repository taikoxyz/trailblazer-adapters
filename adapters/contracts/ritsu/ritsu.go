// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ritsu

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

// IPoolTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type IPoolTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// RitsuMetaData contains all meta data concerning the Ritsu contract.
var RitsuMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"_amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burnSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_protocolFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"permit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RitsuABI is the input ABI used to generate the binding from.
// Deprecated: Use RitsuMetaData.ABI instead.
var RitsuABI = RitsuMetaData.ABI

// Ritsu is an auto generated Go binding around an Ethereum contract.
type Ritsu struct {
	RitsuCaller     // Read-only binding to the contract
	RitsuTransactor // Write-only binding to the contract
	RitsuFilterer   // Log filterer for contract events
}

// RitsuCaller is an auto generated read-only Go binding around an Ethereum contract.
type RitsuCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RitsuTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RitsuTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RitsuFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RitsuFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RitsuSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RitsuSession struct {
	Contract     *Ritsu            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RitsuCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RitsuCallerSession struct {
	Contract *RitsuCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RitsuTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RitsuTransactorSession struct {
	Contract     *RitsuTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RitsuRaw is an auto generated low-level Go binding around an Ethereum contract.
type RitsuRaw struct {
	Contract *Ritsu // Generic contract binding to access the raw methods on
}

// RitsuCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RitsuCallerRaw struct {
	Contract *RitsuCaller // Generic read-only contract binding to access the raw methods on
}

// RitsuTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RitsuTransactorRaw struct {
	Contract *RitsuTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRitsu creates a new instance of Ritsu, bound to a specific deployed contract.
func NewRitsu(address common.Address, backend bind.ContractBackend) (*Ritsu, error) {
	contract, err := bindRitsu(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ritsu{RitsuCaller: RitsuCaller{contract: contract}, RitsuTransactor: RitsuTransactor{contract: contract}, RitsuFilterer: RitsuFilterer{contract: contract}}, nil
}

// NewRitsuCaller creates a new read-only instance of Ritsu, bound to a specific deployed contract.
func NewRitsuCaller(address common.Address, caller bind.ContractCaller) (*RitsuCaller, error) {
	contract, err := bindRitsu(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RitsuCaller{contract: contract}, nil
}

// NewRitsuTransactor creates a new write-only instance of Ritsu, bound to a specific deployed contract.
func NewRitsuTransactor(address common.Address, transactor bind.ContractTransactor) (*RitsuTransactor, error) {
	contract, err := bindRitsu(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RitsuTransactor{contract: contract}, nil
}

// NewRitsuFilterer creates a new log filterer instance of Ritsu, bound to a specific deployed contract.
func NewRitsuFilterer(address common.Address, filterer bind.ContractFilterer) (*RitsuFilterer, error) {
	contract, err := bindRitsu(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RitsuFilterer{contract: contract}, nil
}

// bindRitsu binds a generic wrapper to an already deployed contract.
func bindRitsu(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RitsuMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ritsu *RitsuRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ritsu.Contract.RitsuCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ritsu *RitsuRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ritsu.Contract.RitsuTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ritsu *RitsuRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ritsu.Contract.RitsuTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ritsu *RitsuCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ritsu.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ritsu *RitsuTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ritsu.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ritsu *RitsuTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ritsu.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ritsu *RitsuCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ritsu *RitsuSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ritsu.Contract.DOMAINSEPARATOR(&_Ritsu.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ritsu *RitsuCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ritsu.Contract.DOMAINSEPARATOR(&_Ritsu.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ritsu *RitsuCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ritsu *RitsuSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.Allowance(&_Ritsu.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Ritsu *RitsuCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.Allowance(&_Ritsu.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ritsu *RitsuCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ritsu *RitsuSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.BalanceOf(&_Ritsu.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Ritsu *RitsuCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.BalanceOf(&_Ritsu.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ritsu *RitsuCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ritsu *RitsuSession) Decimals() (uint8, error) {
	return _Ritsu.Contract.Decimals(&_Ritsu.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ritsu *RitsuCallerSession) Decimals() (uint8, error) {
	return _Ritsu.Contract.Decimals(&_Ritsu.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_Ritsu *RitsuCaller) GetAmountIn(opts *bind.CallOpts, _tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getAmountIn", _tokenOut, _amountOut, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_Ritsu *RitsuSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _Ritsu.Contract.GetAmountIn(&_Ritsu.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_Ritsu *RitsuCallerSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _Ritsu.Contract.GetAmountIn(&_Ritsu.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_Ritsu *RitsuCaller) GetAmountOut(opts *bind.CallOpts, _tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getAmountOut", _tokenIn, _amountIn, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_Ritsu *RitsuSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _Ritsu.Contract.GetAmountOut(&_Ritsu.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_Ritsu *RitsuCallerSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _Ritsu.Contract.GetAmountOut(&_Ritsu.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Ritsu *RitsuCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Ritsu *RitsuSession) GetAssets() ([]common.Address, error) {
	return _Ritsu.Contract.GetAssets(&_Ritsu.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_Ritsu *RitsuCallerSession) GetAssets() ([]common.Address, error) {
	return _Ritsu.Contract.GetAssets(&_Ritsu.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_Ritsu *RitsuCaller) GetProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_Ritsu *RitsuSession) GetProtocolFee() (*big.Int, error) {
	return _Ritsu.Contract.GetProtocolFee(&_Ritsu.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_Ritsu *RitsuCallerSession) GetProtocolFee() (*big.Int, error) {
	return _Ritsu.Contract.GetProtocolFee(&_Ritsu.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_Ritsu *RitsuCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_Ritsu *RitsuSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _Ritsu.Contract.GetReserves(&_Ritsu.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_Ritsu *RitsuCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _Ritsu.Contract.GetReserves(&_Ritsu.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_Ritsu *RitsuCaller) GetSwapFee(opts *bind.CallOpts, _sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "getSwapFee", _sender, _tokenIn, _tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_Ritsu *RitsuSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _Ritsu.Contract.GetSwapFee(&_Ritsu.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_Ritsu *RitsuCallerSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _Ritsu.Contract.GetSwapFee(&_Ritsu.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_Ritsu *RitsuCaller) InvariantLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "invariantLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_Ritsu *RitsuSession) InvariantLast() (*big.Int, error) {
	return _Ritsu.Contract.InvariantLast(&_Ritsu.CallOpts)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_Ritsu *RitsuCallerSession) InvariantLast() (*big.Int, error) {
	return _Ritsu.Contract.InvariantLast(&_Ritsu.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_Ritsu *RitsuCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_Ritsu *RitsuSession) Master() (common.Address, error) {
	return _Ritsu.Contract.Master(&_Ritsu.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_Ritsu *RitsuCallerSession) Master() (common.Address, error) {
	return _Ritsu.Contract.Master(&_Ritsu.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ritsu *RitsuCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ritsu *RitsuSession) Name() (string, error) {
	return _Ritsu.Contract.Name(&_Ritsu.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ritsu *RitsuCallerSession) Name() (string, error) {
	return _Ritsu.Contract.Name(&_Ritsu.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Ritsu *RitsuCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Ritsu *RitsuSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.Nonces(&_Ritsu.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Ritsu *RitsuCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Ritsu.Contract.Nonces(&_Ritsu.CallOpts, arg0)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_Ritsu *RitsuCaller) PoolType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_Ritsu *RitsuSession) PoolType() (uint16, error) {
	return _Ritsu.Contract.PoolType(&_Ritsu.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_Ritsu *RitsuCallerSession) PoolType() (uint16, error) {
	return _Ritsu.Contract.PoolType(&_Ritsu.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Ritsu *RitsuCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Ritsu *RitsuSession) Reserve0() (*big.Int, error) {
	return _Ritsu.Contract.Reserve0(&_Ritsu.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_Ritsu *RitsuCallerSession) Reserve0() (*big.Int, error) {
	return _Ritsu.Contract.Reserve0(&_Ritsu.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Ritsu *RitsuCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Ritsu *RitsuSession) Reserve1() (*big.Int, error) {
	return _Ritsu.Contract.Reserve1(&_Ritsu.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_Ritsu *RitsuCallerSession) Reserve1() (*big.Int, error) {
	return _Ritsu.Contract.Reserve1(&_Ritsu.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Ritsu *RitsuCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Ritsu *RitsuSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Ritsu.Contract.SupportsInterface(&_Ritsu.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_Ritsu *RitsuCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Ritsu.Contract.SupportsInterface(&_Ritsu.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ritsu *RitsuCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ritsu *RitsuSession) Symbol() (string, error) {
	return _Ritsu.Contract.Symbol(&_Ritsu.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ritsu *RitsuCallerSession) Symbol() (string, error) {
	return _Ritsu.Contract.Symbol(&_Ritsu.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ritsu *RitsuCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ritsu *RitsuSession) Token0() (common.Address, error) {
	return _Ritsu.Contract.Token0(&_Ritsu.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ritsu *RitsuCallerSession) Token0() (common.Address, error) {
	return _Ritsu.Contract.Token0(&_Ritsu.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ritsu *RitsuCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ritsu *RitsuSession) Token1() (common.Address, error) {
	return _Ritsu.Contract.Token1(&_Ritsu.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ritsu *RitsuCallerSession) Token1() (common.Address, error) {
	return _Ritsu.Contract.Token1(&_Ritsu.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ritsu *RitsuCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ritsu *RitsuSession) TotalSupply() (*big.Int, error) {
	return _Ritsu.Contract.TotalSupply(&_Ritsu.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ritsu *RitsuCallerSession) TotalSupply() (*big.Int, error) {
	return _Ritsu.Contract.TotalSupply(&_Ritsu.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Ritsu *RitsuCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ritsu.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Ritsu *RitsuSession) Vault() (common.Address, error) {
	return _Ritsu.Contract.Vault(&_Ritsu.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Ritsu *RitsuCallerSession) Vault() (common.Address, error) {
	return _Ritsu.Contract.Vault(&_Ritsu.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Ritsu *RitsuSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.Approve(&_Ritsu.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.Approve(&_Ritsu.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_Ritsu *RitsuTransactor) Burn(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "burn", _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_Ritsu *RitsuSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Burn(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_Ritsu *RitsuTransactorSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Burn(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuTransactor) BurnSingle(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "burnSingle", _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.BurnSingle(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuTransactorSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.BurnSingle(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_Ritsu *RitsuTransactor) Mint(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "mint", _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_Ritsu *RitsuSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Mint(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_Ritsu *RitsuTransactorSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Mint(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Ritsu *RitsuTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "permit", _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Ritsu *RitsuSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Permit(&_Ritsu.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Ritsu *RitsuTransactorSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Permit(&_Ritsu.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_Ritsu *RitsuTransactor) Permit2(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "permit2", _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_Ritsu *RitsuSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Permit2(&_Ritsu.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_Ritsu *RitsuTransactorSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Permit2(&_Ritsu.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuTransactor) Swap(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "swap", _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Swap(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_Ritsu *RitsuTransactorSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _Ritsu.Contract.Swap(&_Ritsu.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.Transfer(&_Ritsu.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.Transfer(&_Ritsu.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.TransferFrom(&_Ritsu.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Ritsu *RitsuTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ritsu.Contract.TransferFrom(&_Ritsu.TransactOpts, _from, _to, _amount)
}

// RitsuApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ritsu contract.
type RitsuApprovalIterator struct {
	Event *RitsuApproval // Event containing the contract specifics and raw log

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
func (it *RitsuApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuApproval)
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
		it.Event = new(RitsuApproval)
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
func (it *RitsuApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuApproval represents a Approval event raised by the Ritsu contract.
type RitsuApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Ritsu *RitsuFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RitsuApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RitsuApprovalIterator{contract: _Ritsu.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Ritsu *RitsuFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RitsuApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuApproval)
				if err := _Ritsu.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Ritsu *RitsuFilterer) ParseApproval(log types.Log) (*RitsuApproval, error) {
	event := new(RitsuApproval)
	if err := _Ritsu.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RitsuBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Ritsu contract.
type RitsuBurnIterator struct {
	Event *RitsuBurn // Event containing the contract specifics and raw log

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
func (it *RitsuBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuBurn)
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
		it.Event = new(RitsuBurn)
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
func (it *RitsuBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuBurn represents a Burn event raised by the Ritsu contract.
type RitsuBurn struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*RitsuBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RitsuBurnIterator{contract: _Ritsu.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *RitsuBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuBurn)
				if err := _Ritsu.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) ParseBurn(log types.Log) (*RitsuBurn, error) {
	event := new(RitsuBurn)
	if err := _Ritsu.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RitsuMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Ritsu contract.
type RitsuMintIterator struct {
	Event *RitsuMint // Event containing the contract specifics and raw log

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
func (it *RitsuMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuMint)
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
		it.Event = new(RitsuMint)
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
func (it *RitsuMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuMint represents a Mint event raised by the Ritsu contract.
type RitsuMint struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*RitsuMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RitsuMintIterator{contract: _Ritsu.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *RitsuMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuMint)
				if err := _Ritsu.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_Ritsu *RitsuFilterer) ParseMint(log types.Log) (*RitsuMint, error) {
	event := new(RitsuMint)
	if err := _Ritsu.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RitsuSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Ritsu contract.
type RitsuSwapIterator struct {
	Event *RitsuSwap // Event containing the contract specifics and raw log

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
func (it *RitsuSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuSwap)
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
		it.Event = new(RitsuSwap)
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
func (it *RitsuSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuSwap represents a Swap event raised by the Ritsu contract.
type RitsuSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Ritsu *RitsuFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*RitsuSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RitsuSwapIterator{contract: _Ritsu.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Ritsu *RitsuFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *RitsuSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuSwap)
				if err := _Ritsu.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Ritsu *RitsuFilterer) ParseSwap(log types.Log) (*RitsuSwap, error) {
	event := new(RitsuSwap)
	if err := _Ritsu.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RitsuSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Ritsu contract.
type RitsuSyncIterator struct {
	Event *RitsuSync // Event containing the contract specifics and raw log

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
func (it *RitsuSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuSync)
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
		it.Event = new(RitsuSync)
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
func (it *RitsuSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuSync represents a Sync event raised by the Ritsu contract.
type RitsuSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Ritsu *RitsuFilterer) FilterSync(opts *bind.FilterOpts) (*RitsuSyncIterator, error) {

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &RitsuSyncIterator{contract: _Ritsu.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Ritsu *RitsuFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *RitsuSync) (event.Subscription, error) {

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuSync)
				if err := _Ritsu.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_Ritsu *RitsuFilterer) ParseSync(log types.Log) (*RitsuSync, error) {
	event := new(RitsuSync)
	if err := _Ritsu.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RitsuTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ritsu contract.
type RitsuTransferIterator struct {
	Event *RitsuTransfer // Event containing the contract specifics and raw log

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
func (it *RitsuTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RitsuTransfer)
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
		it.Event = new(RitsuTransfer)
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
func (it *RitsuTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RitsuTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RitsuTransfer represents a Transfer event raised by the Ritsu contract.
type RitsuTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ritsu *RitsuFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RitsuTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RitsuTransferIterator{contract: _Ritsu.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ritsu *RitsuFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RitsuTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ritsu.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RitsuTransfer)
				if err := _Ritsu.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ritsu *RitsuFilterer) ParseTransfer(log types.Log) (*RitsuTransfer, error) {
	event := new(RitsuTransfer)
	if err := _Ritsu.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
