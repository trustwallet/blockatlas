package blockatlas

import "errors"

// ErrSourceConn signals that the connection to the source API failed
var ErrSourceConn = errors.New("connection to servers failed")

// ErrInvalidAddr signals that the requested address is invalid
var ErrInvalidAddr = errors.New("invalid address")

// ErrNotFound signals that the resource has not been found
var ErrNotFound = errors.New("not found")
