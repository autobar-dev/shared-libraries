package emailrepository

import (
	"net/http"
)

type EmailRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}
