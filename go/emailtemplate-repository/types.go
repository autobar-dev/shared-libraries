package emailtemplaterepository

import (
	"net/http"
)

type EmailTemplateRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type RenderedTemplate struct {
	Plain string `json:"plain"`
	Html  string `json:"html"`
}
