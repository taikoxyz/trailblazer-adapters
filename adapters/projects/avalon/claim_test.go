package avalon_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/avalon"
)

func TestClaimIndexer(t *testing.T) {
	ethereumRPC := "https://ethereum-rpc.publicnode.com"
	blocknumber := int64(21743609)
	testClaimAddress := "0x9c9a26f011a89f920f86fc48e2ed3f0fae71683b"

	ctx := context.Background()

	client, err := ethclient.Dial(ethereumRPC)
	require.NoError(t, err)

	indexer := avalon.NewClaimIndexer(client, []common.Address{common.HexToAddress(testClaimAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	// https://www.oklink.com/de/eth/tx/0xb3648f17578f1d791696677ababa4d3fda6b46a2a82fad7b6ed2d15d7b817e4d
	ps, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	assert.Len(t, ps, 1)
	assert.Equal(t, common.HexToAddress("0x2557ac54165134d7efd5ab94b750e9e04147beb1"), ps[0].User)
	assert.Equal(t, big.NewInt(4000000000000000000), ps[0].TokenAmount)
	assert.Equal(t, avalon.AvlTokenDecimal, ps[0].TokenDecimals)
	assert.Equal(t, common.HexToAddress(avalon.AvlTokenAddress), ps[0].Token)
	assert.Equal(t, uint64(1738315631), ps[0].BlockTime)
	assert.Equal(t, uint64(blocknumber), ps[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0xb3648f17578f1d791696677ababa4d3fda6b46a2a82fad7b6ed2d15d7b817e4d"), ps[0].TxHash)
}
