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
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
}

type Rate struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Rate      float64   `json:"rate"`
	UpdatedAt time.Time `json:"updated_at"`
}
