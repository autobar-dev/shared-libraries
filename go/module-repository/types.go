package modulerepository

import (
	"net/http"
	"time"
)

type ModuleRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type Module struct {
	Id              int32          `json:"id"`
	SerialNumber    string         `json:"serial_number"`
	StationId       *string        `json:"station_id"`
	ProductId       *string        `json:"product_id"`
	Enabled         bool           `json:"enabled"`
	Prices          map[string]int `json:"prices"`
	DisplayCurrency string         `json:"display_currency"`
	DisplayUnit     DisplayUnit    `json:"display_unit"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type DisplayUnit struct {
	Id                     int32     `json:"id"`
	Symbol                 string    `json:"symbol"`
	DivisorFromMillilitres float64   `json:"divisor_from_millilitres"`
	DecimalsDisplayed      int32     `json:"decimals_displayed"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
