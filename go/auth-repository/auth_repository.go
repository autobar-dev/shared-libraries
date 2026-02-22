package authrepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewAuthRepository(service_url string, microservice_name string) *AuthRepository {
	client := &http.Client{}

	return &AuthRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (ar *AuthRepository) LoginUser(email string, password string, remember_me bool) (*Tokens, error) {
	url := fmt.Sprintf("%s/user/login", ar.service_url)

	body := &ServiceUserLoginRequestBody{
		Email:      email,
		Password:   password,
		RememberMe: remember_me,
	}

	var sulr ServiceUserLoginResponse
	res, err := sharedutils.NewPostRequest(ar.http_client, ar.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&sulr); err != nil {
		return nil, err
	}

	if sulr.Status == "error" {
		return nil, fmt.Errorf("error while parsing user login response: %s", *sulr.Error)
	}

	return sulr.Data, nil
}

func (ar *AuthRepository) RegisterUser(surrb *ServiceUserRegisterRequestBody) error {
	url := fmt.Sprintf("%s/user/register", ar.service_url)

	var surr ServiceUserRegisterResponse
	res, err := sharedutils.NewPostRequest(ar.http_client, ar.microservice_name, url, surrb)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&surr); err != nil {
		return err
	}

	if surr.Status == "error" {
		return fmt.Errorf("error while parsing response: %s", *surr.Error)
	}

	return nil
}

func (ar *AuthRepository) LoginModule(serial_number string, private_key string) (*Tokens, error) {
	url := fmt.Sprintf("%s/module/login", ar.service_url)

	body := &ServiceModuleLoginRequestBody{
		SerialNumber: serial_number,
		PrivateKey:   private_key,
	}

	var smlr ServiceModuleLoginResponse
	res, err := sharedutils.NewPostRequest(ar.http_client, ar.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&smlr); err != nil {
		return nil, err
	}

	if smlr.Status == "error" {
		return nil, fmt.Errorf("error while parsing module login response: %s", *smlr.Error)
	}

	return smlr.Data, nil
}

func (ar *AuthRepository) RegisterModule(serial_number string) (private_key *string, err error) {
	url := fmt.Sprintf("%s/module/register", ar.service_url)

	body := &ServiceModuleRegisterRequestBody{
		SerialNumber: serial_number,
	}

	response, err := sharedutils.NewPostRequest(ar.http_client, ar.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var smrr ServiceModuleRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&smrr); err != nil {
		return nil, err
	}

	if smrr.Status == "error" {
		return nil, fmt.Errorf("error while parsing module register response: %s", *smrr.Error)
	}

	return &smrr.Data.PrivateKey, nil
}
