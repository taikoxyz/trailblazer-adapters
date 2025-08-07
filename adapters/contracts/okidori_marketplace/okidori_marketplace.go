// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package okidori_marketplace

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

// OkidoriListing is an auto generated low-level Go binding around an user-defined struct.
type OkidoriListing struct {
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
}

// OkidoriMarketplaceMetaData contains all meta data concerning the OkidoriMarketplace contract.
var OkidoriMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"CurrencyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"CurrencyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeBasisPointsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ListingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"NFTListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"NFTSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"royaltyBasisPoints\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetCollectionRoyalties\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"addAllowedCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedCurrencies\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structOkidori.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"listNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"removeAllowedCurrency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"royaltyBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setCollectionRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setCollectionRoyaltiesOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setFeeBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OkidoriMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use OkidoriMarketplaceMetaData.ABI instead.
var OkidoriMarketplaceABI = OkidoriMarketplaceMetaData.ABI

// OkidoriMarketplace is an auto generated Go binding around an Ethereum contract.
type OkidoriMarketplace struct {
	OkidoriMarketplaceCaller     // Read-only binding to the contract
	OkidoriMarketplaceTransactor // Write-only binding to the contract
	OkidoriMarketplaceFilterer   // Log filterer for contract events
}

// OkidoriMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type OkidoriMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkidoriMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OkidoriMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkidoriMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OkidoriMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OkidoriMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OkidoriMarketplaceSession struct {
	Contract     *OkidoriMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OkidoriMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OkidoriMarketplaceCallerSession struct {
	Contract *OkidoriMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OkidoriMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OkidoriMarketplaceTransactorSession struct {
	Contract     *OkidoriMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OkidoriMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type OkidoriMarketplaceRaw struct {
	Contract *OkidoriMarketplace // Generic contract binding to access the raw methods on
}

// OkidoriMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OkidoriMarketplaceCallerRaw struct {
	Contract *OkidoriMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// OkidoriMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OkidoriMarketplaceTransactorRaw struct {
	Contract *OkidoriMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOkidoriMarketplace creates a new instance of OkidoriMarketplace, bound to a specific deployed contract.
func NewOkidoriMarketplace(address common.Address, backend bind.ContractBackend) (*OkidoriMarketplace, error) {
	contract, err := bindOkidoriMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplace{OkidoriMarketplaceCaller: OkidoriMarketplaceCaller{contract: contract}, OkidoriMarketplaceTransactor: OkidoriMarketplaceTransactor{contract: contract}, OkidoriMarketplaceFilterer: OkidoriMarketplaceFilterer{contract: contract}}, nil
}

// NewOkidoriMarketplaceCaller creates a new read-only instance of OkidoriMarketplace, bound to a specific deployed contract.
func NewOkidoriMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*OkidoriMarketplaceCaller, error) {
	contract, err := bindOkidoriMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceCaller{contract: contract}, nil
}

// NewOkidoriMarketplaceTransactor creates a new write-only instance of OkidoriMarketplace, bound to a specific deployed contract.
func NewOkidoriMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*OkidoriMarketplaceTransactor, error) {
	contract, err := bindOkidoriMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceTransactor{contract: contract}, nil
}

// NewOkidoriMarketplaceFilterer creates a new log filterer instance of OkidoriMarketplace, bound to a specific deployed contract.
func NewOkidoriMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*OkidoriMarketplaceFilterer, error) {
	contract, err := bindOkidoriMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceFilterer{contract: contract}, nil
}

// bindOkidoriMarketplace binds a generic wrapper to an already deployed contract.
func bindOkidoriMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OkidoriMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OkidoriMarketplace *OkidoriMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OkidoriMarketplace.Contract.OkidoriMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OkidoriMarketplace *OkidoriMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.OkidoriMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OkidoriMarketplace *OkidoriMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.OkidoriMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OkidoriMarketplace *OkidoriMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OkidoriMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _OkidoriMarketplace.Contract.UPGRADEINTERFACEVERSION(&_OkidoriMarketplace.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _OkidoriMarketplace.Contract.UPGRADEINTERFACEVERSION(&_OkidoriMarketplace.CallOpts)
}

// AllowedCurrencies is a free data retrieval call binding the contract method 0x86e306b5.
//
// Solidity: function allowedCurrencies(address ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) AllowedCurrencies(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "allowedCurrencies", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedCurrencies is a free data retrieval call binding the contract method 0x86e306b5.
//
// Solidity: function allowedCurrencies(address ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) AllowedCurrencies(arg0 common.Address) (bool, error) {
	return _OkidoriMarketplace.Contract.AllowedCurrencies(&_OkidoriMarketplace.CallOpts, arg0)
}

// AllowedCurrencies is a free data retrieval call binding the contract method 0x86e306b5.
//
// Solidity: function allowedCurrencies(address ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) AllowedCurrencies(arg0 common.Address) (bool, error) {
	return _OkidoriMarketplace.Contract.AllowedCurrencies(&_OkidoriMarketplace.CallOpts, arg0)
}

// FeeBasisPoints is a free data retrieval call binding the contract method 0xb8606eef.
//
// Solidity: function feeBasisPoints() view returns(uint256)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) FeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "feeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBasisPoints is a free data retrieval call binding the contract method 0xb8606eef.
//
// Solidity: function feeBasisPoints() view returns(uint256)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) FeeBasisPoints() (*big.Int, error) {
	return _OkidoriMarketplace.Contract.FeeBasisPoints(&_OkidoriMarketplace.CallOpts)
}

// FeeBasisPoints is a free data retrieval call binding the contract method 0xb8606eef.
//
// Solidity: function feeBasisPoints() view returns(uint256)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) FeeBasisPoints() (*big.Int, error) {
	return _OkidoriMarketplace.Contract.FeeBasisPoints(&_OkidoriMarketplace.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,address,uint256,uint256,address))
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (OkidoriListing, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "getListing", listingId)

	if err != nil {
		return *new(OkidoriListing), err
	}

	out0 := *abi.ConvertType(out[0], new(OkidoriListing)).(*OkidoriListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,address,uint256,uint256,address))
func (_OkidoriMarketplace *OkidoriMarketplaceSession) GetListing(listingId *big.Int) (OkidoriListing, error) {
	return _OkidoriMarketplace.Contract.GetListing(&_OkidoriMarketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((address,address,uint256,uint256,address))
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) GetListing(listingId *big.Int) (OkidoriListing, error) {
	return _OkidoriMarketplace.Contract.GetListing(&_OkidoriMarketplace.CallOpts, listingId)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) IsActive(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "isActive", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) IsActive(arg0 *big.Int) (bool, error) {
	return _OkidoriMarketplace.Contract.IsActive(&_OkidoriMarketplace.CallOpts, arg0)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 ) view returns(bool)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) IsActive(arg0 *big.Int) (bool, error) {
	return _OkidoriMarketplace.Contract.IsActive(&_OkidoriMarketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
}, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Seller          common.Address
		ContractAddress common.Address
		TokenId         *big.Int
		Price           *big.Int
		Currency        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) Listings(arg0 *big.Int) (struct {
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
}, error) {
	return _OkidoriMarketplace.Contract.Listings(&_OkidoriMarketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
}, error) {
	return _OkidoriMarketplace.Contract.Listings(&_OkidoriMarketplace.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) Owner() (common.Address, error) {
	return _OkidoriMarketplace.Contract.Owner(&_OkidoriMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) Owner() (common.Address, error) {
	return _OkidoriMarketplace.Contract.Owner(&_OkidoriMarketplace.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) PendingOwner() (common.Address, error) {
	return _OkidoriMarketplace.Contract.PendingOwner(&_OkidoriMarketplace.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) PendingOwner() (common.Address, error) {
	return _OkidoriMarketplace.Contract.PendingOwner(&_OkidoriMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) ProxiableUUID() ([32]byte, error) {
	return _OkidoriMarketplace.Contract.ProxiableUUID(&_OkidoriMarketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _OkidoriMarketplace.Contract.ProxiableUUID(&_OkidoriMarketplace.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0x9fa6b4a0.
//
// Solidity: function royalties(address ) view returns(uint256 royaltyBasisPoints, address receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceCaller) Royalties(opts *bind.CallOpts, arg0 common.Address) (struct {
	RoyaltyBasisPoints *big.Int
	Receiver           common.Address
}, error) {
	var out []interface{}
	err := _OkidoriMarketplace.contract.Call(opts, &out, "royalties", arg0)

	outstruct := new(struct {
		RoyaltyBasisPoints *big.Int
		Receiver           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoyaltyBasisPoints = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Receiver = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Royalties is a free data retrieval call binding the contract method 0x9fa6b4a0.
//
// Solidity: function royalties(address ) view returns(uint256 royaltyBasisPoints, address receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceSession) Royalties(arg0 common.Address) (struct {
	RoyaltyBasisPoints *big.Int
	Receiver           common.Address
}, error) {
	return _OkidoriMarketplace.Contract.Royalties(&_OkidoriMarketplace.CallOpts, arg0)
}

// Royalties is a free data retrieval call binding the contract method 0x9fa6b4a0.
//
// Solidity: function royalties(address ) view returns(uint256 royaltyBasisPoints, address receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceCallerSession) Royalties(arg0 common.Address) (struct {
	RoyaltyBasisPoints *big.Int
	Receiver           common.Address
}, error) {
	return _OkidoriMarketplace.Contract.Royalties(&_OkidoriMarketplace.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) AcceptOwnership() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.AcceptOwnership(&_OkidoriMarketplace.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.AcceptOwnership(&_OkidoriMarketplace.TransactOpts)
}

// AddAllowedCurrency is a paid mutator transaction binding the contract method 0xf2298cf4.
//
// Solidity: function addAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) AddAllowedCurrency(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "addAllowedCurrency", tokenAddress)
}

// AddAllowedCurrency is a paid mutator transaction binding the contract method 0xf2298cf4.
//
// Solidity: function addAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) AddAllowedCurrency(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.AddAllowedCurrency(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// AddAllowedCurrency is a paid mutator transaction binding the contract method 0xf2298cf4.
//
// Solidity: function addAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) AddAllowedCurrency(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.AddAllowedCurrency(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 listingId) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) BuyNFT(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "buyNFT", listingId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 listingId) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) BuyNFT(listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.BuyNFT(&_OkidoriMarketplace.TransactOpts, listingId)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 listingId) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) BuyNFT(listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.BuyNFT(&_OkidoriMarketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "cancelListing", listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.CancelListing(&_OkidoriMarketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.CancelListing(&_OkidoriMarketplace.TransactOpts, listingId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) Initialize() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.Initialize(&_OkidoriMarketplace.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) Initialize() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.Initialize(&_OkidoriMarketplace.TransactOpts)
}

// ListNFT is a paid mutator transaction binding the contract method 0xd68cdc52.
//
// Solidity: function listNFT(address contractAddress, uint256 tokenId, uint256 price, address currency) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) ListNFT(opts *bind.TransactOpts, contractAddress common.Address, tokenId *big.Int, price *big.Int, currency common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "listNFT", contractAddress, tokenId, price, currency)
}

// ListNFT is a paid mutator transaction binding the contract method 0xd68cdc52.
//
// Solidity: function listNFT(address contractAddress, uint256 tokenId, uint256 price, address currency) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) ListNFT(contractAddress common.Address, tokenId *big.Int, price *big.Int, currency common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.ListNFT(&_OkidoriMarketplace.TransactOpts, contractAddress, tokenId, price, currency)
}

// ListNFT is a paid mutator transaction binding the contract method 0xd68cdc52.
//
// Solidity: function listNFT(address contractAddress, uint256 tokenId, uint256 price, address currency) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) ListNFT(contractAddress common.Address, tokenId *big.Int, price *big.Int, currency common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.ListNFT(&_OkidoriMarketplace.TransactOpts, contractAddress, tokenId, price, currency)
}

// RemoveAllowedCurrency is a paid mutator transaction binding the contract method 0x84ba316d.
//
// Solidity: function removeAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) RemoveAllowedCurrency(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "removeAllowedCurrency", tokenAddress)
}

// RemoveAllowedCurrency is a paid mutator transaction binding the contract method 0x84ba316d.
//
// Solidity: function removeAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) RemoveAllowedCurrency(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.RemoveAllowedCurrency(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// RemoveAllowedCurrency is a paid mutator transaction binding the contract method 0x84ba316d.
//
// Solidity: function removeAllowedCurrency(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) RemoveAllowedCurrency(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.RemoveAllowedCurrency(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.RenounceOwnership(&_OkidoriMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.RenounceOwnership(&_OkidoriMarketplace.TransactOpts)
}

// SetCollectionRoyalties is a paid mutator transaction binding the contract method 0x7f69254a.
//
// Solidity: function setCollectionRoyalties(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) SetCollectionRoyalties(opts *bind.TransactOpts, contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "setCollectionRoyalties", contractAddress, amount, receiver)
}

// SetCollectionRoyalties is a paid mutator transaction binding the contract method 0x7f69254a.
//
// Solidity: function setCollectionRoyalties(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) SetCollectionRoyalties(contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetCollectionRoyalties(&_OkidoriMarketplace.TransactOpts, contractAddress, amount, receiver)
}

// SetCollectionRoyalties is a paid mutator transaction binding the contract method 0x7f69254a.
//
// Solidity: function setCollectionRoyalties(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) SetCollectionRoyalties(contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetCollectionRoyalties(&_OkidoriMarketplace.TransactOpts, contractAddress, amount, receiver)
}

// SetCollectionRoyaltiesOwner is a paid mutator transaction binding the contract method 0x36439391.
//
// Solidity: function setCollectionRoyaltiesOwner(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) SetCollectionRoyaltiesOwner(opts *bind.TransactOpts, contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "setCollectionRoyaltiesOwner", contractAddress, amount, receiver)
}

// SetCollectionRoyaltiesOwner is a paid mutator transaction binding the contract method 0x36439391.
//
// Solidity: function setCollectionRoyaltiesOwner(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) SetCollectionRoyaltiesOwner(contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetCollectionRoyaltiesOwner(&_OkidoriMarketplace.TransactOpts, contractAddress, amount, receiver)
}

// SetCollectionRoyaltiesOwner is a paid mutator transaction binding the contract method 0x36439391.
//
// Solidity: function setCollectionRoyaltiesOwner(address contractAddress, uint256 amount, address receiver) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) SetCollectionRoyaltiesOwner(contractAddress common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetCollectionRoyaltiesOwner(&_OkidoriMarketplace.TransactOpts, contractAddress, amount, receiver)
}

// SetFeeBasisPoints is a paid mutator transaction binding the contract method 0x0165dd27.
//
// Solidity: function setFeeBasisPoints(uint256 newBasisPoints) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) SetFeeBasisPoints(opts *bind.TransactOpts, newBasisPoints *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "setFeeBasisPoints", newBasisPoints)
}

// SetFeeBasisPoints is a paid mutator transaction binding the contract method 0x0165dd27.
//
// Solidity: function setFeeBasisPoints(uint256 newBasisPoints) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) SetFeeBasisPoints(newBasisPoints *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetFeeBasisPoints(&_OkidoriMarketplace.TransactOpts, newBasisPoints)
}

// SetFeeBasisPoints is a paid mutator transaction binding the contract method 0x0165dd27.
//
// Solidity: function setFeeBasisPoints(uint256 newBasisPoints) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) SetFeeBasisPoints(newBasisPoints *big.Int) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.SetFeeBasisPoints(&_OkidoriMarketplace.TransactOpts, newBasisPoints)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.TransferOwnership(&_OkidoriMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.TransferOwnership(&_OkidoriMarketplace.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.UpgradeToAndCall(&_OkidoriMarketplace.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.UpgradeToAndCall(&_OkidoriMarketplace.TransactOpts, newImplementation, data)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactor) WithdrawFees(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.contract.Transact(opts, "withdrawFees", tokenAddress)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceSession) WithdrawFees(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.WithdrawFees(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address tokenAddress) returns()
func (_OkidoriMarketplace *OkidoriMarketplaceTransactorSession) WithdrawFees(tokenAddress common.Address) (*types.Transaction, error) {
	return _OkidoriMarketplace.Contract.WithdrawFees(&_OkidoriMarketplace.TransactOpts, tokenAddress)
}

// OkidoriMarketplaceCurrencyAddedIterator is returned from FilterCurrencyAdded and is used to iterate over the raw logs and unpacked data for CurrencyAdded events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceCurrencyAddedIterator struct {
	Event *OkidoriMarketplaceCurrencyAdded // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceCurrencyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceCurrencyAdded)
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
		it.Event = new(OkidoriMarketplaceCurrencyAdded)
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
func (it *OkidoriMarketplaceCurrencyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceCurrencyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceCurrencyAdded represents a CurrencyAdded event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceCurrencyAdded struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCurrencyAdded is a free log retrieval operation binding the contract event 0xe0390a89516146ccfb796a9fbfc3f7646282194c554d05020484599ee992dd5a.
//
// Solidity: event CurrencyAdded(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterCurrencyAdded(opts *bind.FilterOpts, currency []common.Address) (*OkidoriMarketplaceCurrencyAddedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "CurrencyAdded", currencyRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceCurrencyAddedIterator{contract: _OkidoriMarketplace.contract, event: "CurrencyAdded", logs: logs, sub: sub}, nil
}

// WatchCurrencyAdded is a free log subscription operation binding the contract event 0xe0390a89516146ccfb796a9fbfc3f7646282194c554d05020484599ee992dd5a.
//
// Solidity: event CurrencyAdded(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchCurrencyAdded(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceCurrencyAdded, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "CurrencyAdded", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceCurrencyAdded)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "CurrencyAdded", log); err != nil {
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

// ParseCurrencyAdded is a log parse operation binding the contract event 0xe0390a89516146ccfb796a9fbfc3f7646282194c554d05020484599ee992dd5a.
//
// Solidity: event CurrencyAdded(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseCurrencyAdded(log types.Log) (*OkidoriMarketplaceCurrencyAdded, error) {
	event := new(OkidoriMarketplaceCurrencyAdded)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "CurrencyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceCurrencyRemovedIterator is returned from FilterCurrencyRemoved and is used to iterate over the raw logs and unpacked data for CurrencyRemoved events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceCurrencyRemovedIterator struct {
	Event *OkidoriMarketplaceCurrencyRemoved // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceCurrencyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceCurrencyRemoved)
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
		it.Event = new(OkidoriMarketplaceCurrencyRemoved)
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
func (it *OkidoriMarketplaceCurrencyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceCurrencyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceCurrencyRemoved represents a CurrencyRemoved event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceCurrencyRemoved struct {
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCurrencyRemoved is a free log retrieval operation binding the contract event 0xa40d69111be14f29022626d38310e47cc2d7f4cb728961509c2f65a4bee08c5b.
//
// Solidity: event CurrencyRemoved(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterCurrencyRemoved(opts *bind.FilterOpts, currency []common.Address) (*OkidoriMarketplaceCurrencyRemovedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "CurrencyRemoved", currencyRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceCurrencyRemovedIterator{contract: _OkidoriMarketplace.contract, event: "CurrencyRemoved", logs: logs, sub: sub}, nil
}

// WatchCurrencyRemoved is a free log subscription operation binding the contract event 0xa40d69111be14f29022626d38310e47cc2d7f4cb728961509c2f65a4bee08c5b.
//
// Solidity: event CurrencyRemoved(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchCurrencyRemoved(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceCurrencyRemoved, currency []common.Address) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "CurrencyRemoved", currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceCurrencyRemoved)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "CurrencyRemoved", log); err != nil {
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

// ParseCurrencyRemoved is a log parse operation binding the contract event 0xa40d69111be14f29022626d38310e47cc2d7f4cb728961509c2f65a4bee08c5b.
//
// Solidity: event CurrencyRemoved(address indexed currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseCurrencyRemoved(log types.Log) (*OkidoriMarketplaceCurrencyRemoved, error) {
	event := new(OkidoriMarketplaceCurrencyRemoved)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "CurrencyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceFeeBasisPointsSetIterator is returned from FilterFeeBasisPointsSet and is used to iterate over the raw logs and unpacked data for FeeBasisPointsSet events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceFeeBasisPointsSetIterator struct {
	Event *OkidoriMarketplaceFeeBasisPointsSet // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceFeeBasisPointsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceFeeBasisPointsSet)
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
		it.Event = new(OkidoriMarketplaceFeeBasisPointsSet)
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
func (it *OkidoriMarketplaceFeeBasisPointsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceFeeBasisPointsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceFeeBasisPointsSet represents a FeeBasisPointsSet event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceFeeBasisPointsSet struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeBasisPointsSet is a free log retrieval operation binding the contract event 0xc5c39a79f3f590646ae56276b3837de05ba12bacf6e6b08368377e426059ad4c.
//
// Solidity: event FeeBasisPointsSet(uint256 indexed amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterFeeBasisPointsSet(opts *bind.FilterOpts, amount []*big.Int) (*OkidoriMarketplaceFeeBasisPointsSetIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "FeeBasisPointsSet", amountRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceFeeBasisPointsSetIterator{contract: _OkidoriMarketplace.contract, event: "FeeBasisPointsSet", logs: logs, sub: sub}, nil
}

// WatchFeeBasisPointsSet is a free log subscription operation binding the contract event 0xc5c39a79f3f590646ae56276b3837de05ba12bacf6e6b08368377e426059ad4c.
//
// Solidity: event FeeBasisPointsSet(uint256 indexed amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchFeeBasisPointsSet(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceFeeBasisPointsSet, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "FeeBasisPointsSet", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceFeeBasisPointsSet)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "FeeBasisPointsSet", log); err != nil {
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

// ParseFeeBasisPointsSet is a log parse operation binding the contract event 0xc5c39a79f3f590646ae56276b3837de05ba12bacf6e6b08368377e426059ad4c.
//
// Solidity: event FeeBasisPointsSet(uint256 indexed amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseFeeBasisPointsSet(log types.Log) (*OkidoriMarketplaceFeeBasisPointsSet, error) {
	event := new(OkidoriMarketplaceFeeBasisPointsSet)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "FeeBasisPointsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceFeeWithdrawnIterator is returned from FilterFeeWithdrawn and is used to iterate over the raw logs and unpacked data for FeeWithdrawn events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceFeeWithdrawnIterator struct {
	Event *OkidoriMarketplaceFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceFeeWithdrawn)
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
		it.Event = new(OkidoriMarketplaceFeeWithdrawn)
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
func (it *OkidoriMarketplaceFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceFeeWithdrawn represents a FeeWithdrawn event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceFeeWithdrawn struct {
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawn is a free log retrieval operation binding the contract event 0x00ed5939179dc194223f0edd1517ecee2210b22da7f82c8e4b1795e93b9f06aa.
//
// Solidity: event FeeWithdrawn(address indexed to, address indexed token, uint256 amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterFeeWithdrawn(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*OkidoriMarketplaceFeeWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "FeeWithdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceFeeWithdrawnIterator{contract: _OkidoriMarketplace.contract, event: "FeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawn is a free log subscription operation binding the contract event 0x00ed5939179dc194223f0edd1517ecee2210b22da7f82c8e4b1795e93b9f06aa.
//
// Solidity: event FeeWithdrawn(address indexed to, address indexed token, uint256 amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceFeeWithdrawn, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "FeeWithdrawn", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceFeeWithdrawn)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
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

// ParseFeeWithdrawn is a log parse operation binding the contract event 0x00ed5939179dc194223f0edd1517ecee2210b22da7f82c8e4b1795e93b9f06aa.
//
// Solidity: event FeeWithdrawn(address indexed to, address indexed token, uint256 amount)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseFeeWithdrawn(log types.Log) (*OkidoriMarketplaceFeeWithdrawn, error) {
	event := new(OkidoriMarketplaceFeeWithdrawn)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceInitializedIterator struct {
	Event *OkidoriMarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceInitialized)
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
		it.Event = new(OkidoriMarketplaceInitialized)
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
func (it *OkidoriMarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceInitialized represents a Initialized event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*OkidoriMarketplaceInitializedIterator, error) {

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceInitializedIterator{contract: _OkidoriMarketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceInitialized)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseInitialized(log types.Log) (*OkidoriMarketplaceInitialized, error) {
	event := new(OkidoriMarketplaceInitialized)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceListingCancelledIterator is returned from FilterListingCancelled and is used to iterate over the raw logs and unpacked data for ListingCancelled events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceListingCancelledIterator struct {
	Event *OkidoriMarketplaceListingCancelled // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceListingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceListingCancelled)
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
		it.Event = new(OkidoriMarketplaceListingCancelled)
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
func (it *OkidoriMarketplaceListingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceListingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceListingCancelled represents a ListingCancelled event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceListingCancelled struct {
	ListingId       *big.Int
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterListingCancelled is a free log retrieval operation binding the contract event 0x61b65dee2812bcf80c52c603ed6479cde390ec97008da9021789097ffe9b9de7.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterListingCancelled(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, contractAddress []common.Address) (*OkidoriMarketplaceListingCancelledIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "ListingCancelled", listingIdRule, sellerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceListingCancelledIterator{contract: _OkidoriMarketplace.contract, event: "ListingCancelled", logs: logs, sub: sub}, nil
}

// WatchListingCancelled is a free log subscription operation binding the contract event 0x61b65dee2812bcf80c52c603ed6479cde390ec97008da9021789097ffe9b9de7.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchListingCancelled(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceListingCancelled, listingId []*big.Int, seller []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "ListingCancelled", listingIdRule, sellerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceListingCancelled)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
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

// ParseListingCancelled is a log parse operation binding the contract event 0x61b65dee2812bcf80c52c603ed6479cde390ec97008da9021789097ffe9b9de7.
//
// Solidity: event ListingCancelled(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseListingCancelled(log types.Log) (*OkidoriMarketplaceListingCancelled, error) {
	event := new(OkidoriMarketplaceListingCancelled)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceNFTListedIterator is returned from FilterNFTListed and is used to iterate over the raw logs and unpacked data for NFTListed events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceNFTListedIterator struct {
	Event *OkidoriMarketplaceNFTListed // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceNFTListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceNFTListed)
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
		it.Event = new(OkidoriMarketplaceNFTListed)
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
func (it *OkidoriMarketplaceNFTListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceNFTListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceNFTListed represents a NFTListed event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceNFTListed struct {
	ListingId       *big.Int
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNFTListed is a free log retrieval operation binding the contract event 0x6b46a494d716ac0c5c923f09b3c2b579c98fc55156272713ffbf4df5f24de105.
//
// Solidity: event NFTListed(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterNFTListed(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, contractAddress []common.Address) (*OkidoriMarketplaceNFTListedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "NFTListed", listingIdRule, sellerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceNFTListedIterator{contract: _OkidoriMarketplace.contract, event: "NFTListed", logs: logs, sub: sub}, nil
}

// WatchNFTListed is a free log subscription operation binding the contract event 0x6b46a494d716ac0c5c923f09b3c2b579c98fc55156272713ffbf4df5f24de105.
//
// Solidity: event NFTListed(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchNFTListed(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceNFTListed, listingId []*big.Int, seller []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "NFTListed", listingIdRule, sellerRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceNFTListed)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "NFTListed", log); err != nil {
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

// ParseNFTListed is a log parse operation binding the contract event 0x6b46a494d716ac0c5c923f09b3c2b579c98fc55156272713ffbf4df5f24de105.
//
// Solidity: event NFTListed(uint256 indexed listingId, address indexed seller, address indexed contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseNFTListed(log types.Log) (*OkidoriMarketplaceNFTListed, error) {
	event := new(OkidoriMarketplaceNFTListed)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "NFTListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceNFTSoldIterator is returned from FilterNFTSold and is used to iterate over the raw logs and unpacked data for NFTSold events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceNFTSoldIterator struct {
	Event *OkidoriMarketplaceNFTSold // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceNFTSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceNFTSold)
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
		it.Event = new(OkidoriMarketplaceNFTSold)
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
func (it *OkidoriMarketplaceNFTSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceNFTSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceNFTSold represents a NFTSold event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceNFTSold struct {
	ListingId       *big.Int
	Buyer           common.Address
	Seller          common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Currency        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNFTSold is a free log retrieval operation binding the contract event 0x3cb1c49116d70fdce7e3cc17ee7ecf520353af544747a0c7cc45162c80e43dd1.
//
// Solidity: event NFTSold(uint256 indexed listingId, address indexed buyer, address indexed seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterNFTSold(opts *bind.FilterOpts, listingId []*big.Int, buyer []common.Address, seller []common.Address) (*OkidoriMarketplaceNFTSoldIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "NFTSold", listingIdRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceNFTSoldIterator{contract: _OkidoriMarketplace.contract, event: "NFTSold", logs: logs, sub: sub}, nil
}

// WatchNFTSold is a free log subscription operation binding the contract event 0x3cb1c49116d70fdce7e3cc17ee7ecf520353af544747a0c7cc45162c80e43dd1.
//
// Solidity: event NFTSold(uint256 indexed listingId, address indexed buyer, address indexed seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchNFTSold(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceNFTSold, listingId []*big.Int, buyer []common.Address, seller []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "NFTSold", listingIdRule, buyerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceNFTSold)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "NFTSold", log); err != nil {
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

// ParseNFTSold is a log parse operation binding the contract event 0x3cb1c49116d70fdce7e3cc17ee7ecf520353af544747a0c7cc45162c80e43dd1.
//
// Solidity: event NFTSold(uint256 indexed listingId, address indexed buyer, address indexed seller, address contractAddress, uint256 tokenId, uint256 price, address currency)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseNFTSold(log types.Log) (*OkidoriMarketplaceNFTSold, error) {
	event := new(OkidoriMarketplaceNFTSold)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "NFTSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceOwnershipTransferStartedIterator struct {
	Event *OkidoriMarketplaceOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceOwnershipTransferStarted)
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
		it.Event = new(OkidoriMarketplaceOwnershipTransferStarted)
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
func (it *OkidoriMarketplaceOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OkidoriMarketplaceOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceOwnershipTransferStartedIterator{contract: _OkidoriMarketplace.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceOwnershipTransferStarted)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseOwnershipTransferStarted(log types.Log) (*OkidoriMarketplaceOwnershipTransferStarted, error) {
	event := new(OkidoriMarketplaceOwnershipTransferStarted)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceOwnershipTransferredIterator struct {
	Event *OkidoriMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceOwnershipTransferred)
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
		it.Event = new(OkidoriMarketplaceOwnershipTransferred)
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
func (it *OkidoriMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OkidoriMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceOwnershipTransferredIterator{contract: _OkidoriMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceOwnershipTransferred)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*OkidoriMarketplaceOwnershipTransferred, error) {
	event := new(OkidoriMarketplaceOwnershipTransferred)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceSetCollectionRoyaltiesIterator is returned from FilterSetCollectionRoyalties and is used to iterate over the raw logs and unpacked data for SetCollectionRoyalties events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceSetCollectionRoyaltiesIterator struct {
	Event *OkidoriMarketplaceSetCollectionRoyalties // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceSetCollectionRoyaltiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceSetCollectionRoyalties)
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
		it.Event = new(OkidoriMarketplaceSetCollectionRoyalties)
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
func (it *OkidoriMarketplaceSetCollectionRoyaltiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceSetCollectionRoyaltiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceSetCollectionRoyalties represents a SetCollectionRoyalties event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceSetCollectionRoyalties struct {
	ContractAddress    common.Address
	RoyaltyBasisPoints *big.Int
	Receiver           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetCollectionRoyalties is a free log retrieval operation binding the contract event 0x722629ff07d664f1b76adca27d3cec23e66855f6474077ec3e08ecf25d354711.
//
// Solidity: event SetCollectionRoyalties(address indexed contractAddress, uint256 indexed royaltyBasisPoints, address indexed receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterSetCollectionRoyalties(opts *bind.FilterOpts, contractAddress []common.Address, royaltyBasisPoints []*big.Int, receiver []common.Address) (*OkidoriMarketplaceSetCollectionRoyaltiesIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var royaltyBasisPointsRule []interface{}
	for _, royaltyBasisPointsItem := range royaltyBasisPoints {
		royaltyBasisPointsRule = append(royaltyBasisPointsRule, royaltyBasisPointsItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "SetCollectionRoyalties", contractAddressRule, royaltyBasisPointsRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceSetCollectionRoyaltiesIterator{contract: _OkidoriMarketplace.contract, event: "SetCollectionRoyalties", logs: logs, sub: sub}, nil
}

// WatchSetCollectionRoyalties is a free log subscription operation binding the contract event 0x722629ff07d664f1b76adca27d3cec23e66855f6474077ec3e08ecf25d354711.
//
// Solidity: event SetCollectionRoyalties(address indexed contractAddress, uint256 indexed royaltyBasisPoints, address indexed receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchSetCollectionRoyalties(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceSetCollectionRoyalties, contractAddress []common.Address, royaltyBasisPoints []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var royaltyBasisPointsRule []interface{}
	for _, royaltyBasisPointsItem := range royaltyBasisPoints {
		royaltyBasisPointsRule = append(royaltyBasisPointsRule, royaltyBasisPointsItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "SetCollectionRoyalties", contractAddressRule, royaltyBasisPointsRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceSetCollectionRoyalties)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "SetCollectionRoyalties", log); err != nil {
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

// ParseSetCollectionRoyalties is a log parse operation binding the contract event 0x722629ff07d664f1b76adca27d3cec23e66855f6474077ec3e08ecf25d354711.
//
// Solidity: event SetCollectionRoyalties(address indexed contractAddress, uint256 indexed royaltyBasisPoints, address indexed receiver)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseSetCollectionRoyalties(log types.Log) (*OkidoriMarketplaceSetCollectionRoyalties, error) {
	event := new(OkidoriMarketplaceSetCollectionRoyalties)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "SetCollectionRoyalties", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OkidoriMarketplaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceUpgradedIterator struct {
	Event *OkidoriMarketplaceUpgraded // Event containing the contract specifics and raw log

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
func (it *OkidoriMarketplaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OkidoriMarketplaceUpgraded)
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
		it.Event = new(OkidoriMarketplaceUpgraded)
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
func (it *OkidoriMarketplaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OkidoriMarketplaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OkidoriMarketplaceUpgraded represents a Upgraded event raised by the OkidoriMarketplace contract.
type OkidoriMarketplaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OkidoriMarketplaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OkidoriMarketplaceUpgradedIterator{contract: _OkidoriMarketplace.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OkidoriMarketplaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _OkidoriMarketplace.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OkidoriMarketplaceUpgraded)
				if err := _OkidoriMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_OkidoriMarketplace *OkidoriMarketplaceFilterer) ParseUpgraded(log types.Log) (*OkidoriMarketplaceUpgraded, error) {
	event := new(OkidoriMarketplaceUpgraded)
	if err := _OkidoriMarketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
