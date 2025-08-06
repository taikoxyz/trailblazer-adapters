package loopring_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopring"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestLockIndexer(t *testing.T) {
	blocknumber := int64(496553)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := loopring.NewLockIndexer(client, []common.Address{common.HexToAddress(loopring.LockAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	locks, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, locks, 1)
	assert.Equal(t, common.HexToAddress("0x6b1029C9AE8Aa5EEA9e045E8ba3C93d380D5BDDa"), locks[0].User)
	assert.Equal(t, big.NewInt(1000000000000000000), locks[0].TokenAmount)
	assert.Equal(t, adapters.TaikoTokenDecimals, locks[0].TokenDecimals)
	assert.Equal(t, common.HexToAddress(adapters.TaikoTokenAddress), locks[0].Token)
	assert.Equal(t, uint64(1730112839), locks[0].BlockTime)
	assert.Equal(t, uint64(259200), locks[0].Duration)
	assert.Equal(t, uint64(blocknumber), locks[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x604cc9466718ffd00787cbe7b0aefaa3b9ab5da030e59b60e6dec4fa31922ad9"), locks[0].TxHash)
}
