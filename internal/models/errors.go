package models

import "errors"

var (
	ErrUserNotFound       = errors.New("User not found in the database.")
	ErrInvalidCredentials = errors.New("Credentials provided do not match.")
	ErrInvalidToken       = errors.New("Invalid Token Provided.")
)
