package service

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrForbidden = errors.New("forbidden")

	ErrInvalidData = errors.New("invalid data")

	ErrInvalidPathParam = errors.New("invalid path param")

	ErrIdIncorrect = errors.New("id is not provided or incorrect ")
)
