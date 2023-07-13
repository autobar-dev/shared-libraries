package authrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func NewAuthRepository(service_url string) *AuthRepository {
	return &AuthRepository{
		service_url: service_url,
	}
}

func (ar AuthRepository) LoginModule(serial_number string, private_key string, remember_me *bool) (*string, error) {
	url := fmt.Sprintf("%s/module/login", ar.service_url)

	body := &ServiceModuleLoginRequestBody{
		SerialNumber: serial_number,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var smlr ServiceModuleLoginResponse
	if err := json.NewDecoder(response.Body).Decode(&smlr); err != nil {
		return nil, err
	}

	if smlr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing module login response: %s", *smlr.Error))
	}

	return &smlr.Data.SessionId, nil
}

func (ar AuthRepository) RegisterModule(serial_number string) (*ServiceModule, error) {
	url := fmt.Sprintf("%s/module/register", ar.service_url)

	body := &ServiceModuleRegisterRequestBody{
		SerialNumber: serial_number,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var smrr ServiceModuleRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&smrr); err != nil {
		return nil, err
	}

	if smrr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing module register response: %s", *smrr.Error))
	}

	return smrr.Data, nil
}

func (ar AuthRepository) GetAllSessionsForClient(session string) (*[]ServiceSessionInfo, error) {
	url := fmt.Sprintf("%s/session/all-for-client?session_id=%s", ar.service_url, session)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var sasfcr ServiceAllSessionsForClientResponse
	if err := json.NewDecoder(response.Body).Decode(&sasfcr); err != nil {
		return nil, err
	}

	if sasfcr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing all sessions for client response: %s", *sasfcr.Error))
	}

	return sasfcr.Data, nil
}

func (ar AuthRepository) RemoveSession(session string) error {
	url := fmt.Sprintf("%s/session/remove", ar.service_url)

	body := &ServiceRemoveSessionRequestBody{
		SessionId: session,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var srsr ServiceRemoveSessionResponse
	if err := json.NewDecoder(response.Body).Decode(&srsr); err != nil {
		return err
	}

	if srsr.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing module login response: %s", *srsr.Error))
	}

	return nil
}

func (ar AuthRepository) RemoveSessionByInternalId(internal_id int) error {
	url := fmt.Sprintf("%s/session/remove-by-internal-id", ar.service_url)

	body := &ServiceRemoveSessionByInternalIdRequestBody{
		InternalId: internal_id,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var srsbiir ServiceRemoveSessionByInternalIdResponse
	if err := json.NewDecoder(response.Body).Decode(&srsbiir); err != nil {
		return err
	}

	if srsbiir.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing remove session by internal id response: %s", *srsbiir.Error))
	}

	return nil
}

func (ar AuthRepository) VerifySession(session string) (*ServiceSessionData, error) {
	url := fmt.Sprintf("%s/session/verify?session_id=%s", ar.service_url, session)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var svsr ServiceVerifySessionResponse
	if err := json.NewDecoder(response.Body).Decode(&svsr); err != nil {
		return nil, err
	}

	if svsr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing session verify response: %s", *svsr.Error))
	}

	return svsr.Data, nil
}

func (ar AuthRepository) LoginUser(email string, password string, remember_me *bool) (*string, error) {
	url := fmt.Sprintf("%s/user/login", ar.service_url)

	body := &ServiceUserLoginRequestBody{
		Email:      email,
		Password:   password,
		RememberMe: remember_me,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var sulr ServiceUserLoginResponse
	if err := json.NewDecoder(response.Body).Decode(&sulr); err != nil {
		return nil, err
	}

	if sulr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing user login response: %s", *sulr.Error))
	}

	return &sulr.Data.SessionId, nil
}

func (ar AuthRepository) RegisterUser(email string, password string, auto_login *bool, remember_me *bool) (*string, error) {
	url := fmt.Sprintf("%s/user/register", ar.service_url)

	body := &ServiceUserRegisterRequestBody{
		Email:      email,
		Password:   password,
		AutoLogin:  auto_login,
		RememberMe: remember_me,
	}
	body_json, _ := json.Marshal(body)

	response, err := http.Post(url, "application/json", bytes.NewReader(body_json))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var surr ServiceUserRegisterResponse
	if err := json.NewDecoder(response.Body).Decode(&surr); err != nil {
		return nil, err
	}

	if surr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing user register response: %s", *surr.Error))
	}

	return &surr.Data.SessionId, nil
}
