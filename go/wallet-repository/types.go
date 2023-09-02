package walletrepository

import (
	"net/http"
)

type WalletRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type Wallet struct {
	Id           int    `json:"id"`
	UserId       string `json:"user_id"`
	CurrencyCode string `json:"currency_code"`
	Balance      int    `json:"balance"`
}
