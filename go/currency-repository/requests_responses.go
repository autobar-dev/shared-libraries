package currencyrepository

// Get currency
type ServiceGetCurrencyResponse struct {
	Status string    `json:"status"`
	Error  *string   `json:"error"`
	Data   *Currency `json:"data"`
}
