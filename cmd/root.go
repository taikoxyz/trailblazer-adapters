package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

var adapter string
var blockNumber int64

var ProcessCmd = &cobra.Command{
	Use:   "process",
	Short: "Process specified blocks with a given adapter",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("https://rpc.taiko.xyz")
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}

		switch adapter {
		case "NewTransactionSender":
			validSenders := map[string]struct{}{
				common.HexToAddress("0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5").Hex(): {},
			}
			processor := blocks.NewTransactionSender(validSenders)
			blockNumberBig := big.NewInt(blockNumber)
			block, err := client.BlockByNumber(context.Background(), blockNumberBig)
			if err != nil {
				log.Fatalf("Failed to fetch the block: %v", err)
			}

			senders, err := processor.ProcessBlock(context.Background(), block, client)
			if err != nil {
				log.Fatalf("Failed to process the block: %v", err)
			}

			fmt.Printf("Senders: %v\n", senders)

		case "DotTaikoIndexer":
			processor := logs.NewDotTaikoIndexer()

			chainID, err := client.ChainID(context.Background())
			if err != nil {
				log.Fatalf("Failed to fetch the chain ID: %v", err)
			}
			query := ethereum.FilterQuery{
				Addresses: processor.Addresses,
				FromBlock: big.NewInt(blockNumber),
				ToBlock:   big.NewInt(blockNumber),
			}
			logs, err := client.FilterLogs(context.Background(), query)
			if err != nil {
				log.Fatalf("Failed to fetch the logs: %v", err)
			}
			senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
			if err != nil {
				log.Fatalf("Failed to process the logs: %v", err)
			}

			fmt.Printf("Senders: %v\n", senders)

		default:
			log.Fatalf("Adapter %s is not supported", adapter)
		}
	},
}

func init() {
	ProcessCmd.Flags().StringVarP(&adapter, "adapter", "a", "", "Adapter to use (required)")
	ProcessCmd.Flags().Int64VarP(&blockNumber, "block", "b", 0, "Block number to process (required)")
	ProcessCmd.MarkFlagRequired("adapter")
	ProcessCmd.MarkFlagRequired("block")
}
