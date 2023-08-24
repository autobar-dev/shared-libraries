package userrepository

import (
	"net/http"
)

type UserRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type User struct {
	Id string `json:"id"`
}
