package nitro

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/nitro"
)

var (
	fundsPaidHash = crypto.Keccak256Hash([]byte("FundsPaid(bytes32,address,uint256)")) // 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464
)


type IRelayIndexer struct {
	TargetAddresses []common.Address
}

type IRelayEvent struct {
	MessageHash [32]byte `json:"messageHash"`
	Forwarder   common.Address  `json:"forwarder"`
	Nonce       *big.Int   `json:"nonce"`
}

func NewIRelayIndexer() *IRelayIndexer {
	return &IRelayIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0x7bd616192fb2b364f9d29b2026165281a5f2ff2f"), // taiko asset fwd contract
	}}
}

func (indexer *IRelayIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *IRelayIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if !indexer.isRelevantLog(vLog.Topics[0]) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		result = append(result, *transferData)
	}
	return result, nil
}

func (indexer *IRelayIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == fundsPaidHash.Hex()
}

func (indexer *IRelayIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {	
	var iRelayEvent IRelayEvent

	nitroForwarderABI, err := abi.JSON(strings.NewReader(nitro.ABI))
	if err != nil {
		return nil, err
	}

	err = nitroForwarderABI.UnpackIntoInterface(&iRelayEvent, "FundsPaid", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        iRelayEvent.Forwarder,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
