package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Params map[string]interface{}

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{})
	err := InitSentry()
	if err != nil {
		Error(err)
	}
}

type Msg struct {
	Message string
	Params  map[string]interface{}
}

func (msg *Msg) String() string {
	if len(msg.Params) > 0 {
		return fmt.Sprintf("%s - %v", msg.Message, msg.Params)
	}
	return fmt.Sprintf("%s", msg.Message)
}

func Info(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Info with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.Params).Info(msg.Message)
}

func Debug(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Debug with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.Params).Debug(msg.Message)
}

func Warn(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Warn with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.Params).Warn(msg.Message)
}

func getMessage(args ...interface{}) *Msg {
	msg := &Msg{}
	var generic []string
	var message []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case string:
			message = append(message, arg)
		case Params:
			msg.Params = arg
		case map[string]interface{}:
			msg.Params = arg
		default:
			generic = append(generic, fmt.Sprintf("%v", arg))
		}
	}
	if len(message) > 0 {
		msg.Message = strings.Join(message[:], ": ")
	}
	if len(generic) > 0 {
		msg.Params["objects"] = strings.Join(generic[:], " | ")
	}
	return msg
}
