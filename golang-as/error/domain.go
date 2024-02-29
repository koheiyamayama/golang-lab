package error

import "errors"

type DomainError error

var (
	ErrNotFound       = DomainError(errors.New("not found"))
	ErrUnauthorized   = DomainError(errors.New("unauthorized"))
	ErrInternalServer = DomainError(errors.New("internal server error"))
	ErrBadRequest     = DomainError(errors.New("bad request"))
)
