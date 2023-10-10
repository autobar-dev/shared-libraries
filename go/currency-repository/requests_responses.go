package currencyrepository

// Get currency
type ServiceGetCurrencyResponse struct {
	Status string    `json:"status"`
	Error  *string   `json:"error"`
	Data   *Currency `json:"data"`
}

// Get rate
type ServiceGetRateResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Rate   `json:"data"`
}
