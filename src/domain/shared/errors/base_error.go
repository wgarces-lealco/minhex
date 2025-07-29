package errors

import "fmt"

type DomainError struct {
	Code    string
	Message string
	Domain  string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s:%s] %s", e.Domain, e.Code, e.Message)
}

func NewDomainError(domain, code, message string) *DomainError {
	return &DomainError{
		Domain:  domain,
		Code:    code,
		Message: message,
	}
}
