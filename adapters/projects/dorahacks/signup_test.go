package dorahacks_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/dorahacks"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestSignupIndexer(t *testing.T) {
	blocknumber := int64(861132)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := dorahacks.NewSignupIndexer(client, []common.Address{common.HexToAddress(dorahacks.VotingAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ws, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, ws, 1)
	assert.Equal(t, common.HexToAddress("0xA9d1f06b6AF87B312aBbb523A5e41995F199930b"), ws[0].User)
	assert.Equal(t, uint64(861132), ws[0].BlockNumber)
	assert.Equal(t, uint64(1739528687), ws[0].BlockTime)
	assert.Equal(t, common.HexToHash("0x04f0a71af85058eeedd078cb2a21e18c82dbf93d814dcf0e051ad6c891116b60"), ws[0].TxHash)
}
