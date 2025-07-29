package errors

import "minhex/src/domain/shared/errors"

var (
	ErrUserNotFound      = errors.NewDomainError("users", "USER_NOT_FOUND", "User not found")
	ErrUserAlreadyExists = errors.NewDomainError("users", "USER_ALREADY_EXISTS", "User already exists")
	ErrInvalidEmail      = errors.NewDomainError("users", "INVALID_EMAIL", "Invalid email format")
	ErrInvalidName       = errors.NewDomainError("users", "INVALID_NAME", "Name cannot be empty")
)
