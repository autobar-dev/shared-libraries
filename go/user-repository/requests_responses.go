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

// Locale
type ServiceGetLocaleRequestQuery struct {
	Code string `url:"code"`
}

type ServiceGetLocaleResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Locale `json:"data"`
}

// Role
type ServiceGetRoleRequestQuery struct {
	Name string `url:"name"`
}

type ServiceGetRoleResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Role   `json:"data"`
}
