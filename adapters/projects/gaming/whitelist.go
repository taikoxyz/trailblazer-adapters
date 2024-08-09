package gaming

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

const (
	// Define your addresses in a constant
	astra         = "0xAstra"
	crackAndStack = "Crack&Stack"
	looperLands   = "LooperLands"
	ultiverse     = "Ultiverse"
	zypherNetwork = "ZypherNetwork"
	evmWarfare    = "EVMWarfare"
)

// GetAddresses returns the map of addresses
func GetAddresses() map[string][]common.Address {
	return map[string][]common.Address{
		astra: {
			common.HexToAddress("0x90CE48ED68C6FCAe6F13b445F1573f003cF1804d"),
			common.HexToAddress("0x34723B92aE9708BA33843120A86035D049dA7dfA"),
		},
		crackAndStack: {
			common.HexToAddress("0x7ddB8A975778a434dE03dd666F11Ce962DCdD290"),
			common.HexToAddress("0x2c301eBfB0bb42Af519377578099b63E921515B7"),
			common.HexToAddress("0x6C8865042792B5E919fC95bF771ccaDF6F0cfA22"),
			common.HexToAddress("0xD31A4be996b7E1cc20974181127E6fCA15649913"),
			common.HexToAddress("0xA9EC1fEEE212851c829B028F094156CD04A3a547"),
			common.HexToAddress("0x1aca21a2a2a070d3536a69733c7044fedeb88f5a"),
			common.HexToAddress("0xb64C1461453DAdD104A583dCCeef30ce296fde20"),
			common.HexToAddress("0xD8F7cd7d919c5266777FB83542F956dD30E80187"),
			common.HexToAddress("0x009C32F03d6eEa4F6DA9DD3f8EC7Dc85824Ae0e6"),
		},
		looperLands: {
			common.HexToAddress("0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424"),
		},
		ultiverse: {
			common.HexToAddress("0xa30c146A816931f6134dD8676B9dE8E77496954d"),
			common.HexToAddress("0xE3b5500039F401e48627e8025b37d4871cF34f36"),
			common.HexToAddress("0x06F9914838903162515aFa67D5b99Ada0F9791cc"),
		},
		zypherNetwork: {
			common.HexToAddress("0xd4629d312CdC663D062F3Fbc322534A9Df0151bC"),
		},
		evmWarfare: {
			common.HexToAddress("0xA9E8e13BA0309486b2F710B8201B534D26633a89"),
		},
	}
}

// GamingWhitelist stores valid recipients
type GamingWhitelist struct {
	ValidRecipient map[string]struct{}
}

// NewGamingWhitelist initializes a new GamingWhitelist
func NewGamingWhitelist() *GamingWhitelist {
	whitelist := &GamingWhitelist{
		ValidRecipient: make(map[string]struct{}),
	}
	addresses := GetAddresses()
	for _, addrList := range addresses {
		for _, addr := range addrList {
			whitelist.ValidRecipient[addr.Hex()] = struct{}{}
		}
	}
	return whitelist
}

// ProcessBlock processes a block and returns the whitelist entries
func (indexer *GamingWhitelist) ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) ([]adapters.Whitelist, error) {
	txs := block.Transactions()
	var result []adapters.Whitelist

	for _, tx := range txs {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}
		if receipt.Status == 0 {
			continue
		}
		sender, err := client.TransactionSender(ctx, tx, block.Hash(), receipt.TransactionIndex)
		if err != nil {
			return nil, err
		}
		to := tx.To()
		if to == nil {
			continue
		}
		if _, exists := indexer.ValidRecipient[to.Hex()]; exists {
			result = append(result, adapters.Whitelist{
				User:        sender,
				BlockNumber: block.Number().Uint64(),
				Time:        block.Time(),
			})
		}
	}
	return result, nil
}
