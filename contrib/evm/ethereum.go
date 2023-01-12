package evm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/willyrgf/cll-fw-demo/pkg/integration"
)

type Provider struct {
	client client.Client
}

func NewProvider(c *client.Client) Provider {
	return Provider{
		client: &client.NullClient{},
	}
}

func (p *Provider) Balance(ctx context.Context, address integration.Address) (*big.Int, error) {
	account, blockNumber := translateAddress(address)

	assetEth, err := p.client.GetEthBalance(ctx, account, blockNumber)

	return translateBalance(assetEth), err
}

func (p *Provider) ChainID(ctx context.Context) (uint64, error) {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) SendTx(ctx context.Context, tx integration.Transaction) integration.TransactionResponse {
	ethereumTx := translateTx(tx)

	err := p.client.SendTransaction(ctx, &ethereumTx)

	return integration.TransactionResponse{Err: err}
}

func translateBalance(assetEth *assets.Eth) *big.Int {
	// translate assetEth -> balance in big.Int
	return &big.Int{}
}

func translateAddress(address integration.Address) (common.Address, *big.Int) {
	// translate address -> account and blockNumber
	return common.Address{}, &big.Int{}
}

func translateTx(tx integration.Transaction) types.Transaction {
	// translate tx -> types.Transaction
	return types.Transaction{}
}
