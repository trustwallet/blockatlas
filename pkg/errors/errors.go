package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Params map[string]interface{}

// Error represents a error's specification.
type Error struct {
	Err   error
	Type  Type
	meta  map[string]interface{}
	stack []string
}

var (
	_ error = (*Error)(nil)
)

func (e *Error) isEmpty() bool {
	return e.meta == nil && e.Type == TypeNone && e.Err == nil
}

func (e *Error) Error() string {
	r, err := e.MarshalJSON()
	if err != nil {
		return e.Err.Error()
	}
	return string(r)
}

func (e *Error) String() string {
	msg := e.Err.Error()
	if e.Type != TypeNone {
		msg = fmt.Sprintf("%s | Type: %s", msg, e.Type.String())
	}
	if len(e.Meta()) > 0 {
		msg = fmt.Sprintf("%s | Meta: %s", msg, e.Meta())
	}
	if len(e.stack) > 0 {
		msg = fmt.Sprintf("%s | Stack: %s", msg, e.stack)
	}
	return msg
}

// SetMeta sets the error's meta data.
func (e *Error) SetMeta(data Params) *Error {
	e.meta = data
	return e
}

func (e *Error) Meta() string {
	r, err := json.Marshal(e.meta)
	if err != nil {
		return ""
	}
	return string(r)
}

// JSON creates a properly formatted JSON
func (e *Error) JSON() interface{} {
	p := Params{}
	if e.meta != nil {
		p["meta"] = e.meta
	}
	if e.Err != nil {
		p["error"] = e.Err.Error()
	}
	if e.Type != TypeNone {
		p["type"] = e.Type.String()
	}
	if len(e.stack) > 0 {
		p["stack"] = e.stack
	}
	return p
}

// MarshalJSON implements the json.Marshaller interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.JSON())
}

func (e *Error) PushToSentry() *Error {
	SendError(e)
	return e
}

// T create a new error with runtime stack trace.
func T(args ...interface{}) *Error {
	e := E(args...)
	for i := 1; i <= 5; i++ {
		_, fn, line, ok := runtime.Caller(i)
		if ok {
			e.stack = append(e.stack, fmt.Sprintf("%s:%d", fn, line))
		}
	}
	return e
}

// E create a new error.
func E(args ...interface{}) *Error {
	e := &Error{Type: TypeNone, meta: make(Params)}
	var message []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case string:
			message = append(message, arg)
		case error:
			message = append([]string{arg.Error()}, message...)
		case Type:
			e.Type = arg
		case Params:
			appendMap(e.meta, arg)
		case map[string]interface{}:
			appendMap(e.meta, arg)
		default:
			continue
		}
	}
	if len(message) > 0 {
		msg := strings.Join(message[:], ": ")
		e.Err = errors.New(msg)
	}
	return e
}
