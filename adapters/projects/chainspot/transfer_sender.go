package chainspot

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/chainspot"
)

var logChainspotSigHash = crypto.Keccak256Hash([]byte("Chainspot(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"))

type ChainspotIndexer struct {
	TargetAddresses []common.Address
}

type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

type ChainspotEvent struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
}

type ChainspotLoyaltyNFTEvent struct {
	NFTAddress   common.Address
	User         common.Address
	Level        uint8
}

func NewChainspotIndexer() *ChainspotIndexer {
	return &ChainspotIndexer{
		TargetAddresses: []common.Address{
			common.HexToAddress("0x3f96aF2AF6f644D5Fd1FC2d5A016CcE991198103"),
		},
	}
}

func (indexer *ChainspotIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *ChainspotIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
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

func (indexer *ChainspotIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logChainspotSigHash.Hex()
}

func (indexer *ChainspotIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
    var ChainspotEvent ChainspotEvent

    ChainspotABI, err := abi.JSON(strings.NewReader(chainspot.ABI))
    if err != nil {
        return nil, err
    }

    err = ChainspotABI.UnpackIntoInterface(&ChainspotEvent, "Chainspot", vLog.Data)
    if err != nil {
        return nil, err
    }

    block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
    if err != nil {
        return nil, err
    }

    return &adapters.Whitelist{
        User:        ChainspotEvent.Recipient,
        Time:        block.Time(),
        BlockNumber: block.Number().Uint64(),
    }, nil
}


func (indexer *ChainspotIndexer) ProcessNFTEvent(ctx context.Context, vLog types.Log) (*ChainspotLoyaltyNFTEvent, error) {
	var event ChainspotLoyaltyNFTEvent
	return &event, nil
}

