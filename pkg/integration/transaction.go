package integration

type Transaction struct {
	// any generic type to be translate by network,
	// maybe it can be a generic based on provider
	Body map[string]any
}

type TransactionResponse struct {
	// any generic type to be translate by network,
	// maybe it can be a generic based on provider
	Body map[string]any
	Err  error
}
