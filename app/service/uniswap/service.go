package uniswap

import (
	"bytes"
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/futuramacoder/uniswap-api/app/pkg/client/eth"
)

var (
	feeMultiplier = big.NewInt(997)
	feeDivisor    = big.NewInt(1000)
)

var bigIntPool = sync.Pool{
	New: func() interface{} { return new(big.Int) },
}

func getBigInt() *big.Int  { return bigIntPool.Get().(*big.Int) }
func putBigInt(b *big.Int) { bigIntPool.Put(b) }

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

	return OutputAmount(amount, res0, res1), nil
}

func OutputAmount(amountIn, reserve0, reserve1 *big.Int) *big.Int {
	tmpA := getBigInt()
	tmpB := getBigInt()
	tmpC := getBigInt()
	tmpD := getBigInt()
	defer func() {
		putBigInt(tmpA)
		putBigInt(tmpB)
		putBigInt(tmpC)
		putBigInt(tmpD)
	}()

	tmpA.Mul(amountIn, feeMultiplier)
	tmpB.Mul(tmpA, reserve1)
	tmpC.Mul(reserve0, feeDivisor)
	tmpD.Add(tmpC, tmpA)

	return new(big.Int).Div(tmpB, tmpD)
}

func (s *Service) sortTokens(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	if bytes.Compare(tkn0.Bytes(), tkn1.Bytes()) > 0 {
		return tkn1, tkn0
	}
	return tkn0, tkn1
}
