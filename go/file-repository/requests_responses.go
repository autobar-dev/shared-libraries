package filerepository

// Get file
type ServiceGetFileResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *File   `json:"data"`
}

// Delete file
type ServiceDeleteFileRequestBody struct {
	Id string `json:"id"`
}

type ServiceDeleteFileResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}
