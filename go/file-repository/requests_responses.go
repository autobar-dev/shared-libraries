package filerepository

// Get file
type ServiceGetFileResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *File   `json:"data"`
}

// Delete file
type ServiceDeleteFileRequestBody struct {
	Id int `json:"id"`
}

type ServiceDeleteFileResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

// Upload file
type ServiceUploadFileResponseData struct {
	Id int `json:"id"`
}

type ServiceUploadFileResponse struct {
	Status string                         `json:"status"`
	Error  *string                        `json:"error"`
	Data   *ServiceUploadFileResponseData `json:"data"`
}
