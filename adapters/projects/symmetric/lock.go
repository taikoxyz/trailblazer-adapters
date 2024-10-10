package symmetric

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/balancer_token"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/balancer_vault"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/symmetric"
)

const (
	// https://taikoscan.io/address/0xc0A740cDd1C647d9c77367E47f0D0c253140E6e3
	LockAddress string = "0xc0A740cDd1C647d9c77367E47f0D0c253140E6e3"

	logDepositSignature string = "Deposit(address,uint256,uint256,int128,uint256)"
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
		if !indexer.isDeposit(l) {
			continue
		}

		var depositEvent struct {
			Value *big.Int
			Type  *big.Int
			Ts    *big.Int
		}

		user := common.BytesToAddress(l.Topics[1].Bytes()[12:])

		symmetricABI, err := abi.JSON(strings.NewReader(symmetric.SymmetricABI))
		if err != nil {
			return nil, err
		}

		err = symmetricABI.UnpackIntoInterface(&depositEvent, "Deposit", l.Data)
		if err != nil {
			return nil, err
		}

		symmetric, err := symmetric.NewSymmetric(l.Address, indexer.client)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		token, err := symmetric.Token(&bind.CallOpts{
			BlockNumber: block.Number(),
		})
		if err != nil {
			return nil, err
		}

		balancerToken, err := balancer_token.NewBalancerToken(token, indexer.client)
		if err != nil {
			return nil, err
		}

		totalSupply, err := balancerToken.TotalSupply(&bind.CallOpts{
			BlockNumber: block.Number(),
		})
		if err != nil {
			return nil, err
		}

		poolId, err := balancerToken.GetPoolId(&bind.CallOpts{
			BlockNumber: block.Number(),
		})
		if err != nil {
			return nil, err
		}

		vault, err := balancerToken.GetVault(&bind.CallOpts{
			BlockNumber: block.Number(),
		})
		if err != nil {
			return nil, err
		}

		balancerVault, err := balancer_vault.NewBalancerVault(vault, indexer.client)
		if err != nil {
			return nil, err
		}

		poolTokens, err := balancerVault.GetPoolTokens(&bind.CallOpts{
			BlockNumber: block.Number(),
		}, poolId)
		if err != nil {
			return nil, err
		}

		for i, token := range poolTokens.Tokens {
			tokenAmount := poolTokens.Balances[i]
			userAmount := new(big.Int).Mul(tokenAmount, depositEvent.Value)
			userAmount.Div(userAmount, totalSupply)
			lock := &adapters.Lock{
				User:          user,
				TokenAmount:   userAmount,
				TokenDecimals: adapters.TaikoTokenDecimals,
				Token:         token,
				Time:          block.Time(),
				BlockNumber:   block.NumberU64(),
				TxHash:        l.TxHash,
			}
			locks = append(locks, *lock)
		}
	}

	return locks, nil
}

func (indexer *LockIndexer) isDeposit(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logDepositSignature)).Hex()
}
