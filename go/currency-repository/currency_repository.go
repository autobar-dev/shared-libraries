package currencyrepository

import (
	"encoding/json"
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

func (cr *CurrencyRepository) GetCurrencyByCode(
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
		return nil, fmt.Errorf("error while getting currency response: %s", *sgcr.Error)
	}

	return sgcr.Data, nil
}

func (cr *CurrencyRepository) GetRate(
	from_code string,
	to_code string,
) (*Rate, error) {
	url := fmt.Sprintf("%s/rate/?from=%s&to=%s", cr.service_url, from_code, to_code)

	res, err := sharedutils.NewGetRequest(cr.http_client, cr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgrr ServiceGetRateResponse
	if err := json.NewDecoder(res.Body).Decode(&sgrr); err != nil {
		return nil, err
	}

	if sgrr.Status == "error" {
		return nil, fmt.Errorf("error while getting rate response: %s", *sgrr.Error)
	}

	return sgrr.Data, nil
}
