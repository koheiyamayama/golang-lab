package error

import "errors"

type InfraError error

var (
	ErrOpenAPINotFound   = InfraError(errors.New("not found"))
	ErrOpenAPIInternal   = InfraError(errors.New("internal server error"))
	ErrOpenAPIBadRequest = InfraError(errors.New("bad request"))
)
