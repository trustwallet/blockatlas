package errors

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"time"
)

func InitSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              viper.GetString("sentry.dsn"),
		AttachStacktrace: true,
	})
	if err != nil {
		return E(err, "InitSentry failed")
	}
	return nil
}

func SendError(err error) {
	sentry.CaptureException(err)
}

func SendFatal(err error) {
	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
}

func SendMessage(msg string) {
	sentry.CaptureMessage(msg)
}
