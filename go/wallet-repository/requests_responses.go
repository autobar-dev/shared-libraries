package walletrepository

// Create wallet
type ServiceCreateWalletRequestBody struct {
	UserId       string `json:"user_id"`
	CurrencyCode string `json:"currency_code"`
}

type ServiceCreateWalletResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Wallet `json:"data"`
}
