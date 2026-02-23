package userrepository

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
	"github.com/google/go-querystring/query"
)

func NewUserRepository(service_url string, microservice_name string) *UserRepository {
	client := &http.Client{}

	return &UserRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (ur *UserRepository) GetUserById(id string) (*User, error) {
	query_values, _ := query.Values(&ServiceGetUserByIdRequestQuery{
		Id: id,
	})
	query_string := query_values.Encode()

	url := fmt.Sprintf("%s/?%s", ur.service_url, query_string)

	res, err := sharedutils.NewGetRequest(ur.http_client, ur.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgubir ServiceGetUserByIdResponse
	if err := json.NewDecoder(res.Body).Decode(&sgubir); err != nil {
		return nil, err
	}

	if sgubir.Status == "error" {
		return nil, fmt.Errorf("error while parsing get user by id response: %s", *sgubir.Error)
	}

	return sgubir.Data, nil
}

func (ur *UserRepository) GetLocale(code string) (*Locale, error) {
	query_values, _ := query.Values(&ServiceGetLocaleRequestQuery{
		Code: code,
	})
	query_string := query_values.Encode()

	url := fmt.Sprintf("%s/locale?%s", ur.service_url, query_string)

	res, err := sharedutils.NewGetRequest(ur.http_client, ur.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sglr ServiceGetLocaleResponse
	if err := json.NewDecoder(res.Body).Decode(&sglr); err != nil {
		return nil, err
	}

	if sglr.Status == "error" {
		return nil, fmt.Errorf("error while parsing get locale response: %s", *sglr.Error)
	}

	return sglr.Data, nil
}

func (ur *UserRepository) GetRole(name string) (*Role, error) {
	query_values, _ := query.Values(&ServiceGetRoleRequestQuery{
		Name: &name,
	})
	query_string := query_values.Encode()

	url := fmt.Sprintf("%s/role?%s", ur.service_url, query_string)

	res, err := sharedutils.NewGetRequest(ur.http_client, ur.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgr ServiceGetRoleResponse
	if err := json.NewDecoder(res.Body).Decode(&sgr); err != nil {
		return nil, err
	}

	if sgr.Status == "error" {
		return nil, fmt.Errorf("error while parsing get role response: %s", *sgr.Error)
	}

	return sgr.Data, nil
}

func (ur *UserRepository) GetRoleById(id int) (*Role, error) {
	query_values, _ := query.Values(&ServiceGetRoleRequestQuery{
		Id: &id,
	})
	query_string := query_values.Encode()

	url := fmt.Sprintf("%s/role?%s", ur.service_url, query_string)

	res, err := sharedutils.NewGetRequest(ur.http_client, ur.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgr ServiceGetRoleResponse
	if err := json.NewDecoder(res.Body).Decode(&sgr); err != nil {
		return nil, err
	}

	if sgr.Status == "error" {
		return nil, fmt.Errorf("error while parsing get role response: %s", *sgr.Error)
	}

	return sgr.Data, nil
}
