package productrepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewProductRepository(service_url string, microservice_name string) *ProductRepository {
	client := &http.Client{}

	return &ProductRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (pr *ProductRepository) GetProductById(
	id string,
) (*Product, error) {
	url := fmt.Sprintf("%s/?id=%s", pr.service_url, id)

	res, err := sharedutils.NewGetRequest(pr.http_client, pr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgpr ServiceGetProductResponse
	if err := json.NewDecoder(res.Body).Decode(&sgpr); err != nil {
		return nil, err
	}

	if sgpr.Status == "error" {
		return nil, fmt.Errorf("error while parsing get product response: %s", *sgpr.Error)
	}

	return sgpr.Data.Product, nil
}
