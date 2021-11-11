# sentrusfasthttp

## Example

```go
import (
	"github.com/orandin/sentrus"
	sentrusfasthttp "github.com/orandin/sentrus/fasthttp"
)

// Set up Sentry here

// Add hook
logrus.AddHook(sentrus.NewHook(
    []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},
    sentrus.WithCustomCaptureLog(sentrusfasthttp.CaptureLog),
))

sentryHandler := sentryfasthttp.New(sentryfasthttp.Options{
    Repanic: true,
    WaitForDelivery: true,
})

enhanceSentryEvent := func (handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if hub := sentry.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(ctx)
	}
}

defaultHandler := enhanceSentryEvent(func (ctx *fasthttp.RequestCtx) {
	logHandler := logrus.WithContext(ctx)
	logHandler.Warn("Ceci est un test")

	ctx.SetStatusCode(fasthttp.StatusOK)
})

fooHandler := enhanceSentryEvent(func (ctx *fasthttp.RequestCtx) {
	panic("y tho")
})

fastHTTPHandler := func (ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/foo":
		fooHandler(ctx)
	default:
		defaultHandler(ctx)
	}
}

fmt.Println("Listening and serving HTTP on :3000")

if err := fasthttp.ListenAndServe(":3000", sentryHandler.Handle(fastHTTPHandler)); err != nil {
    panic(err)
}
```

