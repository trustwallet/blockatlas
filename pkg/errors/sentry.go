package errors

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"time"
)

func InitSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              viper.GetString("sentry.dsn"),
		AttachStacktrace: true,
	})
	if err != nil {
		logger.Error("Set Sentry DSN error", err)
	}
	return err
}

func SendError(err error) {
	sentry.CaptureException(err)
}

func SendFatal(err error) {
	sentry.CaptureException(err)
	if sentry.Flush(time.Second * 5) {
		logger.Info("All sentry queued events delivered!")
	} else {
		logger.Info("Sentry flush timeout reached")
	}
}

func SendMessage(msg string) {
	sentry.CaptureMessage(msg)
}
