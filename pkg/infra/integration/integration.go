package integration

import (
	"context"
	"math/big"
)

type Reader interface {
	Balance(ctx context.Context, address Address) (*big.Int, error)
	ChainID(ctx context.Context) (uint64, error)
}

type Writer interface {
	SendTx(ctx context.Context, tx Transaction) TransactionResponse
}
