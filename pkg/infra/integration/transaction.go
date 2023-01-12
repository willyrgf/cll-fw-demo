package integration

type Transaction struct {
	Body map[string]any // any generic type to be translate by network
}

type TransactionResponse struct {
	Body map[string]any // any generic type to be translate by network
	Err  error
}
