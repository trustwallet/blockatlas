package storage

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("storage: obj not found")
	ErrNotStored    = errors.New("storage: not stored")
	ErrNotUpdated   = errors.New("storage: not updated")
	ErrNotDeleted   = errors.New("storage: not deleted")
	ErrAlreadyExist = errors.New("storage: object already exist")
)
