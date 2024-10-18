package transactionsender_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	transactionsender "github.com/taikoxyz/trailblazer-adapters/adapters/transaction_sender"
)

func TestProcessor(t *testing.T) {
	taikoRPC := "https://rpc.taiko.xyz"
	blocknumber := int64(447484)

	ctx := context.Background()

	client, err := ethclient.Dial(taikoRPC)
	require.NoError(t, err)

	processor := transactionsender.New(client)

	block, err := adapters.GetBlock(ctx, client, blocknumber)
	require.NoError(t, err)

	whitelist, err := processor.Process(ctx, block)
	assert.NoError(t, err)
	assert.Len(t, whitelist, 1)
	assert.Equal(t, common.HexToAddress("0x62e949C4150b79d63263430B6C5298550e63A17E"), whitelist[0].User)
	assert.Equal(t, uint64(1728471443), whitelist[0].Time)
	assert.Equal(t, uint64(blocknumber), whitelist[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x2cf888a4b1dad0da32d1c415aaef9ed536428d841008e79cffa14e3ed05e9d00"), whitelist[0].TxHash)
}
