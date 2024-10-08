package drips

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/drips"
)

var (
	logDepositSigHash       = crypto.Keccak256Hash([]byte("Deposit(address,uint256)"))
	taiko                   = common.HexToAddress("0xa9d23408b9ba935c230493c40c73824df71a0975")
	taikoDecimals     uint8 = 18
)

type DripIndexer struct {
	address common.Address
}

func NewDripsIndexer(address common.Address) *DripIndexer {
	return &DripIndexer{
		address: address,
	}
}

func (indexer *DripIndexer) Address() []common.Address {
	return []common.Address{indexer.address}
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *DripIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Lock, error) {
	var result []adapters.Lock
	for _, vLog := range logs {
		if !isDeposit(vLog) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		if transferData != nil {
			result = append(result, *transferData)
		}
	}
	return result, nil
}

func isDeposit(vLog types.Log) bool {
	return vLog.Topics[0].Hex() == logDepositSigHash.Hex()
}

func (indexer *DripIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Lock, error) {
	user := common.BytesToAddress(vLog.Topics[1].Bytes()[12:])
	var depositEvent struct {
		Assets *big.Int
	}

	dripsABI, err := abi.JSON(strings.NewReader(drips.DripsABI))
	if err != nil {
		return nil, err
	}

	err = dripsABI.UnpackIntoInterface(&depositEvent, "Deposit", vLog.Data)
	if err != nil {
		return nil, err
	}

	// Fetch the block details
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	// Return the LPTransfer struct with calculated values
	return &adapters.Lock{
		User:          user,
		TokenAmount:   depositEvent.Assets,
		TokenDecimals: taikoDecimals,
		Token:         taiko,
		Time:          block.Time(),
		BlockNumber:   block.Number().Uint64(),
	}, nil
}
