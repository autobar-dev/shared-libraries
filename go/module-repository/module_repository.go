package modulerepository

import (
	"net/http"
)

func NewModuleRepository(service_url string, microservice_name string) *ModuleRepository {
	client := &http.Client{}

	return &ModuleRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}
