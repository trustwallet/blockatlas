package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
	if len(e.Params) > 0 {
		return fmt.Sprintf("%s: %s\n%s", e.Message, e.Err, e.Params)
	}
	return fmt.Sprintf("%s: %s", e.Message, e.Err)
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Error with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.Params).Error(e)
	SendError(e)
}

func Fatal(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Fatal with no arguments")
	}
	e := getError(args...)
	SendFatal(e)
	log.WithFields(e.Params).Fatal(e)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Panic with no arguments")
	}
	e := getError(args...)
	SendFatal(e)
	log.WithFields(e.Params).Panic(e)
}

func getError(args ...interface{}) *Err {
	e := &Err{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case string:
			e.Message = arg
		case error:
			e.Err = arg
		case Params:
			e.Params = arg
		case map[string]interface{}:
			e.Params = arg
		default:
			continue
		}
	}
	return e
}
