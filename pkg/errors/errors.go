package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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

// SetType sets the error.
func (e *Error) SetError(err error) *Error {
	e.Err = err
	return e
}

// SetType sets the error from message.
func (e *Error) SetErrorFromString(err string) *Error {
	e.Err = errors.New(err)
	return e
}

// SetType sets the error's type.
func (e *Error) SetType(flags Type) *Error {
	e.Type = flags
	return e
}

// SetMeta sets the error's meta data.
func (e *Error) SetMeta(data interface{}) *Error {
	e.meta = data
	return e
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

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

// JSON creates a properly formatted JSON
func (msg *Error) JSON() interface{} {
	json := H{}
	if msg.meta != nil {
		value := reflect.ValueOf(msg.Meta)
		switch value.Kind() {
		case reflect.Struct:
			return msg.Meta
		case reflect.Map:
			for _, key := range value.MapKeys() {
				json[key.String()] = value.MapIndex(key).Interface()
			}
		default:
			json["meta"] = msg.Meta()
		}
	}
	if _, ok := json["error"]; !ok {
		json["error"] = msg.Error()
	}
	if _, ok := json["type"]; !ok {
		json["type"] = msg.Type
	}
	return json
}

// MarshalJSON implements the json.Marshaller interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.JSON())
}
