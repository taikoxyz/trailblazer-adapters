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
)

func TestLockIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(445262)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := symmetric.NewLockIndexer(client, []common.Address{common.HexToAddress(symmetric.LockAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	locks, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, locks, 2)
	txHash := common.HexToHash("0x27058b3cfbd15e3466cdd1b86ab387a473f2ca01c0697eb6b47ae7b33fdf97b0")
	user := common.HexToAddress("0x7255Db0d1C1B93Fb756157074fa0613Aa6878F31")
	time := 1728397319
	expectedLocks := []adapters.Lock{
		{
			User:          user,
			TokenAmount:   big.NewInt(45884668717994),
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.WETHAddress),
			Time:          uint64(time),
			BlockNumber:   uint64(blocknumber),
			TxHash:        txHash,
		},
		{
			User:          user,
			TokenAmount:   big.NewInt(287649277362130342),
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			Time:          uint64(time),
			BlockNumber:   uint64(blocknumber),
			TxHash:        txHash,
		},
	}

	for i, lock := range locks {
		assert.Equal(t, expectedLocks[i].User, lock.User)
		assert.Equal(t, expectedLocks[i].TokenAmount, lock.TokenAmount)
		assert.Equal(t, expectedLocks[i].TokenDecimals, lock.TokenDecimals)
		assert.Equal(t, expectedLocks[i].Token, lock.Token)
		assert.Equal(t, expectedLocks[i].Time, lock.Time)
		assert.Equal(t, expectedLocks[i].BlockNumber, lock.BlockNumber)
		assert.Equal(t, expectedLocks[i].TxHash, lock.TxHash)
	}
}
