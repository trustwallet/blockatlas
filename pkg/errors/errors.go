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
	meta  interface{}
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
func (e *Error) SetMeta(data interface{}) *Error {
	e.meta = data
	return e
}

func (e *Error) Meta() string {
	var meta string
	switch arg := e.meta.(type) {
	case nil:
		return ""
	case string:
		return meta
	case Params:
		r, err := json.Marshal(arg)
		if err != nil {
			return ""
		}
		return string(r)
	case map[string]interface{}:
		r, err := json.Marshal(arg)
		if err != nil {
			return ""
		}
		return string(r)
	default:
		return fmt.Sprintf("%v", arg)
	}
}

// JSON creates a properly formatted JSON
func (e *Error) JSON() interface{} {
	p := Params{}
	if e.meta != nil {
		switch arg := e.meta.(type) {
		case Params:
			p = arg
		case map[string]interface{}:
			p = arg
		default:
			p["meta"] = e.Meta()
		}
	}
	if _, ok := p["error"]; !ok {
		p["error"] = e.Err.Error()
	}
	if _, ok := p["type"]; !ok && e.Type != TypeNone {
		p["type"] = e.Type.String()
	}
	if _, ok := p["stack"]; !ok {
		p["stack"] = e.stack
	}
	return p
}

// MarshalJSON implements the json.Marshaller interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.JSON())
}

func E(args ...interface{}) *Error {
	e := &Error{Type: TypeNone}
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
		default:
			e.meta = arg
		}
	}
	if len(message) > 0 {
		msg := strings.Join(message[:], ": ")
		e.Err = errors.New(msg)
	}

	for i := 1; i <= 5; i++ {
		_, fn, line, ok := runtime.Caller(i)
		if ok {
			e.stack = append(e.stack, fmt.Sprintf("%s:%d", fn, line))
		}
	}
	return e
}
