package polaris_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/polaris"
	"github.com/taikoxyz/trailblazer-adapters/testutils"
)

func TestLPTransferIndexer(t *testing.T) {
	blocknumber := int64(408155)

	ctx := context.Background()

	client, err := ethclient.Dial(testutils.TaikoRPC)
	require.NoError(t, err)

	indexer := polaris.NewLPTransferIndexer(client, []common.Address{common.HexToAddress(polaris.VaultAddress)})

	logs, err := adapters.GetLogs(ctx, client, indexer.Addresses(), blocknumber)
	require.NoError(t, err)

	transfers, err := indexer.Index(ctx, logs...)
	assert.NoError(t, err)
	txHash := common.HexToHash("0x53cadfc76c846db7bc14bcdca5b97fb52ba3cae7e4f641e339d9964a27da8cc1")
	user := common.HexToAddress("0xb535d4D9126c58d8dA8fE5775088aF77DE37F5D7")
	time := 1727179415
	expectedTransfers := []adapters.LPTransfer{
		{
			Metadata: adapters.Metadata{
				BlockTime:   uint64(time),
				BlockNumber: uint64(blocknumber),
				TxHash:      txHash,
			},
			From:           user,
			To:             common.HexToAddress(polaris.VaultAddress),
			Token0Amount:   big.NewInt(223714285714285710),
			Token0Decimals: uint8(18),
			Token0:         common.HexToAddress("0x7824693029834D294A92c54138e81e94d8D0Fd06"),
			Token1Amount:   big.NewInt(300000000000000000),
			Token1Decimals: uint8(18),
			Token1:         common.HexToAddress("0xA9d23408b9bA935c230493c40C73824Df71A0975"),
		},
	}

	for i, transfer := range transfers {
		assert.Equal(t, expectedTransfers[i].From, transfer.From)
		assert.Equal(t, expectedTransfers[i].To, transfer.To)
		assert.Equal(t, expectedTransfers[i].Token0Amount, transfer.Token0Amount)
		assert.Equal(t, expectedTransfers[i].Token0Decimals, transfer.Token0Decimals)
		assert.Equal(t, expectedTransfers[i].Token0, transfer.Token0)
		assert.Equal(t, expectedTransfers[i].Token1Amount, transfer.Token1Amount)
		assert.Equal(t, expectedTransfers[i].Token1Decimals, transfer.Token1Decimals)
		assert.Equal(t, expectedTransfers[i].Token1, transfer.Token1)
		assert.Equal(t, expectedTransfers[i].BlockTime, transfer.BlockTime)
		assert.Equal(t, expectedTransfers[i].BlockNumber, transfer.BlockNumber)
		assert.Equal(t, expectedTransfers[i].TxHash, transfer.TxHash)
	}
}
