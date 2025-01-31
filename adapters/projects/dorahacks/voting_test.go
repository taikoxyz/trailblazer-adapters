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
)

func TestVotingIndexer(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(789673)
	testVotingAddress := "0x649Ea09c01FB8f5d54262948F39c7648cFd97641"

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	indexer := dorahacks.NewVotingIndexer(client, []common.Address{common.HexToAddress(testVotingAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ws, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, ws, 2)
	assert.Equal(t, common.HexToAddress("0xCc31c9cb03bFb995fF9ac62Da615Fb62b670aA1B"), ws[0].User)
	assert.Equal(t, uint64(789673), ws[0].BlockNumber)
	assert.Equal(t, uint64(1737736247), ws[0].Time)
	assert.Equal(t, common.HexToHash("0xfd4271026b789a73c1cd689e1d46949f582c782cff47b4c46703abbf229bc57b"), ws[0].TxHash)

	assert.Equal(t, common.HexToAddress("0xCc31c9cb03bFb995fF9ac62Da615Fb62b670aA1B"), ws[1].User)
	assert.Equal(t, uint64(789673), ws[1].BlockNumber)
	assert.Equal(t, uint64(1737736247), ws[1].Time)
	assert.Equal(t, common.HexToHash("0xfd4271026b789a73c1cd689e1d46949f582c782cff47b4c46703abbf229bc57b"), ws[1].TxHash)
}
