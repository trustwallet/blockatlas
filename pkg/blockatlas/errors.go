package blockatlas

import "errors"

var (
	// ErrSourceConn signals that the connection to the source API failed
	ErrSourceConn = errors.New("connection to servers failed")

	// ErrInvalidAddr signals that the requested address is invalid
	ErrInvalidAddr = errors.New("invalid address")

	// ErrNotFound signals that the resource has not been found
	ErrNotFound = errors.New("not found")

	// ErrInvalidKey signals that the requested key is invalid
	ErrInvalidKey = errors.New("invalid key")
)
