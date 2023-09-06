package currencyrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewCurrencyRepository(service_url string, microservice_name string) *CurrencyRepository {
	client := &http.Client{}

	return &CurrencyRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (cr CurrencyRepository) GetCurrencyByCode(
	code string,
) (*Currency, error) {
	url := fmt.Sprintf("%s/currency/?code=%s", cr.service_url, code)

	res, err := sharedutils.NewGetRequest(cr.http_client, cr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgcr ServiceGetCurrencyResponse
	if err := json.NewDecoder(res.Body).Decode(&sgcr); err != nil {
		return nil, err
	}

	if sgcr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing get currency response: %s", *sgcr.Error))
	}

	return sgcr.Data, nil
}
