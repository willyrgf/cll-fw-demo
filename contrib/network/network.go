package network

import (
	solanaClient "github.com/smartcontractkit/chainlink-solana/pkg/solana/client"
	"github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/willyrgf/cll-fw-demo/contrib/network/evm"
	"github.com/willyrgf/cll-fw-demo/contrib/network/nonevm"
	"github.com/willyrgf/cll-fw-demo/pkg/integration"
)

type Network interface {
	integration.Reader
	integration.Writer
}

func New(networkName string) Network {
	var network Network

	switch networkName {
	case "solana":
		network = nonevm.NewProvider(&solanaClient.Client{})
	case "ethereum":
		network = evm.NewProvider(&client.NullClient{})
	default:
		panic("network doesnt supported")
	}

	return network
}
