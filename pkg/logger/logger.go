package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"os"
	"strings"
)

type Params map[string]interface{}

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	err := errors.InitSentry()
	if err != nil {
		Error(err)
	}
}

type message struct {
	message string
	params  map[string]interface{}
}

func (msg *message) String() string {
	if len(msg.params) > 0 {
		return fmt.Sprintf("%s - %v", msg.message, msg.params)
	}
	return fmt.Sprintf("%s", msg.message)
}

func Info(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Info with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.params).Info(msg.message)
}

func Debug(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Debug with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.params).Debug(msg.message)
}

func Warn(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Warn with no arguments")
	}
	msg := getMessage(args...)
	log.WithFields(msg.params).Warn(msg.message)
}

func getMessage(args ...interface{}) *message {
	msg := &message{params: make(Params), message: ""}
	var generic []string
	var message []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case error:
			continue
		case string:
			message = append(message, arg)
		case Params:
			appendMap(msg.params, arg)
		case map[string]interface{}:
			appendMap(msg.params, arg)
		default:
			generic = append(generic, fmt.Sprintf("%s", arg))
		}
	}
	if len(message) > 0 {
		msg.message = strings.Join(message[:], ": ")
	}
	if len(generic) > 0 {
		msg.params["objects"] = strings.Join(generic[:], " | ")
	}
	return msg
}

func appendMap(root map[string]interface{}, tmp map[string]interface{}) {
	for k, v := range tmp {
		root[k] = v
	}
}
