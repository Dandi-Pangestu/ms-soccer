package domains

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Your requested item is not found")
	ErrConflict            = errors.New("Your item already exist")
	ErrBadRequest          = errors.New("Given param is not valid")
	ErrUnprocessableEntity = errors.New("Unprocessable entity")
	ErrUnauthenticate      = errors.New("Unauthenticate")
	ErrUnauthorized        = errors.New("Unauthorized")
)
