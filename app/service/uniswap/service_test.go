package uniswap

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/futuramacoder/uniswap-api/app/pkg/client/eth"
)

// mockClient реализует только нужный нам метод из eth.Client
type mockClient struct{}

func (m *mockClient) GetReserves(ctx context.Context, pool common.Address) (*eth.PairReserves, error) {
	reserve0, _ := big.NewInt(0).SetString("6897994292349957088357", 0)
	reserve1, _ := big.NewInt(0).SetString("17580630745241", 0)
	return &eth.PairReserves{
		Reserve0: reserve0,
		Reserve1: reserve1,
	}, nil
}

func TestGetOutputAmount(t *testing.T) {
	client := &mockClient{}
	service := NewService(client)

	token0 := "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
	token1 := "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	pool := "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"
	inputAmount, _ := big.NewInt(0).SetString("10000000000000000000", 0)

	out, err := service.GetOutputAmount(context.Background(), token0, token1, pool, inputAmount)

	assert.NoError(t, err)
	assert.NotNil(t, out)

	expected, _ := big.NewInt(0).SetString("25373450283", 0)

	assert.Equal(t, expected.String(), out.String())
}
