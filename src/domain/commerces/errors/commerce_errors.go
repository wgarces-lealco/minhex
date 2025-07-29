package errors

import "minhex/src/domain/shared/errors"

var (
	ErrCommerceNotFound      = errors.NewDomainError("commerces", "COMMERCE_NOT_FOUND", "Commerce not found")
	ErrCommerceAlreadyExists = errors.NewDomainError("commerces", "COMMERCE_ALREADY_EXISTS", "Commerce already exists")
	ErrInvalidCommerceName   = errors.NewDomainError("commerces", "INVALID_COMMERCE_NAME", "Commerce name cannot be empty")
	ErrInvalidOwner          = errors.NewDomainError("commerces", "INVALID_OWNER", "Invalid owner ID")
	ErrCommerceAlreadyActive = errors.NewDomainError("commerces", "COMMERCE_ALREADY_ACTIVE", "Commerce is already active")
)
