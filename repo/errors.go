package repo

import "errors"

var (
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrRecordNotFound = errors.New("record not found")
)
