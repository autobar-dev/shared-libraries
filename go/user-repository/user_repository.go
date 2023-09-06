package userrepository

import (
	"encoding/json"
	"errors"
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

	url := fmt.Sprintf("%s/get-by-id?%s", ur.service_url, query_string)

	res, err := sharedutils.NewGetRequest(ur.http_client, ur.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgubir ServiceGetUserByIdResponse
	if err := json.NewDecoder(res.Body).Decode(&sgubir); err != nil {
		return nil, err
	}

	if sgubir.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing get user by id response: %s", *sgubir.Error))
	}

	return sgubir.Data, nil
}
