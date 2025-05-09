package loopring_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopring"
)

func TestDepositRequestedIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(1116640)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := loopring.NewDepositRequestedIndexer(client, []common.Address{common.HexToAddress(loopring.DepositRequestedAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ws, err := indexer.Index(ctx, logs...)
	require.NoError(t, err)
	require.Len(t, ws, 1)
	assert.Equal(t, common.HexToAddress("0x4e79125DDebA405d82326e85C80603e15B3a3f80"), ws[0].User)
	assert.Equal(t, uint64(1745954111), ws[0].Time)
	assert.Equal(t, uint64(blocknumber), ws[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x87cfc6f982c745bf400bb5dd824ae24e2440d6796237b213782ed20eca5cd34d"), ws[0].TxHash)
}
