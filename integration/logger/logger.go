package logger

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/integration/config"
	"os"
	"time"
)

func getFormatter() (log.Formatter, error) {
	switch config.Configuration.Log.Formatter {
	case "text":
		return &log.TextFormatter{}, nil
	case "json":
		return &log.JSONFormatter{}, nil
	default:
		return nil, errors.New("Wrong formatter, valid ones: 'text' or 'json'")
	}
}

func InitLogger() error {
	formatter, err := getFormatter()
	if err != nil {
		return err
	}
	log.SetFormatter(formatter)

	if config.Configuration.Log.File_Path != "" {
		f, err := os.OpenFile(config.Configuration.Log.File_Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Errorf("Failed to open file %s (going forward without writing to file), error: %s", config.Configuration.Log.File_Path, err)
		} else {
			log.SetOutput(f)
			log.Infof("%s opened, log will be written to it.", config.Configuration.Log.File_Path)
		}
	}
	return nil
}

func TimeTrack(name, method, url string, t time.Duration) {
	log.WithFields(log.Fields{
		"nano":    t.Nanoseconds(),
		"seconds": t.Seconds(),
		"url":     url,
		"method":  method,
	}).Infof("%s time track", name)
}

func NewRequest(name, url string) {
	log.WithFields(log.Fields{
		"url": url,
	}).Infof("%s new request", name)
}

func GetFiles(count int) {
	log.WithFields(log.Fields{
		"count": count,
	}).Info("new files founded")
}

func GetCoins(count int) {
	log.WithFields(log.Fields{
		"count": count,
	}).Info("coins founded")
}

func FileTest(file string) {
	log.WithFields(log.Fields{
		"file": file,
	}).Info("starting test from file")
}

func Test(path, method string, code int) {
	log.WithFields(log.Fields{
		"path":   path,
		"method": method,
		"code":   code,
	}).Info("starting test")
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Info(args ...interface{}) {
	log.WithFields(log.Fields{}).Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.WithFields(log.Fields{}).Infof(format, args...)
}
