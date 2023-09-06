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

// Get wallet
type ServiceGetWalletResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Wallet `json:"data"`
}

// Create transaction
type ServiceTransactionDepositRequestBody struct {
	UserId string `json:"user_id"`
	Value  int64  `json:"value"`
}

type ServiceTransactionWithdrawRequestBody struct {
	UserId string `json:"user_id"`
	Value  int64  `json:"value"`
}

type ServiceTransactionPurchaseRequestBody struct {
	UserId string `json:"user_id"`
	Value  int64  `json:"value"`
}

type ServiceTransactionRefundRequestBody struct {
	UserId string `json:"user_id"`
	Value  int64  `json:"value"`
}

type ServiceTransactionCurrencyChangeRequestBody struct {
	UserId string `json:"user_id"`
	Value  int64  `json:"value"`
}

type ServiceCreateTransactionResponse struct {
	Status string       `json:"status"`
	Error  *string      `json:"error"`
	Data   *Transaction `json:"data"`
}
