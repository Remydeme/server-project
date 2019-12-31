package db

import (
	"errors"
)

var (
	ErrorObjectAlreadyExists = errors.New("Object already exists duplicate object are forbidden")
	ErrorOccured             = errors.New("An error occured impossible to add new object in database")
	ErrorValidation          = errors.New("Internal error validation error")
)
