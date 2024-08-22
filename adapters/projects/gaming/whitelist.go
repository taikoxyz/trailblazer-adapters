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
			common.HexToAddress("0xF8F1B21615BDbEA8D142cfaf4828EA0236Cab115"),
			common.HexToAddress("0x12689b6ddE632E69fBAA70d066f86aC9fDd33dd1"),
		},
		"LooperLands": {
            common.HexToAddress("0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424"),
            common.HexToAddress("0x563Be8E43692bDC27525afF630fDaE22D5A83530"),
            common.HexToAddress("0x6f2D2EFC2773923f0d7De4eac1F845C34A094690"),
            common.HexToAddress("0xCA99F9DbF4A13D4de05B41a68041dcE7929cb5e0"),
            common.HexToAddress("0xE8e9B020B72352d55410a147aB49f53ff479A40c"),
            common.HexToAddress("0x56b0D8d04de22f2539945258ddB288C123026775"),
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
			common.HexToAddress("0xc24728996C974788027871C4b1902d9ac70dBCDd"),
      		common.HexToAddress("0xf394504D9B224923C2b569B719EE99640fd585dA"),
      		common.HexToAddress("0x02E857a5F93F531aAC2bAD56113B81155DbcF236"),
      		common.HexToAddress("0x810C42A71358dc1e0Ecc32815DadD90c586AfD1c"),
      		common.HexToAddress("0xdb2E1287AAC9974AB28a66fABF9bCB34C5f37712"),
      		common.HexToAddress("0xE026FB242D9523dc8E8d8833F7309dbdbED59d3d"),
      		common.HexToAddress("0x94d266D7f7F0548282439c0c62962B048ad86d1a"),
     		common.HexToAddress("0x2b7DcAeAe477B3Bb4b8a3a043611fE06804F6f9F"),
     		common.HexToAddress("0x67C620BbA764b2a45DF5369E7A03d6bD7A210D47"),
     		common.HexToAddress("0xd1aFA3A14572cBb11d9aAb0E6D9b3F7780689fd0"),
      		common.HexToAddress("0x842856460c527d37D4B5fbBDc3a0a2c20DAF3281"),
      		common.HexToAddress("0x5baDC9630a34aC5Af3BBa448A23C681dAEDA679B"),
      		common.HexToAddress("0x565571650491C471F31eEfF60e382D3022ccB717"),
      		common.HexToAddress("0xeE80b4711c316D66bE0A0865838a8A45e71618E7"),
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
