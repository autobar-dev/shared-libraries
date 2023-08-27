package sharedutils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func NewPostRequest(
	client *http.Client,
	sender_microservice_name string,
	url string,
	body interface{},
) (*http.Response, error) {
	body_json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", sender_microservice_name)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
