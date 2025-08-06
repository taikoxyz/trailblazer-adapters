package symmetric_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/symmetric"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestLockIndexer(t *testing.T) {
	blocknumber := int64(667176)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := symmetric.NewLockIndexer(client, []common.Address{common.HexToAddress(symmetric.LockAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	locks, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, locks, 2)
	txHash := common.HexToHash("0x6de4309af3be9f566195c9c33f3ac6d46e7084856988f5695b5ec3d080511882")
	user := common.HexToAddress("0xFb4D6288D1c51292dC9899f5F876A2Cf1f9fef43")
	time := 1734021131
	duration := 8407669
	expectedLocks := []adapters.Lock{
		{
			Metadata: adapters.Metadata{
				BlockTime:   uint64(time),
				BlockNumber: uint64(blocknumber),
				TxHash:      txHash,
			},
			User:          user,
			TokenAmount:   big.NewInt(413346247810476),
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.WETHAddress),
			Duration:      uint64(duration),
		},
		{
			Metadata: adapters.Metadata{
				BlockTime:   uint64(time),
				BlockNumber: uint64(blocknumber),
				TxHash:      txHash,
			},
			User:          user,
			TokenAmount:   big.NewInt(2926854880867952131),
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			Duration:      uint64(duration),
		},
	}

	for i, lock := range locks {
		assert.Equal(t, expectedLocks[i].User, lock.User)
		assert.Equal(t, expectedLocks[i].TokenAmount, lock.TokenAmount)
		assert.Equal(t, expectedLocks[i].TokenDecimals, lock.TokenDecimals)
		assert.Equal(t, expectedLocks[i].Duration, lock.Duration)
		assert.Equal(t, expectedLocks[i].Token, lock.Token)
		assert.Equal(t, expectedLocks[i].BlockTime, lock.BlockTime)
		assert.Equal(t, expectedLocks[i].BlockNumber, lock.BlockNumber)
		assert.Equal(t, expectedLocks[i].TxHash, lock.TxHash)
	}
}
