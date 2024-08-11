package gaming

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

var (
	Addresses = map[string][]common.Address{
		"bullishs": {
			common.HexToAddress("0xAd2803cDBfFc8bae61eeC0A1CF849877f0E742Bd"),
		},
		"0xAstra": {
			common.HexToAddress("0x90CE48ED68C6FCAe6F13b445F1573f003cF1804d"),
			common.HexToAddress("0x34723B92aE9708BA33843120A86035D049dA7dfA"),
		},
		"21Blackjack": {
			common.HexToAddress("0x8C5720982b54874F53F2514BbD2382ADce98a0ca"),
			common.HexToAddress("0x78adDA11Bfc437DeC4a39318FF7e52Cf00DC062c"),
		},
		"Crack&Stack": {
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
		"LooperLands": {
			common.HexToAddress("0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424"),
		},
		"StupidMonkeys": {
			common.HexToAddress("0xCA99F9DbF4A13D4de05B41a68041dcE7929cb5e0"),
		},
		"SwordsAndDungeons": {
			common.HexToAddress("0xac67aB5D44d5A7DbB55b8D9d1893889B7DaE3a36"),
			common.HexToAddress("0xc8b15Aa9BD24961E2ee4f3C67aEA4EafcB071108"),
			common.HexToAddress("0xEF9bF561ccbCBeDA1A71341A2162e5e6fF28c929"),
		},
		"Ultiverse": {
			common.HexToAddress("0xa30c146A816931f6134dD8676B9dE8E77496954d"),
			common.HexToAddress("0xE3b5500039F401e48627e8025b37d4871cF34f36"),
			common.HexToAddress("0x06F9914838903162515aFa67D5b99Ada0F9791cc"),
		},
		"ZypherNetwork": {
			common.HexToAddress("0xd4629d312CdC663D062F3Fbc322534A9Df0151bC"),
		},
		"EVMWarfare": {
			common.HexToAddress("0xA9E8e13BA0309486b2F710B8201B534D26633a89"),
		},
		"Brigade": {
			common.HexToAddress("0x8a93AAE6D94680658012B887BfDd981A17661Ef4"),
		},
		"Robots.Farm": {
			common.HexToAddress("0xED2cC316E81d574330E60aE9593a8dd34b9a4C41"),
		},
		"WorldOfDypians": {
			common.HexToAddress("0xCb2Eb4ba62346751F36bA652010b553759141AEE"),
			common.HexToAddress("0xaf33f679be47733bd3abb5b0b977b6ba3ed8d01e"),
			common.HexToAddress("0x620655Ee8320bA51cf4cc06bf6a7C14022271764"),
		},
	}
)

type GamingWhitelist struct {
	ValidRecipient map[string]struct{}
}

func NewGamingWhitelist() *GamingWhitelist {
	whitelist := &GamingWhitelist{
		ValidRecipient: map[string]struct{}{},
	}
	for _, addrList := range Addresses {
		for _, addr := range addrList {
			whitelist.ValidRecipient[addr.Hex()] = struct{}{}
		}
	}
	return whitelist
}

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
