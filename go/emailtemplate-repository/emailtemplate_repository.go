package emailtemplaterepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
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

	res, err := sharedutils.NewPostRequest(etr.http_client, etr.microservice_name, url, body)
	if err != nil {
		return nil, err
	}

	var srtr ServiceRenderTemplateResponse
	if err := json.NewDecoder(res.Body).Decode(&srtr); err != nil {
		return nil, err
	}

	if srtr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing render template response: %s", *srtr.Error))
	}

	return srtr.Data, nil
}
