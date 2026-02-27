package productrepository

import (
	"net/http"
	"time"
)

type ProductRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type ProductBadgeType string

const (
	ProductBadgeTypePrimary   ProductBadgeType = "primary"
	ProductBadgeTypeSecondary ProductBadgeType = "secondary"
)

type ProductBadge struct {
	Type  ProductBadgeType `json:"type"`
	Label string           `json:"label"`
	Value *string          `json:"value"`
}

type Product struct {
	Id           int               `json:"id"`
	Slug         string            `json:"slug"`
	ImageFileId  int               `json:"image_file_id"`
	Names        map[string]string `json:"names"`
	Descriptions map[string]string `json:"descriptions"`
	Enabled      bool              `json:"enabled"`
	Badges       []ProductBadge    `json:"badges"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type ProductResponseType string

const (
	RedirectProductResponseType ProductResponseType = "redirect"
	DataProductResponseType     ProductResponseType = "product"
)
