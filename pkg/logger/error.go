package logger

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type errMessage struct {
	*message
	err error
}

func Error(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Error with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.params).Error(e.err)
}

func Fatal(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Fatal with no arguments")
	}
	e := getError(args...)
	errors.SendFatal(e.err)
	log.WithFields(e.params).Fatal(e.err)
}

func Panic(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Panic with no arguments")
	}
	e := getError(args...)
	errors.SendFatal(e.err)
	log.WithFields(e.params).Panic(e.err)
}

func getError(args ...interface{}) *errMessage {
	msg := getMessage(args...)
	err := &errMessage{message: msg}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *errors.Error:
			err.err = arg
		case error:
			err.err = errors.E(arg)
		case nil:
			continue
		default:
			continue
		}
	}
	if err.err == nil {
		err.err = errors.E(msg.message, msg.params)
	}
	return err
}
