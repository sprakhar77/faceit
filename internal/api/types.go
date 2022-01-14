package api

import (
	"net/mail"
)

// GetAllUsersRequest is the request structure for /GET: /users
type GetAllUsersRequest struct {
	Country string  `json:"country"`
	Limit   *uint64 `json:"limit"`
	Offset  *uint64 `json:"offset"`
}

// CreateUserRequest is the request structure for /POST: /users
type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nickname,omitempty"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

// Validate the parameters of incoming CreateUserRequest
func (r *CreateUserRequest) Validate() error {
	_, err := mail.ParseAddress(r.Email)
	return err
}

// UpdateUserRequest is the request structure for /PUT: /users
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nickname,omitempty"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

// Validate the parameters of incoming UpdateUserRequest
func (r *UpdateUserRequest) Validate() error {
	_, err := mail.ParseAddress(r.Email)
	return err
}
