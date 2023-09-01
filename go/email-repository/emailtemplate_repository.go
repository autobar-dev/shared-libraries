package emailrepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewEmailRepository(service_url string, microservice_name string) *EmailRepository {
	client := &http.Client{}

	return &EmailRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (er EmailRepository) Send(
	from string,
	to string,
	subject string,
	message_plain string,
	message_html string,
) error {
	url := fmt.Sprintf("%s/send", er.service_url)

	body := &ServiceSendRequestBody{
		From:    from,
		To:      to,
		Subject: subject,
		Message: ServiceSendRequestBodyMessage{
			Plain: message_plain,
			Html:  message_html,
		},
	}

	res, err := sharedutils.NewPostRequest(er.http_client, er.microservice_name, url, body)
	if err != nil {
		return err
	}

	var ssr ServiceSendResponse
	if err := json.NewDecoder(res.Body).Decode(&ssr); err != nil {
		return err
	}

	if ssr.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing send response: %s", *ssr.Error))
	}

	return nil
}
