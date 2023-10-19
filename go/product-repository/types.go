package productrepository

import (
	"net/http"
	"time"

	filerepository "github.com/autobar-dev/shared-libraries/go/file-repository"
)

type ProductRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type Product struct {
	Id           string              `json:"id"`
	Names        map[string]string   `json:"names"`
	Descriptions map[string]string   `json:"descriptions"`
	Cover        filerepository.File `json:"cover"`
	Enabled      bool                `json:"enabled"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
}

type ProductResponseType string

const (
	RedirectProductResponseType ProductResponseType = "redirect"
	DataProductResponseType     ProductResponseType = "product"
)
