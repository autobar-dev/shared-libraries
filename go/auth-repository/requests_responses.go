package authrepository

// Refresh tokens
type ServiceRefreshTokensRequestBody struct {
	RefreshToken string `json:"refresh_token"`
}

type ServiceRefreshTokensResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

// Login user
type ServiceUserLoginRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type ServiceUserLoginResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Tokens `json:"data"`
}

// Register user
type ServiceUserRegisterRequestBody struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ServiceUserRegisterResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *string `json:"data"` // always `nil`
}

// Module login
type ServiceModuleLoginRequestBody struct {
	SerialNumber string `json:"serial_number"`
	PrivateKey   string `json:"private_key"`
}

type ServiceModuleLoginResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *Tokens `json:"data"`
}

// Module registration
type ServiceModuleRegisterRequestBody struct {
	SerialNumber string `json:"serial_number"`
}

type ServiceModuleRegisterResponseData struct {
	PrivateKey string `json:"private_key"`
}

type ServiceModuleRegisterResponse struct {
	Status string                             `json:"status"`
	Error  *string                            `json:"error"`
	Data   *ServiceModuleRegisterResponseData `json:"data"`
}
