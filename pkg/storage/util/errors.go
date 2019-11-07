package util

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
)

var (
	ErrNotFound     = errors.E("storage: obj not found")
	ErrNotStored    = errors.E("storage: not stored")
	ErrNotUpdated   = errors.E("storage: not updated")
	ErrNotDeleted   = errors.E("storage: not deleted")
	ErrInvalidType  = errors.E("storage: invalid type")
	ErrAlreadyExist = errors.E("storage: object already exist")
)
