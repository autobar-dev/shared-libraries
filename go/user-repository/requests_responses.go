package userrepository

// Render email template
type ServiceGetUserRequestBody struct{}

type ServiceGetUserResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *User   `json:"data"`
}
