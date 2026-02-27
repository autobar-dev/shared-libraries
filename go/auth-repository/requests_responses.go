package authrepository

import "time"

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
	UserId      string `json:"user_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	LocaleId    int    `json:"locale_id"`
	RoleId      int    `json:"role_id"`
}

type ServiceUserRegisterResponse struct {
	Status string  `json:"status"`
	Error  *string `json:"error"`
	Data   *string `json:"data"` // always `nil`
}

// Module login challenge
type ServiceModuleLoginChallengeRequestBody struct {
	SerialNumber string `json:"serial_number"`
}

type ServiceModuleLoginChallengeResponseData struct {
	Challenge string    `json:"challenge"`
	ExpiresAt time.Time `json:"expires_at"`
}

type ServiceModuleLoginChallengeResponse struct {
	Status string                                   `json:"status"`
	Error  *string                                  `json:"error"`
	Data   *ServiceModuleLoginChallengeResponseData `json:"data"`
}

// Module login
type ServiceModuleLoginRequestBody struct {
	CertificateBase64 string `json:"certificate_base64"`
	Challenge         string `json:"challenge"`
	SignatureBase64   string `json:"signature_base64"`
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
	CertificateBase64 string `json:"certificate_base64"`
	PrivateKeyBase64  string `json:"private_key_base64"`
}

type ServiceModuleRegisterResponse struct {
	Status string                             `json:"status"`
	Error  *string                            `json:"error"`
	Data   *ServiceModuleRegisterResponseData `json:"data"`
}
