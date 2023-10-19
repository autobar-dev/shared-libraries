package filerepository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	sharedutils "github.com/autobar-dev/shared-libraries/go/shared-utils"
)

func NewFileRepository(service_url string, microservice_name string) *FileRepository {
	client := &http.Client{}

	return &FileRepository{
		service_url:       service_url,
		http_client:       client,
		microservice_name: microservice_name,
	}
}

func (fr *FileRepository) GetFile(
	id string,
) (*File, error) {
	url := fmt.Sprintf("%s/?id=%s", fr.service_url, id)

	res, err := sharedutils.NewGetRequest(fr.http_client, fr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgfr ServiceGetFileResponse
	if err := json.NewDecoder(res.Body).Decode(&sgfr); err != nil {
		return nil, err
	}

	if sgfr.Status == "error" {
		return nil, errors.New(fmt.Sprintf("error while parsing get file response: %s", *sgfr.Error))
	}

	return sgfr.Data, nil
}

func (fr *FileRepository) DeleteFile(
	id string,
) error {
	url := fmt.Sprintf("%s/delete", fr.service_url)

	sdfrb := &ServiceDeleteFileRequestBody{
		Id: id,
	}
	res, err := sharedutils.NewDeleteRequest(fr.http_client, fr.microservice_name, url, sdfrb)
	if err != nil {
		return err
	}

	var sdfr ServiceDeleteFileResponse
	if err := json.NewDecoder(res.Body).Decode(&sdfr); err != nil {
		return err
	}

	if sdfr.Status == "error" {
		return errors.New(fmt.Sprintf("error while parsing delete file response: %s", *sdfr.Error))
	}

	return nil
}
