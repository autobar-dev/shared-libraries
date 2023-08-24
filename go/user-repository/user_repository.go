package userrepository

import (
	"net/http"
)

func NewUserRepository(service_url string, microservice_name string) *UserRepository {
	client := &http.Client{}

	return &UserRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}
