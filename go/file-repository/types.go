package filerepository

import (
	"net/http"
	"time"
)

type FileRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type File struct {
	Id        string    `json:"id"`
	Extension string    `json:"extension"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
