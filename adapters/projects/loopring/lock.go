package loopring

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/loopring_lock"
)

const (
	// https://taikoscan.io/address/0xaD32A362645Ac9139CFb5Ba3A2A46fC4c378812B
	LockAddress string = "0xaD32A362645Ac9139CFb5Ba3A2A46fC4c378812B"

	logDepositWithDurationSignature string = "Deposited(address,address,address,uint96,uint256)"
)

type LockIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewLockIndexer(client *ethclient.Client, addresses []common.Address) *LockIndexer {
	return &LockIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Lock] = &LockIndexer{}

func (indexer *LockIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *LockIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Lock, error) {
	var locks []adapters.Lock

	for _, l := range logs {
		if !indexer.isDepositWithDuration(l) {
			continue
		}

		var depositWithDurationEvent struct {
			From     common.Address
			To       common.Address
			Token    common.Address
			Amount   *big.Int
			Duration *big.Int
		}

		loopringABI, err := abi.JSON(strings.NewReader(loopring_lock.LoopringLockABI))
		if err != nil {
			return nil, err
		}

		err = loopringABI.UnpackIntoInterface(&depositWithDurationEvent, "Deposited", l.Data)
		if err != nil {
			return nil, err
		}

		if depositWithDurationEvent.Token != common.HexToAddress(adapters.TaikoTokenAddress) {
			continue
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		lock := &adapters.Lock{
			Metadata: adapters.Metadata{
				BlockTime:   block.Time(),
				BlockNumber: block.NumberU64(),
				TxHash:      l.TxHash,
			},
			User:          depositWithDurationEvent.To,
			TokenAmount:   depositWithDurationEvent.Amount,
			TokenDecimals: adapters.TaikoTokenDecimals,
			Token:         common.HexToAddress(adapters.TaikoTokenAddress),
			Duration:      depositWithDurationEvent.Duration.Uint64(),
		}

		locks = append(locks, *lock)
	}

	return locks, nil
}

func (indexer *LockIndexer) isDepositWithDuration(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logDepositWithDurationSignature)).Hex()
}
