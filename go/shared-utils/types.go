package sharedutils

type ServiceResponseStatus string

const (
	OkServiceResponse    ServiceResponseStatus = "ok"
	ErrorServiceResponse ServiceResponseStatus = "error"
)

type ServiceResponse struct {
	Status ServiceResponseStatus `json:"status"`
	Error  *string               `json:"error"`
}
