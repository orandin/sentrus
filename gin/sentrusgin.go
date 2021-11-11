package sentrusgin

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/orandin/sentrus"
)

func CaptureLog(entry *logrus.Entry, defaultHub *sentry.Hub, tags map[string]string) error {
	hubToUse := defaultHub

	if ctx, ok := entry.Context.(*gin.Context); ok {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hubToUse = hub
		}
	}

	return sentrus.DefaultCaptureLog(entry, hubToUse, tags)
}
