package sentry

import (
	"time"

	sentry "github.com/getsentry/sentry-go"
)

func InitSentry(sentryDsn string) error {
	if sentryDsn == "" {
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDsn,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		return err
	}
	defer sentry.Flush(2 * time.Second)

	return nil
}
