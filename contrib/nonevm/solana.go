package nonevm

import (
	"context"
	"math/big"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/client"
	"github.com/willyrgf/cll-fw-demo/pkg/integration"
)

type Provider struct {
	client *client.Client
}

func NewProvider(c *client.Client) Provider {
	return Provider{
		client: c,
	}
}

func (p *Provider) Balance(ctx context.Context, address integration.Address) (*big.Int, error) {
	solanaAddress := translateAddress(address)
	b, err := p.client.Balance(solanaAddress)
	return big.NewInt(int64(b)), err
}

func (p *Provider) ChainID(ctx context.Context) (uint64, error) {
	panic("not implemented") // TODO: Implement
}

func (p *Provider) SendTx(ctx context.Context, tx integration.Transaction) integration.TransactionResponse {
	solanaTx := translateTx(tx)

	solanaSignature, _ := p.client.SendTx(ctx, &solanaTx)

	transactionResponse := translateSolanaSignature(solanaSignature)

	return transactionResponse
}

func translateAddress(address integration.Address) solana.PublicKey {
	// translate address -> solana.PublicKey
	return solana.PublicKey{}
}

func translateTx(tx integration.Transaction) solana.Transaction {
	// translate tx -> solana.Transaction
	return solana.Transaction{}
}

func translateSolanaSignature(s solana.Signature) integration.TransactionResponse {
	// translate s -> TransactionResponse
	return integration.TransactionResponse{}
}
