package authrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func NewAuthRepository(service_url string, microservice_name string) *AuthRepository {
	client := &http.Client{}

	return &AuthRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (ar AuthRepository) LoginUser(email string, password string, remember_me bool) (*Tokens, error) {
	url := fmt.Sprintf("%s/user/login", ar.service_url)

	body := &ServiceUserLoginRequestBody{
		Email:      email,
		Password:   password,
		RememberMe: remember_me,
	}
	body_json, _ := json.Marshal(body)
	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", ar.microservice_name)

	response, err := ar.http_client.Do(req)
	if err != nil {
		return nil, err
	}

	var sulr ServiceUserLoginResponse
	if err := json.NewDecoder(response.Body).Decode(&sulr); err != nil {
		return nil, err
	}

	if sulr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing user login response: %s", *sulr.Error))
	}

	return sulr.Data, nil
}

func (ar AuthRepository) RegisterUser(user_id string, email string, password string) error {
	url := fmt.Sprintf("%s/user/register", ar.service_url)

	body := &ServiceUserRegisterRequestBody{
		UserId:   user_id,
		Email:    email,
		Password: password,
	}
	body_json, _ := json.Marshal(body)
	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return err
	}

	req.Header.Add("X-Internal", ar.microservice_name)

	response, err := ar.http_client.Do(req)
	if err != nil {
		return err
	}

	var surr ServiceUserRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&surr); err != nil {
		return err
	}

	if surr.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing user register response: %s", *surr.Error))
	}

	return nil
}

func (ar AuthRepository) LoginModule(serial_number string, private_key string) (*Tokens, error) {
	url := fmt.Sprintf("%s/module/login", ar.service_url)

	body := &ServiceModuleLoginRequestBody{
		SerialNumber: serial_number,
		PrivateKey:   private_key,
	}
	body_json, _ := json.Marshal(body)
	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", ar.microservice_name)

	response, err := ar.http_client.Do(req)
	if err != nil {
		return nil, err
	}

	var smlr ServiceModuleLoginResponse
	if err := json.NewDecoder(response.Body).Decode(&smlr); err != nil {
		return nil, err
	}

	if smlr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing module login response: %s", *smlr.Error))
	}

	return smlr.Data, nil
}

func (ar AuthRepository) RegisterModule(serial_number string) (private_key *string, err error) {
	url := fmt.Sprintf("%s/module/register", ar.service_url)

	body := &ServiceModuleRegisterRequestBody{
		SerialNumber: serial_number,
	}
	body_json, _ := json.Marshal(body)
	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", ar.microservice_name)

	response, err := ar.http_client.Do(req)
	if err != nil {
		return nil, err
	}

	var smrr ServiceModuleRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&smrr); err != nil {
		return nil, err
	}

	if smrr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing module register response: %s", *smrr.Error))
	}

	return &smrr.Data.PrivateKey, nil
}
