// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package symmetric

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Amount *big.Int
	End    *big.Int
}

// SymmetricMetaData contains all meta data concerning the Symmetric contract.
var SymmetricMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TotalUnlock\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"locktime\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"type\",\"type\":\"int128\",\"indexed\":false},{\"name\":\"ts\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ts\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"name\":\"prevSupply\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_admin_addr\",\"type\":\"address\"},{\"name\":\"_admin_unlock_all\",\"type\":\"address\"},{\"name\":\"_max_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_transfer_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_smart_wallet_checker\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_smart_wallet_checker\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_min_unlock_timestamp\",\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_all_unlock\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_last_user_slope\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_point_history__ts\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"locked__end\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"checkpoint\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deposit_for\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"create_lock\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increase_amount\",\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increase_unlock_time\",\"inputs\":[{\"name\":\"_unlock_time\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"_t\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOfAt\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"_block\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[{\"name\":\"t\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupplyAt\",\"inputs\":[{\"name\":\"_block\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"MAXTIME\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"supply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"locked\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"amount\",\"type\":\"int128\"},{\"name\":\"end\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"point_history\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"bias\",\"type\":\"int128\"},{\"name\":\"slope\",\"type\":\"int128\"},{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"blk\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_point_history\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"name\":\"bias\",\"type\":\"int128\"},{\"name\":\"slope\",\"type\":\"int128\"},{\"name\":\"ts\",\"type\":\"uint256\"},{\"name\":\"blk\",\"type\":\"uint256\"}]}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"user_point_epoch\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"slope_changes\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_smart_wallet_checker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"smart_wallet_checker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_unlock_all\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"all_unlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"min_unlock_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// SymmetricABI is the input ABI used to generate the binding from.
// Deprecated: Use SymmetricMetaData.ABI instead.
var SymmetricABI = SymmetricMetaData.ABI

// Symmetric is an auto generated Go binding around an Ethereum contract.
type Symmetric struct {
	SymmetricCaller     // Read-only binding to the contract
	SymmetricTransactor // Write-only binding to the contract
	SymmetricFilterer   // Log filterer for contract events
}

// SymmetricCaller is an auto generated read-only Go binding around an Ethereum contract.
type SymmetricCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymmetricTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SymmetricTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymmetricFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SymmetricFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymmetricSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SymmetricSession struct {
	Contract     *Symmetric        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SymmetricCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SymmetricCallerSession struct {
	Contract *SymmetricCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SymmetricTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SymmetricTransactorSession struct {
	Contract     *SymmetricTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SymmetricRaw is an auto generated low-level Go binding around an Ethereum contract.
type SymmetricRaw struct {
	Contract *Symmetric // Generic contract binding to access the raw methods on
}

// SymmetricCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SymmetricCallerRaw struct {
	Contract *SymmetricCaller // Generic read-only contract binding to access the raw methods on
}

// SymmetricTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SymmetricTransactorRaw struct {
	Contract *SymmetricTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSymmetric creates a new instance of Symmetric, bound to a specific deployed contract.
func NewSymmetric(address common.Address, backend bind.ContractBackend) (*Symmetric, error) {
	contract, err := bindSymmetric(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Symmetric{SymmetricCaller: SymmetricCaller{contract: contract}, SymmetricTransactor: SymmetricTransactor{contract: contract}, SymmetricFilterer: SymmetricFilterer{contract: contract}}, nil
}

// NewSymmetricCaller creates a new read-only instance of Symmetric, bound to a specific deployed contract.
func NewSymmetricCaller(address common.Address, caller bind.ContractCaller) (*SymmetricCaller, error) {
	contract, err := bindSymmetric(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SymmetricCaller{contract: contract}, nil
}

// NewSymmetricTransactor creates a new write-only instance of Symmetric, bound to a specific deployed contract.
func NewSymmetricTransactor(address common.Address, transactor bind.ContractTransactor) (*SymmetricTransactor, error) {
	contract, err := bindSymmetric(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SymmetricTransactor{contract: contract}, nil
}

// NewSymmetricFilterer creates a new log filterer instance of Symmetric, bound to a specific deployed contract.
func NewSymmetricFilterer(address common.Address, filterer bind.ContractFilterer) (*SymmetricFilterer, error) {
	contract, err := bindSymmetric(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SymmetricFilterer{contract: contract}, nil
}

// bindSymmetric binds a generic wrapper to an already deployed contract.
func bindSymmetric(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SymmetricMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Symmetric *SymmetricRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Symmetric.Contract.SymmetricCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Symmetric *SymmetricRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.Contract.SymmetricTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Symmetric *SymmetricRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Symmetric.Contract.SymmetricTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Symmetric *SymmetricCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Symmetric.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Symmetric *SymmetricTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Symmetric *SymmetricTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Symmetric.Contract.contract.Transact(opts, method, params...)
}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Symmetric *SymmetricCaller) MAXTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "MAXTIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Symmetric *SymmetricSession) MAXTIME() (*big.Int, error) {
	return _Symmetric.Contract.MAXTIME(&_Symmetric.CallOpts)
}

// MAXTIME is a free data retrieval call binding the contract method 0xee00ef3a.
//
// Solidity: function MAXTIME() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) MAXTIME() (*big.Int, error) {
	return _Symmetric.Contract.MAXTIME(&_Symmetric.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Symmetric *SymmetricCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Symmetric *SymmetricSession) TOKEN() (common.Address, error) {
	return _Symmetric.Contract.TOKEN(&_Symmetric.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Symmetric *SymmetricCallerSession) TOKEN() (common.Address, error) {
	return _Symmetric.Contract.TOKEN(&_Symmetric.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Symmetric *SymmetricCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Symmetric *SymmetricSession) Admin() (common.Address, error) {
	return _Symmetric.Contract.Admin(&_Symmetric.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Symmetric *SymmetricCallerSession) Admin() (common.Address, error) {
	return _Symmetric.Contract.Admin(&_Symmetric.CallOpts)
}

// AdminUnlockAll is a free data retrieval call binding the contract method 0x14261425.
//
// Solidity: function admin_unlock_all() view returns(address)
func (_Symmetric *SymmetricCaller) AdminUnlockAll(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "admin_unlock_all")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminUnlockAll is a free data retrieval call binding the contract method 0x14261425.
//
// Solidity: function admin_unlock_all() view returns(address)
func (_Symmetric *SymmetricSession) AdminUnlockAll() (common.Address, error) {
	return _Symmetric.Contract.AdminUnlockAll(&_Symmetric.CallOpts)
}

// AdminUnlockAll is a free data retrieval call binding the contract method 0x14261425.
//
// Solidity: function admin_unlock_all() view returns(address)
func (_Symmetric *SymmetricCallerSession) AdminUnlockAll() (common.Address, error) {
	return _Symmetric.Contract.AdminUnlockAll(&_Symmetric.CallOpts)
}

// AllUnlock is a free data retrieval call binding the contract method 0xb3d8f7e7.
//
// Solidity: function all_unlock() view returns(bool)
func (_Symmetric *SymmetricCaller) AllUnlock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "all_unlock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllUnlock is a free data retrieval call binding the contract method 0xb3d8f7e7.
//
// Solidity: function all_unlock() view returns(bool)
func (_Symmetric *SymmetricSession) AllUnlock() (bool, error) {
	return _Symmetric.Contract.AllUnlock(&_Symmetric.CallOpts)
}

// AllUnlock is a free data retrieval call binding the contract method 0xb3d8f7e7.
//
// Solidity: function all_unlock() view returns(bool)
func (_Symmetric *SymmetricCallerSession) AllUnlock() (bool, error) {
	return _Symmetric.Contract.AllUnlock(&_Symmetric.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Symmetric *SymmetricCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Symmetric *SymmetricSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOf(&_Symmetric.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOf(&_Symmetric.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Symmetric *SymmetricCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "balanceOf0", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Symmetric *SymmetricSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOf0(&_Symmetric.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOf0(&_Symmetric.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOfAt(&_Symmetric.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.BalanceOfAt(&_Symmetric.CallOpts, addr, _block)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Symmetric *SymmetricCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Symmetric *SymmetricSession) Decimals() (*big.Int, error) {
	return _Symmetric.Contract.Decimals(&_Symmetric.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) Decimals() (*big.Int, error) {
	return _Symmetric.Contract.Decimals(&_Symmetric.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Symmetric *SymmetricCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Symmetric *SymmetricSession) Epoch() (*big.Int, error) {
	return _Symmetric.Contract.Epoch(&_Symmetric.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) Epoch() (*big.Int, error) {
	return _Symmetric.Contract.Epoch(&_Symmetric.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Symmetric *SymmetricCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Symmetric *SymmetricSession) FutureAdmin() (common.Address, error) {
	return _Symmetric.Contract.FutureAdmin(&_Symmetric.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_Symmetric *SymmetricCallerSession) FutureAdmin() (common.Address, error) {
	return _Symmetric.Contract.FutureAdmin(&_Symmetric.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "future_smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Symmetric.Contract.FutureSmartWalletChecker(&_Symmetric.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _Symmetric.Contract.FutureSmartWalletChecker(&_Symmetric.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Symmetric *SymmetricCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Symmetric *SymmetricSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.GetLastUserSlope(&_Symmetric.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_Symmetric *SymmetricCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.GetLastUserSlope(&_Symmetric.CallOpts, addr)
}

// IsInitialized is a free data retrieval call binding the contract method 0x9a01873c.
//
// Solidity: function is_initialized() view returns(bool)
func (_Symmetric *SymmetricCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "is_initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x9a01873c.
//
// Solidity: function is_initialized() view returns(bool)
func (_Symmetric *SymmetricSession) IsInitialized() (bool, error) {
	return _Symmetric.Contract.IsInitialized(&_Symmetric.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x9a01873c.
//
// Solidity: function is_initialized() view returns(bool)
func (_Symmetric *SymmetricCallerSession) IsInitialized() (bool, error) {
	return _Symmetric.Contract.IsInitialized(&_Symmetric.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns((int128,uint256))
func (_Symmetric *SymmetricCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (Struct1, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "locked", arg0)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns((int128,uint256))
func (_Symmetric *SymmetricSession) Locked(arg0 common.Address) (Struct1, error) {
	return _Symmetric.Contract.Locked(&_Symmetric.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns((int128,uint256))
func (_Symmetric *SymmetricCallerSession) Locked(arg0 common.Address) (Struct1, error) {
	return _Symmetric.Contract.Locked(&_Symmetric.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Symmetric *SymmetricCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Symmetric *SymmetricSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.LockedEnd(&_Symmetric.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _Symmetric.Contract.LockedEnd(&_Symmetric.CallOpts, _addr)
}

// MinUnlockTimestamp is a free data retrieval call binding the contract method 0x8d0644df.
//
// Solidity: function min_unlock_timestamp() view returns(uint256)
func (_Symmetric *SymmetricCaller) MinUnlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "min_unlock_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUnlockTimestamp is a free data retrieval call binding the contract method 0x8d0644df.
//
// Solidity: function min_unlock_timestamp() view returns(uint256)
func (_Symmetric *SymmetricSession) MinUnlockTimestamp() (*big.Int, error) {
	return _Symmetric.Contract.MinUnlockTimestamp(&_Symmetric.CallOpts)
}

// MinUnlockTimestamp is a free data retrieval call binding the contract method 0x8d0644df.
//
// Solidity: function min_unlock_timestamp() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) MinUnlockTimestamp() (*big.Int, error) {
	return _Symmetric.Contract.MinUnlockTimestamp(&_Symmetric.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Symmetric *SymmetricCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Symmetric *SymmetricSession) Name() (string, error) {
	return _Symmetric.Contract.Name(&_Symmetric.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Symmetric *SymmetricCallerSession) Name() (string, error) {
	return _Symmetric.Contract.Name(&_Symmetric.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (Struct0, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "point_history", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricSession) PointHistory(arg0 *big.Int) (Struct0, error) {
	return _Symmetric.Contract.PointHistory(&_Symmetric.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricCallerSession) PointHistory(arg0 *big.Int) (Struct0, error) {
	return _Symmetric.Contract.PointHistory(&_Symmetric.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Symmetric *SymmetricCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Symmetric *SymmetricSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.SlopeChanges(&_Symmetric.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_Symmetric *SymmetricCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.SlopeChanges(&_Symmetric.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricSession) SmartWalletChecker() (common.Address, error) {
	return _Symmetric.Contract.SmartWalletChecker(&_Symmetric.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_Symmetric *SymmetricCallerSession) SmartWalletChecker() (common.Address, error) {
	return _Symmetric.Contract.SmartWalletChecker(&_Symmetric.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Symmetric *SymmetricCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Symmetric *SymmetricSession) Supply() (*big.Int, error) {
	return _Symmetric.Contract.Supply(&_Symmetric.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) Supply() (*big.Int, error) {
	return _Symmetric.Contract.Supply(&_Symmetric.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Symmetric *SymmetricCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Symmetric *SymmetricSession) Symbol() (string, error) {
	return _Symmetric.Contract.Symbol(&_Symmetric.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Symmetric *SymmetricCallerSession) Symbol() (string, error) {
	return _Symmetric.Contract.Symbol(&_Symmetric.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Symmetric *SymmetricCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Symmetric *SymmetricSession) Token() (common.Address, error) {
	return _Symmetric.Contract.Token(&_Symmetric.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Symmetric *SymmetricCallerSession) Token() (common.Address, error) {
	return _Symmetric.Contract.Token(&_Symmetric.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Symmetric *SymmetricCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Symmetric *SymmetricSession) TotalSupply() (*big.Int, error) {
	return _Symmetric.Contract.TotalSupply(&_Symmetric.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Symmetric *SymmetricCallerSession) TotalSupply() (*big.Int, error) {
	return _Symmetric.Contract.TotalSupply(&_Symmetric.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Symmetric *SymmetricCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "totalSupply0", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Symmetric *SymmetricSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.TotalSupply0(&_Symmetric.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.TotalSupply0(&_Symmetric.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.TotalSupplyAt(&_Symmetric.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.TotalSupplyAt(&_Symmetric.CallOpts, _block)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Symmetric *SymmetricCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Symmetric *SymmetricSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Symmetric.Contract.UserPointEpoch(&_Symmetric.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _Symmetric.Contract.UserPointEpoch(&_Symmetric.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (Struct0, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (Struct0, error) {
	return _Symmetric.Contract.UserPointHistory(&_Symmetric.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns((int128,int128,uint256,uint256))
func (_Symmetric *SymmetricCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (Struct0, error) {
	return _Symmetric.Contract.UserPointHistory(&_Symmetric.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Symmetric *SymmetricCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Symmetric.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Symmetric *SymmetricSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.UserPointHistoryTs(&_Symmetric.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_Symmetric *SymmetricCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _Symmetric.Contract.UserPointHistoryTs(&_Symmetric.CallOpts, _addr, _idx)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Symmetric *SymmetricTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Symmetric *SymmetricSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Symmetric.Contract.ApplySmartWalletChecker(&_Symmetric.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_Symmetric *SymmetricTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _Symmetric.Contract.ApplySmartWalletChecker(&_Symmetric.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Symmetric *SymmetricTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Symmetric *SymmetricSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Symmetric.Contract.ApplyTransferOwnership(&_Symmetric.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Symmetric *SymmetricTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Symmetric.Contract.ApplyTransferOwnership(&_Symmetric.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Symmetric *SymmetricTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Symmetric *SymmetricSession) Checkpoint() (*types.Transaction, error) {
	return _Symmetric.Contract.Checkpoint(&_Symmetric.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_Symmetric *SymmetricTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _Symmetric.Contract.Checkpoint(&_Symmetric.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Symmetric *SymmetricTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Symmetric *SymmetricSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Symmetric.Contract.CommitSmartWalletChecker(&_Symmetric.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_Symmetric *SymmetricTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _Symmetric.Contract.CommitSmartWalletChecker(&_Symmetric.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Symmetric *SymmetricTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Symmetric *SymmetricSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Symmetric.Contract.CommitTransferOwnership(&_Symmetric.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_Symmetric *SymmetricTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _Symmetric.Contract.CommitTransferOwnership(&_Symmetric.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Symmetric *SymmetricTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Symmetric *SymmetricSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.CreateLock(&_Symmetric.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_Symmetric *SymmetricTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.CreateLock(&_Symmetric.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Symmetric *SymmetricTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Symmetric *SymmetricSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.DepositFor(&_Symmetric.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_Symmetric *SymmetricTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.DepositFor(&_Symmetric.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Symmetric *SymmetricTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Symmetric *SymmetricSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.IncreaseAmount(&_Symmetric.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_Symmetric *SymmetricTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.IncreaseAmount(&_Symmetric.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Symmetric *SymmetricTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Symmetric *SymmetricSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.IncreaseUnlockTime(&_Symmetric.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_Symmetric *SymmetricTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.IncreaseUnlockTime(&_Symmetric.TransactOpts, _unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d26ce81.
//
// Solidity: function initialize(address _token_addr, string _name, string _symbol, address _admin_addr, address _admin_unlock_all, uint256 _max_time) returns()
func (_Symmetric *SymmetricTransactor) Initialize(opts *bind.TransactOpts, _token_addr common.Address, _name string, _symbol string, _admin_addr common.Address, _admin_unlock_all common.Address, _max_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "initialize", _token_addr, _name, _symbol, _admin_addr, _admin_unlock_all, _max_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d26ce81.
//
// Solidity: function initialize(address _token_addr, string _name, string _symbol, address _admin_addr, address _admin_unlock_all, uint256 _max_time) returns()
func (_Symmetric *SymmetricSession) Initialize(_token_addr common.Address, _name string, _symbol string, _admin_addr common.Address, _admin_unlock_all common.Address, _max_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.Initialize(&_Symmetric.TransactOpts, _token_addr, _name, _symbol, _admin_addr, _admin_unlock_all, _max_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x8d26ce81.
//
// Solidity: function initialize(address _token_addr, string _name, string _symbol, address _admin_addr, address _admin_unlock_all, uint256 _max_time) returns()
func (_Symmetric *SymmetricTransactorSession) Initialize(_token_addr common.Address, _name string, _symbol string, _admin_addr common.Address, _admin_unlock_all common.Address, _max_time *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.Initialize(&_Symmetric.TransactOpts, _token_addr, _name, _symbol, _admin_addr, _admin_unlock_all, _max_time)
}

// SetAllUnlock is a paid mutator transaction binding the contract method 0x66e01ca9.
//
// Solidity: function set_all_unlock() returns()
func (_Symmetric *SymmetricTransactor) SetAllUnlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "set_all_unlock")
}

// SetAllUnlock is a paid mutator transaction binding the contract method 0x66e01ca9.
//
// Solidity: function set_all_unlock() returns()
func (_Symmetric *SymmetricSession) SetAllUnlock() (*types.Transaction, error) {
	return _Symmetric.Contract.SetAllUnlock(&_Symmetric.TransactOpts)
}

// SetAllUnlock is a paid mutator transaction binding the contract method 0x66e01ca9.
//
// Solidity: function set_all_unlock() returns()
func (_Symmetric *SymmetricTransactorSession) SetAllUnlock() (*types.Transaction, error) {
	return _Symmetric.Contract.SetAllUnlock(&_Symmetric.TransactOpts)
}

// SetMinUnlockTimestamp is a paid mutator transaction binding the contract method 0x81e35933.
//
// Solidity: function set_min_unlock_timestamp(uint256 _timestamp) returns()
func (_Symmetric *SymmetricTransactor) SetMinUnlockTimestamp(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "set_min_unlock_timestamp", _timestamp)
}

// SetMinUnlockTimestamp is a paid mutator transaction binding the contract method 0x81e35933.
//
// Solidity: function set_min_unlock_timestamp(uint256 _timestamp) returns()
func (_Symmetric *SymmetricSession) SetMinUnlockTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.SetMinUnlockTimestamp(&_Symmetric.TransactOpts, _timestamp)
}

// SetMinUnlockTimestamp is a paid mutator transaction binding the contract method 0x81e35933.
//
// Solidity: function set_min_unlock_timestamp(uint256 _timestamp) returns()
func (_Symmetric *SymmetricTransactorSession) SetMinUnlockTimestamp(_timestamp *big.Int) (*types.Transaction, error) {
	return _Symmetric.Contract.SetMinUnlockTimestamp(&_Symmetric.TransactOpts, _timestamp)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Symmetric *SymmetricTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Symmetric.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Symmetric *SymmetricSession) Withdraw() (*types.Transaction, error) {
	return _Symmetric.Contract.Withdraw(&_Symmetric.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Symmetric *SymmetricTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Symmetric.Contract.Withdraw(&_Symmetric.TransactOpts)
}

// SymmetricApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the Symmetric contract.
type SymmetricApplyOwnershipIterator struct {
	Event *SymmetricApplyOwnership // Event containing the contract specifics and raw log

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
func (it *SymmetricApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricApplyOwnership)
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
		it.Event = new(SymmetricApplyOwnership)
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
func (it *SymmetricApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricApplyOwnership represents a ApplyOwnership event raised by the Symmetric contract.
type SymmetricApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Symmetric *SymmetricFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*SymmetricApplyOwnershipIterator, error) {

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &SymmetricApplyOwnershipIterator{contract: _Symmetric.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Symmetric *SymmetricFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *SymmetricApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricApplyOwnership)
				if err := _Symmetric.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_Symmetric *SymmetricFilterer) ParseApplyOwnership(log types.Log) (*SymmetricApplyOwnership, error) {
	event := new(SymmetricApplyOwnership)
	if err := _Symmetric.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymmetricCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the Symmetric contract.
type SymmetricCommitOwnershipIterator struct {
	Event *SymmetricCommitOwnership // Event containing the contract specifics and raw log

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
func (it *SymmetricCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricCommitOwnership)
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
		it.Event = new(SymmetricCommitOwnership)
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
func (it *SymmetricCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricCommitOwnership represents a CommitOwnership event raised by the Symmetric contract.
type SymmetricCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Symmetric *SymmetricFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*SymmetricCommitOwnershipIterator, error) {

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &SymmetricCommitOwnershipIterator{contract: _Symmetric.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Symmetric *SymmetricFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *SymmetricCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricCommitOwnership)
				if err := _Symmetric.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_Symmetric *SymmetricFilterer) ParseCommitOwnership(log types.Log) (*SymmetricCommitOwnership, error) {
	event := new(SymmetricCommitOwnership)
	if err := _Symmetric.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymmetricDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Symmetric contract.
type SymmetricDepositIterator struct {
	Event *SymmetricDeposit // Event containing the contract specifics and raw log

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
func (it *SymmetricDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricDeposit)
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
		it.Event = new(SymmetricDeposit)
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
func (it *SymmetricDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricDeposit represents a Deposit event raised by the Symmetric contract.
type SymmetricDeposit struct {
	Provider common.Address
	Value    *big.Int
	Locktime *big.Int
	Arg3     *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Symmetric *SymmetricFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*SymmetricDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &SymmetricDepositIterator{contract: _Symmetric.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Symmetric *SymmetricFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SymmetricDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricDeposit)
				if err := _Symmetric.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_Symmetric *SymmetricFilterer) ParseDeposit(log types.Log) (*SymmetricDeposit, error) {
	event := new(SymmetricDeposit)
	if err := _Symmetric.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymmetricSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the Symmetric contract.
type SymmetricSupplyIterator struct {
	Event *SymmetricSupply // Event containing the contract specifics and raw log

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
func (it *SymmetricSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricSupply)
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
		it.Event = new(SymmetricSupply)
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
func (it *SymmetricSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricSupply represents a Supply event raised by the Symmetric contract.
type SymmetricSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Symmetric *SymmetricFilterer) FilterSupply(opts *bind.FilterOpts) (*SymmetricSupplyIterator, error) {

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &SymmetricSupplyIterator{contract: _Symmetric.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Symmetric *SymmetricFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *SymmetricSupply) (event.Subscription, error) {

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricSupply)
				if err := _Symmetric.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_Symmetric *SymmetricFilterer) ParseSupply(log types.Log) (*SymmetricSupply, error) {
	event := new(SymmetricSupply)
	if err := _Symmetric.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymmetricTotalUnlockIterator is returned from FilterTotalUnlock and is used to iterate over the raw logs and unpacked data for TotalUnlock events raised by the Symmetric contract.
type SymmetricTotalUnlockIterator struct {
	Event *SymmetricTotalUnlock // Event containing the contract specifics and raw log

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
func (it *SymmetricTotalUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricTotalUnlock)
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
		it.Event = new(SymmetricTotalUnlock)
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
func (it *SymmetricTotalUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricTotalUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricTotalUnlock represents a TotalUnlock event raised by the Symmetric contract.
type SymmetricTotalUnlock struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTotalUnlock is a free log retrieval operation binding the contract event 0x7ca88488f569b55d1fb073429f75839e505182c1f6d26e5caed22e9828df430c.
//
// Solidity: event TotalUnlock(bool status)
func (_Symmetric *SymmetricFilterer) FilterTotalUnlock(opts *bind.FilterOpts) (*SymmetricTotalUnlockIterator, error) {

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "TotalUnlock")
	if err != nil {
		return nil, err
	}
	return &SymmetricTotalUnlockIterator{contract: _Symmetric.contract, event: "TotalUnlock", logs: logs, sub: sub}, nil
}

// WatchTotalUnlock is a free log subscription operation binding the contract event 0x7ca88488f569b55d1fb073429f75839e505182c1f6d26e5caed22e9828df430c.
//
// Solidity: event TotalUnlock(bool status)
func (_Symmetric *SymmetricFilterer) WatchTotalUnlock(opts *bind.WatchOpts, sink chan<- *SymmetricTotalUnlock) (event.Subscription, error) {

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "TotalUnlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricTotalUnlock)
				if err := _Symmetric.contract.UnpackLog(event, "TotalUnlock", log); err != nil {
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

// ParseTotalUnlock is a log parse operation binding the contract event 0x7ca88488f569b55d1fb073429f75839e505182c1f6d26e5caed22e9828df430c.
//
// Solidity: event TotalUnlock(bool status)
func (_Symmetric *SymmetricFilterer) ParseTotalUnlock(log types.Log) (*SymmetricTotalUnlock, error) {
	event := new(SymmetricTotalUnlock)
	if err := _Symmetric.contract.UnpackLog(event, "TotalUnlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymmetricWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Symmetric contract.
type SymmetricWithdrawIterator struct {
	Event *SymmetricWithdraw // Event containing the contract specifics and raw log

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
func (it *SymmetricWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymmetricWithdraw)
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
		it.Event = new(SymmetricWithdraw)
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
func (it *SymmetricWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymmetricWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymmetricWithdraw represents a Withdraw event raised by the Symmetric contract.
type SymmetricWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Symmetric *SymmetricFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*SymmetricWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Symmetric.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &SymmetricWithdrawIterator{contract: _Symmetric.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Symmetric *SymmetricFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SymmetricWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Symmetric.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymmetricWithdraw)
				if err := _Symmetric.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_Symmetric *SymmetricFilterer) ParseWithdraw(log types.Log) (*SymmetricWithdraw, error) {
	event := new(SymmetricWithdraw)
	if err := _Symmetric.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
