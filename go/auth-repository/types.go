package authrepository

import (
	"net/http"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenOwnerType string

const (
	UserTokenOwnerType   TokenOwnerType = "user"
	ModuleTokenOwnerType TokenOwnerType = "module"
)

type AccessTokenPayload struct {
	ClientType TokenOwnerType
	Identifier string
	Role       *string
}

type ServiceModule struct {
	Id           int       `json:"id"`
	SerialNumber string    `json:"serial_number"`
	PrivateKey   string    `json:"private_key"`
	CreatedAt    time.Time `json:"created_at"`
}

type AuthRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}
