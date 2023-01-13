package main

import (
	"context"

	"github.com/willyrgf/cll-fw-demo/contrib/network"
	"github.com/willyrgf/cll-fw-demo/pkg/infra/logger"
	"github.com/willyrgf/cll-fw-demo/pkg/integration"
	"go.opentelemetry.io/otel/trace"
)

// run the same operation in n networks
func main() {
	logger := logger.New(logger.WithTrace(trace.NewNoopTracerProvider().Tracer("general")))
	ctx := context.TODO()

	networkToRun := []string{"solana", "ethereum"}

	for _, networkName := range networkToRun {
		logger.Infof("main(): runSomethingPerNetwork(ctx, networkName: %+v)", networkName)
		runSomethingPerNetwork(ctx, networkName)
	}
}

func runSomethingPerNetwork(ctx context.Context, networkName string) {
	// setup the network provider
	n := network.New(networkName)

	// get a balance from the provider network based on address
	b, err := n.Balance(ctx, integration.Address{})
	if err != nil {
		panic("error handling")
	}

	// create some transaction to be executeted on provider
	tx := integration.Transaction{Body: map[string]any{"balance": b}}

	// send this transaction to provider
	txResponse := n.SendTx(ctx, tx)

	// do something with the response
	_ = txResponse
}
