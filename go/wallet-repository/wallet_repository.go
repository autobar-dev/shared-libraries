package walletrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewWalletRepository(service_url string, microservice_name string) *WalletRepository {
	client := &http.Client{}

	return &WalletRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (wr WalletRepository) Create(
	user_id string,
	currency_code string,
) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet/create", wr.service_url)

	body := &ServiceCreateWalletRequestBody{
		UserId:       user_id,
		CurrencyCode: currency_code,
	}

	res, err := sharedutils.NewPostRequest(wr.http_client, wr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var scwr ServiceCreateWalletResponse
	if err := json.NewDecoder(res.Body).Decode(&scwr); err != nil {
		return nil, err
	}

	if scwr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing create wallet response: %s", *scwr.Error))
	}

	return scwr.Data, nil
}
