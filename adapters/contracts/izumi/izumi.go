// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package izumi

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

// LiquidityManagerAddLiquidityParam is an auto generated low-level Go binding around an user-defined struct.
type LiquidityManagerAddLiquidityParam struct {
	Lid        *big.Int
	XLim       *big.Int
	YLim       *big.Int
	AmountXMin *big.Int
	AmountYMin *big.Int
	Deadline   *big.Int
}

// LiquidityManagerMintParam is an auto generated low-level Go binding around an user-defined struct.
type LiquidityManagerMintParam struct {
	Miner      common.Address
	TokenX     common.Address
	TokenY     common.Address
	Fee        *big.Int
	Pl         *big.Int
	Pr         *big.Int
	XLim       *big.Int
	YLim       *big.Int
	AmountXMin *big.Int
	AmountYMin *big.Int
	Deadline   *big.Int
}

// IzumiMetaData contains all meta data concerning the Izumi contract.
var IzumiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDelta\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityDelta\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"name\":\"DecLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lid\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"xLim\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"yLim\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountXMin\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountYMin\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityManager.AddLiquidityParam\",\"name\":\"addLiquidityParam\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidityDelta\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lid\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lid\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"amountXLim\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountYLim\",\"type\":\"uint128\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"initialPoint\",\"type\":\"int24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lid\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidDelta\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amountXMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountYMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"decLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"liquidities\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"leftPt\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"rightPt\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainTokenX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainTokenY\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"poolId\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"pl\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"pr\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"xLim\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"yLim\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountXMin\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountYMin\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLiquidityManager.MintParam\",\"name\":\"mintParam\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lid\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintDepositCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolIds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"poolMetas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IzumiABI is the input ABI used to generate the binding from.
// Deprecated: Use IzumiMetaData.ABI instead.
var IzumiABI = IzumiMetaData.ABI

// Izumi is an auto generated Go binding around an Ethereum contract.
type Izumi struct {
	IzumiCaller     // Read-only binding to the contract
	IzumiTransactor // Write-only binding to the contract
	IzumiFilterer   // Log filterer for contract events
}

// IzumiCaller is an auto generated read-only Go binding around an Ethereum contract.
type IzumiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IzumiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IzumiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IzumiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IzumiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IzumiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IzumiSession struct {
	Contract     *Izumi            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IzumiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IzumiCallerSession struct {
	Contract *IzumiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IzumiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IzumiTransactorSession struct {
	Contract     *IzumiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IzumiRaw is an auto generated low-level Go binding around an Ethereum contract.
type IzumiRaw struct {
	Contract *Izumi // Generic contract binding to access the raw methods on
}

// IzumiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IzumiCallerRaw struct {
	Contract *IzumiCaller // Generic read-only contract binding to access the raw methods on
}

// IzumiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IzumiTransactorRaw struct {
	Contract *IzumiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIzumi creates a new instance of Izumi, bound to a specific deployed contract.
func NewIzumi(address common.Address, backend bind.ContractBackend) (*Izumi, error) {
	contract, err := bindIzumi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Izumi{IzumiCaller: IzumiCaller{contract: contract}, IzumiTransactor: IzumiTransactor{contract: contract}, IzumiFilterer: IzumiFilterer{contract: contract}}, nil
}

// NewIzumiCaller creates a new read-only instance of Izumi, bound to a specific deployed contract.
func NewIzumiCaller(address common.Address, caller bind.ContractCaller) (*IzumiCaller, error) {
	contract, err := bindIzumi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IzumiCaller{contract: contract}, nil
}

// NewIzumiTransactor creates a new write-only instance of Izumi, bound to a specific deployed contract.
func NewIzumiTransactor(address common.Address, transactor bind.ContractTransactor) (*IzumiTransactor, error) {
	contract, err := bindIzumi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IzumiTransactor{contract: contract}, nil
}

// NewIzumiFilterer creates a new log filterer instance of Izumi, bound to a specific deployed contract.
func NewIzumiFilterer(address common.Address, filterer bind.ContractFilterer) (*IzumiFilterer, error) {
	contract, err := bindIzumi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IzumiFilterer{contract: contract}, nil
}

// bindIzumi binds a generic wrapper to an already deployed contract.
func bindIzumi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IzumiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Izumi *IzumiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Izumi.Contract.IzumiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Izumi *IzumiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Izumi.Contract.IzumiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Izumi *IzumiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Izumi.Contract.IzumiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Izumi *IzumiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Izumi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Izumi *IzumiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Izumi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Izumi *IzumiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Izumi.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Izumi *IzumiCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Izumi *IzumiSession) WETH9() (common.Address, error) {
	return _Izumi.Contract.WETH9(&_Izumi.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Izumi *IzumiCallerSession) WETH9() (common.Address, error) {
	return _Izumi.Contract.WETH9(&_Izumi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Izumi *IzumiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Izumi *IzumiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Izumi.Contract.BalanceOf(&_Izumi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Izumi *IzumiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Izumi.Contract.BalanceOf(&_Izumi.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Izumi *IzumiCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Izumi *IzumiSession) BaseURI() (string, error) {
	return _Izumi.Contract.BaseURI(&_Izumi.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Izumi *IzumiCallerSession) BaseURI() (string, error) {
	return _Izumi.Contract.BaseURI(&_Izumi.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Izumi *IzumiCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Izumi *IzumiSession) Factory() (common.Address, error) {
	return _Izumi.Contract.Factory(&_Izumi.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Izumi *IzumiCallerSession) Factory() (common.Address, error) {
	return _Izumi.Contract.Factory(&_Izumi.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Izumi *IzumiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Izumi *IzumiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Izumi.Contract.GetApproved(&_Izumi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Izumi *IzumiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Izumi.Contract.GetApproved(&_Izumi.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Izumi *IzumiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Izumi *IzumiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Izumi.Contract.IsApprovedForAll(&_Izumi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Izumi *IzumiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Izumi.Contract.IsApprovedForAll(&_Izumi.CallOpts, owner, operator)
}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Izumi *IzumiCaller) Liquidities(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "liquidities", arg0)

	outstruct := new(struct {
		LeftPt           *big.Int
		RightPt          *big.Int
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		RemainTokenX     *big.Int
		RemainTokenY     *big.Int
		PoolId           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LeftPt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RightPt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RemainTokenX = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RemainTokenY = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PoolId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Izumi *IzumiSession) Liquidities(arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	return _Izumi.Contract.Liquidities(&_Izumi.CallOpts, arg0)
}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Izumi *IzumiCallerSession) Liquidities(arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	return _Izumi.Contract.Liquidities(&_Izumi.CallOpts, arg0)
}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Izumi *IzumiCaller) LiquidityNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "liquidityNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Izumi *IzumiSession) LiquidityNum() (*big.Int, error) {
	return _Izumi.Contract.LiquidityNum(&_Izumi.CallOpts)
}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Izumi *IzumiCallerSession) LiquidityNum() (*big.Int, error) {
	return _Izumi.Contract.LiquidityNum(&_Izumi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Izumi *IzumiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Izumi *IzumiSession) Name() (string, error) {
	return _Izumi.Contract.Name(&_Izumi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Izumi *IzumiCallerSession) Name() (string, error) {
	return _Izumi.Contract.Name(&_Izumi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Izumi *IzumiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Izumi *IzumiSession) Owner() (common.Address, error) {
	return _Izumi.Contract.Owner(&_Izumi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Izumi *IzumiCallerSession) Owner() (common.Address, error) {
	return _Izumi.Contract.Owner(&_Izumi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Izumi *IzumiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Izumi *IzumiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Izumi.Contract.OwnerOf(&_Izumi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Izumi *IzumiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Izumi.Contract.OwnerOf(&_Izumi.CallOpts, tokenId)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Izumi *IzumiCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Izumi *IzumiSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Izumi.Contract.Pool(&_Izumi.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Izumi *IzumiCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Izumi.Contract.Pool(&_Izumi.CallOpts, tokenX, tokenY, fee)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Izumi *IzumiCaller) PoolIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "poolIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Izumi *IzumiSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _Izumi.Contract.PoolIds(&_Izumi.CallOpts, arg0)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Izumi *IzumiCallerSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _Izumi.Contract.PoolIds(&_Izumi.CallOpts, arg0)
}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Izumi *IzumiCaller) PoolMetas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "poolMetas", arg0)

	outstruct := new(struct {
		TokenX common.Address
		TokenY common.Address
		Fee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenX = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenY = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Izumi *IzumiSession) PoolMetas(arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	return _Izumi.Contract.PoolMetas(&_Izumi.CallOpts, arg0)
}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Izumi *IzumiCallerSession) PoolMetas(arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	return _Izumi.Contract.PoolMetas(&_Izumi.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Izumi *IzumiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Izumi *IzumiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Izumi.Contract.SupportsInterface(&_Izumi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Izumi *IzumiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Izumi.Contract.SupportsInterface(&_Izumi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Izumi *IzumiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Izumi *IzumiSession) Symbol() (string, error) {
	return _Izumi.Contract.Symbol(&_Izumi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Izumi *IzumiCallerSession) Symbol() (string, error) {
	return _Izumi.Contract.Symbol(&_Izumi.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Izumi *IzumiCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Izumi *IzumiSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Izumi.Contract.TokenByIndex(&_Izumi.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Izumi *IzumiCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Izumi.Contract.TokenByIndex(&_Izumi.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Izumi *IzumiCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Izumi *IzumiSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Izumi.Contract.TokenOfOwnerByIndex(&_Izumi.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Izumi *IzumiCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Izumi.Contract.TokenOfOwnerByIndex(&_Izumi.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Izumi *IzumiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Izumi *IzumiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Izumi.Contract.TokenURI(&_Izumi.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Izumi *IzumiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Izumi.Contract.TokenURI(&_Izumi.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Izumi *IzumiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Izumi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Izumi *IzumiSession) TotalSupply() (*big.Int, error) {
	return _Izumi.Contract.TotalSupply(&_Izumi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Izumi *IzumiCallerSession) TotalSupply() (*big.Int, error) {
	return _Izumi.Contract.TotalSupply(&_Izumi.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactor) AddLiquidity(opts *bind.TransactOpts, addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "addLiquidity", addLiquidityParam)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiSession) AddLiquidity(addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Izumi.Contract.AddLiquidity(&_Izumi.TransactOpts, addLiquidityParam)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactorSession) AddLiquidity(addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Izumi.Contract.AddLiquidity(&_Izumi.TransactOpts, addLiquidityParam)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Izumi *IzumiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Approve(&_Izumi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Approve(&_Izumi.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Izumi *IzumiTransactor) Burn(opts *bind.TransactOpts, lid *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "burn", lid)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Izumi *IzumiSession) Burn(lid *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Burn(&_Izumi.TransactOpts, lid)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Izumi *IzumiTransactorSession) Burn(lid *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Burn(&_Izumi.TransactOpts, lid)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "collect", recipient, lid, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiSession) Collect(recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Collect(&_Izumi.TransactOpts, recipient, lid, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactorSession) Collect(recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.Collect(&_Izumi.TransactOpts, recipient, lid, amountXLim, amountYLim)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Izumi *IzumiTransactor) CreatePool(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "createPool", tokenX, tokenY, fee, initialPoint)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Izumi *IzumiSession) CreatePool(tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.CreatePool(&_Izumi.TransactOpts, tokenX, tokenY, fee, initialPoint)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Izumi *IzumiTransactorSession) CreatePool(tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.CreatePool(&_Izumi.TransactOpts, tokenX, tokenY, fee, initialPoint)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactor) DecLiquidity(opts *bind.TransactOpts, lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "decLiquidity", lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiSession) DecLiquidity(lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.DecLiquidity(&_Izumi.TransactOpts, lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactorSession) DecLiquidity(lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.DecLiquidity(&_Izumi.TransactOpts, lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactor) Mint(opts *bind.TransactOpts, mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "mint", mintParam)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiSession) Mint(mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Izumi.Contract.Mint(&_Izumi.TransactOpts, mintParam)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiTransactorSession) Mint(mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Izumi.Contract.Mint(&_Izumi.TransactOpts, mintParam)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Izumi *IzumiTransactor) MintDepositCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "mintDepositCallback", x, y, data)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Izumi *IzumiSession) MintDepositCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Izumi.Contract.MintDepositCallback(&_Izumi.TransactOpts, x, y, data)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Izumi *IzumiTransactorSession) MintDepositCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Izumi.Contract.MintDepositCallback(&_Izumi.TransactOpts, x, y, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Izumi *IzumiTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Izumi *IzumiSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Izumi.Contract.Multicall(&_Izumi.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Izumi *IzumiTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Izumi.Contract.Multicall(&_Izumi.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Izumi *IzumiTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Izumi *IzumiSession) RefundETH() (*types.Transaction, error) {
	return _Izumi.Contract.RefundETH(&_Izumi.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Izumi *IzumiTransactorSession) RefundETH() (*types.Transaction, error) {
	return _Izumi.Contract.RefundETH(&_Izumi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Izumi *IzumiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Izumi *IzumiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Izumi.Contract.RenounceOwnership(&_Izumi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Izumi *IzumiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Izumi.Contract.RenounceOwnership(&_Izumi.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.SafeTransferFrom(&_Izumi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.SafeTransferFrom(&_Izumi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Izumi *IzumiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Izumi *IzumiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Izumi.Contract.SafeTransferFrom0(&_Izumi.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Izumi *IzumiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Izumi.Contract.SafeTransferFrom0(&_Izumi.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Izumi *IzumiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Izumi *IzumiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Izumi.Contract.SetApprovalForAll(&_Izumi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Izumi *IzumiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Izumi.Contract.SetApprovalForAll(&_Izumi.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Izumi *IzumiTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Izumi *IzumiSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Izumi.Contract.SetBaseURI(&_Izumi.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Izumi *IzumiTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Izumi.Contract.SetBaseURI(&_Izumi.TransactOpts, newBaseURI)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "sweepToken", token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.SweepToken(&_Izumi.TransactOpts, token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiTransactorSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.SweepToken(&_Izumi.TransactOpts, token, minAmount, recipient)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.TransferFrom(&_Izumi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Izumi *IzumiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Izumi.Contract.TransferFrom(&_Izumi.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Izumi *IzumiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Izumi *IzumiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.TransferOwnership(&_Izumi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Izumi *IzumiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.TransferOwnership(&_Izumi.TransactOpts, newOwner)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiTransactor) UnwrapWETH9(opts *bind.TransactOpts, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.contract.Transact(opts, "unwrapWETH9", minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.UnwrapWETH9(&_Izumi.TransactOpts, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Izumi *IzumiTransactorSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Izumi.Contract.UnwrapWETH9(&_Izumi.TransactOpts, minAmount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Izumi *IzumiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Izumi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Izumi *IzumiSession) Receive() (*types.Transaction, error) {
	return _Izumi.Contract.Receive(&_Izumi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Izumi *IzumiTransactorSession) Receive() (*types.Transaction, error) {
	return _Izumi.Contract.Receive(&_Izumi.TransactOpts)
}

// IzumiAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Izumi contract.
type IzumiAddLiquidityIterator struct {
	Event *IzumiAddLiquidity // Event containing the contract specifics and raw log

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
func (it *IzumiAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiAddLiquidity)
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
		it.Event = new(IzumiAddLiquidity)
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
func (it *IzumiAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiAddLiquidity represents a AddLiquidity event raised by the Izumi contract.
type IzumiAddLiquidity struct {
	NftId          *big.Int
	Pool           common.Address
	LiquidityDelta *big.Int
	AmountX        *big.Int
	AmountY        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) FilterAddLiquidity(opts *bind.FilterOpts, nftId []*big.Int) (*IzumiAddLiquidityIterator, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "AddLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return &IzumiAddLiquidityIterator{contract: _Izumi.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *IzumiAddLiquidity, nftId []*big.Int) (event.Subscription, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "AddLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiAddLiquidity)
				if err := _Izumi.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) ParseAddLiquidity(log types.Log) (*IzumiAddLiquidity, error) {
	event := new(IzumiAddLiquidity)
	if err := _Izumi.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IzumiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Izumi contract.
type IzumiApprovalIterator struct {
	Event *IzumiApproval // Event containing the contract specifics and raw log

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
func (it *IzumiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiApproval)
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
		it.Event = new(IzumiApproval)
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
func (it *IzumiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiApproval represents a Approval event raised by the Izumi contract.
type IzumiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IzumiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IzumiApprovalIterator{contract: _Izumi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IzumiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiApproval)
				if err := _Izumi.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) ParseApproval(log types.Log) (*IzumiApproval, error) {
	event := new(IzumiApproval)
	if err := _Izumi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IzumiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Izumi contract.
type IzumiApprovalForAllIterator struct {
	Event *IzumiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IzumiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiApprovalForAll)
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
		it.Event = new(IzumiApprovalForAll)
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
func (it *IzumiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiApprovalForAll represents a ApprovalForAll event raised by the Izumi contract.
type IzumiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Izumi *IzumiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IzumiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IzumiApprovalForAllIterator{contract: _Izumi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Izumi *IzumiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IzumiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiApprovalForAll)
				if err := _Izumi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Izumi *IzumiFilterer) ParseApprovalForAll(log types.Log) (*IzumiApprovalForAll, error) {
	event := new(IzumiApprovalForAll)
	if err := _Izumi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IzumiDecLiquidityIterator is returned from FilterDecLiquidity and is used to iterate over the raw logs and unpacked data for DecLiquidity events raised by the Izumi contract.
type IzumiDecLiquidityIterator struct {
	Event *IzumiDecLiquidity // Event containing the contract specifics and raw log

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
func (it *IzumiDecLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiDecLiquidity)
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
		it.Event = new(IzumiDecLiquidity)
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
func (it *IzumiDecLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiDecLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiDecLiquidity represents a DecLiquidity event raised by the Izumi contract.
type IzumiDecLiquidity struct {
	NftId          *big.Int
	Pool           common.Address
	LiquidityDelta *big.Int
	AmountX        *big.Int
	AmountY        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDecLiquidity is a free log retrieval operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) FilterDecLiquidity(opts *bind.FilterOpts, nftId []*big.Int) (*IzumiDecLiquidityIterator, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "DecLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return &IzumiDecLiquidityIterator{contract: _Izumi.contract, event: "DecLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecLiquidity is a free log subscription operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) WatchDecLiquidity(opts *bind.WatchOpts, sink chan<- *IzumiDecLiquidity, nftId []*big.Int) (event.Subscription, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "DecLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiDecLiquidity)
				if err := _Izumi.contract.UnpackLog(event, "DecLiquidity", log); err != nil {
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

// ParseDecLiquidity is a log parse operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Izumi *IzumiFilterer) ParseDecLiquidity(log types.Log) (*IzumiDecLiquidity, error) {
	event := new(IzumiDecLiquidity)
	if err := _Izumi.contract.UnpackLog(event, "DecLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IzumiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Izumi contract.
type IzumiOwnershipTransferredIterator struct {
	Event *IzumiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IzumiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiOwnershipTransferred)
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
		it.Event = new(IzumiOwnershipTransferred)
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
func (it *IzumiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiOwnershipTransferred represents a OwnershipTransferred event raised by the Izumi contract.
type IzumiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Izumi *IzumiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IzumiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IzumiOwnershipTransferredIterator{contract: _Izumi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Izumi *IzumiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IzumiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiOwnershipTransferred)
				if err := _Izumi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Izumi *IzumiFilterer) ParseOwnershipTransferred(log types.Log) (*IzumiOwnershipTransferred, error) {
	event := new(IzumiOwnershipTransferred)
	if err := _Izumi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IzumiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Izumi contract.
type IzumiTransferIterator struct {
	Event *IzumiTransfer // Event containing the contract specifics and raw log

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
func (it *IzumiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IzumiTransfer)
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
		it.Event = new(IzumiTransfer)
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
func (it *IzumiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IzumiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IzumiTransfer represents a Transfer event raised by the Izumi contract.
type IzumiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IzumiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Izumi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IzumiTransferIterator{contract: _Izumi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IzumiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Izumi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IzumiTransfer)
				if err := _Izumi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Izumi *IzumiFilterer) ParseTransfer(log types.Log) (*IzumiTransfer, error) {
	event := new(IzumiTransfer)
	if err := _Izumi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
