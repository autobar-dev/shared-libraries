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

type TransactionType string

const (
	TransactionTypeDeposit        TransactionType = "deposit"
	TransactionTypeWithdraw       TransactionType = "withdraw"
	TransactionTypePurchase       TransactionType = "purchase"
	TransactionTypeRefund         TransactionType = "refund"
	TransactionTypeCurrencyChange TransactionType = "currency_change"
)

type Transaction struct {
	Id              string          `json:"id"`
	WalletId        int             `json:"wallet_id"`
	TransactionType TransactionType `json:"transaction_type"`
	Value           int             `json:"value"`
	CurrencyCode    string          `json:"currency_code"`
	CreatedAt       string          `json:"created_at"`
}
