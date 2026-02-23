package currencyrepository

import (
	"net/http"
	"time"
)

type CurrencyRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type Currency struct {
	Id               int       `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	MinorUnitDivisor int       `json:"minor_unit_divisor"`
	Symbol           *string   `json:"symbol"`
	Enabled          bool      `json:"enabled"`
	CreatedAt        time.Time `json:"created_at"`
}

type Rate struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Rate      float64   `json:"rate"`
	FetchedAt time.Time `json:"fetched_at"`
}
