// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package drips

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

// StakingERC20Init is an auto generated low-level Go binding around an user-defined struct.
type StakingERC20Init struct {
	Admin          common.Address
	Operator       common.Address
	Pauser         common.Address
	Asset          common.Address
	Cooldown       *big.Int
	MinStake       *big.Int
	MaxStakeSupply *big.Int
	Allocator      common.Address
}

// DripsMetaData contains all meta data concerning the Drips contract.
var DripsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressZeroNotExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Cooldown\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"}],\"name\":\"DepositAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositOverBond\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DurationNotExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"EnumerableMapNonexistentKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ExceededWithdrawal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientWithdrawableBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DepositWithDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"setterSelector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"setterSignature\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ProtocolConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"name\":\"UserLockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestUnlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"}],\"name\":\"UserUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"addLockDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"deposited\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durations\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getUserLockUpArrays\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getUserLockUps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakeSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"}],\"internalType\":\"structStakingERC20.Init\",\"name\":\"init\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"contractIPauser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"removeLockDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCooldown\",\"type\":\"uint256\"}],\"name\":\"setCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxStakeSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxStakeSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinStake\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"unlockLockups\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"updateLockUps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"userStakeCooldown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DripsABI is the input ABI used to generate the binding from.
// Deprecated: Use DripsMetaData.ABI instead.
var DripsABI = DripsMetaData.ABI

// Drips is an auto generated Go binding around an Ethereum contract.
type Drips struct {
	DripsCaller     // Read-only binding to the contract
	DripsTransactor // Write-only binding to the contract
	DripsFilterer   // Log filterer for contract events
}

// DripsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DripsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DripsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DripsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DripsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DripsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DripsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DripsSession struct {
	Contract     *Drips            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DripsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DripsCallerSession struct {
	Contract *DripsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DripsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DripsTransactorSession struct {
	Contract     *DripsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DripsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DripsRaw struct {
	Contract *Drips // Generic contract binding to access the raw methods on
}

// DripsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DripsCallerRaw struct {
	Contract *DripsCaller // Generic read-only contract binding to access the raw methods on
}

// DripsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DripsTransactorRaw struct {
	Contract *DripsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDrips creates a new instance of Drips, bound to a specific deployed contract.
func NewDrips(address common.Address, backend bind.ContractBackend) (*Drips, error) {
	contract, err := bindDrips(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Drips{DripsCaller: DripsCaller{contract: contract}, DripsTransactor: DripsTransactor{contract: contract}, DripsFilterer: DripsFilterer{contract: contract}}, nil
}

// NewDripsCaller creates a new read-only instance of Drips, bound to a specific deployed contract.
func NewDripsCaller(address common.Address, caller bind.ContractCaller) (*DripsCaller, error) {
	contract, err := bindDrips(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DripsCaller{contract: contract}, nil
}

// NewDripsTransactor creates a new write-only instance of Drips, bound to a specific deployed contract.
func NewDripsTransactor(address common.Address, transactor bind.ContractTransactor) (*DripsTransactor, error) {
	contract, err := bindDrips(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DripsTransactor{contract: contract}, nil
}

// NewDripsFilterer creates a new log filterer instance of Drips, bound to a specific deployed contract.
func NewDripsFilterer(address common.Address, filterer bind.ContractFilterer) (*DripsFilterer, error) {
	contract, err := bindDrips(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DripsFilterer{contract: contract}, nil
}

// bindDrips binds a generic wrapper to an already deployed contract.
func bindDrips(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DripsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Drips *DripsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Drips.Contract.DripsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Drips *DripsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Drips.Contract.DripsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Drips *DripsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Drips.Contract.DripsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Drips *DripsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Drips.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Drips *DripsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Drips.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Drips *DripsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Drips.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Drips *DripsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Drips *DripsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Drips.Contract.DEFAULTADMINROLE(&_Drips.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Drips *DripsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Drips.Contract.DEFAULTADMINROLE(&_Drips.CallOpts)
}

// STAKINGOPERATORROLE is a free data retrieval call binding the contract method 0x97e3b7f4.
//
// Solidity: function STAKING_OPERATOR_ROLE() view returns(bytes32)
func (_Drips *DripsCaller) STAKINGOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "STAKING_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGOPERATORROLE is a free data retrieval call binding the contract method 0x97e3b7f4.
//
// Solidity: function STAKING_OPERATOR_ROLE() view returns(bytes32)
func (_Drips *DripsSession) STAKINGOPERATORROLE() ([32]byte, error) {
	return _Drips.Contract.STAKINGOPERATORROLE(&_Drips.CallOpts)
}

// STAKINGOPERATORROLE is a free data retrieval call binding the contract method 0x97e3b7f4.
//
// Solidity: function STAKING_OPERATOR_ROLE() view returns(bytes32)
func (_Drips *DripsCallerSession) STAKINGOPERATORROLE() ([32]byte, error) {
	return _Drips.Contract.STAKINGOPERATORROLE(&_Drips.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_Drips *DripsCaller) Allocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "allocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_Drips *DripsSession) Allocator() (common.Address, error) {
	return _Drips.Contract.Allocator(&_Drips.CallOpts)
}

// Allocator is a free data retrieval call binding the contract method 0xaa5dcecc.
//
// Solidity: function allocator() view returns(address)
func (_Drips *DripsCallerSession) Allocator() (common.Address, error) {
	return _Drips.Contract.Allocator(&_Drips.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Drips *DripsCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Drips *DripsSession) Asset() (common.Address, error) {
	return _Drips.Contract.Asset(&_Drips.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Drips *DripsCallerSession) Asset() (common.Address, error) {
	return _Drips.Contract.Asset(&_Drips.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Drips *DripsCaller) Cooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "cooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Drips *DripsSession) Cooldown() (*big.Int, error) {
	return _Drips.Contract.Cooldown(&_Drips.CallOpts)
}

// Cooldown is a free data retrieval call binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() view returns(uint256)
func (_Drips *DripsCallerSession) Cooldown() (*big.Int, error) {
	return _Drips.Contract.Cooldown(&_Drips.CallOpts)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address account) view returns(uint256)
func (_Drips *DripsCaller) Deposited(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "deposited", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address account) view returns(uint256)
func (_Drips *DripsSession) Deposited(account common.Address) (*big.Int, error) {
	return _Drips.Contract.Deposited(&_Drips.CallOpts, account)
}

// Deposited is a free data retrieval call binding the contract method 0xcb13cddb.
//
// Solidity: function deposited(address account) view returns(uint256)
func (_Drips *DripsCallerSession) Deposited(account common.Address) (*big.Int, error) {
	return _Drips.Contract.Deposited(&_Drips.CallOpts, account)
}

// Durations is a free data retrieval call binding the contract method 0x4a41d89d.
//
// Solidity: function durations() view returns(uint256[])
func (_Drips *DripsCaller) Durations(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "durations")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Durations is a free data retrieval call binding the contract method 0x4a41d89d.
//
// Solidity: function durations() view returns(uint256[])
func (_Drips *DripsSession) Durations() ([]*big.Int, error) {
	return _Drips.Contract.Durations(&_Drips.CallOpts)
}

// Durations is a free data retrieval call binding the contract method 0x4a41d89d.
//
// Solidity: function durations() view returns(uint256[])
func (_Drips *DripsCallerSession) Durations() ([]*big.Int, error) {
	return _Drips.Contract.Durations(&_Drips.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Drips *DripsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Drips *DripsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Drips.Contract.GetRoleAdmin(&_Drips.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Drips *DripsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Drips.Contract.GetRoleAdmin(&_Drips.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Drips *DripsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Drips *DripsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Drips.Contract.GetRoleMember(&_Drips.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Drips *DripsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Drips.Contract.GetRoleMember(&_Drips.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Drips *DripsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Drips *DripsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Drips.Contract.GetRoleMemberCount(&_Drips.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Drips *DripsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Drips.Contract.GetRoleMemberCount(&_Drips.CallOpts, role)
}

// GetUserLockUpArrays is a free data retrieval call binding the contract method 0x5f019d51.
//
// Solidity: function getUserLockUpArrays(address owner) view returns(uint256[], uint256[], uint256[])
func (_Drips *DripsCaller) GetUserLockUpArrays(opts *bind.CallOpts, owner common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "getUserLockUpArrays", owner)

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, err

}

// GetUserLockUpArrays is a free data retrieval call binding the contract method 0x5f019d51.
//
// Solidity: function getUserLockUpArrays(address owner) view returns(uint256[], uint256[], uint256[])
func (_Drips *DripsSession) GetUserLockUpArrays(owner common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	return _Drips.Contract.GetUserLockUpArrays(&_Drips.CallOpts, owner)
}

// GetUserLockUpArrays is a free data retrieval call binding the contract method 0x5f019d51.
//
// Solidity: function getUserLockUpArrays(address owner) view returns(uint256[], uint256[], uint256[])
func (_Drips *DripsCallerSession) GetUserLockUpArrays(owner common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	return _Drips.Contract.GetUserLockUpArrays(&_Drips.CallOpts, owner)
}

// GetUserLockUps is a free data retrieval call binding the contract method 0x0689ef28.
//
// Solidity: function getUserLockUps(address owner) view returns(uint256)
func (_Drips *DripsCaller) GetUserLockUps(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "getUserLockUps", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLockUps is a free data retrieval call binding the contract method 0x0689ef28.
//
// Solidity: function getUserLockUps(address owner) view returns(uint256)
func (_Drips *DripsSession) GetUserLockUps(owner common.Address) (*big.Int, error) {
	return _Drips.Contract.GetUserLockUps(&_Drips.CallOpts, owner)
}

// GetUserLockUps is a free data retrieval call binding the contract method 0x0689ef28.
//
// Solidity: function getUserLockUps(address owner) view returns(uint256)
func (_Drips *DripsCallerSession) GetUserLockUps(owner common.Address) (*big.Int, error) {
	return _Drips.Contract.GetUserLockUps(&_Drips.CallOpts, owner)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Drips *DripsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Drips *DripsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Drips.Contract.HasRole(&_Drips.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Drips *DripsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Drips.Contract.HasRole(&_Drips.CallOpts, role, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Drips *DripsCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Drips *DripsSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Drips.Contract.MaxDeposit(&_Drips.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Drips *DripsCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Drips.Contract.MaxDeposit(&_Drips.CallOpts, arg0)
}

// MaxStakeSupply is a free data retrieval call binding the contract method 0x74e71acb.
//
// Solidity: function maxStakeSupply() view returns(uint256)
func (_Drips *DripsCaller) MaxStakeSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "maxStakeSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeSupply is a free data retrieval call binding the contract method 0x74e71acb.
//
// Solidity: function maxStakeSupply() view returns(uint256)
func (_Drips *DripsSession) MaxStakeSupply() (*big.Int, error) {
	return _Drips.Contract.MaxStakeSupply(&_Drips.CallOpts)
}

// MaxStakeSupply is a free data retrieval call binding the contract method 0x74e71acb.
//
// Solidity: function maxStakeSupply() view returns(uint256)
func (_Drips *DripsCallerSession) MaxStakeSupply() (*big.Int, error) {
	return _Drips.Contract.MaxStakeSupply(&_Drips.CallOpts)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Drips *DripsCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Drips *DripsSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Drips.Contract.MaxWithdraw(&_Drips.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Drips *DripsCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Drips.Contract.MaxWithdraw(&_Drips.CallOpts, owner)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Drips *DripsCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Drips *DripsSession) MinStake() (*big.Int, error) {
	return _Drips.Contract.MinStake(&_Drips.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() view returns(uint256)
func (_Drips *DripsCallerSession) MinStake() (*big.Int, error) {
	return _Drips.Contract.MinStake(&_Drips.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Drips *DripsCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Drips *DripsSession) Pauser() (common.Address, error) {
	return _Drips.Contract.Pauser(&_Drips.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Drips *DripsCallerSession) Pauser() (common.Address, error) {
	return _Drips.Contract.Pauser(&_Drips.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Drips *DripsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Drips *DripsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Drips.Contract.SupportsInterface(&_Drips.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Drips *DripsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Drips.Contract.SupportsInterface(&_Drips.CallOpts, interfaceId)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Drips *DripsCaller) TotalDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "totalDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Drips *DripsSession) TotalDeposit() (*big.Int, error) {
	return _Drips.Contract.TotalDeposit(&_Drips.CallOpts)
}

// TotalDeposit is a free data retrieval call binding the contract method 0xf6153ccd.
//
// Solidity: function totalDeposit() view returns(uint256)
func (_Drips *DripsCallerSession) TotalDeposit() (*big.Int, error) {
	return _Drips.Contract.TotalDeposit(&_Drips.CallOpts)
}

// TotalParticipants is a free data retrieval call binding the contract method 0xa26dbf26.
//
// Solidity: function totalParticipants() view returns(uint256)
func (_Drips *DripsCaller) TotalParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "totalParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalParticipants is a free data retrieval call binding the contract method 0xa26dbf26.
//
// Solidity: function totalParticipants() view returns(uint256)
func (_Drips *DripsSession) TotalParticipants() (*big.Int, error) {
	return _Drips.Contract.TotalParticipants(&_Drips.CallOpts)
}

// TotalParticipants is a free data retrieval call binding the contract method 0xa26dbf26.
//
// Solidity: function totalParticipants() view returns(uint256)
func (_Drips *DripsCallerSession) TotalParticipants() (*big.Int, error) {
	return _Drips.Contract.TotalParticipants(&_Drips.CallOpts)
}

// UserStakeCooldown is a free data retrieval call binding the contract method 0xa0edb2df.
//
// Solidity: function userStakeCooldown(address depositor) view returns(bool, uint256)
func (_Drips *DripsCaller) UserStakeCooldown(opts *bind.CallOpts, depositor common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _Drips.contract.Call(opts, &out, "userStakeCooldown", depositor)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// UserStakeCooldown is a free data retrieval call binding the contract method 0xa0edb2df.
//
// Solidity: function userStakeCooldown(address depositor) view returns(bool, uint256)
func (_Drips *DripsSession) UserStakeCooldown(depositor common.Address) (bool, *big.Int, error) {
	return _Drips.Contract.UserStakeCooldown(&_Drips.CallOpts, depositor)
}

// UserStakeCooldown is a free data retrieval call binding the contract method 0xa0edb2df.
//
// Solidity: function userStakeCooldown(address depositor) view returns(bool, uint256)
func (_Drips *DripsCallerSession) UserStakeCooldown(depositor common.Address) (bool, *big.Int, error) {
	return _Drips.Contract.UserStakeCooldown(&_Drips.CallOpts, depositor)
}

// AddLockDuration is a paid mutator transaction binding the contract method 0xef9beeee.
//
// Solidity: function addLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsTransactor) AddLockDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "addLockDuration", duration)
}

// AddLockDuration is a paid mutator transaction binding the contract method 0xef9beeee.
//
// Solidity: function addLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsSession) AddLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.AddLockDuration(&_Drips.TransactOpts, duration)
}

// AddLockDuration is a paid mutator transaction binding the contract method 0xef9beeee.
//
// Solidity: function addLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsTransactorSession) AddLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.AddLockDuration(&_Drips.TransactOpts, duration)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) payable returns(uint256)
func (_Drips *DripsTransactor) Deposit(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "deposit", arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) payable returns(uint256)
func (_Drips *DripsSession) Deposit(arg0 *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.Deposit(&_Drips.TransactOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 ) payable returns(uint256)
func (_Drips *DripsTransactorSession) Deposit(arg0 *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.Deposit(&_Drips.TransactOpts, arg0)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 assets, uint256 duration) returns(uint256)
func (_Drips *DripsTransactor) Deposit0(opts *bind.TransactOpts, assets *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "deposit0", assets, duration)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 assets, uint256 duration) returns(uint256)
func (_Drips *DripsSession) Deposit0(assets *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.Deposit0(&_Drips.TransactOpts, assets, duration)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 assets, uint256 duration) returns(uint256)
func (_Drips *DripsTransactorSession) Deposit0(assets *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.Deposit0(&_Drips.TransactOpts, assets, duration)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Drips *DripsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Drips *DripsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.Contract.GrantRole(&_Drips.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Drips *DripsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.Contract.GrantRole(&_Drips.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeca19736.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256,uint256,address) init) returns()
func (_Drips *DripsTransactor) Initialize(opts *bind.TransactOpts, init StakingERC20Init) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "initialize", init)
}

// Initialize is a paid mutator transaction binding the contract method 0xeca19736.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256,uint256,address) init) returns()
func (_Drips *DripsSession) Initialize(init StakingERC20Init) (*types.Transaction, error) {
	return _Drips.Contract.Initialize(&_Drips.TransactOpts, init)
}

// Initialize is a paid mutator transaction binding the contract method 0xeca19736.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256,uint256,address) init) returns()
func (_Drips *DripsTransactorSession) Initialize(init StakingERC20Init) (*types.Transaction, error) {
	return _Drips.Contract.Initialize(&_Drips.TransactOpts, init)
}

// RemoveLockDuration is a paid mutator transaction binding the contract method 0x9163b7eb.
//
// Solidity: function removeLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsTransactor) RemoveLockDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "removeLockDuration", duration)
}

// RemoveLockDuration is a paid mutator transaction binding the contract method 0x9163b7eb.
//
// Solidity: function removeLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsSession) RemoveLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.RemoveLockDuration(&_Drips.TransactOpts, duration)
}

// RemoveLockDuration is a paid mutator transaction binding the contract method 0x9163b7eb.
//
// Solidity: function removeLockDuration(uint256 duration) returns(bool)
func (_Drips *DripsTransactorSession) RemoveLockDuration(duration *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.RemoveLockDuration(&_Drips.TransactOpts, duration)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Drips *DripsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Drips *DripsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Drips.Contract.RenounceRole(&_Drips.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Drips *DripsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Drips.Contract.RenounceRole(&_Drips.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Drips *DripsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Drips *DripsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.Contract.RevokeRole(&_Drips.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Drips *DripsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Drips.Contract.RevokeRole(&_Drips.TransactOpts, role, account)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_Drips *DripsTransactor) SetCooldown(opts *bind.TransactOpts, newCooldown *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "setCooldown", newCooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_Drips *DripsSession) SetCooldown(newCooldown *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetCooldown(&_Drips.TransactOpts, newCooldown)
}

// SetCooldown is a paid mutator transaction binding the contract method 0x4fc3f41a.
//
// Solidity: function setCooldown(uint256 newCooldown) returns()
func (_Drips *DripsTransactorSession) SetCooldown(newCooldown *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetCooldown(&_Drips.TransactOpts, newCooldown)
}

// SetMaxStakeSupply is a paid mutator transaction binding the contract method 0x24221016.
//
// Solidity: function setMaxStakeSupply(uint256 newMaxStakeSupply) returns()
func (_Drips *DripsTransactor) SetMaxStakeSupply(opts *bind.TransactOpts, newMaxStakeSupply *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "setMaxStakeSupply", newMaxStakeSupply)
}

// SetMaxStakeSupply is a paid mutator transaction binding the contract method 0x24221016.
//
// Solidity: function setMaxStakeSupply(uint256 newMaxStakeSupply) returns()
func (_Drips *DripsSession) SetMaxStakeSupply(newMaxStakeSupply *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetMaxStakeSupply(&_Drips.TransactOpts, newMaxStakeSupply)
}

// SetMaxStakeSupply is a paid mutator transaction binding the contract method 0x24221016.
//
// Solidity: function setMaxStakeSupply(uint256 newMaxStakeSupply) returns()
func (_Drips *DripsTransactorSession) SetMaxStakeSupply(newMaxStakeSupply *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetMaxStakeSupply(&_Drips.TransactOpts, newMaxStakeSupply)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 newMinStake) returns()
func (_Drips *DripsTransactor) SetMinStake(opts *bind.TransactOpts, newMinStake *big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "setMinStake", newMinStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 newMinStake) returns()
func (_Drips *DripsSession) SetMinStake(newMinStake *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetMinStake(&_Drips.TransactOpts, newMinStake)
}

// SetMinStake is a paid mutator transaction binding the contract method 0x8c80fd90.
//
// Solidity: function setMinStake(uint256 newMinStake) returns()
func (_Drips *DripsTransactorSession) SetMinStake(newMinStake *big.Int) (*types.Transaction, error) {
	return _Drips.Contract.SetMinStake(&_Drips.TransactOpts, newMinStake)
}

// UnlockLockups is a paid mutator transaction binding the contract method 0xba61ecae.
//
// Solidity: function unlockLockups(address[] users, uint256[] amounts) returns()
func (_Drips *DripsTransactor) UnlockLockups(opts *bind.TransactOpts, users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "unlockLockups", users, amounts)
}

// UnlockLockups is a paid mutator transaction binding the contract method 0xba61ecae.
//
// Solidity: function unlockLockups(address[] users, uint256[] amounts) returns()
func (_Drips *DripsSession) UnlockLockups(users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Drips.Contract.UnlockLockups(&_Drips.TransactOpts, users, amounts)
}

// UnlockLockups is a paid mutator transaction binding the contract method 0xba61ecae.
//
// Solidity: function unlockLockups(address[] users, uint256[] amounts) returns()
func (_Drips *DripsTransactorSession) UnlockLockups(users []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Drips.Contract.UnlockLockups(&_Drips.TransactOpts, users, amounts)
}

// UpdateLockUps is a paid mutator transaction binding the contract method 0xc64db12e.
//
// Solidity: function updateLockUps(address owner) returns()
func (_Drips *DripsTransactor) UpdateLockUps(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "updateLockUps", owner)
}

// UpdateLockUps is a paid mutator transaction binding the contract method 0xc64db12e.
//
// Solidity: function updateLockUps(address owner) returns()
func (_Drips *DripsSession) UpdateLockUps(owner common.Address) (*types.Transaction, error) {
	return _Drips.Contract.UpdateLockUps(&_Drips.TransactOpts, owner)
}

// UpdateLockUps is a paid mutator transaction binding the contract method 0xc64db12e.
//
// Solidity: function updateLockUps(address owner) returns()
func (_Drips *DripsTransactorSession) UpdateLockUps(owner common.Address) (*types.Transaction, error) {
	return _Drips.Contract.UpdateLockUps(&_Drips.TransactOpts, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_Drips *DripsTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Drips.contract.Transact(opts, "withdraw", assets, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_Drips *DripsSession) Withdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Drips.Contract.Withdraw(&_Drips.TransactOpts, assets, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 assets, address receiver) returns(uint256)
func (_Drips *DripsTransactorSession) Withdraw(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Drips.Contract.Withdraw(&_Drips.TransactOpts, assets, receiver)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Drips *DripsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Drips.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Drips *DripsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Drips.Contract.Fallback(&_Drips.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Drips *DripsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Drips.Contract.Fallback(&_Drips.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Drips *DripsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Drips.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Drips *DripsSession) Receive() (*types.Transaction, error) {
	return _Drips.Contract.Receive(&_Drips.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Drips *DripsTransactorSession) Receive() (*types.Transaction, error) {
	return _Drips.Contract.Receive(&_Drips.TransactOpts)
}

// DripsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Drips contract.
type DripsDepositIterator struct {
	Event *DripsDeposit // Event containing the contract specifics and raw log

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
func (it *DripsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsDeposit)
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
		it.Event = new(DripsDeposit)
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
func (it *DripsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsDeposit represents a Deposit event raised by the Drips contract.
type DripsDeposit struct {
	Sender common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 assets)
func (_Drips *DripsFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*DripsDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &DripsDepositIterator{contract: _Drips.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 assets)
func (_Drips *DripsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DripsDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsDeposit)
				if err := _Drips.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 assets)
func (_Drips *DripsFilterer) ParseDeposit(log types.Log) (*DripsDeposit, error) {
	event := new(DripsDeposit)
	if err := _Drips.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsDepositWithDurationIterator is returned from FilterDepositWithDuration and is used to iterate over the raw logs and unpacked data for DepositWithDuration events raised by the Drips contract.
type DripsDepositWithDurationIterator struct {
	Event *DripsDepositWithDuration // Event containing the contract specifics and raw log

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
func (it *DripsDepositWithDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsDepositWithDuration)
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
		it.Event = new(DripsDepositWithDuration)
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
func (it *DripsDepositWithDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsDepositWithDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsDepositWithDuration represents a DepositWithDuration event raised by the Drips contract.
type DripsDepositWithDuration struct {
	Owner     common.Address
	LockStart *big.Int
	Amount    *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositWithDuration is a free log retrieval operation binding the contract event 0x71a810dd86885cd9fea6e9b97fe8287ec71dd6475db601e9b925679c9f3ec3d1.
//
// Solidity: event DepositWithDuration(address indexed owner, uint256 lockStart, uint256 amount, uint256 duration)
func (_Drips *DripsFilterer) FilterDepositWithDuration(opts *bind.FilterOpts, owner []common.Address) (*DripsDepositWithDurationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "DepositWithDuration", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DripsDepositWithDurationIterator{contract: _Drips.contract, event: "DepositWithDuration", logs: logs, sub: sub}, nil
}

// WatchDepositWithDuration is a free log subscription operation binding the contract event 0x71a810dd86885cd9fea6e9b97fe8287ec71dd6475db601e9b925679c9f3ec3d1.
//
// Solidity: event DepositWithDuration(address indexed owner, uint256 lockStart, uint256 amount, uint256 duration)
func (_Drips *DripsFilterer) WatchDepositWithDuration(opts *bind.WatchOpts, sink chan<- *DripsDepositWithDuration, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "DepositWithDuration", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsDepositWithDuration)
				if err := _Drips.contract.UnpackLog(event, "DepositWithDuration", log); err != nil {
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

// ParseDepositWithDuration is a log parse operation binding the contract event 0x71a810dd86885cd9fea6e9b97fe8287ec71dd6475db601e9b925679c9f3ec3d1.
//
// Solidity: event DepositWithDuration(address indexed owner, uint256 lockStart, uint256 amount, uint256 duration)
func (_Drips *DripsFilterer) ParseDepositWithDuration(log types.Log) (*DripsDepositWithDuration, error) {
	event := new(DripsDepositWithDuration)
	if err := _Drips.contract.UnpackLog(event, "DepositWithDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Drips contract.
type DripsInitializedIterator struct {
	Event *DripsInitialized // Event containing the contract specifics and raw log

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
func (it *DripsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsInitialized)
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
		it.Event = new(DripsInitialized)
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
func (it *DripsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsInitialized represents a Initialized event raised by the Drips contract.
type DripsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Drips *DripsFilterer) FilterInitialized(opts *bind.FilterOpts) (*DripsInitializedIterator, error) {

	logs, sub, err := _Drips.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DripsInitializedIterator{contract: _Drips.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Drips *DripsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DripsInitialized) (event.Subscription, error) {

	logs, sub, err := _Drips.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsInitialized)
				if err := _Drips.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Drips *DripsFilterer) ParseInitialized(log types.Log) (*DripsInitialized, error) {
	event := new(DripsInitialized)
	if err := _Drips.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsProtocolConfigChangedIterator is returned from FilterProtocolConfigChanged and is used to iterate over the raw logs and unpacked data for ProtocolConfigChanged events raised by the Drips contract.
type DripsProtocolConfigChangedIterator struct {
	Event *DripsProtocolConfigChanged // Event containing the contract specifics and raw log

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
func (it *DripsProtocolConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsProtocolConfigChanged)
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
		it.Event = new(DripsProtocolConfigChanged)
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
func (it *DripsProtocolConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsProtocolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsProtocolConfigChanged represents a ProtocolConfigChanged event raised by the Drips contract.
type DripsProtocolConfigChanged struct {
	SetterSelector  [4]byte
	SetterSignature string
	Value           []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProtocolConfigChanged is a free log retrieval operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_Drips *DripsFilterer) FilterProtocolConfigChanged(opts *bind.FilterOpts, setterSelector [][4]byte) (*DripsProtocolConfigChangedIterator, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return &DripsProtocolConfigChangedIterator{contract: _Drips.contract, event: "ProtocolConfigChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolConfigChanged is a free log subscription operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_Drips *DripsFilterer) WatchProtocolConfigChanged(opts *bind.WatchOpts, sink chan<- *DripsProtocolConfigChanged, setterSelector [][4]byte) (event.Subscription, error) {

	var setterSelectorRule []interface{}
	for _, setterSelectorItem := range setterSelector {
		setterSelectorRule = append(setterSelectorRule, setterSelectorItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "ProtocolConfigChanged", setterSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsProtocolConfigChanged)
				if err := _Drips.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
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

// ParseProtocolConfigChanged is a log parse operation binding the contract event 0x01d854e8dde9402801a4c6f2840193465752abfad61e0bb7c4258d526ae42e74.
//
// Solidity: event ProtocolConfigChanged(bytes4 indexed setterSelector, string setterSignature, bytes value)
func (_Drips *DripsFilterer) ParseProtocolConfigChanged(log types.Log) (*DripsProtocolConfigChanged, error) {
	event := new(DripsProtocolConfigChanged)
	if err := _Drips.contract.UnpackLog(event, "ProtocolConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Drips contract.
type DripsRoleAdminChangedIterator struct {
	Event *DripsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DripsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsRoleAdminChanged)
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
		it.Event = new(DripsRoleAdminChanged)
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
func (it *DripsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsRoleAdminChanged represents a RoleAdminChanged event raised by the Drips contract.
type DripsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Drips *DripsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DripsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DripsRoleAdminChangedIterator{contract: _Drips.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Drips *DripsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DripsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsRoleAdminChanged)
				if err := _Drips.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Drips *DripsFilterer) ParseRoleAdminChanged(log types.Log) (*DripsRoleAdminChanged, error) {
	event := new(DripsRoleAdminChanged)
	if err := _Drips.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Drips contract.
type DripsRoleGrantedIterator struct {
	Event *DripsRoleGranted // Event containing the contract specifics and raw log

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
func (it *DripsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsRoleGranted)
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
		it.Event = new(DripsRoleGranted)
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
func (it *DripsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsRoleGranted represents a RoleGranted event raised by the Drips contract.
type DripsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Drips *DripsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DripsRoleGrantedIterator, error) {

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

	logs, sub, err := _Drips.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DripsRoleGrantedIterator{contract: _Drips.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Drips *DripsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DripsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Drips.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsRoleGranted)
				if err := _Drips.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Drips *DripsFilterer) ParseRoleGranted(log types.Log) (*DripsRoleGranted, error) {
	event := new(DripsRoleGranted)
	if err := _Drips.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Drips contract.
type DripsRoleRevokedIterator struct {
	Event *DripsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DripsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsRoleRevoked)
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
		it.Event = new(DripsRoleRevoked)
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
func (it *DripsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsRoleRevoked represents a RoleRevoked event raised by the Drips contract.
type DripsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Drips *DripsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DripsRoleRevokedIterator, error) {

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

	logs, sub, err := _Drips.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DripsRoleRevokedIterator{contract: _Drips.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Drips *DripsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DripsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Drips.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsRoleRevoked)
				if err := _Drips.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Drips *DripsFilterer) ParseRoleRevoked(log types.Log) (*DripsRoleRevoked, error) {
	event := new(DripsRoleRevoked)
	if err := _Drips.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsUserLockUpdatedIterator is returned from FilterUserLockUpdated and is used to iterate over the raw logs and unpacked data for UserLockUpdated events raised by the Drips contract.
type DripsUserLockUpdatedIterator struct {
	Event *DripsUserLockUpdated // Event containing the contract specifics and raw log

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
func (it *DripsUserLockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsUserLockUpdated)
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
		it.Event = new(DripsUserLockUpdated)
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
func (it *DripsUserLockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsUserLockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsUserLockUpdated represents a UserLockUpdated event raised by the Drips contract.
type DripsUserLockUpdated struct {
	Owner        common.Address
	LockedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserLockUpdated is a free log retrieval operation binding the contract event 0x9e819531b9f4f6660e839d7b1ddb93441f1a62e8248416d7807881afecfa2498.
//
// Solidity: event UserLockUpdated(address indexed owner, uint256 lockedAmount)
func (_Drips *DripsFilterer) FilterUserLockUpdated(opts *bind.FilterOpts, owner []common.Address) (*DripsUserLockUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "UserLockUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DripsUserLockUpdatedIterator{contract: _Drips.contract, event: "UserLockUpdated", logs: logs, sub: sub}, nil
}

// WatchUserLockUpdated is a free log subscription operation binding the contract event 0x9e819531b9f4f6660e839d7b1ddb93441f1a62e8248416d7807881afecfa2498.
//
// Solidity: event UserLockUpdated(address indexed owner, uint256 lockedAmount)
func (_Drips *DripsFilterer) WatchUserLockUpdated(opts *bind.WatchOpts, sink chan<- *DripsUserLockUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "UserLockUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsUserLockUpdated)
				if err := _Drips.contract.UnpackLog(event, "UserLockUpdated", log); err != nil {
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

// ParseUserLockUpdated is a log parse operation binding the contract event 0x9e819531b9f4f6660e839d7b1ddb93441f1a62e8248416d7807881afecfa2498.
//
// Solidity: event UserLockUpdated(address indexed owner, uint256 lockedAmount)
func (_Drips *DripsFilterer) ParseUserLockUpdated(log types.Log) (*DripsUserLockUpdated, error) {
	event := new(DripsUserLockUpdated)
	if err := _Drips.contract.UnpackLog(event, "UserLockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsUserUnlockedIterator is returned from FilterUserUnlocked and is used to iterate over the raw logs and unpacked data for UserUnlocked events raised by the Drips contract.
type DripsUserUnlockedIterator struct {
	Event *DripsUserUnlocked // Event containing the contract specifics and raw log

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
func (it *DripsUserUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsUserUnlocked)
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
		it.Event = new(DripsUserUnlocked)
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
func (it *DripsUserUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsUserUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsUserUnlocked represents a UserUnlocked event raised by the Drips contract.
type DripsUserUnlocked struct {
	Owner         common.Address
	RequestUnlock *big.Int
	UnlockAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserUnlocked is a free log retrieval operation binding the contract event 0x8334525b087bd40a461fad9f34d3dd47831c2639f0ff68c8013c0e88c08f3117.
//
// Solidity: event UserUnlocked(address indexed owner, uint256 requestUnlock, uint256 unlockAmount)
func (_Drips *DripsFilterer) FilterUserUnlocked(opts *bind.FilterOpts, owner []common.Address) (*DripsUserUnlockedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "UserUnlocked", ownerRule)
	if err != nil {
		return nil, err
	}
	return &DripsUserUnlockedIterator{contract: _Drips.contract, event: "UserUnlocked", logs: logs, sub: sub}, nil
}

// WatchUserUnlocked is a free log subscription operation binding the contract event 0x8334525b087bd40a461fad9f34d3dd47831c2639f0ff68c8013c0e88c08f3117.
//
// Solidity: event UserUnlocked(address indexed owner, uint256 requestUnlock, uint256 unlockAmount)
func (_Drips *DripsFilterer) WatchUserUnlocked(opts *bind.WatchOpts, sink chan<- *DripsUserUnlocked, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "UserUnlocked", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsUserUnlocked)
				if err := _Drips.contract.UnpackLog(event, "UserUnlocked", log); err != nil {
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

// ParseUserUnlocked is a log parse operation binding the contract event 0x8334525b087bd40a461fad9f34d3dd47831c2639f0ff68c8013c0e88c08f3117.
//
// Solidity: event UserUnlocked(address indexed owner, uint256 requestUnlock, uint256 unlockAmount)
func (_Drips *DripsFilterer) ParseUserUnlocked(log types.Log) (*DripsUserUnlocked, error) {
	event := new(DripsUserUnlocked)
	if err := _Drips.contract.UnpackLog(event, "UserUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DripsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Drips contract.
type DripsWithdrawIterator struct {
	Event *DripsWithdraw // Event containing the contract specifics and raw log

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
func (it *DripsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DripsWithdraw)
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
		it.Event = new(DripsWithdraw)
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
func (it *DripsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DripsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DripsWithdraw represents a Withdraw event raised by the Drips contract.
type DripsWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Assets   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, uint256 assets)
func (_Drips *DripsFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*DripsWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Drips.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &DripsWithdrawIterator{contract: _Drips.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, uint256 assets)
func (_Drips *DripsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *DripsWithdraw, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Drips.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DripsWithdraw)
				if err := _Drips.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, uint256 assets)
func (_Drips *DripsFilterer) ParseWithdraw(log types.Log) (*DripsWithdraw, error) {
	event := new(DripsWithdraw)
	if err := _Drips.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
