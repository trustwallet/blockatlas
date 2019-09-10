package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Err struct {
	Message string
	Params  map[string]interface{}
	Err     error
}

var (
	_ error = (*Err)(nil)
)

func (e *Err) Error() string {
	msg := e.Message
	if e.Err != nil {
		msg = fmt.Sprintf("%s: %s", msg, e.Err)
	}
	if len(e.Params) > 0 {
		msg = fmt.Sprintf("%s - %s", msg, e.Params)
	}
	return msg
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Error with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.Params).Error(e.Message)
	SendError(e)
}

func Fatal(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Fatal with no arguments")
	}
	e := getError(args...)
	SendFatal(e)
	log.WithFields(e.Params).Fatal(e.Message)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Panic with no arguments")
	}
	e := getError(args...)
	SendFatal(e)
	log.WithFields(e.Params).Panic(e.Message)
}

func getError(args ...interface{}) *Err {
	e := &Err{Params: make(Params)}
	var message []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case string:
			message = append(message, arg)
		case error:
			e.Err = arg
		case Params:
			appendMap(e.Params, arg)
		case map[string]interface{}:
			appendMap(e.Params, arg)
		default:
			continue
		}
	}
	if len(message) > 0 {
		e.Message = strings.Join(message[:], ": ")
	}
	return e
}
