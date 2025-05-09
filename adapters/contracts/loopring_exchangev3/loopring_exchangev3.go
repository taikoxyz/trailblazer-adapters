// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package loopring_exchangev3

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

// ExchangeDataAccountLeaf is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataAccountLeaf struct {
	AccountID  uint32
	Owner      common.Address
	PubKeyX    *big.Int
	PubKeyY    *big.Int
	Nonce      uint32
	FeeBipsAMM *big.Int
}

// ExchangeDataBalanceLeaf is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBalanceLeaf struct {
	TokenID     uint16
	Balance     *big.Int
	WeightAMM   *big.Int
	StorageRoot *big.Int
}

// ExchangeDataBlock is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBlock struct {
	BlockType             uint8
	BlockSize             uint16
	BlockVersion          uint8
	Data                  []byte
	Proof                 [8]*big.Int
	StoreBlockInfoOnchain bool
	AuxiliaryData         []byte
	OffchainData          []byte
}

// ExchangeDataBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBlockInfo struct {
	Timestamp     uint32
	BlockDataHash [28]byte
}

// ExchangeDataConstants is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataConstants struct {
	SNARKSCALARFIELD                         *big.Int
	MAXOPENFORCEDREQUESTS                    *big.Int
	MAXAGEFORCEDREQUESTUNTILWITHDRAWMODE     *big.Int
	TIMESTAMPHALFWINDOWSIZEINSECONDS         *big.Int
	MAXNUMACCOUNTS                           *big.Int
	MAXNUMTOKENS                             *big.Int
	MINAGEPROTOCOLFEESUNTILUPDATED           *big.Int
	MINTIMEINSHUTDOWN                        *big.Int
	TXDATAAVAILABILITYSIZE                   *big.Int
	MAXAGEDEPOSITUNTILWITHDRAWABLEUPPERBOUND *big.Int
}

// ExchangeDataMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataMerkleProof struct {
	AccountLeaf        ExchangeDataAccountLeaf
	BalanceLeaf        ExchangeDataBalanceLeaf
	Nft                ExchangeDataNft
	AccountMerkleProof [48]*big.Int
	BalanceMerkleProof [24]*big.Int
}

// ExchangeDataNft is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataNft struct {
	Minter         common.Address
	NftType        uint8
	Token          common.Address
	NftID          *big.Int
	CreatorFeeBips uint8
}

// LoopringExchangev3MetaData contains all meta data concerning the LoopringExchangev3 contract.
var LoopringExchangev3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicDataHash\",\"type\":\"bytes32\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"DepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesisMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"ExchangeCloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"}],\"name\":\"ForcedWithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"NFTDepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NftWithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NftWithdrawalFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"takerFeeBips\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"makerFeeBips\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"previousTakerFeeBips\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"previousMakerFeeBips\",\"type\":\"uint8\"}],\"name\":\"ProtocolFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"tokenId\",\"type\":\"uint16\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawalModeActivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"approveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"transactionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"approveTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnExchangeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"enumExchangeData.NftType\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"depositNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"}],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"}],\"name\":\"forceWithdrawByTokenID\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAgentRegistry\",\"outputs\":[{\"internalType\":\"contractIAgentRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmmFeeBips\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAmountWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumExchangeData.NftType\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"getAmountWithdrawableNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockIdx\",\"type\":\"uint256\"}],\"name\":\"getBlockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes28\",\"name\":\"blockDataHash\",\"type\":\"bytes28\"}],\"internalType\":\"structExchangeData.BlockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConstants\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"SNARK_SCALAR_FIELD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_OPEN_FORCED_REQUESTS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_AGE_FORCED_REQUEST_UNTIL_WITHDRAW_MODE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TIMESTAMP_HALF_WINDOW_SIZE_IN_SECONDS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_NUM_ACCOUNTS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_NUM_TOKENS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MIN_AGE_PROTOCOL_FEES_UNTIL_UPDATED\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MIN_TIME_IN_SHUTDOWN\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TX_DATA_AVAILABILITY_SIZE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_AGE_DEPOSIT_UNTIL_WITHDRAWABLE_UPPERBOUND\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.Constants\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxAgeDepositUntilWithdrawable\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumAvailableForcedSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getPendingDepositAmount\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumExchangeData.NftType\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"}],\"name\":\"getPendingNFTDepositAmount\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getProtocolFeeLastWithdrawnTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeValues\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"syncedAt\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"takerFeeBips\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"makerFeeBips\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"previousTakerFeeBips\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"previousMakerFeeBips\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"}],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"storageID\",\"type\":\"uint32\"}],\"name\":\"getWithdrawalRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_loopring\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isForcedWithdrawalPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInWithdrawalMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"isTransactionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isUserOrAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWithdrawnInWithdrawalMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"}],\"name\":\"notifyForcedRequestTooOld\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onchainTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshBlockVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_agentRegistry\",\"type\":\"address\"}],\"name\":\"setAgentRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowOnchainTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_feeBips\",\"type\":\"uint8\"}],\"name\":\"setAmmFeeBips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"name\":\"setDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setMaxAgeDepositUntilWithdrawable\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"storageID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"setWithdrawalRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"blockType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"blockVersion\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bool\",\"name\":\"storeBlockInfoOnchain\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"auxiliaryData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainData\",\"type\":\"bytes\"}],\"internalType\":\"structExchangeData.Block[]\",\"name\":\"blocks\",\"type\":\"tuple[]\"}],\"name\":\"submitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawExchangeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawExchangeStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"withdrawFromApprovedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"minters\",\"type\":\"address[]\"},{\"internalType\":\"enumExchangeData.NftType[]\",\"name\":\"nftTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIDs\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFromApprovedWithdrawalsNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawFromDepositRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"feeBipsAMM\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.AccountLeaf\",\"name\":\"accountLeaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"tokenID\",\"type\":\"uint16\"},{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"weightAMM\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storageRoot\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.BalanceLeaf\",\"name\":\"balanceLeaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"enumExchangeData.NftType\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"creatorFeeBips\",\"type\":\"uint8\"}],\"internalType\":\"structExchangeData.Nft\",\"name\":\"nft\",\"type\":\"tuple\"},{\"internalType\":\"uint256[48]\",\"name\":\"accountMerkleProof\",\"type\":\"uint256[48]\"},{\"internalType\":\"uint256[24]\",\"name\":\"balanceMerkleProof\",\"type\":\"uint256[24]\"}],\"internalType\":\"structExchangeData.MerkleProof\",\"name\":\"merkleProof\",\"type\":\"tuple\"}],\"name\":\"withdrawFromMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumExchangeData.NftType\",\"name\":\"nftType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"}],\"name\":\"withdrawFromNFTDepositRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFees\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// LoopringExchangev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use LoopringExchangev3MetaData.ABI instead.
var LoopringExchangev3ABI = LoopringExchangev3MetaData.ABI

// LoopringExchangev3 is an auto generated Go binding around an Ethereum contract.
type LoopringExchangev3 struct {
	LoopringExchangev3Caller     // Read-only binding to the contract
	LoopringExchangev3Transactor // Write-only binding to the contract
	LoopringExchangev3Filterer   // Log filterer for contract events
}

// LoopringExchangev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type LoopringExchangev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringExchangev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LoopringExchangev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringExchangev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoopringExchangev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopringExchangev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoopringExchangev3Session struct {
	Contract     *LoopringExchangev3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LoopringExchangev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoopringExchangev3CallerSession struct {
	Contract *LoopringExchangev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LoopringExchangev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoopringExchangev3TransactorSession struct {
	Contract     *LoopringExchangev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LoopringExchangev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type LoopringExchangev3Raw struct {
	Contract *LoopringExchangev3 // Generic contract binding to access the raw methods on
}

// LoopringExchangev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoopringExchangev3CallerRaw struct {
	Contract *LoopringExchangev3Caller // Generic read-only contract binding to access the raw methods on
}

// LoopringExchangev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoopringExchangev3TransactorRaw struct {
	Contract *LoopringExchangev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLoopringExchangev3 creates a new instance of LoopringExchangev3, bound to a specific deployed contract.
func NewLoopringExchangev3(address common.Address, backend bind.ContractBackend) (*LoopringExchangev3, error) {
	contract, err := bindLoopringExchangev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3{LoopringExchangev3Caller: LoopringExchangev3Caller{contract: contract}, LoopringExchangev3Transactor: LoopringExchangev3Transactor{contract: contract}, LoopringExchangev3Filterer: LoopringExchangev3Filterer{contract: contract}}, nil
}

// NewLoopringExchangev3Caller creates a new read-only instance of LoopringExchangev3, bound to a specific deployed contract.
func NewLoopringExchangev3Caller(address common.Address, caller bind.ContractCaller) (*LoopringExchangev3Caller, error) {
	contract, err := bindLoopringExchangev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3Caller{contract: contract}, nil
}

// NewLoopringExchangev3Transactor creates a new write-only instance of LoopringExchangev3, bound to a specific deployed contract.
func NewLoopringExchangev3Transactor(address common.Address, transactor bind.ContractTransactor) (*LoopringExchangev3Transactor, error) {
	contract, err := bindLoopringExchangev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3Transactor{contract: contract}, nil
}

// NewLoopringExchangev3Filterer creates a new log filterer instance of LoopringExchangev3, bound to a specific deployed contract.
func NewLoopringExchangev3Filterer(address common.Address, filterer bind.ContractFilterer) (*LoopringExchangev3Filterer, error) {
	contract, err := bindLoopringExchangev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3Filterer{contract: contract}, nil
}

// bindLoopringExchangev3 binds a generic wrapper to an already deployed contract.
func bindLoopringExchangev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoopringExchangev3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoopringExchangev3 *LoopringExchangev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoopringExchangev3.Contract.LoopringExchangev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoopringExchangev3 *LoopringExchangev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.LoopringExchangev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoopringExchangev3 *LoopringExchangev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.LoopringExchangev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoopringExchangev3 *LoopringExchangev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoopringExchangev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoopringExchangev3 *LoopringExchangev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoopringExchangev3 *LoopringExchangev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Session) DomainSeparator() ([32]byte, error) {
	return _LoopringExchangev3.Contract.DomainSeparator(&_LoopringExchangev3.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) DomainSeparator() ([32]byte, error) {
	return _LoopringExchangev3.Contract.DomainSeparator(&_LoopringExchangev3.CallOpts)
}

// GetAgentRegistry is a free data retrieval call binding the contract method 0x91cae372.
//
// Solidity: function getAgentRegistry() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetAgentRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getAgentRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgentRegistry is a free data retrieval call binding the contract method 0x91cae372.
//
// Solidity: function getAgentRegistry() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetAgentRegistry() (common.Address, error) {
	return _LoopringExchangev3.Contract.GetAgentRegistry(&_LoopringExchangev3.CallOpts)
}

// GetAgentRegistry is a free data retrieval call binding the contract method 0x91cae372.
//
// Solidity: function getAgentRegistry() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetAgentRegistry() (common.Address, error) {
	return _LoopringExchangev3.Contract.GetAgentRegistry(&_LoopringExchangev3.CallOpts)
}

// GetAmmFeeBips is a free data retrieval call binding the contract method 0xfcd7810c.
//
// Solidity: function getAmmFeeBips() view returns(uint8)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetAmmFeeBips(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getAmmFeeBips")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAmmFeeBips is a free data retrieval call binding the contract method 0xfcd7810c.
//
// Solidity: function getAmmFeeBips() view returns(uint8)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetAmmFeeBips() (uint8, error) {
	return _LoopringExchangev3.Contract.GetAmmFeeBips(&_LoopringExchangev3.CallOpts)
}

// GetAmmFeeBips is a free data retrieval call binding the contract method 0xfcd7810c.
//
// Solidity: function getAmmFeeBips() view returns(uint8)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetAmmFeeBips() (uint8, error) {
	return _LoopringExchangev3.Contract.GetAmmFeeBips(&_LoopringExchangev3.CallOpts)
}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetAmountWithdrawable(opts *bind.CallOpts, owner common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getAmountWithdrawable", owner, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetAmountWithdrawable(owner common.Address, token common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetAmountWithdrawable(&_LoopringExchangev3.CallOpts, owner, token)
}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetAmountWithdrawable(owner common.Address, token common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetAmountWithdrawable(&_LoopringExchangev3.CallOpts, owner, token)
}

// GetAmountWithdrawableNFT is a free data retrieval call binding the contract method 0x05987d57.
//
// Solidity: function getAmountWithdrawableNFT(address owner, address token, uint8 nftType, uint256 nftID, address minter) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetAmountWithdrawableNFT(opts *bind.CallOpts, owner common.Address, token common.Address, nftType uint8, nftID *big.Int, minter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getAmountWithdrawableNFT", owner, token, nftType, nftID, minter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountWithdrawableNFT is a free data retrieval call binding the contract method 0x05987d57.
//
// Solidity: function getAmountWithdrawableNFT(address owner, address token, uint8 nftType, uint256 nftID, address minter) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetAmountWithdrawableNFT(owner common.Address, token common.Address, nftType uint8, nftID *big.Int, minter common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetAmountWithdrawableNFT(&_LoopringExchangev3.CallOpts, owner, token, nftType, nftID, minter)
}

// GetAmountWithdrawableNFT is a free data retrieval call binding the contract method 0x05987d57.
//
// Solidity: function getAmountWithdrawableNFT(address owner, address token, uint8 nftType, uint256 nftID, address minter) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetAmountWithdrawableNFT(owner common.Address, token common.Address, nftType uint8, nftID *big.Int, minter common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetAmountWithdrawableNFT(&_LoopringExchangev3.CallOpts, owner, token, nftType, nftID, minter)
}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetBlockHeight() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetBlockHeight(&_LoopringExchangev3.CallOpts)
}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetBlockHeight() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetBlockHeight(&_LoopringExchangev3.CallOpts)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetBlockInfo(opts *bind.CallOpts, blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getBlockInfo", blockIdx)

	if err != nil {
		return *new(ExchangeDataBlockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeDataBlockInfo)).(*ExchangeDataBlockInfo)

	return out0, err

}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_LoopringExchangev3 *LoopringExchangev3Session) GetBlockInfo(blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	return _LoopringExchangev3.Contract.GetBlockInfo(&_LoopringExchangev3.CallOpts, blockIdx)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetBlockInfo(blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	return _LoopringExchangev3.Contract.GetBlockInfo(&_LoopringExchangev3.CallOpts, blockIdx)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetConstants(opts *bind.CallOpts) (ExchangeDataConstants, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getConstants")

	if err != nil {
		return *new(ExchangeDataConstants), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeDataConstants)).(*ExchangeDataConstants)

	return out0, err

}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LoopringExchangev3 *LoopringExchangev3Session) GetConstants() (ExchangeDataConstants, error) {
	return _LoopringExchangev3.Contract.GetConstants(&_LoopringExchangev3.CallOpts)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetConstants() (ExchangeDataConstants, error) {
	return _LoopringExchangev3.Contract.GetConstants(&_LoopringExchangev3.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetDepositContract() (common.Address, error) {
	return _LoopringExchangev3.Contract.GetDepositContract(&_LoopringExchangev3.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetDepositContract() (common.Address, error) {
	return _LoopringExchangev3.Contract.GetDepositContract(&_LoopringExchangev3.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetDomainSeparator() ([32]byte, error) {
	return _LoopringExchangev3.Contract.GetDomainSeparator(&_LoopringExchangev3.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetDomainSeparator() ([32]byte, error) {
	return _LoopringExchangev3.Contract.GetDomainSeparator(&_LoopringExchangev3.CallOpts)
}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetExchangeStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getExchangeStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetExchangeStake() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetExchangeStake(&_LoopringExchangev3.CallOpts)
}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetExchangeStake() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetExchangeStake(&_LoopringExchangev3.CallOpts)
}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetMaxAgeDepositUntilWithdrawable(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getMaxAgeDepositUntilWithdrawable")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetMaxAgeDepositUntilWithdrawable() (uint32, error) {
	return _LoopringExchangev3.Contract.GetMaxAgeDepositUntilWithdrawable(&_LoopringExchangev3.CallOpts)
}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetMaxAgeDepositUntilWithdrawable() (uint32, error) {
	return _LoopringExchangev3.Contract.GetMaxAgeDepositUntilWithdrawable(&_LoopringExchangev3.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetMerkleRoot() ([32]byte, error) {
	return _LoopringExchangev3.Contract.GetMerkleRoot(&_LoopringExchangev3.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetMerkleRoot() ([32]byte, error) {
	return _LoopringExchangev3.Contract.GetMerkleRoot(&_LoopringExchangev3.CallOpts)
}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetNumAvailableForcedSlots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getNumAvailableForcedSlots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetNumAvailableForcedSlots() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetNumAvailableForcedSlots(&_LoopringExchangev3.CallOpts)
}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetNumAvailableForcedSlots() (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetNumAvailableForcedSlots(&_LoopringExchangev3.CallOpts)
}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetPendingDepositAmount(opts *bind.CallOpts, owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getPendingDepositAmount", owner, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetPendingDepositAmount(owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetPendingDepositAmount(&_LoopringExchangev3.CallOpts, owner, tokenAddress)
}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetPendingDepositAmount(owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetPendingDepositAmount(&_LoopringExchangev3.CallOpts, owner, tokenAddress)
}

// GetPendingNFTDepositAmount is a free data retrieval call binding the contract method 0x7d363601.
//
// Solidity: function getPendingNFTDepositAmount(address owner, address token, uint8 nftType, uint256 nftID) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetPendingNFTDepositAmount(opts *bind.CallOpts, owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getPendingNFTDepositAmount", owner, token, nftType, nftID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingNFTDepositAmount is a free data retrieval call binding the contract method 0x7d363601.
//
// Solidity: function getPendingNFTDepositAmount(address owner, address token, uint8 nftType, uint256 nftID) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetPendingNFTDepositAmount(owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetPendingNFTDepositAmount(&_LoopringExchangev3.CallOpts, owner, token, nftType, nftID)
}

// GetPendingNFTDepositAmount is a free data retrieval call binding the contract method 0x7d363601.
//
// Solidity: function getPendingNFTDepositAmount(address owner, address token, uint8 nftType, uint256 nftID) view returns(uint96)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetPendingNFTDepositAmount(owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetPendingNFTDepositAmount(&_LoopringExchangev3.CallOpts, owner, token, nftType, nftID)
}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetProtocolFeeLastWithdrawnTime(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getProtocolFeeLastWithdrawnTime", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetProtocolFeeLastWithdrawnTime(tokenAddress common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetProtocolFeeLastWithdrawnTime(&_LoopringExchangev3.CallOpts, tokenAddress)
}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetProtocolFeeLastWithdrawnTime(tokenAddress common.Address) (*big.Int, error) {
	return _LoopringExchangev3.Contract.GetProtocolFeeLastWithdrawnTime(&_LoopringExchangev3.CallOpts, tokenAddress)
}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetProtocolFeeValues(opts *bind.CallOpts) (struct {
	SyncedAt             uint32
	TakerFeeBips         uint8
	MakerFeeBips         uint8
	PreviousTakerFeeBips uint8
	PreviousMakerFeeBips uint8
}, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getProtocolFeeValues")

	outstruct := new(struct {
		SyncedAt             uint32
		TakerFeeBips         uint8
		MakerFeeBips         uint8
		PreviousTakerFeeBips uint8
		PreviousMakerFeeBips uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SyncedAt = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TakerFeeBips = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.MakerFeeBips = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.PreviousTakerFeeBips = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.PreviousMakerFeeBips = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetProtocolFeeValues() (struct {
	SyncedAt             uint32
	TakerFeeBips         uint8
	MakerFeeBips         uint8
	PreviousTakerFeeBips uint8
	PreviousMakerFeeBips uint8
}, error) {
	return _LoopringExchangev3.Contract.GetProtocolFeeValues(&_LoopringExchangev3.CallOpts)
}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetProtocolFeeValues() (struct {
	SyncedAt             uint32
	TakerFeeBips         uint8
	MakerFeeBips         uint8
	PreviousTakerFeeBips uint8
	PreviousMakerFeeBips uint8
}, error) {
	return _LoopringExchangev3.Contract.GetProtocolFeeValues(&_LoopringExchangev3.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0xef365218.
//
// Solidity: function getTokenAddress(uint16 tokenID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetTokenAddress(opts *bind.CallOpts, tokenID uint16) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getTokenAddress", tokenID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0xef365218.
//
// Solidity: function getTokenAddress(uint16 tokenID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetTokenAddress(tokenID uint16) (common.Address, error) {
	return _LoopringExchangev3.Contract.GetTokenAddress(&_LoopringExchangev3.CallOpts, tokenID)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0xef365218.
//
// Solidity: function getTokenAddress(uint16 tokenID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetTokenAddress(tokenID uint16) (common.Address, error) {
	return _LoopringExchangev3.Contract.GetTokenAddress(&_LoopringExchangev3.CallOpts, tokenID)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetTokenID(opts *bind.CallOpts, tokenAddress common.Address) (uint16, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getTokenID", tokenAddress)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetTokenID(tokenAddress common.Address) (uint16, error) {
	return _LoopringExchangev3.Contract.GetTokenID(&_LoopringExchangev3.CallOpts, tokenAddress)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetTokenID(tokenAddress common.Address) (uint16, error) {
	return _LoopringExchangev3.Contract.GetTokenID(&_LoopringExchangev3.CallOpts, tokenAddress)
}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0x1ef36835.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) GetWithdrawalRecipient(opts *bind.CallOpts, from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "getWithdrawalRecipient", from, to, token, amount, storageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0x1ef36835.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) GetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	return _LoopringExchangev3.Contract.GetWithdrawalRecipient(&_LoopringExchangev3.CallOpts, from, to, token, amount, storageID)
}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0x1ef36835.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID) view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) GetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	return _LoopringExchangev3.Contract.GetWithdrawalRecipient(&_LoopringExchangev3.CallOpts, from, to, token, amount, storageID)
}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsForcedWithdrawalPending(opts *bind.CallOpts, accountID uint32, token common.Address) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isForcedWithdrawalPending", accountID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsForcedWithdrawalPending(accountID uint32, token common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsForcedWithdrawalPending(&_LoopringExchangev3.CallOpts, accountID, token)
}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsForcedWithdrawalPending(accountID uint32, token common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsForcedWithdrawalPending(&_LoopringExchangev3.CallOpts, accountID, token)
}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsInWithdrawalMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isInWithdrawalMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsInWithdrawalMode() (bool, error) {
	return _LoopringExchangev3.Contract.IsInWithdrawalMode(&_LoopringExchangev3.CallOpts)
}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsInWithdrawalMode() (bool, error) {
	return _LoopringExchangev3.Contract.IsInWithdrawalMode(&_LoopringExchangev3.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsShutdown() (bool, error) {
	return _LoopringExchangev3.Contract.IsShutdown(&_LoopringExchangev3.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsShutdown() (bool, error) {
	return _LoopringExchangev3.Contract.IsShutdown(&_LoopringExchangev3.CallOpts)
}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsTransactionApproved(opts *bind.CallOpts, owner common.Address, transactionHash [32]byte) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isTransactionApproved", owner, transactionHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsTransactionApproved(owner common.Address, transactionHash [32]byte) (bool, error) {
	return _LoopringExchangev3.Contract.IsTransactionApproved(&_LoopringExchangev3.CallOpts, owner, transactionHash)
}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsTransactionApproved(owner common.Address, transactionHash [32]byte) (bool, error) {
	return _LoopringExchangev3.Contract.IsTransactionApproved(&_LoopringExchangev3.CallOpts, owner, transactionHash)
}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsUserOrAgent(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isUserOrAgent", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsUserOrAgent(owner common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsUserOrAgent(&_LoopringExchangev3.CallOpts, owner)
}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsUserOrAgent(owner common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsUserOrAgent(&_LoopringExchangev3.CallOpts, owner)
}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Caller) IsWithdrawnInWithdrawalMode(opts *bind.CallOpts, accountID uint32, token common.Address) (bool, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "isWithdrawnInWithdrawalMode", accountID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3Session) IsWithdrawnInWithdrawalMode(accountID uint32, token common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsWithdrawnInWithdrawalMode(&_LoopringExchangev3.CallOpts, accountID, token)
}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) IsWithdrawnInWithdrawalMode(accountID uint32, token common.Address) (bool, error) {
	return _LoopringExchangev3.Contract.IsWithdrawnInWithdrawalMode(&_LoopringExchangev3.CallOpts, accountID, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) Owner() (common.Address, error) {
	return _LoopringExchangev3.Contract.Owner(&_LoopringExchangev3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) Owner() (common.Address, error) {
	return _LoopringExchangev3.Contract.Owner(&_LoopringExchangev3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3Session) PendingOwner() (common.Address, error) {
	return _LoopringExchangev3.Contract.PendingOwner(&_LoopringExchangev3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) PendingOwner() (common.Address, error) {
	return _LoopringExchangev3.Contract.PendingOwner(&_LoopringExchangev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoopringExchangev3 *LoopringExchangev3Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LoopringExchangev3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoopringExchangev3 *LoopringExchangev3Session) Version() (string, error) {
	return _LoopringExchangev3.Contract.Version(&_LoopringExchangev3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_LoopringExchangev3 *LoopringExchangev3CallerSession) Version() (string, error) {
	return _LoopringExchangev3.Contract.Version(&_LoopringExchangev3.CallOpts)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) ApproveTransaction(opts *bind.TransactOpts, owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "approveTransaction", owner, transactionHash)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) ApproveTransaction(owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ApproveTransaction(&_LoopringExchangev3.TransactOpts, owner, transactionHash)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) ApproveTransaction(owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ApproveTransaction(&_LoopringExchangev3.TransactOpts, owner, transactionHash)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) ApproveTransactions(opts *bind.TransactOpts, owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "approveTransactions", owners, transactionHashes)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) ApproveTransactions(owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ApproveTransactions(&_LoopringExchangev3.TransactOpts, owners, transactionHashes)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) ApproveTransactions(owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ApproveTransactions(&_LoopringExchangev3.TransactOpts, owners, transactionHashes)
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) BurnExchangeStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "burnExchangeStake")
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) BurnExchangeStake() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.BurnExchangeStake(&_LoopringExchangev3.TransactOpts)
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) BurnExchangeStake() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.BurnExchangeStake(&_LoopringExchangev3.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) ClaimOwnership() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ClaimOwnership(&_LoopringExchangev3.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ClaimOwnership(&_LoopringExchangev3.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb1a417f4.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, bytes extraData) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) Deposit(opts *bind.TransactOpts, from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "deposit", from, to, tokenAddress, amount, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0xb1a417f4.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, bytes extraData) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Deposit(&_LoopringExchangev3.TransactOpts, from, to, tokenAddress, amount, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0xb1a417f4.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint96 amount, bytes extraData) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Deposit(&_LoopringExchangev3.TransactOpts, from, to, tokenAddress, amount, extraData)
}

// DepositNFT is a paid mutator transaction binding the contract method 0x7d54f248.
//
// Solidity: function depositNFT(address from, address to, uint8 nftType, address tokenAddress, uint256 nftID, uint96 amount, bytes extraData) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) DepositNFT(opts *bind.TransactOpts, from common.Address, to common.Address, nftType uint8, tokenAddress common.Address, nftID *big.Int, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "depositNFT", from, to, nftType, tokenAddress, nftID, amount, extraData)
}

// DepositNFT is a paid mutator transaction binding the contract method 0x7d54f248.
//
// Solidity: function depositNFT(address from, address to, uint8 nftType, address tokenAddress, uint256 nftID, uint96 amount, bytes extraData) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) DepositNFT(from common.Address, to common.Address, nftType uint8, tokenAddress common.Address, nftID *big.Int, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.DepositNFT(&_LoopringExchangev3.TransactOpts, from, to, nftType, tokenAddress, nftID, amount, extraData)
}

// DepositNFT is a paid mutator transaction binding the contract method 0x7d54f248.
//
// Solidity: function depositNFT(address from, address to, uint8 nftType, address tokenAddress, uint256 nftID, uint96 amount, bytes extraData) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) DepositNFT(from common.Address, to common.Address, nftType uint8, tokenAddress common.Address, nftID *big.Int, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.DepositNFT(&_LoopringExchangev3.TransactOpts, from, to, nftType, tokenAddress, nftID, amount, extraData)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) ForceWithdraw(opts *bind.TransactOpts, owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "forceWithdraw", owner, token, accountID)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) ForceWithdraw(owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ForceWithdraw(&_LoopringExchangev3.TransactOpts, owner, token, accountID)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) ForceWithdraw(owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ForceWithdraw(&_LoopringExchangev3.TransactOpts, owner, token, accountID)
}

// ForceWithdrawByTokenID is a paid mutator transaction binding the contract method 0xcdb4e8c3.
//
// Solidity: function forceWithdrawByTokenID(address owner, uint16 tokenID, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) ForceWithdrawByTokenID(opts *bind.TransactOpts, owner common.Address, tokenID uint16, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "forceWithdrawByTokenID", owner, tokenID, accountID)
}

// ForceWithdrawByTokenID is a paid mutator transaction binding the contract method 0xcdb4e8c3.
//
// Solidity: function forceWithdrawByTokenID(address owner, uint16 tokenID, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) ForceWithdrawByTokenID(owner common.Address, tokenID uint16, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ForceWithdrawByTokenID(&_LoopringExchangev3.TransactOpts, owner, tokenID, accountID)
}

// ForceWithdrawByTokenID is a paid mutator transaction binding the contract method 0xcdb4e8c3.
//
// Solidity: function forceWithdrawByTokenID(address owner, uint16 tokenID, uint32 accountID) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) ForceWithdrawByTokenID(owner common.Address, tokenID uint16, accountID uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.ForceWithdrawByTokenID(&_LoopringExchangev3.TransactOpts, owner, tokenID, accountID)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) Initialize(opts *bind.TransactOpts, _loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "initialize", _loopring, _owner, _genesisMerkleRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) Initialize(_loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Initialize(&_LoopringExchangev3.TransactOpts, _loopring, _owner, _genesisMerkleRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) Initialize(_loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Initialize(&_LoopringExchangev3.TransactOpts, _loopring, _owner, _genesisMerkleRoot)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0xde6ff7cd.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, uint16 tokenID) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) NotifyForcedRequestTooOld(opts *bind.TransactOpts, accountID uint32, tokenID uint16) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "notifyForcedRequestTooOld", accountID, tokenID)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0xde6ff7cd.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, uint16 tokenID) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) NotifyForcedRequestTooOld(accountID uint32, tokenID uint16) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.NotifyForcedRequestTooOld(&_LoopringExchangev3.TransactOpts, accountID, tokenID)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0xde6ff7cd.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, uint16 tokenID) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) NotifyForcedRequestTooOld(accountID uint32, tokenID uint16) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.NotifyForcedRequestTooOld(&_LoopringExchangev3.TransactOpts, accountID, tokenID)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC1155BatchReceived(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC1155BatchReceived(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC1155Received(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC1155Received(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC721Received(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnERC721Received(&_LoopringExchangev3.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) OnchainTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "onchainTransferFrom", from, to, token, amount)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) OnchainTransferFrom(from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnchainTransferFrom(&_LoopringExchangev3.TransactOpts, from, to, token, amount)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) OnchainTransferFrom(from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.OnchainTransferFrom(&_LoopringExchangev3.TransactOpts, from, to, token, amount)
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) RefreshBlockVerifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "refreshBlockVerifier")
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) RefreshBlockVerifier() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RefreshBlockVerifier(&_LoopringExchangev3.TransactOpts)
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) RefreshBlockVerifier() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RefreshBlockVerifier(&_LoopringExchangev3.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) RegisterToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "registerToken", tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3Session) RegisterToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RegisterToken(&_LoopringExchangev3.TransactOpts, tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint16)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) RegisterToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RegisterToken(&_LoopringExchangev3.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) RenounceOwnership() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RenounceOwnership(&_LoopringExchangev3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.RenounceOwnership(&_LoopringExchangev3.TransactOpts)
}

// SetAgentRegistry is a paid mutator transaction binding the contract method 0x28342ecf.
//
// Solidity: function setAgentRegistry(address _agentRegistry) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetAgentRegistry(opts *bind.TransactOpts, _agentRegistry common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setAgentRegistry", _agentRegistry)
}

// SetAgentRegistry is a paid mutator transaction binding the contract method 0x28342ecf.
//
// Solidity: function setAgentRegistry(address _agentRegistry) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SetAgentRegistry(_agentRegistry common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAgentRegistry(&_LoopringExchangev3.TransactOpts, _agentRegistry)
}

// SetAgentRegistry is a paid mutator transaction binding the contract method 0x28342ecf.
//
// Solidity: function setAgentRegistry(address _agentRegistry) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetAgentRegistry(_agentRegistry common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAgentRegistry(&_LoopringExchangev3.TransactOpts, _agentRegistry)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetAllowOnchainTransferFrom(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setAllowOnchainTransferFrom", value)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SetAllowOnchainTransferFrom(value bool) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAllowOnchainTransferFrom(&_LoopringExchangev3.TransactOpts, value)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetAllowOnchainTransferFrom(value bool) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAllowOnchainTransferFrom(&_LoopringExchangev3.TransactOpts, value)
}

// SetAmmFeeBips is a paid mutator transaction binding the contract method 0xcf4bc4c1.
//
// Solidity: function setAmmFeeBips(uint8 _feeBips) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetAmmFeeBips(opts *bind.TransactOpts, _feeBips uint8) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setAmmFeeBips", _feeBips)
}

// SetAmmFeeBips is a paid mutator transaction binding the contract method 0xcf4bc4c1.
//
// Solidity: function setAmmFeeBips(uint8 _feeBips) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SetAmmFeeBips(_feeBips uint8) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAmmFeeBips(&_LoopringExchangev3.TransactOpts, _feeBips)
}

// SetAmmFeeBips is a paid mutator transaction binding the contract method 0xcf4bc4c1.
//
// Solidity: function setAmmFeeBips(uint8 _feeBips) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetAmmFeeBips(_feeBips uint8) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetAmmFeeBips(&_LoopringExchangev3.TransactOpts, _feeBips)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetDepositContract(opts *bind.TransactOpts, _depositContract common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setDepositContract", _depositContract)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SetDepositContract(_depositContract common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetDepositContract(&_LoopringExchangev3.TransactOpts, _depositContract)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetDepositContract(_depositContract common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetDepositContract(&_LoopringExchangev3.TransactOpts, _depositContract)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetMaxAgeDepositUntilWithdrawable(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setMaxAgeDepositUntilWithdrawable", newValue)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3Session) SetMaxAgeDepositUntilWithdrawable(newValue uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetMaxAgeDepositUntilWithdrawable(&_LoopringExchangev3.TransactOpts, newValue)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetMaxAgeDepositUntilWithdrawable(newValue uint32) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetMaxAgeDepositUntilWithdrawable(&_LoopringExchangev3.TransactOpts, newValue)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0xd5b039ce.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID, address newRecipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SetWithdrawalRecipient(opts *bind.TransactOpts, from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "setWithdrawalRecipient", from, to, token, amount, storageID, newRecipient)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0xd5b039ce.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID, address newRecipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetWithdrawalRecipient(&_LoopringExchangev3.TransactOpts, from, to, token, amount, storageID, newRecipient)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0xd5b039ce.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint96 amount, uint32 storageID, address newRecipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SetWithdrawalRecipient(&_LoopringExchangev3.TransactOpts, from, to, token, amount, storageID, newRecipient)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_LoopringExchangev3 *LoopringExchangev3Session) Shutdown() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Shutdown(&_LoopringExchangev3.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) Shutdown() (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.Shutdown(&_LoopringExchangev3.TransactOpts)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) SubmitBlocks(opts *bind.TransactOpts, blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "submitBlocks", blocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) SubmitBlocks(blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SubmitBlocks(&_LoopringExchangev3.TransactOpts, blocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) SubmitBlocks(blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.SubmitBlocks(&_LoopringExchangev3.TransactOpts, blocks)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.TransferOwnership(&_LoopringExchangev3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.TransferOwnership(&_LoopringExchangev3.TransactOpts, newOwner)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawExchangeFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawExchangeFees", token, recipient)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawExchangeFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawExchangeFees(&_LoopringExchangev3.TransactOpts, token, recipient)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawExchangeFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawExchangeFees(&_LoopringExchangev3.TransactOpts, token, recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawExchangeStake(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawExchangeStake", recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawExchangeStake(recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawExchangeStake(&_LoopringExchangev3.TransactOpts, recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawExchangeStake(recipient common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawExchangeStake(&_LoopringExchangev3.TransactOpts, recipient)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawFromApprovedWithdrawals(opts *bind.TransactOpts, owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawFromApprovedWithdrawals", owners, tokens)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawFromApprovedWithdrawals(owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromApprovedWithdrawals(&_LoopringExchangev3.TransactOpts, owners, tokens)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawFromApprovedWithdrawals(owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromApprovedWithdrawals(&_LoopringExchangev3.TransactOpts, owners, tokens)
}

// WithdrawFromApprovedWithdrawalsNFT is a paid mutator transaction binding the contract method 0x327e965e.
//
// Solidity: function withdrawFromApprovedWithdrawalsNFT(address[] owners, address[] minters, uint8[] nftTypes, address[] tokens, uint256[] nftIDs) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawFromApprovedWithdrawalsNFT(opts *bind.TransactOpts, owners []common.Address, minters []common.Address, nftTypes []uint8, tokens []common.Address, nftIDs []*big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawFromApprovedWithdrawalsNFT", owners, minters, nftTypes, tokens, nftIDs)
}

// WithdrawFromApprovedWithdrawalsNFT is a paid mutator transaction binding the contract method 0x327e965e.
//
// Solidity: function withdrawFromApprovedWithdrawalsNFT(address[] owners, address[] minters, uint8[] nftTypes, address[] tokens, uint256[] nftIDs) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawFromApprovedWithdrawalsNFT(owners []common.Address, minters []common.Address, nftTypes []uint8, tokens []common.Address, nftIDs []*big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromApprovedWithdrawalsNFT(&_LoopringExchangev3.TransactOpts, owners, minters, nftTypes, tokens, nftIDs)
}

// WithdrawFromApprovedWithdrawalsNFT is a paid mutator transaction binding the contract method 0x327e965e.
//
// Solidity: function withdrawFromApprovedWithdrawalsNFT(address[] owners, address[] minters, uint8[] nftTypes, address[] tokens, uint256[] nftIDs) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawFromApprovedWithdrawalsNFT(owners []common.Address, minters []common.Address, nftTypes []uint8, tokens []common.Address, nftIDs []*big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromApprovedWithdrawalsNFT(&_LoopringExchangev3.TransactOpts, owners, minters, nftTypes, tokens, nftIDs)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawFromDepositRequest(opts *bind.TransactOpts, owner common.Address, token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawFromDepositRequest", owner, token)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawFromDepositRequest(owner common.Address, token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromDepositRequest(&_LoopringExchangev3.TransactOpts, owner, token)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawFromDepositRequest(owner common.Address, token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromDepositRequest(&_LoopringExchangev3.TransactOpts, owner, token)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x93b80987.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32,uint256),(uint16,uint96,uint256,uint256),(address,uint8,address,uint256,uint8),uint256[48],uint256[24]) merkleProof) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawFromMerkleTree(opts *bind.TransactOpts, merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawFromMerkleTree", merkleProof)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x93b80987.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32,uint256),(uint16,uint96,uint256,uint256),(address,uint8,address,uint256,uint8),uint256[48],uint256[24]) merkleProof) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawFromMerkleTree(merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromMerkleTree(&_LoopringExchangev3.TransactOpts, merkleProof)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x93b80987.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32,uint256),(uint16,uint96,uint256,uint256),(address,uint8,address,uint256,uint8),uint256[48],uint256[24]) merkleProof) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawFromMerkleTree(merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromMerkleTree(&_LoopringExchangev3.TransactOpts, merkleProof)
}

// WithdrawFromNFTDepositRequest is a paid mutator transaction binding the contract method 0x0394bc2b.
//
// Solidity: function withdrawFromNFTDepositRequest(address owner, address token, uint8 nftType, uint256 nftID) returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawFromNFTDepositRequest(opts *bind.TransactOpts, owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawFromNFTDepositRequest", owner, token, nftType, nftID)
}

// WithdrawFromNFTDepositRequest is a paid mutator transaction binding the contract method 0x0394bc2b.
//
// Solidity: function withdrawFromNFTDepositRequest(address owner, address token, uint8 nftType, uint256 nftID) returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawFromNFTDepositRequest(owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromNFTDepositRequest(&_LoopringExchangev3.TransactOpts, owner, token, nftType, nftID)
}

// WithdrawFromNFTDepositRequest is a paid mutator transaction binding the contract method 0x0394bc2b.
//
// Solidity: function withdrawFromNFTDepositRequest(address owner, address token, uint8 nftType, uint256 nftID) returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawFromNFTDepositRequest(owner common.Address, token common.Address, nftType uint8, nftID *big.Int) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawFromNFTDepositRequest(&_LoopringExchangev3.TransactOpts, owner, token, nftType, nftID)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Transactor) WithdrawProtocolFees(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.contract.Transact(opts, "withdrawProtocolFees", token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3Session) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawProtocolFees(&_LoopringExchangev3.TransactOpts, token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_LoopringExchangev3 *LoopringExchangev3TransactorSession) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _LoopringExchangev3.Contract.WithdrawProtocolFees(&_LoopringExchangev3.TransactOpts, token)
}

// LoopringExchangev3BlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the LoopringExchangev3 contract.
type LoopringExchangev3BlockSubmittedIterator struct {
	Event *LoopringExchangev3BlockSubmitted // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3BlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3BlockSubmitted)
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
		it.Event = new(LoopringExchangev3BlockSubmitted)
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
func (it *LoopringExchangev3BlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3BlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3BlockSubmitted represents a BlockSubmitted event raised by the LoopringExchangev3 contract.
type LoopringExchangev3BlockSubmitted struct {
	BlockIdx       *big.Int
	MerkleRoot     [32]byte
	PublicDataHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterBlockSubmitted(opts *bind.FilterOpts, blockIdx []*big.Int) (*LoopringExchangev3BlockSubmittedIterator, error) {

	var blockIdxRule []interface{}
	for _, blockIdxItem := range blockIdx {
		blockIdxRule = append(blockIdxRule, blockIdxItem)
	}

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "BlockSubmitted", blockIdxRule)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3BlockSubmittedIterator{contract: _LoopringExchangev3.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3BlockSubmitted, blockIdx []*big.Int) (event.Subscription, error) {

	var blockIdxRule []interface{}
	for _, blockIdxItem := range blockIdx {
		blockIdxRule = append(blockIdxRule, blockIdxItem)
	}

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "BlockSubmitted", blockIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3BlockSubmitted)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseBlockSubmitted(log types.Log) (*LoopringExchangev3BlockSubmitted, error) {
	event := new(LoopringExchangev3BlockSubmitted)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3DepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the LoopringExchangev3 contract.
type LoopringExchangev3DepositRequestedIterator struct {
	Event *LoopringExchangev3DepositRequested // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3DepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3DepositRequested)
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
		it.Event = new(LoopringExchangev3DepositRequested)
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
func (it *LoopringExchangev3DepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3DepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3DepositRequested represents a DepositRequested event raised by the LoopringExchangev3 contract.
type LoopringExchangev3DepositRequested struct {
	From    common.Address
	To      common.Address
	Token   common.Address
	TokenId uint16
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0x73ff7b101bcdc22f199e8e1dd9893170a683d6897be4f1086ca05705abb886ae.
//
// Solidity: event DepositRequested(address from, address to, address token, uint16 tokenId, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterDepositRequested(opts *bind.FilterOpts) (*LoopringExchangev3DepositRequestedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3DepositRequestedIterator{contract: _LoopringExchangev3.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0x73ff7b101bcdc22f199e8e1dd9893170a683d6897be4f1086ca05705abb886ae.
//
// Solidity: event DepositRequested(address from, address to, address token, uint16 tokenId, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3DepositRequested) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3DepositRequested)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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

// ParseDepositRequested is a log parse operation binding the contract event 0x73ff7b101bcdc22f199e8e1dd9893170a683d6897be4f1086ca05705abb886ae.
//
// Solidity: event DepositRequested(address from, address to, address token, uint16 tokenId, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseDepositRequested(log types.Log) (*LoopringExchangev3DepositRequested, error) {
	event := new(LoopringExchangev3DepositRequested)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3ExchangeClonedIterator is returned from FilterExchangeCloned and is used to iterate over the raw logs and unpacked data for ExchangeCloned events raised by the LoopringExchangev3 contract.
type LoopringExchangev3ExchangeClonedIterator struct {
	Event *LoopringExchangev3ExchangeCloned // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3ExchangeClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3ExchangeCloned)
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
		it.Event = new(LoopringExchangev3ExchangeCloned)
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
func (it *LoopringExchangev3ExchangeClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3ExchangeClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3ExchangeCloned represents a ExchangeCloned event raised by the LoopringExchangev3 contract.
type LoopringExchangev3ExchangeCloned struct {
	ExchangeAddress   common.Address
	Owner             common.Address
	GenesisMerkleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterExchangeCloned is a free log retrieval operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterExchangeCloned(opts *bind.FilterOpts) (*LoopringExchangev3ExchangeClonedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "ExchangeCloned")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3ExchangeClonedIterator{contract: _LoopringExchangev3.contract, event: "ExchangeCloned", logs: logs, sub: sub}, nil
}

// WatchExchangeCloned is a free log subscription operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchExchangeCloned(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3ExchangeCloned) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "ExchangeCloned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3ExchangeCloned)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "ExchangeCloned", log); err != nil {
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

// ParseExchangeCloned is a log parse operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseExchangeCloned(log types.Log) (*LoopringExchangev3ExchangeCloned, error) {
	event := new(LoopringExchangev3ExchangeCloned)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "ExchangeCloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3ForcedWithdrawalRequestedIterator is returned from FilterForcedWithdrawalRequested and is used to iterate over the raw logs and unpacked data for ForcedWithdrawalRequested events raised by the LoopringExchangev3 contract.
type LoopringExchangev3ForcedWithdrawalRequestedIterator struct {
	Event *LoopringExchangev3ForcedWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3ForcedWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3ForcedWithdrawalRequested)
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
		it.Event = new(LoopringExchangev3ForcedWithdrawalRequested)
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
func (it *LoopringExchangev3ForcedWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3ForcedWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3ForcedWithdrawalRequested represents a ForcedWithdrawalRequested event raised by the LoopringExchangev3 contract.
type LoopringExchangev3ForcedWithdrawalRequested struct {
	Owner     common.Address
	TokenID   uint16
	AccountID uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForcedWithdrawalRequested is a free log retrieval operation binding the contract event 0x9cdf733df701225f835cfd5df1c6376ef646298f9c2271e86f12078903e26094.
//
// Solidity: event ForcedWithdrawalRequested(address owner, uint16 tokenID, uint32 accountID)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterForcedWithdrawalRequested(opts *bind.FilterOpts) (*LoopringExchangev3ForcedWithdrawalRequestedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "ForcedWithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3ForcedWithdrawalRequestedIterator{contract: _LoopringExchangev3.contract, event: "ForcedWithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchForcedWithdrawalRequested is a free log subscription operation binding the contract event 0x9cdf733df701225f835cfd5df1c6376ef646298f9c2271e86f12078903e26094.
//
// Solidity: event ForcedWithdrawalRequested(address owner, uint16 tokenID, uint32 accountID)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchForcedWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3ForcedWithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "ForcedWithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3ForcedWithdrawalRequested)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "ForcedWithdrawalRequested", log); err != nil {
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

// ParseForcedWithdrawalRequested is a log parse operation binding the contract event 0x9cdf733df701225f835cfd5df1c6376ef646298f9c2271e86f12078903e26094.
//
// Solidity: event ForcedWithdrawalRequested(address owner, uint16 tokenID, uint32 accountID)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseForcedWithdrawalRequested(log types.Log) (*LoopringExchangev3ForcedWithdrawalRequested, error) {
	event := new(LoopringExchangev3ForcedWithdrawalRequested)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "ForcedWithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3NFTDepositRequestedIterator is returned from FilterNFTDepositRequested and is used to iterate over the raw logs and unpacked data for NFTDepositRequested events raised by the LoopringExchangev3 contract.
type LoopringExchangev3NFTDepositRequestedIterator struct {
	Event *LoopringExchangev3NFTDepositRequested // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3NFTDepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3NFTDepositRequested)
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
		it.Event = new(LoopringExchangev3NFTDepositRequested)
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
func (it *LoopringExchangev3NFTDepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3NFTDepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3NFTDepositRequested represents a NFTDepositRequested event raised by the LoopringExchangev3 contract.
type LoopringExchangev3NFTDepositRequested struct {
	From    common.Address
	To      common.Address
	NftType uint8
	Token   common.Address
	NftID   *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTDepositRequested is a free log retrieval operation binding the contract event 0x59c0422c484ae5502249fd758097d16508ef0e61b5f7f962bef8145a58701287.
//
// Solidity: event NFTDepositRequested(address from, address to, uint8 nftType, address token, uint256 nftID, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterNFTDepositRequested(opts *bind.FilterOpts) (*LoopringExchangev3NFTDepositRequestedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "NFTDepositRequested")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3NFTDepositRequestedIterator{contract: _LoopringExchangev3.contract, event: "NFTDepositRequested", logs: logs, sub: sub}, nil
}

// WatchNFTDepositRequested is a free log subscription operation binding the contract event 0x59c0422c484ae5502249fd758097d16508ef0e61b5f7f962bef8145a58701287.
//
// Solidity: event NFTDepositRequested(address from, address to, uint8 nftType, address token, uint256 nftID, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchNFTDepositRequested(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3NFTDepositRequested) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "NFTDepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3NFTDepositRequested)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "NFTDepositRequested", log); err != nil {
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

// ParseNFTDepositRequested is a log parse operation binding the contract event 0x59c0422c484ae5502249fd758097d16508ef0e61b5f7f962bef8145a58701287.
//
// Solidity: event NFTDepositRequested(address from, address to, uint8 nftType, address token, uint256 nftID, uint96 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseNFTDepositRequested(log types.Log) (*LoopringExchangev3NFTDepositRequested, error) {
	event := new(LoopringExchangev3NFTDepositRequested)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "NFTDepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3NftWithdrawalCompletedIterator is returned from FilterNftWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for NftWithdrawalCompleted events raised by the LoopringExchangev3 contract.
type LoopringExchangev3NftWithdrawalCompletedIterator struct {
	Event *LoopringExchangev3NftWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3NftWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3NftWithdrawalCompleted)
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
		it.Event = new(LoopringExchangev3NftWithdrawalCompleted)
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
func (it *LoopringExchangev3NftWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3NftWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3NftWithdrawalCompleted represents a NftWithdrawalCompleted event raised by the LoopringExchangev3 contract.
type LoopringExchangev3NftWithdrawalCompleted struct {
	Category uint8
	From     common.Address
	To       common.Address
	TokenID  uint16
	Token    common.Address
	NftID    *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNftWithdrawalCompleted is a free log retrieval operation binding the contract event 0xa1af8cedfbd157f967dd83203fd025ece9cd1ee0f63f38632b24f1e03778bc28.
//
// Solidity: event NftWithdrawalCompleted(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterNftWithdrawalCompleted(opts *bind.FilterOpts) (*LoopringExchangev3NftWithdrawalCompletedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "NftWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3NftWithdrawalCompletedIterator{contract: _LoopringExchangev3.contract, event: "NftWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchNftWithdrawalCompleted is a free log subscription operation binding the contract event 0xa1af8cedfbd157f967dd83203fd025ece9cd1ee0f63f38632b24f1e03778bc28.
//
// Solidity: event NftWithdrawalCompleted(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchNftWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3NftWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "NftWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3NftWithdrawalCompleted)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "NftWithdrawalCompleted", log); err != nil {
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

// ParseNftWithdrawalCompleted is a log parse operation binding the contract event 0xa1af8cedfbd157f967dd83203fd025ece9cd1ee0f63f38632b24f1e03778bc28.
//
// Solidity: event NftWithdrawalCompleted(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseNftWithdrawalCompleted(log types.Log) (*LoopringExchangev3NftWithdrawalCompleted, error) {
	event := new(LoopringExchangev3NftWithdrawalCompleted)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "NftWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3NftWithdrawalFailedIterator is returned from FilterNftWithdrawalFailed and is used to iterate over the raw logs and unpacked data for NftWithdrawalFailed events raised by the LoopringExchangev3 contract.
type LoopringExchangev3NftWithdrawalFailedIterator struct {
	Event *LoopringExchangev3NftWithdrawalFailed // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3NftWithdrawalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3NftWithdrawalFailed)
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
		it.Event = new(LoopringExchangev3NftWithdrawalFailed)
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
func (it *LoopringExchangev3NftWithdrawalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3NftWithdrawalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3NftWithdrawalFailed represents a NftWithdrawalFailed event raised by the LoopringExchangev3 contract.
type LoopringExchangev3NftWithdrawalFailed struct {
	Category uint8
	From     common.Address
	To       common.Address
	TokenID  uint16
	Token    common.Address
	NftID    *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNftWithdrawalFailed is a free log retrieval operation binding the contract event 0x7cfb0fa39777da32fa26ee2f1cb14cc1c1ecd01e99d1d6bba446143f96db036f.
//
// Solidity: event NftWithdrawalFailed(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterNftWithdrawalFailed(opts *bind.FilterOpts) (*LoopringExchangev3NftWithdrawalFailedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "NftWithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3NftWithdrawalFailedIterator{contract: _LoopringExchangev3.contract, event: "NftWithdrawalFailed", logs: logs, sub: sub}, nil
}

// WatchNftWithdrawalFailed is a free log subscription operation binding the contract event 0x7cfb0fa39777da32fa26ee2f1cb14cc1c1ecd01e99d1d6bba446143f96db036f.
//
// Solidity: event NftWithdrawalFailed(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchNftWithdrawalFailed(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3NftWithdrawalFailed) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "NftWithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3NftWithdrawalFailed)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "NftWithdrawalFailed", log); err != nil {
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

// ParseNftWithdrawalFailed is a log parse operation binding the contract event 0x7cfb0fa39777da32fa26ee2f1cb14cc1c1ecd01e99d1d6bba446143f96db036f.
//
// Solidity: event NftWithdrawalFailed(uint8 category, address from, address to, uint16 tokenID, address token, uint256 nftID, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseNftWithdrawalFailed(log types.Log) (*LoopringExchangev3NftWithdrawalFailed, error) {
	event := new(LoopringExchangev3NftWithdrawalFailed)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "NftWithdrawalFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoopringExchangev3 contract.
type LoopringExchangev3OwnershipTransferredIterator struct {
	Event *LoopringExchangev3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3OwnershipTransferred)
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
		it.Event = new(LoopringExchangev3OwnershipTransferred)
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
func (it *LoopringExchangev3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3OwnershipTransferred represents a OwnershipTransferred event raised by the LoopringExchangev3 contract.
type LoopringExchangev3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoopringExchangev3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3OwnershipTransferredIterator{contract: _LoopringExchangev3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3OwnershipTransferred)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseOwnershipTransferred(log types.Log) (*LoopringExchangev3OwnershipTransferred, error) {
	event := new(LoopringExchangev3OwnershipTransferred)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3ProtocolFeesUpdatedIterator is returned from FilterProtocolFeesUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeesUpdated events raised by the LoopringExchangev3 contract.
type LoopringExchangev3ProtocolFeesUpdatedIterator struct {
	Event *LoopringExchangev3ProtocolFeesUpdated // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3ProtocolFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3ProtocolFeesUpdated)
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
		it.Event = new(LoopringExchangev3ProtocolFeesUpdated)
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
func (it *LoopringExchangev3ProtocolFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3ProtocolFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3ProtocolFeesUpdated represents a ProtocolFeesUpdated event raised by the LoopringExchangev3 contract.
type LoopringExchangev3ProtocolFeesUpdated struct {
	TakerFeeBips         uint8
	MakerFeeBips         uint8
	PreviousTakerFeeBips uint8
	PreviousMakerFeeBips uint8
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeesUpdated is a free log retrieval operation binding the contract event 0xf953e7a938acd4ba0571a50770e0a53fdb8d8618fe5afb1519ab7c76323ac70b.
//
// Solidity: event ProtocolFeesUpdated(uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterProtocolFeesUpdated(opts *bind.FilterOpts) (*LoopringExchangev3ProtocolFeesUpdatedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "ProtocolFeesUpdated")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3ProtocolFeesUpdatedIterator{contract: _LoopringExchangev3.contract, event: "ProtocolFeesUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeesUpdated is a free log subscription operation binding the contract event 0xf953e7a938acd4ba0571a50770e0a53fdb8d8618fe5afb1519ab7c76323ac70b.
//
// Solidity: event ProtocolFeesUpdated(uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchProtocolFeesUpdated(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3ProtocolFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "ProtocolFeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3ProtocolFeesUpdated)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "ProtocolFeesUpdated", log); err != nil {
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

// ParseProtocolFeesUpdated is a log parse operation binding the contract event 0xf953e7a938acd4ba0571a50770e0a53fdb8d8618fe5afb1519ab7c76323ac70b.
//
// Solidity: event ProtocolFeesUpdated(uint8 takerFeeBips, uint8 makerFeeBips, uint8 previousTakerFeeBips, uint8 previousMakerFeeBips)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseProtocolFeesUpdated(log types.Log) (*LoopringExchangev3ProtocolFeesUpdated, error) {
	event := new(LoopringExchangev3ProtocolFeesUpdated)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "ProtocolFeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3ShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the LoopringExchangev3 contract.
type LoopringExchangev3ShutdownIterator struct {
	Event *LoopringExchangev3Shutdown // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3ShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3Shutdown)
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
		it.Event = new(LoopringExchangev3Shutdown)
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
func (it *LoopringExchangev3ShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3ShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3Shutdown represents a Shutdown event raised by the LoopringExchangev3 contract.
type LoopringExchangev3Shutdown struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterShutdown(opts *bind.FilterOpts) (*LoopringExchangev3ShutdownIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3ShutdownIterator{contract: _LoopringExchangev3.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3Shutdown) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3Shutdown)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseShutdown(log types.Log) (*LoopringExchangev3Shutdown, error) {
	event := new(LoopringExchangev3Shutdown)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3TokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the LoopringExchangev3 contract.
type LoopringExchangev3TokenRegisteredIterator struct {
	Event *LoopringExchangev3TokenRegistered // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3TokenRegistered)
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
		it.Event = new(LoopringExchangev3TokenRegistered)
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
func (it *LoopringExchangev3TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3TokenRegistered represents a TokenRegistered event raised by the LoopringExchangev3 contract.
type LoopringExchangev3TokenRegistered struct {
	Token   common.Address
	TokenId uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0x678eb1f52e3e76c52c1143d5699af1d8a5e03743e8840223ccb27aabf3c44c0a.
//
// Solidity: event TokenRegistered(address token, uint16 tokenId)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterTokenRegistered(opts *bind.FilterOpts) (*LoopringExchangev3TokenRegisteredIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3TokenRegisteredIterator{contract: _LoopringExchangev3.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0x678eb1f52e3e76c52c1143d5699af1d8a5e03743e8840223ccb27aabf3c44c0a.
//
// Solidity: event TokenRegistered(address token, uint16 tokenId)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3TokenRegistered) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3TokenRegistered)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0x678eb1f52e3e76c52c1143d5699af1d8a5e03743e8840223ccb27aabf3c44c0a.
//
// Solidity: event TokenRegistered(address token, uint16 tokenId)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseTokenRegistered(log types.Log) (*LoopringExchangev3TokenRegistered, error) {
	event := new(LoopringExchangev3TokenRegistered)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3TransactionApprovedIterator is returned from FilterTransactionApproved and is used to iterate over the raw logs and unpacked data for TransactionApproved events raised by the LoopringExchangev3 contract.
type LoopringExchangev3TransactionApprovedIterator struct {
	Event *LoopringExchangev3TransactionApproved // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3TransactionApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3TransactionApproved)
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
		it.Event = new(LoopringExchangev3TransactionApproved)
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
func (it *LoopringExchangev3TransactionApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3TransactionApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3TransactionApproved represents a TransactionApproved event raised by the LoopringExchangev3 contract.
type LoopringExchangev3TransactionApproved struct {
	Owner           common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionApproved is a free log retrieval operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterTransactionApproved(opts *bind.FilterOpts) (*LoopringExchangev3TransactionApprovedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3TransactionApprovedIterator{contract: _LoopringExchangev3.contract, event: "TransactionApproved", logs: logs, sub: sub}, nil
}

// WatchTransactionApproved is a free log subscription operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchTransactionApproved(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3TransactionApproved) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3TransactionApproved)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
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

// ParseTransactionApproved is a log parse operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseTransactionApproved(log types.Log) (*LoopringExchangev3TransactionApproved, error) {
	event := new(LoopringExchangev3TransactionApproved)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3WithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalCompletedIterator struct {
	Event *LoopringExchangev3WithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3WithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3WithdrawalCompleted)
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
		it.Event = new(LoopringExchangev3WithdrawalCompleted)
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
func (it *LoopringExchangev3WithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3WithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3WithdrawalCompleted represents a WithdrawalCompleted event raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalCompleted struct {
	Category uint8
	From     common.Address
	To       common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*LoopringExchangev3WithdrawalCompletedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3WithdrawalCompletedIterator{contract: _LoopringExchangev3.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3WithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3WithdrawalCompleted)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
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

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseWithdrawalCompleted(log types.Log) (*LoopringExchangev3WithdrawalCompleted, error) {
	event := new(LoopringExchangev3WithdrawalCompleted)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3WithdrawalFailedIterator is returned from FilterWithdrawalFailed and is used to iterate over the raw logs and unpacked data for WithdrawalFailed events raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalFailedIterator struct {
	Event *LoopringExchangev3WithdrawalFailed // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3WithdrawalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3WithdrawalFailed)
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
		it.Event = new(LoopringExchangev3WithdrawalFailed)
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
func (it *LoopringExchangev3WithdrawalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3WithdrawalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3WithdrawalFailed represents a WithdrawalFailed event raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalFailed struct {
	Category uint8
	From     common.Address
	To       common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFailed is a free log retrieval operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterWithdrawalFailed(opts *bind.FilterOpts) (*LoopringExchangev3WithdrawalFailedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "WithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3WithdrawalFailedIterator{contract: _LoopringExchangev3.contract, event: "WithdrawalFailed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFailed is a free log subscription operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchWithdrawalFailed(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3WithdrawalFailed) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "WithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3WithdrawalFailed)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalFailed", log); err != nil {
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

// ParseWithdrawalFailed is a log parse operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseWithdrawalFailed(log types.Log) (*LoopringExchangev3WithdrawalFailed, error) {
	event := new(LoopringExchangev3WithdrawalFailed)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopringExchangev3WithdrawalModeActivatedIterator is returned from FilterWithdrawalModeActivated and is used to iterate over the raw logs and unpacked data for WithdrawalModeActivated events raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalModeActivatedIterator struct {
	Event *LoopringExchangev3WithdrawalModeActivated // Event containing the contract specifics and raw log

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
func (it *LoopringExchangev3WithdrawalModeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopringExchangev3WithdrawalModeActivated)
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
		it.Event = new(LoopringExchangev3WithdrawalModeActivated)
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
func (it *LoopringExchangev3WithdrawalModeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopringExchangev3WithdrawalModeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopringExchangev3WithdrawalModeActivated represents a WithdrawalModeActivated event raised by the LoopringExchangev3 contract.
type LoopringExchangev3WithdrawalModeActivated struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalModeActivated is a free log retrieval operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) FilterWithdrawalModeActivated(opts *bind.FilterOpts) (*LoopringExchangev3WithdrawalModeActivatedIterator, error) {

	logs, sub, err := _LoopringExchangev3.contract.FilterLogs(opts, "WithdrawalModeActivated")
	if err != nil {
		return nil, err
	}
	return &LoopringExchangev3WithdrawalModeActivatedIterator{contract: _LoopringExchangev3.contract, event: "WithdrawalModeActivated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalModeActivated is a free log subscription operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) WatchWithdrawalModeActivated(opts *bind.WatchOpts, sink chan<- *LoopringExchangev3WithdrawalModeActivated) (event.Subscription, error) {

	logs, sub, err := _LoopringExchangev3.contract.WatchLogs(opts, "WithdrawalModeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopringExchangev3WithdrawalModeActivated)
				if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalModeActivated", log); err != nil {
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

// ParseWithdrawalModeActivated is a log parse operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_LoopringExchangev3 *LoopringExchangev3Filterer) ParseWithdrawalModeActivated(log types.Log) (*LoopringExchangev3WithdrawalModeActivated, error) {
	event := new(LoopringExchangev3WithdrawalModeActivated)
	if err := _LoopringExchangev3.contract.UnpackLog(event, "WithdrawalModeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
