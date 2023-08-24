package userrepository

import (
	"net/http"
	"time"
)

type UserRepository struct {
	service_url       string
	http_client       *http.Client
	microservice_name string
}

type User struct {
	Id                         string    `json:"id"`
	Email                      string    `json:"email"`
	FirstName                  string    `json:"first_name"`
	LastName                   string    `json:"last_name"`
	DateOfBirth                time.Time `json:"date_of_birth"`
	Locale                     string    `json:"locale"`
	IdentityVerificationId     string    `json:"identity_verification_id"`
	IdentityVerificationSource string    `json:"identity_verification_source"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
