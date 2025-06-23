package uniswap

import (
	"bytes"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/futuramacoder/uniswap-api/app/pkg/client/eth"
)

var (
	feeMultiplier = big.NewInt(997)
	feeDivisor    = big.NewInt(1000)
)

type Service struct {
	client eth.Eth
}

func NewService(client eth.Eth) *Service {
	return &Service{client: client}
}

// GetOutputAmount calculates output amount for uniswap v2 pool

func (s *Service) GetOutputAmount(ctx context.Context, token0Addr, token1Addr, poolAddr string, amount *big.Int) (*big.Int, error) {
	token0 := common.HexToAddress(token0Addr)
	token1 := common.HexToAddress(token1Addr)
	poolID := common.HexToAddress(poolAddr)

	reserves, err := s.client.GetReserves(ctx, poolID)
	if err != nil {
		return nil, err
	}

	sorted0, _ := s.sortTokens(token0, token1)

	res0, res1 := reserves.Reserve0, reserves.Reserve1
	if sorted0 != token0 {
		res0, res1 = res1, res0
	}

	return s.outputAmount(amount, res0, res1), nil
}

func (s *Service) outputAmount(amountIn, reserve0, reserve1 *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, feeMultiplier)
	numerator := new(big.Int).Mul(amountInWithFee, reserve1)

	denominator := new(big.Int).Add(
		new(big.Int).Mul(reserve0, feeDivisor),
		amountInWithFee,
	)

	return new(big.Int).Div(numerator, denominator)
}

func (s *Service) sortTokens(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	if bytes.Compare(tkn0.Bytes(), tkn1.Bytes()) > 0 {
		return tkn1, tkn0
	}
	return tkn0, tkn1
}
