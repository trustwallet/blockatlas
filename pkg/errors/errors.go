package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Params map[string]interface{}

// Error represents a error's specification.
type Error struct {
	Err  error
	Type Type
	meta interface{}
}

var (
	_ error = (*Error)(nil)
)

func (e *Error) isEmpty() bool {
	return e.meta == nil && e.Type == TypeNone && e.Err == nil
}

func (e *Error) Error() string {
	r, _ := e.MarshalJSON()
	return string(r)
}

func (e *Error) String() string {
	msg := e.Err.Error()
	if e.Type != TypeNone {
		msg = fmt.Sprintf("%s | Code: %s", msg, e.Type.String())
	}
	if len(e.Meta()) > 0 {
		msg = fmt.Sprintf("%s | Meta: %s", msg, e.Meta())
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
		r, _ := json.Marshal(arg)
		return string(r)
	case map[string]interface{}:
		r, _ := json.Marshal(arg)
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
	if _, ok := p["type"]; !ok {
		p["type"] = e.Type.String()
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
	return e
}
