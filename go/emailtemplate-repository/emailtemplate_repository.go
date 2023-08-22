package emailtemplaterepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func NewEmailTemplateRepository(service_url string, microservice_name string) *EmailTemplateRepository {
	client := &http.Client{}

	return &EmailTemplateRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (etr EmailTemplateRepository) RenderTemplate(
	name string,
	version *string,
	locale string,
	params map[string]interface{},
) (*RenderedTemplate, error) {
	url := fmt.Sprintf("%s/render", etr.service_url)

	body := &ServiceRenderTemplateRequestBody{
		TemplateName:    name,
		TemplateVersion: version,
		Locale:          locale,
		Params:          params,
	}
	body_json, _ := json.Marshal(body)
	body_reader := bytes.NewReader(body_json)

	req, err := http.NewRequest(http.MethodPost, url, body_reader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Internal", etr.microservice_name)

	response, err := etr.http_client.Do(req)
	if err != nil {
		return nil, err
	}

	var srtr ServiceRenderTemplateResponse
	if err := json.NewDecoder(response.Body).Decode(&srtr); err != nil {
		return nil, err
	}

	if srtr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing render template response: %s", *srtr.Error))
	}

	return srtr.Data, nil
}
