package userrepository

// Render email template
type ServiceGetUserByIdRequestQuery struct {
	Id string `url:"id"`
}

type ServiceGetUserByIdResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *User   `json:"data"`
}
