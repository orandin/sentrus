package sentrushttp

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/orandin/sentrus"
)

func CaptureLog(entry *logrus.Entry, defaultHub *sentry.Hub, tags map[string]string) error {
	hubToUse := defaultHub
	if hub := sentry.GetHubFromContext(entry.Context); hub != nil {
		hubToUse = hub
	}

	return sentrus.DefaultCaptureLog(entry, hubToUse, tags)
}
