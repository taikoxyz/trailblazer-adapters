package gaming

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

func Recipients() map[string][]string {
	return map[string][]string{
		"bullishs": {
			"0xAd2803cDBfFc8bae61eeC0A1CF849877f0E742Bd",
		},
		"0xAstra": {
			"0x90CE48ED68C6FCAe6F13b445F1573f003cF1804d",
			"0x34723B92aE9708BA33843120A86035D049dA7dfA",
			"0x445ab115F67b1EA7Fa4B5956DeB03796D6CB8e13",
			"0xF3f9c67CDd87De0597d607B5fC206299eB64fffa",
			"0xd0989635D363Bd85ecCE495fe78dF43A9067867D",
		},
		"21Blackjack": {
			"0x8C5720982b54874F53F2514BbD2382ADce98a0ca",
			"0x78adDA11Bfc437DeC4a39318FF7e52Cf00DC062c",
		},
		"Crack&Stack": {
			"0x7ddB8A975778a434dE03dd666F11Ce962DCdD290",
			"0x2c301eBfB0bb42Af519377578099b63E921515B7",
			"0x6C8865042792B5E919fC95bF771ccaDF6F0cfA22",
			"0xD31A4be996b7E1cc20974181127E6fCA15649913",
			"0xA9EC1fEEE212851c829B028F094156CD04A3a547",
			"0x1aca21a2a2a070d3536a69733c7044fedeb88f5a",
			"0xb64C1461453DAdD104A583dCCeef30ce296fde20",
			"0xD8F7cd7d919c5266777FB83542F956dD30E80187",
			"0x009C32F03d6eEa4F6DA9DD3f8EC7Dc85824Ae0e6",
			"0xF8F1B21615BDbEA8D142cfaf4828EA0236Cab115",
			"0x12689b6ddE632E69fBAA70d066f86aC9fDd33dd1",
		},
		"LooperLands": {
			"0xEe01C4b0538849bF1c66bDFB458a7de11B1d7424",
			"0x563Be8E43692bDC27525afF630fDaE22D5A83530",
			"0x6f2D2EFC2773923f0d7De4eac1F845C34A094690",
			"0xE8e9B020B72352d55410a147aB49f53ff479A40c",
			"0x56b0D8d04de22f2539945258ddB288C123026775",
		},
		"StupidMonkeys": {
			"0xCA99F9DbF4A13D4de05B41a68041dcE7929cb5e0",
		},
		"SwordsAndDungeons": {
			"0xac67aB5D44d5A7DbB55b8D9d1893889B7DaE3a36",
			"0xc8b15Aa9BD24961E2ee4f3C67aEA4EafcB071108",
			"0xEF9bF561ccbCBeDA1A71341A2162e5e6fF28c929",
		},
		"Ultiverse": {
			"0xa30c146A816931f6134dD8676B9dE8E77496954d",
			"0xE3b5500039F401e48627e8025b37d4871cF34f36",
			"0x06F9914838903162515aFa67D5b99Ada0F9791cc",
		},
		"ZypherNetwork": {
			"0xd4629d312CdC663D062F3Fbc322534A9Df0151bC",
		},
		"EVMWarfare": {
			"0xA9E8e13BA0309486b2F710B8201B534D26633a89",
		},
		"Brigade": {
			"0x8a93AAE6D94680658012B887BfDd981A17661Ef4",
			"0x409395BC4b50A9BbD45a943A8B0D6236E0F83540",
			"0x20F50518188FB3c9F5adff472E056291C4B98ecE",
			"0x0158A4055428b67e286b2627e91120b49ca1146c",
			"0x73716C57f87fFd4135453aBCe6cf61Bb0E99C410",
			"0x03376f22eF7d08CEE420D07207f85E52638A9fCd",
			"0x72dCB9a28bB8EA172B58130d9fd17A6dBE7A9E41",
			"0xB66A56126fAd14563e62BA2Cda658cb97F7a90dE",
			"0x4f576c055f06EFFdF165C5bA014f0b827D47E27B",
		},
		"Robots.Farm": {
			"0xED2cC316E81d574330E60aE9593a8dd34b9a4C41",
			"0x972fd05EdA4031Ea59Cb213180eBd23524F6DBA6",
		},
		"WorldOfDypians": {
			"0xCb2Eb4ba62346751F36bA652010b553759141AEE",
			"0xaf33f679be47733bd3abb5b0b977b6ba3ed8d01e",
			"0x620655Ee8320bA51cf4cc06bf6a7C14022271764",
			"0xc24728996C974788027871C4b1902d9ac70dBCDd",
			"0xf394504D9B224923C2b569B719EE99640fd585dA",
			"0x02E857a5F93F531aAC2bAD56113B81155DbcF236",
			"0x810C42A71358dc1e0Ecc32815DadD90c586AfD1c",
			"0xdb2E1287AAC9974AB28a66fABF9bCB34C5f37712",
			"0xE026FB242D9523dc8E8d8833F7309dbdbED59d3d",
			"0x94d266D7f7F0548282439c0c62962B048ad86d1a",
			"0x2b7DcAeAe477B3Bb4b8a3a043611fE06804F6f9F",
			"0x67C620BbA764b2a45DF5369E7A03d6bD7A210D47",
			"0xd1aFA3A14572cBb11d9aAb0E6D9b3F7780689fd0",
			"0x842856460c527d37D4B5fbBDc3a0a2c20DAF3281",
			"0x5baDC9630a34aC5Af3BBa448A23C681dAEDA679B",
			"0x565571650491C471F31eEfF60e382D3022ccB717",
			"0xeE80b4711c316D66bE0A0865838a8A45e71618E7",
			"0xb0944b05c6cac25Ab66f28dC78CFd491fc404534",
			"0xBB62213BD1D8094e91aa9b4b9eEdBC74253f2fF6",
			"0xeFdb62eb5234a68975aCa0ef9aD1f83bbCa1B158",
			"0x5F7EB07f189889Cbc8D1FEa0E8d41b313E463dB1",
			"0xC69478b8379Df1DB75FDb6415F72F7037dAF60c6",
			"0xa8097beda67f7e90b5da088e32816dc73c23588a",
		},
		"XenoBunny": {
			"0x1996E10c64213Fe5E86AC7A7ac03Ec169176E4a7",
			"0xF645df6186AD0DbB088f7ef024C4Be640F1DaCd2",
			"0x553B1A22fB7c148690BC14014a9FDFc12e8Fdfa4",
		},
		"OpenALCHI": {
			"0x9a92338b1CFfdf34b4Cc0192A8882F523b853e44",
			"0xed2147ff41e7B14a9C0c37fc4b7a1Eb7644EFb3C",
			"0x5cF295a3c53E090Cf230E1D40f413b8e3c225C26",
		},
	}
}

type GamingProcessor struct {
	client          *ethclient.Client
	ValidRecipients map[string]struct{}
}

func NewGamingProcessor(client *ethclient.Client) *GamingProcessor {
	processor := &GamingProcessor{
		client:          client,
		ValidRecipients: map[string]struct{}{},
	}
	for _, addrList := range Recipients() {
		for _, addr := range addrList {
			processor.ValidRecipients[addr] = struct{}{}
		}
	}
	return processor
}

var _ adapters.BlockProcessor[adapters.Whitelist] = &GamingProcessor{}

func (indexer *GamingProcessor) Process(ctx context.Context, blocks ...*types.Block) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, b := range blocks {
		txs := b.Transactions()

		for _, tx := range txs {
			to := tx.To()
			if to == nil {
				continue
			}

			receipt, err := indexer.client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return nil, err
			}
			if receipt.Status == 0 {
				continue
			}

			sender, err := indexer.client.TransactionSender(ctx, tx, b.Hash(), receipt.TransactionIndex)
			if err != nil {
				return nil, err
			}

			if _, exists := indexer.ValidRecipients[to.Hex()]; exists {
				whitelist = append(whitelist, adapters.Whitelist{
					User:        sender,
					BlockNumber: b.NumberU64(),
					Time:        b.Time(),
					TxHash:      tx.Hash(),
				})
			}
		}
	}

	return whitelist, nil
}
