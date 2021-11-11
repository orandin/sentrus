package fasthttp

import (
	"github.com/getsentry/sentry-go"
	sentryfasthttp "github.com/getsentry/sentry-go/fasthttp"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"github.com/orandin/sentrus"
)

func CaptureLog(entry *logrus.Entry, defaultHub *sentry.Hub, tags map[string]string) error {
	hubToUse := defaultHub

	if ctx, ok := entry.Context.(*fasthttp.RequestCtx); ok {
		if hub := sentryfasthttp.GetHubFromContext(ctx); hub != nil {
			hubToUse = hub
		}
	}

	return sentrus.DefaultCaptureLog(entry, hubToUse, tags)
}
