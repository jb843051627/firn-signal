package model

import "errors"

var (
	ErrNotFound      = errors.New("firn record not found")
	ErrInvalidState  = errors.New("invalid firn state")
	ErrInvalidInput  = errors.New("invalid firn input")
	ErrIncomplete    = errors.New("incomplete thermal profile")
	ErrQualityBlock  = errors.New("quality assessment blocks release")
	ErrAlreadyExists = errors.New("firn record already exists")
)
