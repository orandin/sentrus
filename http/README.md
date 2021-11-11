# sentrushttp

## Example

```go
import (
    "github.com/orandin/sentrus"
    sentrushttp "github.com/orandin/sentrus/http"
)

// Set up Sentry here

// Add hook
logrus.AddHook(sentrus.NewHook(
    []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},
    sentrus.WithCustomCaptureLog(sentrushttp.CaptureLog),
))

enhanceSentryEvent := func(handler http.HandlerFunc) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
            hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
        }
        handler(rw, r)
	}
}

http.HandleFunc("/", sentryHandler.HandleFunc(
    enhanceSentryEvent(func(rw http.ResponseWriter, r *http.Request) {
        logHandler := logrus.WithContext(r.Context())
        logHandler.Warn("Ceci est un test")

        rw.WriteHeader(http.StatusOK)
    }),
))

http.HandleFunc("/foo", sentryHandler.HandleFunc(
    enhanceSentryEvent(func(rw http.ResponseWriter, r *http.Request) {
        panic("y tho")
    }),
))

fmt.Println("Listening and serving HTTP on :3000")

if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
}
```

