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
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestClaimIndexer(t *testing.T) {
	blocknumber := int64(983938)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := avalon.NewClaimIndexer(client, []common.Address{common.HexToAddress(avalon.ClaimAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	ps, err := indexer.Index(ctx, logs...)
	require.NoError(t, err)
	require.Len(t, ps, 1)
	assert.Equal(t, common.HexToAddress("0x06CDfb6eA4B546887352a83768aFAD9481da1Bff"), ps[0].User)
	assert.Equal(t, big.NewInt(5000000000000000000), ps[0].TokenAmount)
	assert.Equal(t, avalon.AvlTokenDecimal, ps[0].TokenDecimals)
	assert.Equal(t, common.HexToAddress(avalon.AvlTokenAddress), ps[0].Token)
	assert.Equal(t, uint64(1742155139), ps[0].BlockTime)
	assert.Equal(t, uint64(blocknumber), ps[0].BlockNumber)
	assert.Equal(t, common.HexToHash("0x711132b470df725db72dd4d277c0a148333aaf7268fec3432427072063247d77"), ps[0].TxHash)
}
