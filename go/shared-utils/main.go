package sharedutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ServiceResponseType struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

func MakePostRequest(
	client *http.Client,
	sender_microservice_name string,
	url string,
	body interface{},
	response *ServiceResponseType,
) error {
	body_json, err := json.Marshal(body)
	if err != nil {
		return err
	}

	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return err
	}

	req.Header.Add("X-Internal", sender_microservice_name)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	if response.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing response: %s", *response.Error))
	}

	return err
}
