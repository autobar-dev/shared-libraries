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
	Id         int       `json:"id"`
	S3ObjectId string    `json:"s3_object_id"`
	Name       string    `json:"name"`
	Checksum   string    `json:"checksum"`
	Size       int64     `json:"size"`
	CreatedAt  time.Time `json:"created_at"`

	Extension string `json:"extension"`
	Url       string `json:"url"`
}
