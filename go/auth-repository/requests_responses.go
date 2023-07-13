package authrepository

// Module registration
type ServiceModuleRegisterRequestBody struct {
	SerialNumber string `json:"serial_number"`
}

type ServiceModuleRegisterResponse struct {
	Status string         `json:"status"`
	Error  *string        `json:"error"`
	Data   *ServiceModule `json:"data"`
}

// Module login
type ServiceModuleLoginRequestBody struct {
	SerialNumber string `json:"serial_number"`
	PrivateKey   string `json:"private_key"`
	RememberMe   *bool  `json:"remember_me"`
}

type ServiceModuleLoginResponseData struct {
	SessionId string `json:"session_id"`
}

type ServiceModuleLoginResponse struct {
	Status string                          `json:"status"`
	Error  *string                         `json:"error"`
	Data   *ServiceModuleLoginResponseData `json:"data"`
}

// Verify session
type ServiceVerifySessionResponse struct {
	Status string              `json:"status"`
	Error  *string             `json:"error"`
	Data   *ServiceSessionData `json:"data"`
}

// All sessions for client
type ServiceAllSessionsForClientResponse struct {
	Status string                `json:"status"`
	Error  *string               `json:"error"`
	Data   *[]ServiceSessionInfo `json:"data"`
}

// Remove session
type ServiceRemoveSessionRequestBody struct {
	SessionId string `json:"session_id"`
}

type ServiceRemoveSessionResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

// Remove session by internal id
type ServiceRemoveSessionByInternalIdRequestBody struct {
	InternalId int     `json:"internal_id"`
	SessionId  *string `json:"session_id"` // user's session ID to check if allowed
}

type ServiceRemoveSessionByInternalIdResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
}

// Login user
type ServiceUserLoginRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe *bool  `json:"remember_me"`
}

type ServiceUserLoginResponseData struct {
	SessionId string `json:"session_id"`
}

type ServiceUserLoginResponse struct {
	Status string                        `json:"status"`
	Error  *string                       `json:"error"`
	Data   *ServiceUserLoginResponseData `json:"data"`
}

// Register user
type ServiceUserRegisterRequestBody struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	AutoLogin  *bool  `json:"auto_login"`
	RememberMe *bool  `json:"remember_me"`
}

type ServiceUserRegisterResponseData struct {
	SessionId string `json:"session_id"`
}

type ServiceUserRegisterResponse struct {
	Status string                           `json:"status"`
	Error  *string                          `json:"error"`
	Data   *ServiceUserRegisterResponseData `json:"data"`
}
