package filerepository

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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
	id int,
	download bool,
) (*File, error) {
	url := fmt.Sprintf("%s/?id=%d&download=%t", fr.service_url, id, download)

	res, err := sharedutils.NewGetRequest(fr.http_client, fr.microservice_name, url)
	if err != nil {
		return nil, err
	}

	var sgfr ServiceGetFileResponse
	if err := json.NewDecoder(res.Body).Decode(&sgfr); err != nil {
		return nil, err
	}

	if sgfr.Status == "error" {
		return nil, fmt.Errorf("error while parsing get file response: %s", *sgfr.Error)
	}

	return sgfr.Data, nil
}

func (fr *FileRepository) DeleteFile(
	id int,
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
		return fmt.Errorf("error while parsing delete file response: %s", *sdfr.Error)
	}

	return nil
}

func (fr *FileRepository) UploadFile(
	directory string,
	file_header *multipart.FileHeader,
) (int, error) {
	url := fmt.Sprintf("%s/upload", fr.service_url)

	file, err := file_header.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer pw.Close()
		defer writer.Close()

		part, err := writer.CreateFormFile("file", file_header.Filename)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, file); err != nil {
			pw.CloseWithError(err)
			return
		}

		if err := writer.WriteField("directory", directory); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()

	req, err := http.NewRequest("POST", url, pr)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-Internal", fr.microservice_name)

	resp, err := fr.http_client.Do(req)
	if err != nil {
		return 0, err
	}

	var sur ServiceUploadFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&sur); err != nil {
		return 0, err
	}

	if sur.Status == "error" {
		return 0, fmt.Errorf("error while parsing upload file response: %s", *sur.Error)
	}

	return sur.Data.Id, nil
}
