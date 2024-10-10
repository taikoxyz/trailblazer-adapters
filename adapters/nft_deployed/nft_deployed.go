package nftdeployed

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc165"
)

const (
	ERC721InterfaceID  string = "0x80ac58cd" // ERC721 interface ID
	ERC1155InterfaceID string = "0xd9b67a26" // ERC1155 interface ID
)

type Processor struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *Processor {
	return &Processor{
		client: client,
	}
}

var _ adapters.BlockProcessor[adapters.Whitelist] = &Processor{}

func (processor *Processor) Process(ctx context.Context, blocks ...*types.Block) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, b := range blocks {
		txs := b.Transactions()

		for _, tx := range txs {
			to := tx.To()
			if to == nil {
				continue
			}

			receipt, err := processor.client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return nil, err
			}

			sender, err := processor.client.TransactionSender(ctx, tx, b.Hash(), receipt.TransactionIndex)
			if err != nil {
				return nil, err
			}

			// TODO: check on zero address for contract?
			isErc721, err := processor.supportsInterface(receipt.ContractAddress, ERC721InterfaceID)
			if err != nil {
				log.Printf(
					"tx %s supports interface ERC721 for contract %s failed with %v but continue",
					tx.Hash(),
					receipt.ContractAddress.Hex(),
					err,
				)
				// TODO: should not continue?
			}

			if isErc721 {
				whitelist = append(whitelist, adapters.Whitelist{
					User:        sender,
					BlockNumber: b.NumberU64(),
					Time:        b.Time(),
					TxHash:      tx.Hash(),
				})
			}

			isErc1155, err := processor.supportsInterface(receipt.ContractAddress, ERC1155InterfaceID)
			if err != nil {
				log.Printf(
					"tx %s supports interface ERC1155 for contract %s failed with %v but continue",
					tx.Hash(),
					receipt.ContractAddress.Hex(),
					err,
				)
				// TODO: should not continue?
			}

			if isErc1155 {
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

func (processor *Processor) supportsInterface(contractAddress common.Address, interfaceID string) (bool, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc165.ABI))
	if err != nil {
		return false, err
	}

	interfaceIDBytes := common.FromHex(interfaceID)
	if len(interfaceIDBytes) != 4 {
		return false, fmt.Errorf("invalid interface ID length")
	}

	data, err := parsedABI.Pack("supportsInterface", [4]byte{interfaceIDBytes[0], interfaceIDBytes[1], interfaceIDBytes[2], interfaceIDBytes[3]})
	if err != nil {
		return false, err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := processor.client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return false, err
	}

	var supports bool
	err = parsedABI.UnpackIntoInterface(&supports, "supportsInterface", result)
	if err != nil {
		return false, err
	}

	return supports, nil
}
