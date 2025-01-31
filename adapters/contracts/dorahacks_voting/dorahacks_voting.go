// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dorahacks_voting

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

// DomainObjsStateLeaf is an auto generated low-level Go binding around an user-defined struct.
type DomainObjsStateLeaf struct {
	PubKey             IPubKeyPubKey
	VoiceCreditBalance *big.Int
	VoteOptionTreeRoot *big.Int
	Nonce              *big.Int
}

// IMessageMessage is an auto generated low-level Go binding around an user-defined struct.
type IMessageMessage struct {
	Data [7]*big.Int
}

// IPubKeyPubKey is an auto generated low-level Go binding around an user-defined struct.
type IPubKeyPubKey struct {
	X *big.Int
	Y *big.Int
}

// MACIMaciParameters is an auto generated low-level Go binding around an user-defined struct.
type MACIMaciParameters struct {
	StateTreeDepth      *big.Int
	IntStateTreeDepth   *big.Int
	MessageBatchSize    *big.Int
	VoteOptionTreeDepth *big.Int
}

// DorahacksVotingMetaData contains all meta data concerning the DorahacksVoting contract.
var DorahacksVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_msgIdx\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[7]\",\"name\":\"data\",\"type\":\"uint256[7]\"}],\"indexed\":false,\"internalType\":\"structIMessage.Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_encPubKey\",\"type\":\"tuple\"}],\"name\":\"PublishMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_stateIdx\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_userPubKey\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_voiceCreditBalance\",\"type\":\"uint256\"}],\"name\":\"SignUp\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[7]\",\"name\":\"data\",\"type\":\"uint256[7]\"}],\"internalType\":\"structIMessage.Message[]\",\"name\":\"_messages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey[]\",\"name\":\"_encPubKeys\",\"type\":\"tuple[]\"}],\"name\":\"batchPublishMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coordinatorHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentStateCommitment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTallyCommitment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateKeeper\",\"outputs\":[{\"internalType\":\"contractSignUpGatekeeper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"array\",\"type\":\"uint256[2]\"}],\"name\":\"hash2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"array\",\"type\":\"uint256[5]\"}],\"name\":\"hash5\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_right\",\"type\":\"uint256\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[7]\",\"name\":\"data\",\"type\":\"uint256[7]\"}],\"internalType\":\"structIMessage.Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_encPubKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_prevHash\",\"type\":\"uint256\"}],\"name\":\"hashMessageAndEncPubKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"voiceCreditBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOptionTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDomainObjs.StateLeaf\",\"name\":\"_stateLeaf\",\"type\":\"tuple\"}],\"name\":\"hashStateLeaf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"contractVkRegistry\",\"name\":\"_vkRegistry\",\"type\":\"address\"},{\"internalType\":\"contractQuinaryTreeRoot\",\"name\":\"_qtrLib\",\"type\":\"address\"},{\"internalType\":\"contractVerifier\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"contractSignUpGatekeeper\",\"name\":\"_gateKeeper\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intStateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageBatchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOptionTreeDepth\",\"type\":\"uint256\"}],\"internalType\":\"structMACI.MaciParameters\",\"name\":\"_parameters\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_coordinator\",\"type\":\"tuple\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVoteOptions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgChainLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"msgHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numSignUps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intStateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageBatchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOptionTreeDepth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"internalType\":\"enumMACI.Period\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateCommitment\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"_proof\",\"type\":\"uint256[8]\"}],\"name\":\"processMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTallyCommitment\",\"type\":\"uint256\"},{\"internalType\":\"uint256[8]\",\"name\":\"_proof\",\"type\":\"uint256[8]\"}],\"name\":\"processTally\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[7]\",\"name\":\"data\",\"type\":\"uint256[7]\"}],\"internalType\":\"structIMessage.Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_encPubKey\",\"type\":\"tuple\"}],\"name\":\"publishMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"qtrLib\",\"outputs\":[{\"internalType\":\"contractQuinaryTreeRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"result\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intStateTreeDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageBatchSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteOptionTreeDepth\",\"type\":\"uint256\"}],\"internalType\":\"structMACI.MaciParameters\",\"name\":\"_parameters\",\"type\":\"tuple\"}],\"name\":\"setParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"array\",\"type\":\"uint256[]\"}],\"name\":\"sha256Hash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structIPubKey.PubKey\",\"name\":\"_pubKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"signUp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stateIdxInc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"stateOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopProcessingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_results\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"stopTallyingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopTallyingPeriodWithoutResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxVoteOptions\",\"type\":\"uint256\"}],\"name\":\"stopVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vkRegistry\",\"outputs\":[{\"internalType\":\"contractVkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voiceCreditBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DorahacksVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use DorahacksVotingMetaData.ABI instead.
var DorahacksVotingABI = DorahacksVotingMetaData.ABI

// DorahacksVoting is an auto generated Go binding around an Ethereum contract.
type DorahacksVoting struct {
	DorahacksVotingCaller     // Read-only binding to the contract
	DorahacksVotingTransactor // Write-only binding to the contract
	DorahacksVotingFilterer   // Log filterer for contract events
}

// DorahacksVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type DorahacksVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DorahacksVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DorahacksVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DorahacksVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DorahacksVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DorahacksVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DorahacksVotingSession struct {
	Contract     *DorahacksVoting  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DorahacksVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DorahacksVotingCallerSession struct {
	Contract *DorahacksVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DorahacksVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DorahacksVotingTransactorSession struct {
	Contract     *DorahacksVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DorahacksVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type DorahacksVotingRaw struct {
	Contract *DorahacksVoting // Generic contract binding to access the raw methods on
}

// DorahacksVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DorahacksVotingCallerRaw struct {
	Contract *DorahacksVotingCaller // Generic read-only contract binding to access the raw methods on
}

// DorahacksVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DorahacksVotingTransactorRaw struct {
	Contract *DorahacksVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDorahacksVoting creates a new instance of DorahacksVoting, bound to a specific deployed contract.
func NewDorahacksVoting(address common.Address, backend bind.ContractBackend) (*DorahacksVoting, error) {
	contract, err := bindDorahacksVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DorahacksVoting{DorahacksVotingCaller: DorahacksVotingCaller{contract: contract}, DorahacksVotingTransactor: DorahacksVotingTransactor{contract: contract}, DorahacksVotingFilterer: DorahacksVotingFilterer{contract: contract}}, nil
}

// NewDorahacksVotingCaller creates a new read-only instance of DorahacksVoting, bound to a specific deployed contract.
func NewDorahacksVotingCaller(address common.Address, caller bind.ContractCaller) (*DorahacksVotingCaller, error) {
	contract, err := bindDorahacksVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingCaller{contract: contract}, nil
}

// NewDorahacksVotingTransactor creates a new write-only instance of DorahacksVoting, bound to a specific deployed contract.
func NewDorahacksVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*DorahacksVotingTransactor, error) {
	contract, err := bindDorahacksVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingTransactor{contract: contract}, nil
}

// NewDorahacksVotingFilterer creates a new log filterer instance of DorahacksVoting, bound to a specific deployed contract.
func NewDorahacksVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*DorahacksVotingFilterer, error) {
	contract, err := bindDorahacksVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingFilterer{contract: contract}, nil
}

// bindDorahacksVoting binds a generic wrapper to an already deployed contract.
func bindDorahacksVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DorahacksVotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DorahacksVoting *DorahacksVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DorahacksVoting.Contract.DorahacksVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DorahacksVoting *DorahacksVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.DorahacksVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DorahacksVoting *DorahacksVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.DorahacksVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DorahacksVoting *DorahacksVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DorahacksVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DorahacksVoting *DorahacksVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DorahacksVoting *DorahacksVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DorahacksVoting *DorahacksVotingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DorahacksVoting *DorahacksVotingSession) Admin() (common.Address, error) {
	return _DorahacksVoting.Contract.Admin(&_DorahacksVoting.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_DorahacksVoting *DorahacksVotingCallerSession) Admin() (common.Address, error) {
	return _DorahacksVoting.Contract.Admin(&_DorahacksVoting.CallOpts)
}

// CoordinatorHash is a free data retrieval call binding the contract method 0xde61a568.
//
// Solidity: function coordinatorHash() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) CoordinatorHash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "coordinatorHash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoordinatorHash is a free data retrieval call binding the contract method 0xde61a568.
//
// Solidity: function coordinatorHash() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) CoordinatorHash() (*big.Int, error) {
	return _DorahacksVoting.Contract.CoordinatorHash(&_DorahacksVoting.CallOpts)
}

// CoordinatorHash is a free data retrieval call binding the contract method 0xde61a568.
//
// Solidity: function coordinatorHash() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) CoordinatorHash() (*big.Int, error) {
	return _DorahacksVoting.Contract.CoordinatorHash(&_DorahacksVoting.CallOpts)
}

// CurrentStateCommitment is a free data retrieval call binding the contract method 0xec9289a8.
//
// Solidity: function currentStateCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) CurrentStateCommitment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "currentStateCommitment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentStateCommitment is a free data retrieval call binding the contract method 0xec9289a8.
//
// Solidity: function currentStateCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) CurrentStateCommitment() (*big.Int, error) {
	return _DorahacksVoting.Contract.CurrentStateCommitment(&_DorahacksVoting.CallOpts)
}

// CurrentStateCommitment is a free data retrieval call binding the contract method 0xec9289a8.
//
// Solidity: function currentStateCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) CurrentStateCommitment() (*big.Int, error) {
	return _DorahacksVoting.Contract.CurrentStateCommitment(&_DorahacksVoting.CallOpts)
}

// CurrentTallyCommitment is a free data retrieval call binding the contract method 0xc642eb8b.
//
// Solidity: function currentTallyCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) CurrentTallyCommitment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "currentTallyCommitment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTallyCommitment is a free data retrieval call binding the contract method 0xc642eb8b.
//
// Solidity: function currentTallyCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) CurrentTallyCommitment() (*big.Int, error) {
	return _DorahacksVoting.Contract.CurrentTallyCommitment(&_DorahacksVoting.CallOpts)
}

// CurrentTallyCommitment is a free data retrieval call binding the contract method 0xc642eb8b.
//
// Solidity: function currentTallyCommitment() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) CurrentTallyCommitment() (*big.Int, error) {
	return _DorahacksVoting.Contract.CurrentTallyCommitment(&_DorahacksVoting.CallOpts)
}

// GateKeeper is a free data retrieval call binding the contract method 0x45d61ded.
//
// Solidity: function gateKeeper() view returns(address)
func (_DorahacksVoting *DorahacksVotingCaller) GateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "gateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GateKeeper is a free data retrieval call binding the contract method 0x45d61ded.
//
// Solidity: function gateKeeper() view returns(address)
func (_DorahacksVoting *DorahacksVotingSession) GateKeeper() (common.Address, error) {
	return _DorahacksVoting.Contract.GateKeeper(&_DorahacksVoting.CallOpts)
}

// GateKeeper is a free data retrieval call binding the contract method 0x45d61ded.
//
// Solidity: function gateKeeper() view returns(address)
func (_DorahacksVoting *DorahacksVotingCallerSession) GateKeeper() (common.Address, error) {
	return _DorahacksVoting.Contract.GateKeeper(&_DorahacksVoting.CallOpts)
}

// Hash2 is a free data retrieval call binding the contract method 0x62a361bb.
//
// Solidity: function hash2(uint256[2] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) Hash2(opts *bind.CallOpts, array [2]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "hash2", array)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hash2 is a free data retrieval call binding the contract method 0x62a361bb.
//
// Solidity: function hash2(uint256[2] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) Hash2(array [2]*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Hash2(&_DorahacksVoting.CallOpts, array)
}

// Hash2 is a free data retrieval call binding the contract method 0x62a361bb.
//
// Solidity: function hash2(uint256[2] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) Hash2(array [2]*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Hash2(&_DorahacksVoting.CallOpts, array)
}

// Hash5 is a free data retrieval call binding the contract method 0x9cfced97.
//
// Solidity: function hash5(uint256[5] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) Hash5(opts *bind.CallOpts, array [5]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "hash5", array)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hash5 is a free data retrieval call binding the contract method 0x9cfced97.
//
// Solidity: function hash5(uint256[5] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) Hash5(array [5]*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Hash5(&_DorahacksVoting.CallOpts, array)
}

// Hash5 is a free data retrieval call binding the contract method 0x9cfced97.
//
// Solidity: function hash5(uint256[5] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) Hash5(array [5]*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Hash5(&_DorahacksVoting.CallOpts, array)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 _left, uint256 _right) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) HashLeftRight(opts *bind.CallOpts, _left *big.Int, _right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "hashLeftRight", _left, _right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 _left, uint256 _right) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) HashLeftRight(_left *big.Int, _right *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashLeftRight(&_DorahacksVoting.CallOpts, _left, _right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 _left, uint256 _right) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) HashLeftRight(_left *big.Int, _right *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashLeftRight(&_DorahacksVoting.CallOpts, _left, _right)
}

// HashMessageAndEncPubKey is a free data retrieval call binding the contract method 0xfd3dd132.
//
// Solidity: function hashMessageAndEncPubKey((uint256[7]) _message, (uint256,uint256) _encPubKey, uint256 _prevHash) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) HashMessageAndEncPubKey(opts *bind.CallOpts, _message IMessageMessage, _encPubKey IPubKeyPubKey, _prevHash *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "hashMessageAndEncPubKey", _message, _encPubKey, _prevHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashMessageAndEncPubKey is a free data retrieval call binding the contract method 0xfd3dd132.
//
// Solidity: function hashMessageAndEncPubKey((uint256[7]) _message, (uint256,uint256) _encPubKey, uint256 _prevHash) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) HashMessageAndEncPubKey(_message IMessageMessage, _encPubKey IPubKeyPubKey, _prevHash *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashMessageAndEncPubKey(&_DorahacksVoting.CallOpts, _message, _encPubKey, _prevHash)
}

// HashMessageAndEncPubKey is a free data retrieval call binding the contract method 0xfd3dd132.
//
// Solidity: function hashMessageAndEncPubKey((uint256[7]) _message, (uint256,uint256) _encPubKey, uint256 _prevHash) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) HashMessageAndEncPubKey(_message IMessageMessage, _encPubKey IPubKeyPubKey, _prevHash *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashMessageAndEncPubKey(&_DorahacksVoting.CallOpts, _message, _encPubKey, _prevHash)
}

// HashStateLeaf is a free data retrieval call binding the contract method 0x08c5038f.
//
// Solidity: function hashStateLeaf(((uint256,uint256),uint256,uint256,uint256) _stateLeaf) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) HashStateLeaf(opts *bind.CallOpts, _stateLeaf DomainObjsStateLeaf) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "hashStateLeaf", _stateLeaf)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashStateLeaf is a free data retrieval call binding the contract method 0x08c5038f.
//
// Solidity: function hashStateLeaf(((uint256,uint256),uint256,uint256,uint256) _stateLeaf) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) HashStateLeaf(_stateLeaf DomainObjsStateLeaf) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashStateLeaf(&_DorahacksVoting.CallOpts, _stateLeaf)
}

// HashStateLeaf is a free data retrieval call binding the contract method 0x08c5038f.
//
// Solidity: function hashStateLeaf(((uint256,uint256),uint256,uint256,uint256) _stateLeaf) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) HashStateLeaf(_stateLeaf DomainObjsStateLeaf) (*big.Int, error) {
	return _DorahacksVoting.Contract.HashStateLeaf(&_DorahacksVoting.CallOpts, _stateLeaf)
}

// MaxVoteOptions is a free data retrieval call binding the contract method 0x600b6ab3.
//
// Solidity: function maxVoteOptions() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) MaxVoteOptions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "maxVoteOptions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVoteOptions is a free data retrieval call binding the contract method 0x600b6ab3.
//
// Solidity: function maxVoteOptions() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) MaxVoteOptions() (*big.Int, error) {
	return _DorahacksVoting.Contract.MaxVoteOptions(&_DorahacksVoting.CallOpts)
}

// MaxVoteOptions is a free data retrieval call binding the contract method 0x600b6ab3.
//
// Solidity: function maxVoteOptions() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) MaxVoteOptions() (*big.Int, error) {
	return _DorahacksVoting.Contract.MaxVoteOptions(&_DorahacksVoting.CallOpts)
}

// MsgChainLength is a free data retrieval call binding the contract method 0x9be69195.
//
// Solidity: function msgChainLength() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) MsgChainLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "msgChainLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgChainLength is a free data retrieval call binding the contract method 0x9be69195.
//
// Solidity: function msgChainLength() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) MsgChainLength() (*big.Int, error) {
	return _DorahacksVoting.Contract.MsgChainLength(&_DorahacksVoting.CallOpts)
}

// MsgChainLength is a free data retrieval call binding the contract method 0x9be69195.
//
// Solidity: function msgChainLength() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) MsgChainLength() (*big.Int, error) {
	return _DorahacksVoting.Contract.MsgChainLength(&_DorahacksVoting.CallOpts)
}

// MsgHashes is a free data retrieval call binding the contract method 0xfc862175.
//
// Solidity: function msgHashes(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) MsgHashes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "msgHashes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgHashes is a free data retrieval call binding the contract method 0xfc862175.
//
// Solidity: function msgHashes(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) MsgHashes(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.MsgHashes(&_DorahacksVoting.CallOpts, arg0)
}

// MsgHashes is a free data retrieval call binding the contract method 0xfc862175.
//
// Solidity: function msgHashes(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) MsgHashes(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.MsgHashes(&_DorahacksVoting.CallOpts, arg0)
}

// NumSignUps is a free data retrieval call binding the contract method 0x122db153.
//
// Solidity: function numSignUps() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) NumSignUps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "numSignUps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumSignUps is a free data retrieval call binding the contract method 0x122db153.
//
// Solidity: function numSignUps() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) NumSignUps() (*big.Int, error) {
	return _DorahacksVoting.Contract.NumSignUps(&_DorahacksVoting.CallOpts)
}

// NumSignUps is a free data retrieval call binding the contract method 0x122db153.
//
// Solidity: function numSignUps() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) NumSignUps() (*big.Int, error) {
	return _DorahacksVoting.Contract.NumSignUps(&_DorahacksVoting.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(uint256 stateTreeDepth, uint256 intStateTreeDepth, uint256 messageBatchSize, uint256 voteOptionTreeDepth)
func (_DorahacksVoting *DorahacksVotingCaller) Parameters(opts *bind.CallOpts) (struct {
	StateTreeDepth      *big.Int
	IntStateTreeDepth   *big.Int
	MessageBatchSize    *big.Int
	VoteOptionTreeDepth *big.Int
}, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		StateTreeDepth      *big.Int
		IntStateTreeDepth   *big.Int
		MessageBatchSize    *big.Int
		VoteOptionTreeDepth *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StateTreeDepth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IntStateTreeDepth = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MessageBatchSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VoteOptionTreeDepth = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(uint256 stateTreeDepth, uint256 intStateTreeDepth, uint256 messageBatchSize, uint256 voteOptionTreeDepth)
func (_DorahacksVoting *DorahacksVotingSession) Parameters() (struct {
	StateTreeDepth      *big.Int
	IntStateTreeDepth   *big.Int
	MessageBatchSize    *big.Int
	VoteOptionTreeDepth *big.Int
}, error) {
	return _DorahacksVoting.Contract.Parameters(&_DorahacksVoting.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(uint256 stateTreeDepth, uint256 intStateTreeDepth, uint256 messageBatchSize, uint256 voteOptionTreeDepth)
func (_DorahacksVoting *DorahacksVotingCallerSession) Parameters() (struct {
	StateTreeDepth      *big.Int
	IntStateTreeDepth   *big.Int
	MessageBatchSize    *big.Int
	VoteOptionTreeDepth *big.Int
}, error) {
	return _DorahacksVoting.Contract.Parameters(&_DorahacksVoting.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint8)
func (_DorahacksVoting *DorahacksVotingCaller) Period(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint8)
func (_DorahacksVoting *DorahacksVotingSession) Period() (uint8, error) {
	return _DorahacksVoting.Contract.Period(&_DorahacksVoting.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint8)
func (_DorahacksVoting *DorahacksVotingCallerSession) Period() (uint8, error) {
	return _DorahacksVoting.Contract.Period(&_DorahacksVoting.CallOpts)
}

// QtrLib is a free data retrieval call binding the contract method 0xa4a2cf58.
//
// Solidity: function qtrLib() view returns(address)
func (_DorahacksVoting *DorahacksVotingCaller) QtrLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "qtrLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QtrLib is a free data retrieval call binding the contract method 0xa4a2cf58.
//
// Solidity: function qtrLib() view returns(address)
func (_DorahacksVoting *DorahacksVotingSession) QtrLib() (common.Address, error) {
	return _DorahacksVoting.Contract.QtrLib(&_DorahacksVoting.CallOpts)
}

// QtrLib is a free data retrieval call binding the contract method 0xa4a2cf58.
//
// Solidity: function qtrLib() view returns(address)
func (_DorahacksVoting *DorahacksVotingCallerSession) QtrLib() (common.Address, error) {
	return _DorahacksVoting.Contract.QtrLib(&_DorahacksVoting.CallOpts)
}

// Result is a free data retrieval call binding the contract method 0x3c594059.
//
// Solidity: function result(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) Result(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "result", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Result is a free data retrieval call binding the contract method 0x3c594059.
//
// Solidity: function result(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) Result(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Result(&_DorahacksVoting.CallOpts, arg0)
}

// Result is a free data retrieval call binding the contract method 0x3c594059.
//
// Solidity: function result(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) Result(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Result(&_DorahacksVoting.CallOpts, arg0)
}

// Sha256Hash is a free data retrieval call binding the contract method 0x58bfc379.
//
// Solidity: function sha256Hash(uint256[] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) Sha256Hash(opts *bind.CallOpts, array []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "sha256Hash", array)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sha256Hash is a free data retrieval call binding the contract method 0x58bfc379.
//
// Solidity: function sha256Hash(uint256[] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) Sha256Hash(array []*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Sha256Hash(&_DorahacksVoting.CallOpts, array)
}

// Sha256Hash is a free data retrieval call binding the contract method 0x58bfc379.
//
// Solidity: function sha256Hash(uint256[] array) pure returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) Sha256Hash(array []*big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.Sha256Hash(&_DorahacksVoting.CallOpts, array)
}

// StateIdxInc is a free data retrieval call binding the contract method 0x3e6fd0a8.
//
// Solidity: function stateIdxInc(address ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) StateIdxInc(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "stateIdxInc", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateIdxInc is a free data retrieval call binding the contract method 0x3e6fd0a8.
//
// Solidity: function stateIdxInc(address ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) StateIdxInc(arg0 common.Address) (*big.Int, error) {
	return _DorahacksVoting.Contract.StateIdxInc(&_DorahacksVoting.CallOpts, arg0)
}

// StateIdxInc is a free data retrieval call binding the contract method 0x3e6fd0a8.
//
// Solidity: function stateIdxInc(address ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) StateIdxInc(arg0 common.Address) (*big.Int, error) {
	return _DorahacksVoting.Contract.StateIdxInc(&_DorahacksVoting.CallOpts, arg0)
}

// StateOf is a free data retrieval call binding the contract method 0x45b4903a.
//
// Solidity: function stateOf(address _signer) view returns(uint256, uint256)
func (_DorahacksVoting *DorahacksVotingCaller) StateOf(opts *bind.CallOpts, _signer common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "stateOf", _signer)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// StateOf is a free data retrieval call binding the contract method 0x45b4903a.
//
// Solidity: function stateOf(address _signer) view returns(uint256, uint256)
func (_DorahacksVoting *DorahacksVotingSession) StateOf(_signer common.Address) (*big.Int, *big.Int, error) {
	return _DorahacksVoting.Contract.StateOf(&_DorahacksVoting.CallOpts, _signer)
}

// StateOf is a free data retrieval call binding the contract method 0x45b4903a.
//
// Solidity: function stateOf(address _signer) view returns(uint256, uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) StateOf(_signer common.Address) (*big.Int, *big.Int, error) {
	return _DorahacksVoting.Contract.StateOf(&_DorahacksVoting.CallOpts, _signer)
}

// TotalResult is a free data retrieval call binding the contract method 0x2b57f777.
//
// Solidity: function totalResult() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) TotalResult(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "totalResult")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalResult is a free data retrieval call binding the contract method 0x2b57f777.
//
// Solidity: function totalResult() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) TotalResult() (*big.Int, error) {
	return _DorahacksVoting.Contract.TotalResult(&_DorahacksVoting.CallOpts)
}

// TotalResult is a free data retrieval call binding the contract method 0x2b57f777.
//
// Solidity: function totalResult() view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) TotalResult() (*big.Int, error) {
	return _DorahacksVoting.Contract.TotalResult(&_DorahacksVoting.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DorahacksVoting *DorahacksVotingCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DorahacksVoting *DorahacksVotingSession) Verifier() (common.Address, error) {
	return _DorahacksVoting.Contract.Verifier(&_DorahacksVoting.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DorahacksVoting *DorahacksVotingCallerSession) Verifier() (common.Address, error) {
	return _DorahacksVoting.Contract.Verifier(&_DorahacksVoting.CallOpts)
}

// VkRegistry is a free data retrieval call binding the contract method 0x13fb8932.
//
// Solidity: function vkRegistry() view returns(address)
func (_DorahacksVoting *DorahacksVotingCaller) VkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "vkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VkRegistry is a free data retrieval call binding the contract method 0x13fb8932.
//
// Solidity: function vkRegistry() view returns(address)
func (_DorahacksVoting *DorahacksVotingSession) VkRegistry() (common.Address, error) {
	return _DorahacksVoting.Contract.VkRegistry(&_DorahacksVoting.CallOpts)
}

// VkRegistry is a free data retrieval call binding the contract method 0x13fb8932.
//
// Solidity: function vkRegistry() view returns(address)
func (_DorahacksVoting *DorahacksVotingCallerSession) VkRegistry() (common.Address, error) {
	return _DorahacksVoting.Contract.VkRegistry(&_DorahacksVoting.CallOpts)
}

// VoiceCreditBalance is a free data retrieval call binding the contract method 0x4f5e4f41.
//
// Solidity: function voiceCreditBalance(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCaller) VoiceCreditBalance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DorahacksVoting.contract.Call(opts, &out, "voiceCreditBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoiceCreditBalance is a free data retrieval call binding the contract method 0x4f5e4f41.
//
// Solidity: function voiceCreditBalance(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingSession) VoiceCreditBalance(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.VoiceCreditBalance(&_DorahacksVoting.CallOpts, arg0)
}

// VoiceCreditBalance is a free data retrieval call binding the contract method 0x4f5e4f41.
//
// Solidity: function voiceCreditBalance(uint256 ) view returns(uint256)
func (_DorahacksVoting *DorahacksVotingCallerSession) VoiceCreditBalance(arg0 *big.Int) (*big.Int, error) {
	return _DorahacksVoting.Contract.VoiceCreditBalance(&_DorahacksVoting.CallOpts, arg0)
}

// BatchPublishMessage is a paid mutator transaction binding the contract method 0x6b82a35c.
//
// Solidity: function batchPublishMessage((uint256[7])[] _messages, (uint256,uint256)[] _encPubKeys) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) BatchPublishMessage(opts *bind.TransactOpts, _messages []IMessageMessage, _encPubKeys []IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "batchPublishMessage", _messages, _encPubKeys)
}

// BatchPublishMessage is a paid mutator transaction binding the contract method 0x6b82a35c.
//
// Solidity: function batchPublishMessage((uint256[7])[] _messages, (uint256,uint256)[] _encPubKeys) returns()
func (_DorahacksVoting *DorahacksVotingSession) BatchPublishMessage(_messages []IMessageMessage, _encPubKeys []IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.BatchPublishMessage(&_DorahacksVoting.TransactOpts, _messages, _encPubKeys)
}

// BatchPublishMessage is a paid mutator transaction binding the contract method 0x6b82a35c.
//
// Solidity: function batchPublishMessage((uint256[7])[] _messages, (uint256,uint256)[] _encPubKeys) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) BatchPublishMessage(_messages []IMessageMessage, _encPubKeys []IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.BatchPublishMessage(&_DorahacksVoting.TransactOpts, _messages, _encPubKeys)
}

// Init is a paid mutator transaction binding the contract method 0x22deb130.
//
// Solidity: function init(address _admin, address _vkRegistry, address _qtrLib, address _verifier, address _gateKeeper, (uint256,uint256,uint256,uint256) _parameters, (uint256,uint256) _coordinator) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) Init(opts *bind.TransactOpts, _admin common.Address, _vkRegistry common.Address, _qtrLib common.Address, _verifier common.Address, _gateKeeper common.Address, _parameters MACIMaciParameters, _coordinator IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "init", _admin, _vkRegistry, _qtrLib, _verifier, _gateKeeper, _parameters, _coordinator)
}

// Init is a paid mutator transaction binding the contract method 0x22deb130.
//
// Solidity: function init(address _admin, address _vkRegistry, address _qtrLib, address _verifier, address _gateKeeper, (uint256,uint256,uint256,uint256) _parameters, (uint256,uint256) _coordinator) returns()
func (_DorahacksVoting *DorahacksVotingSession) Init(_admin common.Address, _vkRegistry common.Address, _qtrLib common.Address, _verifier common.Address, _gateKeeper common.Address, _parameters MACIMaciParameters, _coordinator IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.Init(&_DorahacksVoting.TransactOpts, _admin, _vkRegistry, _qtrLib, _verifier, _gateKeeper, _parameters, _coordinator)
}

// Init is a paid mutator transaction binding the contract method 0x22deb130.
//
// Solidity: function init(address _admin, address _vkRegistry, address _qtrLib, address _verifier, address _gateKeeper, (uint256,uint256,uint256,uint256) _parameters, (uint256,uint256) _coordinator) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) Init(_admin common.Address, _vkRegistry common.Address, _qtrLib common.Address, _verifier common.Address, _gateKeeper common.Address, _parameters MACIMaciParameters, _coordinator IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.Init(&_DorahacksVoting.TransactOpts, _admin, _vkRegistry, _qtrLib, _verifier, _gateKeeper, _parameters, _coordinator)
}

// ProcessMessage is a paid mutator transaction binding the contract method 0x3fe43c28.
//
// Solidity: function processMessage(uint256 newStateCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) ProcessMessage(opts *bind.TransactOpts, newStateCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "processMessage", newStateCommitment, _proof)
}

// ProcessMessage is a paid mutator transaction binding the contract method 0x3fe43c28.
//
// Solidity: function processMessage(uint256 newStateCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingSession) ProcessMessage(newStateCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.ProcessMessage(&_DorahacksVoting.TransactOpts, newStateCommitment, _proof)
}

// ProcessMessage is a paid mutator transaction binding the contract method 0x3fe43c28.
//
// Solidity: function processMessage(uint256 newStateCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) ProcessMessage(newStateCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.ProcessMessage(&_DorahacksVoting.TransactOpts, newStateCommitment, _proof)
}

// ProcessTally is a paid mutator transaction binding the contract method 0x124eff0d.
//
// Solidity: function processTally(uint256 newTallyCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) ProcessTally(opts *bind.TransactOpts, newTallyCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "processTally", newTallyCommitment, _proof)
}

// ProcessTally is a paid mutator transaction binding the contract method 0x124eff0d.
//
// Solidity: function processTally(uint256 newTallyCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingSession) ProcessTally(newTallyCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.ProcessTally(&_DorahacksVoting.TransactOpts, newTallyCommitment, _proof)
}

// ProcessTally is a paid mutator transaction binding the contract method 0x124eff0d.
//
// Solidity: function processTally(uint256 newTallyCommitment, uint256[8] _proof) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) ProcessTally(newTallyCommitment *big.Int, _proof [8]*big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.ProcessTally(&_DorahacksVoting.TransactOpts, newTallyCommitment, _proof)
}

// PublishMessage is a paid mutator transaction binding the contract method 0x91c6ad27.
//
// Solidity: function publishMessage((uint256[7]) _message, (uint256,uint256) _encPubKey) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) PublishMessage(opts *bind.TransactOpts, _message IMessageMessage, _encPubKey IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "publishMessage", _message, _encPubKey)
}

// PublishMessage is a paid mutator transaction binding the contract method 0x91c6ad27.
//
// Solidity: function publishMessage((uint256[7]) _message, (uint256,uint256) _encPubKey) returns()
func (_DorahacksVoting *DorahacksVotingSession) PublishMessage(_message IMessageMessage, _encPubKey IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.PublishMessage(&_DorahacksVoting.TransactOpts, _message, _encPubKey)
}

// PublishMessage is a paid mutator transaction binding the contract method 0x91c6ad27.
//
// Solidity: function publishMessage((uint256[7]) _message, (uint256,uint256) _encPubKey) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) PublishMessage(_message IMessageMessage, _encPubKey IPubKeyPubKey) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.PublishMessage(&_DorahacksVoting.TransactOpts, _message, _encPubKey)
}

// SetParameters is a paid mutator transaction binding the contract method 0x01865c96.
//
// Solidity: function setParameters((uint256,uint256,uint256,uint256) _parameters) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) SetParameters(opts *bind.TransactOpts, _parameters MACIMaciParameters) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "setParameters", _parameters)
}

// SetParameters is a paid mutator transaction binding the contract method 0x01865c96.
//
// Solidity: function setParameters((uint256,uint256,uint256,uint256) _parameters) returns()
func (_DorahacksVoting *DorahacksVotingSession) SetParameters(_parameters MACIMaciParameters) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.SetParameters(&_DorahacksVoting.TransactOpts, _parameters)
}

// SetParameters is a paid mutator transaction binding the contract method 0x01865c96.
//
// Solidity: function setParameters((uint256,uint256,uint256,uint256) _parameters) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) SetParameters(_parameters MACIMaciParameters) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.SetParameters(&_DorahacksVoting.TransactOpts, _parameters)
}

// SignUp is a paid mutator transaction binding the contract method 0x30c50728.
//
// Solidity: function signUp((uint256,uint256) _pubKey, bytes _data) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) SignUp(opts *bind.TransactOpts, _pubKey IPubKeyPubKey, _data []byte) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "signUp", _pubKey, _data)
}

// SignUp is a paid mutator transaction binding the contract method 0x30c50728.
//
// Solidity: function signUp((uint256,uint256) _pubKey, bytes _data) returns()
func (_DorahacksVoting *DorahacksVotingSession) SignUp(_pubKey IPubKeyPubKey, _data []byte) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.SignUp(&_DorahacksVoting.TransactOpts, _pubKey, _data)
}

// SignUp is a paid mutator transaction binding the contract method 0x30c50728.
//
// Solidity: function signUp((uint256,uint256) _pubKey, bytes _data) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) SignUp(_pubKey IPubKeyPubKey, _data []byte) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.SignUp(&_DorahacksVoting.TransactOpts, _pubKey, _data)
}

// StopProcessingPeriod is a paid mutator transaction binding the contract method 0xc2606f2a.
//
// Solidity: function stopProcessingPeriod() returns()
func (_DorahacksVoting *DorahacksVotingTransactor) StopProcessingPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "stopProcessingPeriod")
}

// StopProcessingPeriod is a paid mutator transaction binding the contract method 0xc2606f2a.
//
// Solidity: function stopProcessingPeriod() returns()
func (_DorahacksVoting *DorahacksVotingSession) StopProcessingPeriod() (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopProcessingPeriod(&_DorahacksVoting.TransactOpts)
}

// StopProcessingPeriod is a paid mutator transaction binding the contract method 0xc2606f2a.
//
// Solidity: function stopProcessingPeriod() returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) StopProcessingPeriod() (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopProcessingPeriod(&_DorahacksVoting.TransactOpts)
}

// StopTallyingPeriod is a paid mutator transaction binding the contract method 0x830e1b22.
//
// Solidity: function stopTallyingPeriod(uint256[] _results, uint256 _salt) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) StopTallyingPeriod(opts *bind.TransactOpts, _results []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "stopTallyingPeriod", _results, _salt)
}

// StopTallyingPeriod is a paid mutator transaction binding the contract method 0x830e1b22.
//
// Solidity: function stopTallyingPeriod(uint256[] _results, uint256 _salt) returns()
func (_DorahacksVoting *DorahacksVotingSession) StopTallyingPeriod(_results []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopTallyingPeriod(&_DorahacksVoting.TransactOpts, _results, _salt)
}

// StopTallyingPeriod is a paid mutator transaction binding the contract method 0x830e1b22.
//
// Solidity: function stopTallyingPeriod(uint256[] _results, uint256 _salt) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) StopTallyingPeriod(_results []*big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopTallyingPeriod(&_DorahacksVoting.TransactOpts, _results, _salt)
}

// StopTallyingPeriodWithoutResults is a paid mutator transaction binding the contract method 0x70ec4335.
//
// Solidity: function stopTallyingPeriodWithoutResults() returns()
func (_DorahacksVoting *DorahacksVotingTransactor) StopTallyingPeriodWithoutResults(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "stopTallyingPeriodWithoutResults")
}

// StopTallyingPeriodWithoutResults is a paid mutator transaction binding the contract method 0x70ec4335.
//
// Solidity: function stopTallyingPeriodWithoutResults() returns()
func (_DorahacksVoting *DorahacksVotingSession) StopTallyingPeriodWithoutResults() (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopTallyingPeriodWithoutResults(&_DorahacksVoting.TransactOpts)
}

// StopTallyingPeriodWithoutResults is a paid mutator transaction binding the contract method 0x70ec4335.
//
// Solidity: function stopTallyingPeriodWithoutResults() returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) StopTallyingPeriodWithoutResults() (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopTallyingPeriodWithoutResults(&_DorahacksVoting.TransactOpts)
}

// StopVotingPeriod is a paid mutator transaction binding the contract method 0x5336a586.
//
// Solidity: function stopVotingPeriod(uint256 _maxVoteOptions) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) StopVotingPeriod(opts *bind.TransactOpts, _maxVoteOptions *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "stopVotingPeriod", _maxVoteOptions)
}

// StopVotingPeriod is a paid mutator transaction binding the contract method 0x5336a586.
//
// Solidity: function stopVotingPeriod(uint256 _maxVoteOptions) returns()
func (_DorahacksVoting *DorahacksVotingSession) StopVotingPeriod(_maxVoteOptions *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopVotingPeriod(&_DorahacksVoting.TransactOpts, _maxVoteOptions)
}

// StopVotingPeriod is a paid mutator transaction binding the contract method 0x5336a586.
//
// Solidity: function stopVotingPeriod(uint256 _maxVoteOptions) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) StopVotingPeriod(_maxVoteOptions *big.Int) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.StopVotingPeriod(&_DorahacksVoting.TransactOpts, _maxVoteOptions)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_DorahacksVoting *DorahacksVotingTransactor) TransferOwnership(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _DorahacksVoting.contract.Transact(opts, "transferOwnership", newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_DorahacksVoting *DorahacksVotingSession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.TransferOwnership(&_DorahacksVoting.TransactOpts, newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newAdmin) returns()
func (_DorahacksVoting *DorahacksVotingTransactorSession) TransferOwnership(newAdmin common.Address) (*types.Transaction, error) {
	return _DorahacksVoting.Contract.TransferOwnership(&_DorahacksVoting.TransactOpts, newAdmin)
}

// DorahacksVotingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DorahacksVoting contract.
type DorahacksVotingOwnershipTransferredIterator struct {
	Event *DorahacksVotingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DorahacksVotingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DorahacksVotingOwnershipTransferred)
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
		it.Event = new(DorahacksVotingOwnershipTransferred)
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
func (it *DorahacksVotingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DorahacksVotingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DorahacksVotingOwnershipTransferred represents a OwnershipTransferred event raised by the DorahacksVoting contract.
type DorahacksVotingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DorahacksVoting *DorahacksVotingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DorahacksVotingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DorahacksVoting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingOwnershipTransferredIterator{contract: _DorahacksVoting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DorahacksVoting *DorahacksVotingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DorahacksVotingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DorahacksVoting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DorahacksVotingOwnershipTransferred)
				if err := _DorahacksVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DorahacksVoting *DorahacksVotingFilterer) ParseOwnershipTransferred(log types.Log) (*DorahacksVotingOwnershipTransferred, error) {
	event := new(DorahacksVotingOwnershipTransferred)
	if err := _DorahacksVoting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DorahacksVotingPublishMessageIterator is returned from FilterPublishMessage and is used to iterate over the raw logs and unpacked data for PublishMessage events raised by the DorahacksVoting contract.
type DorahacksVotingPublishMessageIterator struct {
	Event *DorahacksVotingPublishMessage // Event containing the contract specifics and raw log

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
func (it *DorahacksVotingPublishMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DorahacksVotingPublishMessage)
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
		it.Event = new(DorahacksVotingPublishMessage)
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
func (it *DorahacksVotingPublishMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DorahacksVotingPublishMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DorahacksVotingPublishMessage represents a PublishMessage event raised by the DorahacksVoting contract.
type DorahacksVotingPublishMessage struct {
	MsgIdx    *big.Int
	Message   IMessageMessage
	EncPubKey IPubKeyPubKey
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublishMessage is a free log retrieval operation binding the contract event 0x8bb5a8cf78a5b2f53c73e2feacb1fb3e91c3f03cb15e33f53174db20e37e3928.
//
// Solidity: event PublishMessage(uint256 indexed _msgIdx, (uint256[7]) _message, (uint256,uint256) _encPubKey)
func (_DorahacksVoting *DorahacksVotingFilterer) FilterPublishMessage(opts *bind.FilterOpts, _msgIdx []*big.Int) (*DorahacksVotingPublishMessageIterator, error) {

	var _msgIdxRule []interface{}
	for _, _msgIdxItem := range _msgIdx {
		_msgIdxRule = append(_msgIdxRule, _msgIdxItem)
	}

	logs, sub, err := _DorahacksVoting.contract.FilterLogs(opts, "PublishMessage", _msgIdxRule)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingPublishMessageIterator{contract: _DorahacksVoting.contract, event: "PublishMessage", logs: logs, sub: sub}, nil
}

// WatchPublishMessage is a free log subscription operation binding the contract event 0x8bb5a8cf78a5b2f53c73e2feacb1fb3e91c3f03cb15e33f53174db20e37e3928.
//
// Solidity: event PublishMessage(uint256 indexed _msgIdx, (uint256[7]) _message, (uint256,uint256) _encPubKey)
func (_DorahacksVoting *DorahacksVotingFilterer) WatchPublishMessage(opts *bind.WatchOpts, sink chan<- *DorahacksVotingPublishMessage, _msgIdx []*big.Int) (event.Subscription, error) {

	var _msgIdxRule []interface{}
	for _, _msgIdxItem := range _msgIdx {
		_msgIdxRule = append(_msgIdxRule, _msgIdxItem)
	}

	logs, sub, err := _DorahacksVoting.contract.WatchLogs(opts, "PublishMessage", _msgIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DorahacksVotingPublishMessage)
				if err := _DorahacksVoting.contract.UnpackLog(event, "PublishMessage", log); err != nil {
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

// ParsePublishMessage is a log parse operation binding the contract event 0x8bb5a8cf78a5b2f53c73e2feacb1fb3e91c3f03cb15e33f53174db20e37e3928.
//
// Solidity: event PublishMessage(uint256 indexed _msgIdx, (uint256[7]) _message, (uint256,uint256) _encPubKey)
func (_DorahacksVoting *DorahacksVotingFilterer) ParsePublishMessage(log types.Log) (*DorahacksVotingPublishMessage, error) {
	event := new(DorahacksVotingPublishMessage)
	if err := _DorahacksVoting.contract.UnpackLog(event, "PublishMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DorahacksVotingSignUpIterator is returned from FilterSignUp and is used to iterate over the raw logs and unpacked data for SignUp events raised by the DorahacksVoting contract.
type DorahacksVotingSignUpIterator struct {
	Event *DorahacksVotingSignUp // Event containing the contract specifics and raw log

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
func (it *DorahacksVotingSignUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DorahacksVotingSignUp)
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
		it.Event = new(DorahacksVotingSignUp)
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
func (it *DorahacksVotingSignUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DorahacksVotingSignUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DorahacksVotingSignUp represents a SignUp event raised by the DorahacksVoting contract.
type DorahacksVotingSignUp struct {
	StateIdx           *big.Int
	UserPubKey         IPubKeyPubKey
	VoiceCreditBalance *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSignUp is a free log retrieval operation binding the contract event 0xc7563c66f89e2fb0839e2b64ed54fe4803ff9428777814772ccfe4c385072c4b.
//
// Solidity: event SignUp(uint256 indexed _stateIdx, (uint256,uint256) _userPubKey, uint256 _voiceCreditBalance)
func (_DorahacksVoting *DorahacksVotingFilterer) FilterSignUp(opts *bind.FilterOpts, _stateIdx []*big.Int) (*DorahacksVotingSignUpIterator, error) {

	var _stateIdxRule []interface{}
	for _, _stateIdxItem := range _stateIdx {
		_stateIdxRule = append(_stateIdxRule, _stateIdxItem)
	}

	logs, sub, err := _DorahacksVoting.contract.FilterLogs(opts, "SignUp", _stateIdxRule)
	if err != nil {
		return nil, err
	}
	return &DorahacksVotingSignUpIterator{contract: _DorahacksVoting.contract, event: "SignUp", logs: logs, sub: sub}, nil
}

// WatchSignUp is a free log subscription operation binding the contract event 0xc7563c66f89e2fb0839e2b64ed54fe4803ff9428777814772ccfe4c385072c4b.
//
// Solidity: event SignUp(uint256 indexed _stateIdx, (uint256,uint256) _userPubKey, uint256 _voiceCreditBalance)
func (_DorahacksVoting *DorahacksVotingFilterer) WatchSignUp(opts *bind.WatchOpts, sink chan<- *DorahacksVotingSignUp, _stateIdx []*big.Int) (event.Subscription, error) {

	var _stateIdxRule []interface{}
	for _, _stateIdxItem := range _stateIdx {
		_stateIdxRule = append(_stateIdxRule, _stateIdxItem)
	}

	logs, sub, err := _DorahacksVoting.contract.WatchLogs(opts, "SignUp", _stateIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DorahacksVotingSignUp)
				if err := _DorahacksVoting.contract.UnpackLog(event, "SignUp", log); err != nil {
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

// ParseSignUp is a log parse operation binding the contract event 0xc7563c66f89e2fb0839e2b64ed54fe4803ff9428777814772ccfe4c385072c4b.
//
// Solidity: event SignUp(uint256 indexed _stateIdx, (uint256,uint256) _userPubKey, uint256 _voiceCreditBalance)
func (_DorahacksVoting *DorahacksVotingFilterer) ParseSignUp(log types.Log) (*DorahacksVotingSignUp, error) {
	event := new(DorahacksVotingSignUp)
	if err := _DorahacksVoting.contract.UnpackLog(event, "SignUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
