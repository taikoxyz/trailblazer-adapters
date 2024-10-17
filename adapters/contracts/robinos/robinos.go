// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package robinos

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

// UserSideBetData is an auto generated low-level Go binding around an user-defined struct.
type UserSideBetData struct {
	EventCode           string
	TokenAddress        common.Address
	TeamNames           [2]string
	WinnerSet           bool
	WinningIndex        uint8
	UserTokensDeposited [2]*big.Int
	UserReward          *big.Int
}

// RobinosMetaData contains all meta data concerning the Robinos contract.
var RobinosMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTeamIndex\",\"name\":\"teamIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"NewUniqueWallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerCut\",\"type\":\"uint256\"}],\"name\":\"OwnerCutTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"userRefunds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"RefundDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"userRewards\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"winningUsers\",\"type\":\"address[]\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"SideBetCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"teamA\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"teamB\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"standardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"SideBetEventInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumTeamIndex\",\"name\":\"teamIndex\",\"type\":\"uint8\"}],\"name\":\"WinningTeamSelected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"calculateTotalRewardAndOwnerCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ownerCut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"cancelSideBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"enumTeamIndex\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"endSaleNow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"getAllUniqueWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"getSideBetData\",\"outputs\":[{\"internalType\":\"string[2]\",\"name\":\"teamNames\",\"type\":\"string[2]\"},{\"internalType\":\"address\",\"name\":\"standardTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"winnerSet\",\"type\":\"bool\"},{\"internalType\":\"enumTeamIndex\",\"name\":\"winningIndex\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"ownerCutWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"maxUsers\",\"type\":\"uint256\"}],\"name\":\"getSideBetDepositData\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"allUsers\",\"type\":\"address[]\"},{\"internalType\":\"address[][2]\",\"name\":\"eventUsers\",\"type\":\"address[][2]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"userTokens\",\"type\":\"uint256[][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"totalTokensDeposited\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTotalTokenSpent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUniqueWalletCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSideBets\",\"type\":\"uint256\"}],\"name\":\"getUserSideBetData\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string[2]\",\"name\":\"teamNames\",\"type\":\"string[2]\"},{\"internalType\":\"bool\",\"name\":\"winnerSet\",\"type\":\"bool\"},{\"internalType\":\"enumTeamIndex\",\"name\":\"winningIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256[2]\",\"name\":\"userTokensDeposited\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"userReward\",\"type\":\"uint256\"}],\"internalType\":\"structUserSideBetData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserSideBets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTokensDeposited\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersRewardsClaimedStatuses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"usersRewardsClaimed\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getWinningUsersAndUserRewards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"winningUsers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"userRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"teamNameA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"teamNameB\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"standardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"}],\"name\":\"initializeSideBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"isSaleOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"saleActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"refundTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"enumTeamIndex\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"selectWinningTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"}],\"name\":\"transferOwnerCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventCode\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"updateSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RobinosABI is the input ABI used to generate the binding from.
// Deprecated: Use RobinosMetaData.ABI instead.
var RobinosABI = RobinosMetaData.ABI

// Robinos is an auto generated Go binding around an Ethereum contract.
type Robinos struct {
	RobinosCaller     // Read-only binding to the contract
	RobinosTransactor // Write-only binding to the contract
	RobinosFilterer   // Log filterer for contract events
}

// RobinosCaller is an auto generated read-only Go binding around an Ethereum contract.
type RobinosCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobinosTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RobinosTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobinosFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RobinosFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RobinosSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RobinosSession struct {
	Contract     *Robinos          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RobinosCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RobinosCallerSession struct {
	Contract *RobinosCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RobinosTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RobinosTransactorSession struct {
	Contract     *RobinosTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RobinosRaw is an auto generated low-level Go binding around an Ethereum contract.
type RobinosRaw struct {
	Contract *Robinos // Generic contract binding to access the raw methods on
}

// RobinosCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RobinosCallerRaw struct {
	Contract *RobinosCaller // Generic read-only contract binding to access the raw methods on
}

// RobinosTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RobinosTransactorRaw struct {
	Contract *RobinosTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRobinos creates a new instance of Robinos, bound to a specific deployed contract.
func NewRobinos(address common.Address, backend bind.ContractBackend) (*Robinos, error) {
	contract, err := bindRobinos(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Robinos{RobinosCaller: RobinosCaller{contract: contract}, RobinosTransactor: RobinosTransactor{contract: contract}, RobinosFilterer: RobinosFilterer{contract: contract}}, nil
}

// NewRobinosCaller creates a new read-only instance of Robinos, bound to a specific deployed contract.
func NewRobinosCaller(address common.Address, caller bind.ContractCaller) (*RobinosCaller, error) {
	contract, err := bindRobinos(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RobinosCaller{contract: contract}, nil
}

// NewRobinosTransactor creates a new write-only instance of Robinos, bound to a specific deployed contract.
func NewRobinosTransactor(address common.Address, transactor bind.ContractTransactor) (*RobinosTransactor, error) {
	contract, err := bindRobinos(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RobinosTransactor{contract: contract}, nil
}

// NewRobinosFilterer creates a new log filterer instance of Robinos, bound to a specific deployed contract.
func NewRobinosFilterer(address common.Address, filterer bind.ContractFilterer) (*RobinosFilterer, error) {
	contract, err := bindRobinos(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RobinosFilterer{contract: contract}, nil
}

// bindRobinos binds a generic wrapper to an already deployed contract.
func bindRobinos(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RobinosMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Robinos *RobinosRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Robinos.Contract.RobinosCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Robinos *RobinosRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Robinos.Contract.RobinosTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Robinos *RobinosRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Robinos.Contract.RobinosTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Robinos *RobinosCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Robinos.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Robinos *RobinosTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Robinos.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Robinos *RobinosTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Robinos.Contract.contract.Transact(opts, method, params...)
}

// CalculateTotalRewardAndOwnerCut is a free data retrieval call binding the contract method 0xf1027842.
//
// Solidity: function calculateTotalRewardAndOwnerCut(string eventCode) view returns(uint256 totalReward, uint256 ownerCut)
func (_Robinos *RobinosCaller) CalculateTotalRewardAndOwnerCut(opts *bind.CallOpts, eventCode string) (struct {
	TotalReward *big.Int
	OwnerCut    *big.Int
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "calculateTotalRewardAndOwnerCut", eventCode)

	outstruct := new(struct {
		TotalReward *big.Int
		OwnerCut    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OwnerCut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateTotalRewardAndOwnerCut is a free data retrieval call binding the contract method 0xf1027842.
//
// Solidity: function calculateTotalRewardAndOwnerCut(string eventCode) view returns(uint256 totalReward, uint256 ownerCut)
func (_Robinos *RobinosSession) CalculateTotalRewardAndOwnerCut(eventCode string) (struct {
	TotalReward *big.Int
	OwnerCut    *big.Int
}, error) {
	return _Robinos.Contract.CalculateTotalRewardAndOwnerCut(&_Robinos.CallOpts, eventCode)
}

// CalculateTotalRewardAndOwnerCut is a free data retrieval call binding the contract method 0xf1027842.
//
// Solidity: function calculateTotalRewardAndOwnerCut(string eventCode) view returns(uint256 totalReward, uint256 ownerCut)
func (_Robinos *RobinosCallerSession) CalculateTotalRewardAndOwnerCut(eventCode string) (struct {
	TotalReward *big.Int
	OwnerCut    *big.Int
}, error) {
	return _Robinos.Contract.CalculateTotalRewardAndOwnerCut(&_Robinos.CallOpts, eventCode)
}

// GetAllUniqueWallets is a free data retrieval call binding the contract method 0x3a65f83d.
//
// Solidity: function getAllUniqueWallets(uint256 maxCount) view returns(address[])
func (_Robinos *RobinosCaller) GetAllUniqueWallets(opts *bind.CallOpts, maxCount *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getAllUniqueWallets", maxCount)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUniqueWallets is a free data retrieval call binding the contract method 0x3a65f83d.
//
// Solidity: function getAllUniqueWallets(uint256 maxCount) view returns(address[])
func (_Robinos *RobinosSession) GetAllUniqueWallets(maxCount *big.Int) ([]common.Address, error) {
	return _Robinos.Contract.GetAllUniqueWallets(&_Robinos.CallOpts, maxCount)
}

// GetAllUniqueWallets is a free data retrieval call binding the contract method 0x3a65f83d.
//
// Solidity: function getAllUniqueWallets(uint256 maxCount) view returns(address[])
func (_Robinos *RobinosCallerSession) GetAllUniqueWallets(maxCount *big.Int) ([]common.Address, error) {
	return _Robinos.Contract.GetAllUniqueWallets(&_Robinos.CallOpts, maxCount)
}

// GetSideBetData is a free data retrieval call binding the contract method 0xf0d503de.
//
// Solidity: function getSideBetData(string eventCode) view returns(string[2] teamNames, address standardTokenAddress, bool winnerSet, uint8 winningIndex, bool ownerCutWithdrawn, bool cancelled)
func (_Robinos *RobinosCaller) GetSideBetData(opts *bind.CallOpts, eventCode string) (struct {
	TeamNames            [2]string
	StandardTokenAddress common.Address
	WinnerSet            bool
	WinningIndex         uint8
	OwnerCutWithdrawn    bool
	Cancelled            bool
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getSideBetData", eventCode)

	outstruct := new(struct {
		TeamNames            [2]string
		StandardTokenAddress common.Address
		WinnerSet            bool
		WinningIndex         uint8
		OwnerCutWithdrawn    bool
		Cancelled            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeamNames = *abi.ConvertType(out[0], new([2]string)).(*[2]string)
	outstruct.StandardTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.WinnerSet = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.WinningIndex = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.OwnerCutWithdrawn = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Cancelled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetSideBetData is a free data retrieval call binding the contract method 0xf0d503de.
//
// Solidity: function getSideBetData(string eventCode) view returns(string[2] teamNames, address standardTokenAddress, bool winnerSet, uint8 winningIndex, bool ownerCutWithdrawn, bool cancelled)
func (_Robinos *RobinosSession) GetSideBetData(eventCode string) (struct {
	TeamNames            [2]string
	StandardTokenAddress common.Address
	WinnerSet            bool
	WinningIndex         uint8
	OwnerCutWithdrawn    bool
	Cancelled            bool
}, error) {
	return _Robinos.Contract.GetSideBetData(&_Robinos.CallOpts, eventCode)
}

// GetSideBetData is a free data retrieval call binding the contract method 0xf0d503de.
//
// Solidity: function getSideBetData(string eventCode) view returns(string[2] teamNames, address standardTokenAddress, bool winnerSet, uint8 winningIndex, bool ownerCutWithdrawn, bool cancelled)
func (_Robinos *RobinosCallerSession) GetSideBetData(eventCode string) (struct {
	TeamNames            [2]string
	StandardTokenAddress common.Address
	WinnerSet            bool
	WinningIndex         uint8
	OwnerCutWithdrawn    bool
	Cancelled            bool
}, error) {
	return _Robinos.Contract.GetSideBetData(&_Robinos.CallOpts, eventCode)
}

// GetSideBetDepositData is a free data retrieval call binding the contract method 0xf862945c.
//
// Solidity: function getSideBetDepositData(string eventCode, uint256 maxUsers) view returns(address[] allUsers, address[][2] eventUsers, uint256[][2] userTokens, uint256[2] totalTokensDeposited)
func (_Robinos *RobinosCaller) GetSideBetDepositData(opts *bind.CallOpts, eventCode string, maxUsers *big.Int) (struct {
	AllUsers             []common.Address
	EventUsers           [2][]common.Address
	UserTokens           [2][]*big.Int
	TotalTokensDeposited [2]*big.Int
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getSideBetDepositData", eventCode, maxUsers)

	outstruct := new(struct {
		AllUsers             []common.Address
		EventUsers           [2][]common.Address
		UserTokens           [2][]*big.Int
		TotalTokensDeposited [2]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllUsers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.EventUsers = *abi.ConvertType(out[1], new([2][]common.Address)).(*[2][]common.Address)
	outstruct.UserTokens = *abi.ConvertType(out[2], new([2][]*big.Int)).(*[2][]*big.Int)
	outstruct.TotalTokensDeposited = *abi.ConvertType(out[3], new([2]*big.Int)).(*[2]*big.Int)

	return *outstruct, err

}

// GetSideBetDepositData is a free data retrieval call binding the contract method 0xf862945c.
//
// Solidity: function getSideBetDepositData(string eventCode, uint256 maxUsers) view returns(address[] allUsers, address[][2] eventUsers, uint256[][2] userTokens, uint256[2] totalTokensDeposited)
func (_Robinos *RobinosSession) GetSideBetDepositData(eventCode string, maxUsers *big.Int) (struct {
	AllUsers             []common.Address
	EventUsers           [2][]common.Address
	UserTokens           [2][]*big.Int
	TotalTokensDeposited [2]*big.Int
}, error) {
	return _Robinos.Contract.GetSideBetDepositData(&_Robinos.CallOpts, eventCode, maxUsers)
}

// GetSideBetDepositData is a free data retrieval call binding the contract method 0xf862945c.
//
// Solidity: function getSideBetDepositData(string eventCode, uint256 maxUsers) view returns(address[] allUsers, address[][2] eventUsers, uint256[][2] userTokens, uint256[2] totalTokensDeposited)
func (_Robinos *RobinosCallerSession) GetSideBetDepositData(eventCode string, maxUsers *big.Int) (struct {
	AllUsers             []common.Address
	EventUsers           [2][]common.Address
	UserTokens           [2][]*big.Int
	TotalTokensDeposited [2]*big.Int
}, error) {
	return _Robinos.Contract.GetSideBetDepositData(&_Robinos.CallOpts, eventCode, maxUsers)
}

// GetTotalTokenSpent is a free data retrieval call binding the contract method 0x236d3972.
//
// Solidity: function getTotalTokenSpent(address tokenAddress) view returns(uint256)
func (_Robinos *RobinosCaller) GetTotalTokenSpent(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getTotalTokenSpent", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalTokenSpent is a free data retrieval call binding the contract method 0x236d3972.
//
// Solidity: function getTotalTokenSpent(address tokenAddress) view returns(uint256)
func (_Robinos *RobinosSession) GetTotalTokenSpent(tokenAddress common.Address) (*big.Int, error) {
	return _Robinos.Contract.GetTotalTokenSpent(&_Robinos.CallOpts, tokenAddress)
}

// GetTotalTokenSpent is a free data retrieval call binding the contract method 0x236d3972.
//
// Solidity: function getTotalTokenSpent(address tokenAddress) view returns(uint256)
func (_Robinos *RobinosCallerSession) GetTotalTokenSpent(tokenAddress common.Address) (*big.Int, error) {
	return _Robinos.Contract.GetTotalTokenSpent(&_Robinos.CallOpts, tokenAddress)
}

// GetUniqueWalletCount is a free data retrieval call binding the contract method 0x2ce934b2.
//
// Solidity: function getUniqueWalletCount() view returns(uint256)
func (_Robinos *RobinosCaller) GetUniqueWalletCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getUniqueWalletCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUniqueWalletCount is a free data retrieval call binding the contract method 0x2ce934b2.
//
// Solidity: function getUniqueWalletCount() view returns(uint256)
func (_Robinos *RobinosSession) GetUniqueWalletCount() (*big.Int, error) {
	return _Robinos.Contract.GetUniqueWalletCount(&_Robinos.CallOpts)
}

// GetUniqueWalletCount is a free data retrieval call binding the contract method 0x2ce934b2.
//
// Solidity: function getUniqueWalletCount() view returns(uint256)
func (_Robinos *RobinosCallerSession) GetUniqueWalletCount() (*big.Int, error) {
	return _Robinos.Contract.GetUniqueWalletCount(&_Robinos.CallOpts)
}

// GetUserSideBetData is a free data retrieval call binding the contract method 0x8126e893.
//
// Solidity: function getUserSideBetData(address user, uint256 maxSideBets) view returns((string,address,string[2],bool,uint8,uint256[2],uint256)[])
func (_Robinos *RobinosCaller) GetUserSideBetData(opts *bind.CallOpts, user common.Address, maxSideBets *big.Int) ([]UserSideBetData, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getUserSideBetData", user, maxSideBets)

	if err != nil {
		return *new([]UserSideBetData), err
	}

	out0 := *abi.ConvertType(out[0], new([]UserSideBetData)).(*[]UserSideBetData)

	return out0, err

}

// GetUserSideBetData is a free data retrieval call binding the contract method 0x8126e893.
//
// Solidity: function getUserSideBetData(address user, uint256 maxSideBets) view returns((string,address,string[2],bool,uint8,uint256[2],uint256)[])
func (_Robinos *RobinosSession) GetUserSideBetData(user common.Address, maxSideBets *big.Int) ([]UserSideBetData, error) {
	return _Robinos.Contract.GetUserSideBetData(&_Robinos.CallOpts, user, maxSideBets)
}

// GetUserSideBetData is a free data retrieval call binding the contract method 0x8126e893.
//
// Solidity: function getUserSideBetData(address user, uint256 maxSideBets) view returns((string,address,string[2],bool,uint8,uint256[2],uint256)[])
func (_Robinos *RobinosCallerSession) GetUserSideBetData(user common.Address, maxSideBets *big.Int) ([]UserSideBetData, error) {
	return _Robinos.Contract.GetUserSideBetData(&_Robinos.CallOpts, user, maxSideBets)
}

// GetUserSideBets is a free data retrieval call binding the contract method 0x85bc600f.
//
// Solidity: function getUserSideBets(address user) view returns(bytes32[])
func (_Robinos *RobinosCaller) GetUserSideBets(opts *bind.CallOpts, user common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getUserSideBets", user)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetUserSideBets is a free data retrieval call binding the contract method 0x85bc600f.
//
// Solidity: function getUserSideBets(address user) view returns(bytes32[])
func (_Robinos *RobinosSession) GetUserSideBets(user common.Address) ([][32]byte, error) {
	return _Robinos.Contract.GetUserSideBets(&_Robinos.CallOpts, user)
}

// GetUserSideBets is a free data retrieval call binding the contract method 0x85bc600f.
//
// Solidity: function getUserSideBets(address user) view returns(bytes32[])
func (_Robinos *RobinosCallerSession) GetUserSideBets(user common.Address) ([][32]byte, error) {
	return _Robinos.Contract.GetUserSideBets(&_Robinos.CallOpts, user)
}

// GetUserTokensDeposited is a free data retrieval call binding the contract method 0xff521b1b.
//
// Solidity: function getUserTokensDeposited(string eventCode, address user) view returns(uint256[2])
func (_Robinos *RobinosCaller) GetUserTokensDeposited(opts *bind.CallOpts, eventCode string, user common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getUserTokensDeposited", eventCode, user)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetUserTokensDeposited is a free data retrieval call binding the contract method 0xff521b1b.
//
// Solidity: function getUserTokensDeposited(string eventCode, address user) view returns(uint256[2])
func (_Robinos *RobinosSession) GetUserTokensDeposited(eventCode string, user common.Address) ([2]*big.Int, error) {
	return _Robinos.Contract.GetUserTokensDeposited(&_Robinos.CallOpts, eventCode, user)
}

// GetUserTokensDeposited is a free data retrieval call binding the contract method 0xff521b1b.
//
// Solidity: function getUserTokensDeposited(string eventCode, address user) view returns(uint256[2])
func (_Robinos *RobinosCallerSession) GetUserTokensDeposited(eventCode string, user common.Address) ([2]*big.Int, error) {
	return _Robinos.Contract.GetUserTokensDeposited(&_Robinos.CallOpts, eventCode, user)
}

// GetUsersRewardsClaimedStatuses is a free data retrieval call binding the contract method 0xcf2fe187.
//
// Solidity: function getUsersRewardsClaimedStatuses(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] users, bool[] usersRewardsClaimed)
func (_Robinos *RobinosCaller) GetUsersRewardsClaimedStatuses(opts *bind.CallOpts, eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	Users               []common.Address
	UsersRewardsClaimed []bool
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getUsersRewardsClaimedStatuses", eventCode, startIndex, endIndex)

	outstruct := new(struct {
		Users               []common.Address
		UsersRewardsClaimed []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Users = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UsersRewardsClaimed = *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetUsersRewardsClaimedStatuses is a free data retrieval call binding the contract method 0xcf2fe187.
//
// Solidity: function getUsersRewardsClaimedStatuses(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] users, bool[] usersRewardsClaimed)
func (_Robinos *RobinosSession) GetUsersRewardsClaimedStatuses(eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	Users               []common.Address
	UsersRewardsClaimed []bool
}, error) {
	return _Robinos.Contract.GetUsersRewardsClaimedStatuses(&_Robinos.CallOpts, eventCode, startIndex, endIndex)
}

// GetUsersRewardsClaimedStatuses is a free data retrieval call binding the contract method 0xcf2fe187.
//
// Solidity: function getUsersRewardsClaimedStatuses(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] users, bool[] usersRewardsClaimed)
func (_Robinos *RobinosCallerSession) GetUsersRewardsClaimedStatuses(eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	Users               []common.Address
	UsersRewardsClaimed []bool
}, error) {
	return _Robinos.Contract.GetUsersRewardsClaimedStatuses(&_Robinos.CallOpts, eventCode, startIndex, endIndex)
}

// GetWinningUsersAndUserRewards is a free data retrieval call binding the contract method 0x6c3d200a.
//
// Solidity: function getWinningUsersAndUserRewards(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] winningUsers, uint256[] userRewards)
func (_Robinos *RobinosCaller) GetWinningUsersAndUserRewards(opts *bind.CallOpts, eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	WinningUsers []common.Address
	UserRewards  []*big.Int
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "getWinningUsersAndUserRewards", eventCode, startIndex, endIndex)

	outstruct := new(struct {
		WinningUsers []common.Address
		UserRewards  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WinningUsers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.UserRewards = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetWinningUsersAndUserRewards is a free data retrieval call binding the contract method 0x6c3d200a.
//
// Solidity: function getWinningUsersAndUserRewards(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] winningUsers, uint256[] userRewards)
func (_Robinos *RobinosSession) GetWinningUsersAndUserRewards(eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	WinningUsers []common.Address
	UserRewards  []*big.Int
}, error) {
	return _Robinos.Contract.GetWinningUsersAndUserRewards(&_Robinos.CallOpts, eventCode, startIndex, endIndex)
}

// GetWinningUsersAndUserRewards is a free data retrieval call binding the contract method 0x6c3d200a.
//
// Solidity: function getWinningUsersAndUserRewards(string eventCode, uint256 startIndex, uint256 endIndex) view returns(address[] winningUsers, uint256[] userRewards)
func (_Robinos *RobinosCallerSession) GetWinningUsersAndUserRewards(eventCode string, startIndex *big.Int, endIndex *big.Int) (struct {
	WinningUsers []common.Address
	UserRewards  []*big.Int
}, error) {
	return _Robinos.Contract.GetWinningUsersAndUserRewards(&_Robinos.CallOpts, eventCode, startIndex, endIndex)
}

// IsSaleOn is a free data retrieval call binding the contract method 0x27025b36.
//
// Solidity: function isSaleOn(string eventCode) view returns(bool saleActive, uint256 saleStart, uint256 saleEnd)
func (_Robinos *RobinosCaller) IsSaleOn(opts *bind.CallOpts, eventCode string) (struct {
	SaleActive bool
	SaleStart  *big.Int
	SaleEnd    *big.Int
}, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "isSaleOn", eventCode)

	outstruct := new(struct {
		SaleActive bool
		SaleStart  *big.Int
		SaleEnd    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SaleActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SaleStart = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SaleEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IsSaleOn is a free data retrieval call binding the contract method 0x27025b36.
//
// Solidity: function isSaleOn(string eventCode) view returns(bool saleActive, uint256 saleStart, uint256 saleEnd)
func (_Robinos *RobinosSession) IsSaleOn(eventCode string) (struct {
	SaleActive bool
	SaleStart  *big.Int
	SaleEnd    *big.Int
}, error) {
	return _Robinos.Contract.IsSaleOn(&_Robinos.CallOpts, eventCode)
}

// IsSaleOn is a free data retrieval call binding the contract method 0x27025b36.
//
// Solidity: function isSaleOn(string eventCode) view returns(bool saleActive, uint256 saleStart, uint256 saleEnd)
func (_Robinos *RobinosCallerSession) IsSaleOn(eventCode string) (struct {
	SaleActive bool
	SaleStart  *big.Int
	SaleEnd    *big.Int
}, error) {
	return _Robinos.Contract.IsSaleOn(&_Robinos.CallOpts, eventCode)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Robinos *RobinosCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Robinos.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Robinos *RobinosSession) Owner() (common.Address, error) {
	return _Robinos.Contract.Owner(&_Robinos.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Robinos *RobinosCallerSession) Owner() (common.Address, error) {
	return _Robinos.Contract.Owner(&_Robinos.CallOpts)
}

// CancelSideBet is a paid mutator transaction binding the contract method 0x7ba36f6b.
//
// Solidity: function cancelSideBet(string eventCode) returns()
func (_Robinos *RobinosTransactor) CancelSideBet(opts *bind.TransactOpts, eventCode string) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "cancelSideBet", eventCode)
}

// CancelSideBet is a paid mutator transaction binding the contract method 0x7ba36f6b.
//
// Solidity: function cancelSideBet(string eventCode) returns()
func (_Robinos *RobinosSession) CancelSideBet(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.CancelSideBet(&_Robinos.TransactOpts, eventCode)
}

// CancelSideBet is a paid mutator transaction binding the contract method 0x7ba36f6b.
//
// Solidity: function cancelSideBet(string eventCode) returns()
func (_Robinos *RobinosTransactorSession) CancelSideBet(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.CancelSideBet(&_Robinos.TransactOpts, eventCode)
}

// Deposit is a paid mutator transaction binding the contract method 0x07a795c9.
//
// Solidity: function deposit(string eventCode, uint8 index, uint256 amount) returns()
func (_Robinos *RobinosTransactor) Deposit(opts *bind.TransactOpts, eventCode string, index uint8, amount *big.Int) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "deposit", eventCode, index, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x07a795c9.
//
// Solidity: function deposit(string eventCode, uint8 index, uint256 amount) returns()
func (_Robinos *RobinosSession) Deposit(eventCode string, index uint8, amount *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.Deposit(&_Robinos.TransactOpts, eventCode, index, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x07a795c9.
//
// Solidity: function deposit(string eventCode, uint8 index, uint256 amount) returns()
func (_Robinos *RobinosTransactorSession) Deposit(eventCode string, index uint8, amount *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.Deposit(&_Robinos.TransactOpts, eventCode, index, amount)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x4cbb4c5b.
//
// Solidity: function distributeReward(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosTransactor) DistributeReward(opts *bind.TransactOpts, eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "distributeReward", eventCode, startIndex, endIndex)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x4cbb4c5b.
//
// Solidity: function distributeReward(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosSession) DistributeReward(eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.DistributeReward(&_Robinos.TransactOpts, eventCode, startIndex, endIndex)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x4cbb4c5b.
//
// Solidity: function distributeReward(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosTransactorSession) DistributeReward(eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.DistributeReward(&_Robinos.TransactOpts, eventCode, startIndex, endIndex)
}

// EndSaleNow is a paid mutator transaction binding the contract method 0x517e640a.
//
// Solidity: function endSaleNow(string eventCode) returns()
func (_Robinos *RobinosTransactor) EndSaleNow(opts *bind.TransactOpts, eventCode string) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "endSaleNow", eventCode)
}

// EndSaleNow is a paid mutator transaction binding the contract method 0x517e640a.
//
// Solidity: function endSaleNow(string eventCode) returns()
func (_Robinos *RobinosSession) EndSaleNow(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.EndSaleNow(&_Robinos.TransactOpts, eventCode)
}

// EndSaleNow is a paid mutator transaction binding the contract method 0x517e640a.
//
// Solidity: function endSaleNow(string eventCode) returns()
func (_Robinos *RobinosTransactorSession) EndSaleNow(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.EndSaleNow(&_Robinos.TransactOpts, eventCode)
}

// InitializeSideBet is a paid mutator transaction binding the contract method 0x6c9c77d0.
//
// Solidity: function initializeSideBet(string eventCode, string teamNameA, string teamNameB, address standardToken, uint256 saleStart, uint256 saleEnd) returns()
func (_Robinos *RobinosTransactor) InitializeSideBet(opts *bind.TransactOpts, eventCode string, teamNameA string, teamNameB string, standardToken common.Address, saleStart *big.Int, saleEnd *big.Int) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "initializeSideBet", eventCode, teamNameA, teamNameB, standardToken, saleStart, saleEnd)
}

// InitializeSideBet is a paid mutator transaction binding the contract method 0x6c9c77d0.
//
// Solidity: function initializeSideBet(string eventCode, string teamNameA, string teamNameB, address standardToken, uint256 saleStart, uint256 saleEnd) returns()
func (_Robinos *RobinosSession) InitializeSideBet(eventCode string, teamNameA string, teamNameB string, standardToken common.Address, saleStart *big.Int, saleEnd *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.InitializeSideBet(&_Robinos.TransactOpts, eventCode, teamNameA, teamNameB, standardToken, saleStart, saleEnd)
}

// InitializeSideBet is a paid mutator transaction binding the contract method 0x6c9c77d0.
//
// Solidity: function initializeSideBet(string eventCode, string teamNameA, string teamNameB, address standardToken, uint256 saleStart, uint256 saleEnd) returns()
func (_Robinos *RobinosTransactorSession) InitializeSideBet(eventCode string, teamNameA string, teamNameB string, standardToken common.Address, saleStart *big.Int, saleEnd *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.InitializeSideBet(&_Robinos.TransactOpts, eventCode, teamNameA, teamNameB, standardToken, saleStart, saleEnd)
}

// RefundTokens is a paid mutator transaction binding the contract method 0x7c973bfd.
//
// Solidity: function refundTokens(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosTransactor) RefundTokens(opts *bind.TransactOpts, eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "refundTokens", eventCode, startIndex, endIndex)
}

// RefundTokens is a paid mutator transaction binding the contract method 0x7c973bfd.
//
// Solidity: function refundTokens(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosSession) RefundTokens(eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.RefundTokens(&_Robinos.TransactOpts, eventCode, startIndex, endIndex)
}

// RefundTokens is a paid mutator transaction binding the contract method 0x7c973bfd.
//
// Solidity: function refundTokens(string eventCode, uint256 startIndex, uint256 endIndex) returns()
func (_Robinos *RobinosTransactorSession) RefundTokens(eventCode string, startIndex *big.Int, endIndex *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.RefundTokens(&_Robinos.TransactOpts, eventCode, startIndex, endIndex)
}

// SelectWinningTeam is a paid mutator transaction binding the contract method 0xb7e94b2d.
//
// Solidity: function selectWinningTeam(string eventCode, uint8 index) returns()
func (_Robinos *RobinosTransactor) SelectWinningTeam(opts *bind.TransactOpts, eventCode string, index uint8) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "selectWinningTeam", eventCode, index)
}

// SelectWinningTeam is a paid mutator transaction binding the contract method 0xb7e94b2d.
//
// Solidity: function selectWinningTeam(string eventCode, uint8 index) returns()
func (_Robinos *RobinosSession) SelectWinningTeam(eventCode string, index uint8) (*types.Transaction, error) {
	return _Robinos.Contract.SelectWinningTeam(&_Robinos.TransactOpts, eventCode, index)
}

// SelectWinningTeam is a paid mutator transaction binding the contract method 0xb7e94b2d.
//
// Solidity: function selectWinningTeam(string eventCode, uint8 index) returns()
func (_Robinos *RobinosTransactorSession) SelectWinningTeam(eventCode string, index uint8) (*types.Transaction, error) {
	return _Robinos.Contract.SelectWinningTeam(&_Robinos.TransactOpts, eventCode, index)
}

// TransferOwnerCut is a paid mutator transaction binding the contract method 0xf04d9344.
//
// Solidity: function transferOwnerCut(string eventCode) returns()
func (_Robinos *RobinosTransactor) TransferOwnerCut(opts *bind.TransactOpts, eventCode string) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "transferOwnerCut", eventCode)
}

// TransferOwnerCut is a paid mutator transaction binding the contract method 0xf04d9344.
//
// Solidity: function transferOwnerCut(string eventCode) returns()
func (_Robinos *RobinosSession) TransferOwnerCut(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.TransferOwnerCut(&_Robinos.TransactOpts, eventCode)
}

// TransferOwnerCut is a paid mutator transaction binding the contract method 0xf04d9344.
//
// Solidity: function transferOwnerCut(string eventCode) returns()
func (_Robinos *RobinosTransactorSession) TransferOwnerCut(eventCode string) (*types.Transaction, error) {
	return _Robinos.Contract.TransferOwnerCut(&_Robinos.TransactOpts, eventCode)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Robinos *RobinosTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Robinos *RobinosSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Robinos.Contract.TransferOwnership(&_Robinos.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Robinos *RobinosTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Robinos.Contract.TransferOwnership(&_Robinos.TransactOpts, newOwner)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x88b37782.
//
// Solidity: function updateSale(string eventCode, uint256 start, uint256 end) returns()
func (_Robinos *RobinosTransactor) UpdateSale(opts *bind.TransactOpts, eventCode string, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Robinos.contract.Transact(opts, "updateSale", eventCode, start, end)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x88b37782.
//
// Solidity: function updateSale(string eventCode, uint256 start, uint256 end) returns()
func (_Robinos *RobinosSession) UpdateSale(eventCode string, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.UpdateSale(&_Robinos.TransactOpts, eventCode, start, end)
}

// UpdateSale is a paid mutator transaction binding the contract method 0x88b37782.
//
// Solidity: function updateSale(string eventCode, uint256 start, uint256 end) returns()
func (_Robinos *RobinosTransactorSession) UpdateSale(eventCode string, start *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Robinos.Contract.UpdateSale(&_Robinos.TransactOpts, eventCode, start, end)
}

// RobinosDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Robinos contract.
type RobinosDepositedIterator struct {
	Event *RobinosDeposited // Event containing the contract specifics and raw log

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
func (it *RobinosDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosDeposited)
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
		it.Event = new(RobinosDeposited)
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
func (it *RobinosDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosDeposited represents a Deposited event raised by the Robinos contract.
type RobinosDeposited struct {
	EventCode string
	Amount    *big.Int
	TeamIndex uint8
	From      common.Address
	Token     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xdb9295926151b770fbfa6f8cad5487026d4729f42e86f88fb3179110a72203f2.
//
// Solidity: event Deposited(string eventCode, uint256 amount, uint8 teamIndex, address from, address token)
func (_Robinos *RobinosFilterer) FilterDeposited(opts *bind.FilterOpts) (*RobinosDepositedIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &RobinosDepositedIterator{contract: _Robinos.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xdb9295926151b770fbfa6f8cad5487026d4729f42e86f88fb3179110a72203f2.
//
// Solidity: event Deposited(string eventCode, uint256 amount, uint8 teamIndex, address from, address token)
func (_Robinos *RobinosFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RobinosDeposited) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosDeposited)
				if err := _Robinos.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xdb9295926151b770fbfa6f8cad5487026d4729f42e86f88fb3179110a72203f2.
//
// Solidity: event Deposited(string eventCode, uint256 amount, uint8 teamIndex, address from, address token)
func (_Robinos *RobinosFilterer) ParseDeposited(log types.Log) (*RobinosDeposited, error) {
	event := new(RobinosDeposited)
	if err := _Robinos.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosNewUniqueWalletIterator is returned from FilterNewUniqueWallet and is used to iterate over the raw logs and unpacked data for NewUniqueWallet events raised by the Robinos contract.
type RobinosNewUniqueWalletIterator struct {
	Event *RobinosNewUniqueWallet // Event containing the contract specifics and raw log

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
func (it *RobinosNewUniqueWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosNewUniqueWallet)
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
		it.Event = new(RobinosNewUniqueWallet)
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
func (it *RobinosNewUniqueWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosNewUniqueWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosNewUniqueWallet represents a NewUniqueWallet event raised by the Robinos contract.
type RobinosNewUniqueWallet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewUniqueWallet is a free log retrieval operation binding the contract event 0x0fab3c30cf8d3339a1fa27828814e526eb0a049f7dd6c9647f21644dcaa99715.
//
// Solidity: event NewUniqueWallet(address wallet)
func (_Robinos *RobinosFilterer) FilterNewUniqueWallet(opts *bind.FilterOpts) (*RobinosNewUniqueWalletIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "NewUniqueWallet")
	if err != nil {
		return nil, err
	}
	return &RobinosNewUniqueWalletIterator{contract: _Robinos.contract, event: "NewUniqueWallet", logs: logs, sub: sub}, nil
}

// WatchNewUniqueWallet is a free log subscription operation binding the contract event 0x0fab3c30cf8d3339a1fa27828814e526eb0a049f7dd6c9647f21644dcaa99715.
//
// Solidity: event NewUniqueWallet(address wallet)
func (_Robinos *RobinosFilterer) WatchNewUniqueWallet(opts *bind.WatchOpts, sink chan<- *RobinosNewUniqueWallet) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "NewUniqueWallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosNewUniqueWallet)
				if err := _Robinos.contract.UnpackLog(event, "NewUniqueWallet", log); err != nil {
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

// ParseNewUniqueWallet is a log parse operation binding the contract event 0x0fab3c30cf8d3339a1fa27828814e526eb0a049f7dd6c9647f21644dcaa99715.
//
// Solidity: event NewUniqueWallet(address wallet)
func (_Robinos *RobinosFilterer) ParseNewUniqueWallet(log types.Log) (*RobinosNewUniqueWallet, error) {
	event := new(RobinosNewUniqueWallet)
	if err := _Robinos.contract.UnpackLog(event, "NewUniqueWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosOwnerCutTransferredIterator is returned from FilterOwnerCutTransferred and is used to iterate over the raw logs and unpacked data for OwnerCutTransferred events raised by the Robinos contract.
type RobinosOwnerCutTransferredIterator struct {
	Event *RobinosOwnerCutTransferred // Event containing the contract specifics and raw log

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
func (it *RobinosOwnerCutTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosOwnerCutTransferred)
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
		it.Event = new(RobinosOwnerCutTransferred)
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
func (it *RobinosOwnerCutTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosOwnerCutTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosOwnerCutTransferred represents a OwnerCutTransferred event raised by the Robinos contract.
type RobinosOwnerCutTransferred struct {
	EventCode string
	OwnerCut  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerCutTransferred is a free log retrieval operation binding the contract event 0xb7f17ae9ed81fc79ccf8aad5798d946bbfb4d825d2c6bfb995a726e199fe11db.
//
// Solidity: event OwnerCutTransferred(string eventCode, uint256 ownerCut)
func (_Robinos *RobinosFilterer) FilterOwnerCutTransferred(opts *bind.FilterOpts) (*RobinosOwnerCutTransferredIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "OwnerCutTransferred")
	if err != nil {
		return nil, err
	}
	return &RobinosOwnerCutTransferredIterator{contract: _Robinos.contract, event: "OwnerCutTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerCutTransferred is a free log subscription operation binding the contract event 0xb7f17ae9ed81fc79ccf8aad5798d946bbfb4d825d2c6bfb995a726e199fe11db.
//
// Solidity: event OwnerCutTransferred(string eventCode, uint256 ownerCut)
func (_Robinos *RobinosFilterer) WatchOwnerCutTransferred(opts *bind.WatchOpts, sink chan<- *RobinosOwnerCutTransferred) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "OwnerCutTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosOwnerCutTransferred)
				if err := _Robinos.contract.UnpackLog(event, "OwnerCutTransferred", log); err != nil {
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

// ParseOwnerCutTransferred is a log parse operation binding the contract event 0xb7f17ae9ed81fc79ccf8aad5798d946bbfb4d825d2c6bfb995a726e199fe11db.
//
// Solidity: event OwnerCutTransferred(string eventCode, uint256 ownerCut)
func (_Robinos *RobinosFilterer) ParseOwnerCutTransferred(log types.Log) (*RobinosOwnerCutTransferred, error) {
	event := new(RobinosOwnerCutTransferred)
	if err := _Robinos.contract.UnpackLog(event, "OwnerCutTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Robinos contract.
type RobinosOwnershipRenouncedIterator struct {
	Event *RobinosOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RobinosOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosOwnershipRenounced)
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
		it.Event = new(RobinosOwnershipRenounced)
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
func (it *RobinosOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosOwnershipRenounced represents a OwnershipRenounced event raised by the Robinos contract.
type RobinosOwnershipRenounced struct {
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed ownerAddress)
func (_Robinos *RobinosFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, ownerAddress []common.Address) (*RobinosOwnershipRenouncedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "OwnershipRenounced", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &RobinosOwnershipRenouncedIterator{contract: _Robinos.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed ownerAddress)
func (_Robinos *RobinosFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RobinosOwnershipRenounced, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "OwnershipRenounced", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosOwnershipRenounced)
				if err := _Robinos.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed ownerAddress)
func (_Robinos *RobinosFilterer) ParseOwnershipRenounced(log types.Log) (*RobinosOwnershipRenounced, error) {
	event := new(RobinosOwnershipRenounced)
	if err := _Robinos.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Robinos contract.
type RobinosOwnershipTransferredIterator struct {
	Event *RobinosOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RobinosOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosOwnershipTransferred)
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
		it.Event = new(RobinosOwnershipTransferred)
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
func (it *RobinosOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosOwnershipTransferred represents a OwnershipTransferred event raised by the Robinos contract.
type RobinosOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Robinos *RobinosFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RobinosOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RobinosOwnershipTransferredIterator{contract: _Robinos.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Robinos *RobinosFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RobinosOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosOwnershipTransferred)
				if err := _Robinos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Robinos *RobinosFilterer) ParseOwnershipTransferred(log types.Log) (*RobinosOwnershipTransferred, error) {
	event := new(RobinosOwnershipTransferred)
	if err := _Robinos.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosRefundDistributedIterator is returned from FilterRefundDistributed and is used to iterate over the raw logs and unpacked data for RefundDistributed events raised by the Robinos contract.
type RobinosRefundDistributedIterator struct {
	Event *RobinosRefundDistributed // Event containing the contract specifics and raw log

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
func (it *RobinosRefundDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosRefundDistributed)
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
		it.Event = new(RobinosRefundDistributed)
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
func (it *RobinosRefundDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosRefundDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosRefundDistributed represents a RefundDistributed event raised by the Robinos contract.
type RobinosRefundDistributed struct {
	EventCode   string
	UserRefunds []*big.Int
	Users       []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefundDistributed is a free log retrieval operation binding the contract event 0x5ae4343104a31393a3ac4197c1c7edaf6ed637db94cd079528fa8cc940744c4b.
//
// Solidity: event RefundDistributed(string eventCode, uint256[] userRefunds, address[] users)
func (_Robinos *RobinosFilterer) FilterRefundDistributed(opts *bind.FilterOpts) (*RobinosRefundDistributedIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "RefundDistributed")
	if err != nil {
		return nil, err
	}
	return &RobinosRefundDistributedIterator{contract: _Robinos.contract, event: "RefundDistributed", logs: logs, sub: sub}, nil
}

// WatchRefundDistributed is a free log subscription operation binding the contract event 0x5ae4343104a31393a3ac4197c1c7edaf6ed637db94cd079528fa8cc940744c4b.
//
// Solidity: event RefundDistributed(string eventCode, uint256[] userRefunds, address[] users)
func (_Robinos *RobinosFilterer) WatchRefundDistributed(opts *bind.WatchOpts, sink chan<- *RobinosRefundDistributed) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "RefundDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosRefundDistributed)
				if err := _Robinos.contract.UnpackLog(event, "RefundDistributed", log); err != nil {
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

// ParseRefundDistributed is a log parse operation binding the contract event 0x5ae4343104a31393a3ac4197c1c7edaf6ed637db94cd079528fa8cc940744c4b.
//
// Solidity: event RefundDistributed(string eventCode, uint256[] userRefunds, address[] users)
func (_Robinos *RobinosFilterer) ParseRefundDistributed(log types.Log) (*RobinosRefundDistributed, error) {
	event := new(RobinosRefundDistributed)
	if err := _Robinos.contract.UnpackLog(event, "RefundDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the Robinos contract.
type RobinosRewardDistributedIterator struct {
	Event *RobinosRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RobinosRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosRewardDistributed)
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
		it.Event = new(RobinosRewardDistributed)
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
func (it *RobinosRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosRewardDistributed represents a RewardDistributed event raised by the Robinos contract.
type RobinosRewardDistributed struct {
	EventCode    string
	UserRewards  []*big.Int
	WinningUsers []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0x4e2182888aeacc7ad61d2a4243195e0db70f7db180457fe321c0be71e68ff346.
//
// Solidity: event RewardDistributed(string eventCode, uint256[] userRewards, address[] winningUsers)
func (_Robinos *RobinosFilterer) FilterRewardDistributed(opts *bind.FilterOpts) (*RobinosRewardDistributedIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "RewardDistributed")
	if err != nil {
		return nil, err
	}
	return &RobinosRewardDistributedIterator{contract: _Robinos.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0x4e2182888aeacc7ad61d2a4243195e0db70f7db180457fe321c0be71e68ff346.
//
// Solidity: event RewardDistributed(string eventCode, uint256[] userRewards, address[] winningUsers)
func (_Robinos *RobinosFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *RobinosRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "RewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosRewardDistributed)
				if err := _Robinos.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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

// ParseRewardDistributed is a log parse operation binding the contract event 0x4e2182888aeacc7ad61d2a4243195e0db70f7db180457fe321c0be71e68ff346.
//
// Solidity: event RewardDistributed(string eventCode, uint256[] userRewards, address[] winningUsers)
func (_Robinos *RobinosFilterer) ParseRewardDistributed(log types.Log) (*RobinosRewardDistributed, error) {
	event := new(RobinosRewardDistributed)
	if err := _Robinos.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosSideBetCancelledIterator is returned from FilterSideBetCancelled and is used to iterate over the raw logs and unpacked data for SideBetCancelled events raised by the Robinos contract.
type RobinosSideBetCancelledIterator struct {
	Event *RobinosSideBetCancelled // Event containing the contract specifics and raw log

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
func (it *RobinosSideBetCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosSideBetCancelled)
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
		it.Event = new(RobinosSideBetCancelled)
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
func (it *RobinosSideBetCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosSideBetCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosSideBetCancelled represents a SideBetCancelled event raised by the Robinos contract.
type RobinosSideBetCancelled struct {
	EventCode string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSideBetCancelled is a free log retrieval operation binding the contract event 0xd377c868ac5f15beec1223dcc5ea569a1fe1c7371f4704f4de58222459b674e6.
//
// Solidity: event SideBetCancelled(string eventCode)
func (_Robinos *RobinosFilterer) FilterSideBetCancelled(opts *bind.FilterOpts) (*RobinosSideBetCancelledIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "SideBetCancelled")
	if err != nil {
		return nil, err
	}
	return &RobinosSideBetCancelledIterator{contract: _Robinos.contract, event: "SideBetCancelled", logs: logs, sub: sub}, nil
}

// WatchSideBetCancelled is a free log subscription operation binding the contract event 0xd377c868ac5f15beec1223dcc5ea569a1fe1c7371f4704f4de58222459b674e6.
//
// Solidity: event SideBetCancelled(string eventCode)
func (_Robinos *RobinosFilterer) WatchSideBetCancelled(opts *bind.WatchOpts, sink chan<- *RobinosSideBetCancelled) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "SideBetCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosSideBetCancelled)
				if err := _Robinos.contract.UnpackLog(event, "SideBetCancelled", log); err != nil {
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

// ParseSideBetCancelled is a log parse operation binding the contract event 0xd377c868ac5f15beec1223dcc5ea569a1fe1c7371f4704f4de58222459b674e6.
//
// Solidity: event SideBetCancelled(string eventCode)
func (_Robinos *RobinosFilterer) ParseSideBetCancelled(log types.Log) (*RobinosSideBetCancelled, error) {
	event := new(RobinosSideBetCancelled)
	if err := _Robinos.contract.UnpackLog(event, "SideBetCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosSideBetEventInitializedIterator is returned from FilterSideBetEventInitialized and is used to iterate over the raw logs and unpacked data for SideBetEventInitialized events raised by the Robinos contract.
type RobinosSideBetEventInitializedIterator struct {
	Event *RobinosSideBetEventInitialized // Event containing the contract specifics and raw log

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
func (it *RobinosSideBetEventInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosSideBetEventInitialized)
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
		it.Event = new(RobinosSideBetEventInitialized)
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
func (it *RobinosSideBetEventInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosSideBetEventInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosSideBetEventInitialized represents a SideBetEventInitialized event raised by the Robinos contract.
type RobinosSideBetEventInitialized struct {
	EventCode     string
	TeamA         string
	TeamB         string
	StandardToken common.Address
	StartTime     *big.Int
	EndTime       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSideBetEventInitialized is a free log retrieval operation binding the contract event 0x87383e89c4c65a57b5b8f0e979b41307ac1cca497b4fe28ab47f733a18299c2f.
//
// Solidity: event SideBetEventInitialized(string eventCode, string teamA, string teamB, address standardToken, uint256 startTime, uint256 endTime)
func (_Robinos *RobinosFilterer) FilterSideBetEventInitialized(opts *bind.FilterOpts) (*RobinosSideBetEventInitializedIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "SideBetEventInitialized")
	if err != nil {
		return nil, err
	}
	return &RobinosSideBetEventInitializedIterator{contract: _Robinos.contract, event: "SideBetEventInitialized", logs: logs, sub: sub}, nil
}

// WatchSideBetEventInitialized is a free log subscription operation binding the contract event 0x87383e89c4c65a57b5b8f0e979b41307ac1cca497b4fe28ab47f733a18299c2f.
//
// Solidity: event SideBetEventInitialized(string eventCode, string teamA, string teamB, address standardToken, uint256 startTime, uint256 endTime)
func (_Robinos *RobinosFilterer) WatchSideBetEventInitialized(opts *bind.WatchOpts, sink chan<- *RobinosSideBetEventInitialized) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "SideBetEventInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosSideBetEventInitialized)
				if err := _Robinos.contract.UnpackLog(event, "SideBetEventInitialized", log); err != nil {
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

// ParseSideBetEventInitialized is a log parse operation binding the contract event 0x87383e89c4c65a57b5b8f0e979b41307ac1cca497b4fe28ab47f733a18299c2f.
//
// Solidity: event SideBetEventInitialized(string eventCode, string teamA, string teamB, address standardToken, uint256 startTime, uint256 endTime)
func (_Robinos *RobinosFilterer) ParseSideBetEventInitialized(log types.Log) (*RobinosSideBetEventInitialized, error) {
	event := new(RobinosSideBetEventInitialized)
	if err := _Robinos.contract.UnpackLog(event, "SideBetEventInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RobinosWinningTeamSelectedIterator is returned from FilterWinningTeamSelected and is used to iterate over the raw logs and unpacked data for WinningTeamSelected events raised by the Robinos contract.
type RobinosWinningTeamSelectedIterator struct {
	Event *RobinosWinningTeamSelected // Event containing the contract specifics and raw log

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
func (it *RobinosWinningTeamSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RobinosWinningTeamSelected)
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
		it.Event = new(RobinosWinningTeamSelected)
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
func (it *RobinosWinningTeamSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RobinosWinningTeamSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RobinosWinningTeamSelected represents a WinningTeamSelected event raised by the Robinos contract.
type RobinosWinningTeamSelected struct {
	EventCode string
	TeamIndex uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWinningTeamSelected is a free log retrieval operation binding the contract event 0x8f859fd2e59147ea4807f15bb2c267cfa617db02508e176b4cdabc84cf23d3ce.
//
// Solidity: event WinningTeamSelected(string eventCode, uint8 teamIndex)
func (_Robinos *RobinosFilterer) FilterWinningTeamSelected(opts *bind.FilterOpts) (*RobinosWinningTeamSelectedIterator, error) {

	logs, sub, err := _Robinos.contract.FilterLogs(opts, "WinningTeamSelected")
	if err != nil {
		return nil, err
	}
	return &RobinosWinningTeamSelectedIterator{contract: _Robinos.contract, event: "WinningTeamSelected", logs: logs, sub: sub}, nil
}

// WatchWinningTeamSelected is a free log subscription operation binding the contract event 0x8f859fd2e59147ea4807f15bb2c267cfa617db02508e176b4cdabc84cf23d3ce.
//
// Solidity: event WinningTeamSelected(string eventCode, uint8 teamIndex)
func (_Robinos *RobinosFilterer) WatchWinningTeamSelected(opts *bind.WatchOpts, sink chan<- *RobinosWinningTeamSelected) (event.Subscription, error) {

	logs, sub, err := _Robinos.contract.WatchLogs(opts, "WinningTeamSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RobinosWinningTeamSelected)
				if err := _Robinos.contract.UnpackLog(event, "WinningTeamSelected", log); err != nil {
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

// ParseWinningTeamSelected is a log parse operation binding the contract event 0x8f859fd2e59147ea4807f15bb2c267cfa617db02508e176b4cdabc84cf23d3ce.
//
// Solidity: event WinningTeamSelected(string eventCode, uint8 teamIndex)
func (_Robinos *RobinosFilterer) ParseWinningTeamSelected(log types.Log) (*RobinosWinningTeamSelected, error) {
	event := new(RobinosWinningTeamSelected)
	if err := _Robinos.contract.UnpackLog(event, "WinningTeamSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
