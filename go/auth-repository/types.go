package authrepository

import "time"

type ServiceModule struct {
	Id           int       `json:"id"`
	SerialNumber string    `json:"serial_number"`
	PrivateKey   string    `json:"private_key"`
	CreatedAt    time.Time `json:"created_at"`
}

type ServiceAuthClientType string

const (
	ModuleServiceAuthClientType ServiceAuthClientType = "module"
	UserServiceAuthClientType   ServiceAuthClientType = "user"
)

type ServiceSessionData struct {
	ClientIdentifier string                `json:"client_identifier"`
	ClientType       ServiceAuthClientType `json:"client_type"`
}

type ServiceSessionInfo struct {
	InternalId       int                   `json:"internal_id"`
	ClientIdentifier string                `json:"client_identifier"`
	ClientType       ServiceAuthClientType `json:"client_type"`
	UserAgent        *string               `json:"user_agent"`
	ValidUntil       time.Time             `json:"valid_until"`
	LastUsed         time.Time             `json:"last_used"`
	CreatedAt        time.Time             `json:"created_at"`
}

type AuthRepository struct {
	service_url string
}
