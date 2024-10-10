// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancer_token

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

// IPoolSwapStructsSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IPoolSwapStructsSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// WeightedPoolNewPoolParams is an auto generated low-level Go binding around an user-defined struct.
type WeightedPoolNewPoolParams struct {
	Name              string
	Symbol            string
	Tokens            []common.Address
	NormalizedWeights []*big.Int
	RateProviders     []common.Address
	AssetManagers     []common.Address
	SwapFeePercentage *big.Int
}

// BalancerTokenMetaData contains all meta data concerning the BalancerToken contract.
var BalancerTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"normalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIRateProvider[]\",\"name\":\"rateProviders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structWeightedPool.NewPoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"protocolFeeProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeePercentageCacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getATHRateProduct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActualSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastPostJoinExitInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNormalizedWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProtocolFeePercentageCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolSwapFeeDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateProviders\",\"outputs\":[{\"internalType\":\"contractIRateProvider[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"poolConfig\",\"type\":\"bytes\"}],\"name\":\"setAssetManagerPoolConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateProtocolFeePercentageCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerTokenMetaData.ABI instead.
var BalancerTokenABI = BalancerTokenMetaData.ABI

// BalancerToken is an auto generated Go binding around an Ethereum contract.
type BalancerToken struct {
	BalancerTokenCaller     // Read-only binding to the contract
	BalancerTokenTransactor // Write-only binding to the contract
	BalancerTokenFilterer   // Log filterer for contract events
}

// BalancerTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerTokenSession struct {
	Contract     *BalancerToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerTokenCallerSession struct {
	Contract *BalancerTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BalancerTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerTokenTransactorSession struct {
	Contract     *BalancerTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BalancerTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerTokenRaw struct {
	Contract *BalancerToken // Generic contract binding to access the raw methods on
}

// BalancerTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerTokenCallerRaw struct {
	Contract *BalancerTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerTokenTransactorRaw struct {
	Contract *BalancerTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerToken creates a new instance of BalancerToken, bound to a specific deployed contract.
func NewBalancerToken(address common.Address, backend bind.ContractBackend) (*BalancerToken, error) {
	contract, err := bindBalancerToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerToken{BalancerTokenCaller: BalancerTokenCaller{contract: contract}, BalancerTokenTransactor: BalancerTokenTransactor{contract: contract}, BalancerTokenFilterer: BalancerTokenFilterer{contract: contract}}, nil
}

// NewBalancerTokenCaller creates a new read-only instance of BalancerToken, bound to a specific deployed contract.
func NewBalancerTokenCaller(address common.Address, caller bind.ContractCaller) (*BalancerTokenCaller, error) {
	contract, err := bindBalancerToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenCaller{contract: contract}, nil
}

// NewBalancerTokenTransactor creates a new write-only instance of BalancerToken, bound to a specific deployed contract.
func NewBalancerTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerTokenTransactor, error) {
	contract, err := bindBalancerToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenTransactor{contract: contract}, nil
}

// NewBalancerTokenFilterer creates a new log filterer instance of BalancerToken, bound to a specific deployed contract.
func NewBalancerTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerTokenFilterer, error) {
	contract, err := bindBalancerToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenFilterer{contract: contract}, nil
}

// bindBalancerToken binds a generic wrapper to an already deployed contract.
func bindBalancerToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalancerTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerToken *BalancerTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerToken.Contract.BalancerTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerToken *BalancerTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.Contract.BalancerTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerToken *BalancerTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerToken.Contract.BalancerTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerToken *BalancerTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerToken *BalancerTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerToken *BalancerTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerToken.Contract.contract.Transact(opts, method, params...)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) DELEGATEPROTOCOLSWAPFEESSENTINEL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerToken.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerToken.CallOpts)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerToken.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerToken *BalancerTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerToken *BalancerTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerToken.Contract.DOMAINSEPARATOR(&_BalancerToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerToken *BalancerTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerToken.Contract.DOMAINSEPARATOR(&_BalancerToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerToken *BalancerTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.Allowance(&_BalancerToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.Allowance(&_BalancerToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.BalanceOf(&_BalancerToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.BalanceOf(&_BalancerToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerToken *BalancerTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerToken *BalancerTokenSession) Decimals() (uint8, error) {
	return _BalancerToken.Contract.Decimals(&_BalancerToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerToken *BalancerTokenCallerSession) Decimals() (uint8, error) {
	return _BalancerToken.Contract.Decimals(&_BalancerToken.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetATHRateProduct(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getATHRateProduct")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetATHRateProduct() (*big.Int, error) {
	return _BalancerToken.Contract.GetATHRateProduct(&_BalancerToken.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetATHRateProduct() (*big.Int, error) {
	return _BalancerToken.Contract.GetATHRateProduct(&_BalancerToken.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerToken *BalancerTokenCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerToken *BalancerTokenSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerToken.Contract.GetActionId(&_BalancerToken.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerToken *BalancerTokenCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerToken.Contract.GetActionId(&_BalancerToken.CallOpts, selector)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetActualSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getActualSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetActualSupply() (*big.Int, error) {
	return _BalancerToken.Contract.GetActualSupply(&_BalancerToken.CallOpts)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetActualSupply() (*big.Int, error) {
	return _BalancerToken.Contract.GetActualSupply(&_BalancerToken.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerToken *BalancerTokenCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerToken *BalancerTokenSession) GetAuthorizer() (common.Address, error) {
	return _BalancerToken.Contract.GetAuthorizer(&_BalancerToken.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerToken *BalancerTokenCallerSession) GetAuthorizer() (common.Address, error) {
	return _BalancerToken.Contract.GetAuthorizer(&_BalancerToken.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerToken *BalancerTokenCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerToken *BalancerTokenSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerToken.Contract.GetDomainSeparator(&_BalancerToken.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerToken *BalancerTokenCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerToken.Contract.GetDomainSeparator(&_BalancerToken.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetInvariant() (*big.Int, error) {
	return _BalancerToken.Contract.GetInvariant(&_BalancerToken.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetInvariant() (*big.Int, error) {
	return _BalancerToken.Contract.GetInvariant(&_BalancerToken.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetLastPostJoinExitInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getLastPostJoinExitInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BalancerToken.Contract.GetLastPostJoinExitInvariant(&_BalancerToken.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BalancerToken.Contract.GetLastPostJoinExitInvariant(&_BalancerToken.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.GetNextNonce(&_BalancerToken.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.GetNextNonce(&_BalancerToken.CallOpts, account)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerToken *BalancerTokenCaller) GetNormalizedWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getNormalizedWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerToken *BalancerTokenSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerToken.Contract.GetNormalizedWeights(&_BalancerToken.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerToken *BalancerTokenCallerSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerToken.Contract.GetNormalizedWeights(&_BalancerToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerToken *BalancerTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerToken *BalancerTokenSession) GetOwner() (common.Address, error) {
	return _BalancerToken.Contract.GetOwner(&_BalancerToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerToken *BalancerTokenCallerSession) GetOwner() (common.Address, error) {
	return _BalancerToken.Contract.GetOwner(&_BalancerToken.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerToken *BalancerTokenCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerToken *BalancerTokenSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerToken.Contract.GetPausedState(&_BalancerToken.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerToken *BalancerTokenCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerToken.Contract.GetPausedState(&_BalancerToken.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerToken *BalancerTokenCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerToken *BalancerTokenSession) GetPoolId() ([32]byte, error) {
	return _BalancerToken.Contract.GetPoolId(&_BalancerToken.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerToken *BalancerTokenCallerSession) GetPoolId() ([32]byte, error) {
	return _BalancerToken.Contract.GetPoolId(&_BalancerToken.CallOpts)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetProtocolFeePercentageCache(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getProtocolFeePercentageCache", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerToken.Contract.GetProtocolFeePercentageCache(&_BalancerToken.CallOpts, feeType)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerToken.Contract.GetProtocolFeePercentageCache(&_BalancerToken.CallOpts, feeType)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerToken *BalancerTokenCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerToken *BalancerTokenSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerToken.Contract.GetProtocolFeesCollector(&_BalancerToken.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerToken *BalancerTokenCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerToken.Contract.GetProtocolFeesCollector(&_BalancerToken.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerToken *BalancerTokenCaller) GetProtocolSwapFeeDelegation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getProtocolSwapFeeDelegation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerToken *BalancerTokenSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerToken.Contract.GetProtocolSwapFeeDelegation(&_BalancerToken.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerToken *BalancerTokenCallerSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerToken.Contract.GetProtocolSwapFeeDelegation(&_BalancerToken.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerToken *BalancerTokenCaller) GetRateProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getRateProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerToken *BalancerTokenSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerToken.Contract.GetRateProviders(&_BalancerToken.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerToken *BalancerTokenCallerSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerToken.Contract.GetRateProviders(&_BalancerToken.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerToken *BalancerTokenCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerToken *BalancerTokenSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerToken.Contract.GetScalingFactors(&_BalancerToken.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerToken *BalancerTokenCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerToken.Contract.GetScalingFactors(&_BalancerToken.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerToken.Contract.GetSwapFeePercentage(&_BalancerToken.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerToken.Contract.GetSwapFeePercentage(&_BalancerToken.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerToken *BalancerTokenCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerToken *BalancerTokenSession) GetVault() (common.Address, error) {
	return _BalancerToken.Contract.GetVault(&_BalancerToken.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerToken *BalancerTokenCallerSession) GetVault() (common.Address, error) {
	return _BalancerToken.Contract.GetVault(&_BalancerToken.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerToken *BalancerTokenCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerToken *BalancerTokenSession) InRecoveryMode() (bool, error) {
	return _BalancerToken.Contract.InRecoveryMode(&_BalancerToken.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerToken *BalancerTokenCallerSession) InRecoveryMode() (bool, error) {
	return _BalancerToken.Contract.InRecoveryMode(&_BalancerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerToken *BalancerTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerToken *BalancerTokenSession) Name() (string, error) {
	return _BalancerToken.Contract.Name(&_BalancerToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerToken *BalancerTokenCallerSession) Name() (string, error) {
	return _BalancerToken.Contract.Name(&_BalancerToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerToken *BalancerTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.Nonces(&_BalancerToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerToken.Contract.Nonces(&_BalancerToken.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerToken *BalancerTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerToken *BalancerTokenSession) Symbol() (string, error) {
	return _BalancerToken.Contract.Symbol(&_BalancerToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerToken *BalancerTokenCallerSession) Symbol() (string, error) {
	return _BalancerToken.Contract.Symbol(&_BalancerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenSession) TotalSupply() (*big.Int, error) {
	return _BalancerToken.Contract.TotalSupply(&_BalancerToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerToken *BalancerTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerToken.Contract.TotalSupply(&_BalancerToken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerToken *BalancerTokenCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerToken.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerToken *BalancerTokenSession) Version() (string, error) {
	return _BalancerToken.Contract.Version(&_BalancerToken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerToken *BalancerTokenCallerSession) Version() (string, error) {
	return _BalancerToken.Contract.Version(&_BalancerToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.Approve(&_BalancerToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.Approve(&_BalancerToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.DecreaseAllowance(&_BalancerToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.DecreaseAllowance(&_BalancerToken.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerToken.Contract.DisableRecoveryMode(&_BalancerToken.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerToken.Contract.DisableRecoveryMode(&_BalancerToken.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerToken.Contract.EnableRecoveryMode(&_BalancerToken.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerToken *BalancerTokenTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerToken.Contract.EnableRecoveryMode(&_BalancerToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerToken *BalancerTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerToken *BalancerTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.IncreaseAllowance(&_BalancerToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerToken *BalancerTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.IncreaseAllowance(&_BalancerToken.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnExitPool(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnExitPool(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnJoinPool(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerToken *BalancerTokenTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnJoinPool(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerToken *BalancerTokenTransactor) OnSwap(opts *bind.TransactOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "onSwap", request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerToken *BalancerTokenSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnSwap(&_BalancerToken.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerToken *BalancerTokenTransactorSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.OnSwap(&_BalancerToken.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerToken *BalancerTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerToken *BalancerTokenSession) Pause() (*types.Transaction, error) {
	return _BalancerToken.Contract.Pause(&_BalancerToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerToken *BalancerTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _BalancerToken.Contract.Pause(&_BalancerToken.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerToken *BalancerTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerToken *BalancerTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.Permit(&_BalancerToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerToken *BalancerTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.Permit(&_BalancerToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerToken *BalancerTokenTransactor) QueryExit(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerToken *BalancerTokenSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.QueryExit(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerToken *BalancerTokenTransactorSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.QueryExit(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerToken *BalancerTokenTransactor) QueryJoin(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerToken *BalancerTokenSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.QueryJoin(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerToken *BalancerTokenTransactorSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.QueryJoin(&_BalancerToken.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerToken *BalancerTokenTransactor) SetAssetManagerPoolConfig(opts *bind.TransactOpts, token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "setAssetManagerPoolConfig", token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerToken *BalancerTokenSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.SetAssetManagerPoolConfig(&_BalancerToken.TransactOpts, token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerToken *BalancerTokenTransactorSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerToken.Contract.SetAssetManagerPoolConfig(&_BalancerToken.TransactOpts, token, poolConfig)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerToken *BalancerTokenTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerToken *BalancerTokenSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.SetSwapFeePercentage(&_BalancerToken.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerToken *BalancerTokenTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.SetSwapFeePercentage(&_BalancerToken.TransactOpts, swapFeePercentage)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.Transfer(&_BalancerToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.Transfer(&_BalancerToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.TransferFrom(&_BalancerToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerToken *BalancerTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerToken.Contract.TransferFrom(&_BalancerToken.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerToken *BalancerTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerToken *BalancerTokenSession) Unpause() (*types.Transaction, error) {
	return _BalancerToken.Contract.Unpause(&_BalancerToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerToken *BalancerTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _BalancerToken.Contract.Unpause(&_BalancerToken.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerToken *BalancerTokenTransactor) UpdateProtocolFeePercentageCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerToken.contract.Transact(opts, "updateProtocolFeePercentageCache")
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerToken *BalancerTokenSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerToken.Contract.UpdateProtocolFeePercentageCache(&_BalancerToken.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerToken *BalancerTokenTransactorSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerToken.Contract.UpdateProtocolFeePercentageCache(&_BalancerToken.TransactOpts)
}

// BalancerTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerToken contract.
type BalancerTokenApprovalIterator struct {
	Event *BalancerTokenApproval // Event containing the contract specifics and raw log

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
func (it *BalancerTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenApproval)
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
		it.Event = new(BalancerTokenApproval)
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
func (it *BalancerTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenApproval represents a Approval event raised by the BalancerToken contract.
type BalancerTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerToken *BalancerTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BalancerTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenApprovalIterator{contract: _BalancerToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerToken *BalancerTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenApproval)
				if err := _BalancerToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerToken *BalancerTokenFilterer) ParseApproval(log types.Log) (*BalancerTokenApproval, error) {
	event := new(BalancerTokenApproval)
	if err := _BalancerToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerTokenPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BalancerToken contract.
type BalancerTokenPausedStateChangedIterator struct {
	Event *BalancerTokenPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *BalancerTokenPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenPausedStateChanged)
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
		it.Event = new(BalancerTokenPausedStateChanged)
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
func (it *BalancerTokenPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenPausedStateChanged represents a PausedStateChanged event raised by the BalancerToken contract.
type BalancerTokenPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerToken *BalancerTokenFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BalancerTokenPausedStateChangedIterator, error) {

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerTokenPausedStateChangedIterator{contract: _BalancerToken.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerToken *BalancerTokenFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerTokenPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenPausedStateChanged)
				if err := _BalancerToken.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerToken *BalancerTokenFilterer) ParsePausedStateChanged(log types.Log) (*BalancerTokenPausedStateChanged, error) {
	event := new(BalancerTokenPausedStateChanged)
	if err := _BalancerToken.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerTokenProtocolFeePercentageCacheUpdatedIterator is returned from FilterProtocolFeePercentageCacheUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeePercentageCacheUpdated events raised by the BalancerToken contract.
type BalancerTokenProtocolFeePercentageCacheUpdatedIterator struct {
	Event *BalancerTokenProtocolFeePercentageCacheUpdated // Event containing the contract specifics and raw log

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
func (it *BalancerTokenProtocolFeePercentageCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenProtocolFeePercentageCacheUpdated)
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
		it.Event = new(BalancerTokenProtocolFeePercentageCacheUpdated)
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
func (it *BalancerTokenProtocolFeePercentageCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenProtocolFeePercentageCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenProtocolFeePercentageCacheUpdated represents a ProtocolFeePercentageCacheUpdated event raised by the BalancerToken contract.
type BalancerTokenProtocolFeePercentageCacheUpdated struct {
	FeeType               *big.Int
	ProtocolFeePercentage *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeePercentageCacheUpdated is a free log retrieval operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) FilterProtocolFeePercentageCacheUpdated(opts *bind.FilterOpts, feeType []*big.Int) (*BalancerTokenProtocolFeePercentageCacheUpdatedIterator, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenProtocolFeePercentageCacheUpdatedIterator{contract: _BalancerToken.contract, event: "ProtocolFeePercentageCacheUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeePercentageCacheUpdated is a free log subscription operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) WatchProtocolFeePercentageCacheUpdated(opts *bind.WatchOpts, sink chan<- *BalancerTokenProtocolFeePercentageCacheUpdated, feeType []*big.Int) (event.Subscription, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenProtocolFeePercentageCacheUpdated)
				if err := _BalancerToken.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
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

// ParseProtocolFeePercentageCacheUpdated is a log parse operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) ParseProtocolFeePercentageCacheUpdated(log types.Log) (*BalancerTokenProtocolFeePercentageCacheUpdated, error) {
	event := new(BalancerTokenProtocolFeePercentageCacheUpdated)
	if err := _BalancerToken.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerTokenRecoveryModeStateChangedIterator is returned from FilterRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for RecoveryModeStateChanged events raised by the BalancerToken contract.
type BalancerTokenRecoveryModeStateChangedIterator struct {
	Event *BalancerTokenRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *BalancerTokenRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenRecoveryModeStateChanged)
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
		it.Event = new(BalancerTokenRecoveryModeStateChanged)
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
func (it *BalancerTokenRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenRecoveryModeStateChanged represents a RecoveryModeStateChanged event raised by the BalancerToken contract.
type BalancerTokenRecoveryModeStateChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerToken *BalancerTokenFilterer) FilterRecoveryModeStateChanged(opts *bind.FilterOpts) (*BalancerTokenRecoveryModeStateChangedIterator, error) {

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerTokenRecoveryModeStateChangedIterator{contract: _BalancerToken.contract, event: "RecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerToken *BalancerTokenFilterer) WatchRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerTokenRecoveryModeStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenRecoveryModeStateChanged)
				if err := _BalancerToken.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
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

// ParseRecoveryModeStateChanged is a log parse operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerToken *BalancerTokenFilterer) ParseRecoveryModeStateChanged(log types.Log) (*BalancerTokenRecoveryModeStateChanged, error) {
	event := new(BalancerTokenRecoveryModeStateChanged)
	if err := _BalancerToken.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerTokenSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BalancerToken contract.
type BalancerTokenSwapFeePercentageChangedIterator struct {
	Event *BalancerTokenSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BalancerTokenSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenSwapFeePercentageChanged)
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
		it.Event = new(BalancerTokenSwapFeePercentageChanged)
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
func (it *BalancerTokenSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BalancerToken contract.
type BalancerTokenSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BalancerTokenSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerTokenSwapFeePercentageChangedIterator{contract: _BalancerToken.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BalancerTokenSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenSwapFeePercentageChanged)
				if err := _BalancerToken.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerToken *BalancerTokenFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BalancerTokenSwapFeePercentageChanged, error) {
	event := new(BalancerTokenSwapFeePercentageChanged)
	if err := _BalancerToken.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerToken contract.
type BalancerTokenTransferIterator struct {
	Event *BalancerTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BalancerTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerTokenTransfer)
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
		it.Event = new(BalancerTokenTransfer)
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
func (it *BalancerTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerTokenTransfer represents a Transfer event raised by the BalancerToken contract.
type BalancerTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerToken *BalancerTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BalancerTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BalancerTokenTransferIterator{contract: _BalancerToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerToken *BalancerTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerTokenTransfer)
				if err := _BalancerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BalancerToken *BalancerTokenFilterer) ParseTransfer(log types.Log) (*BalancerTokenTransfer, error) {
	event := new(BalancerTokenTransfer)
	if err := _BalancerToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
