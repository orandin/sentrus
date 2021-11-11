# sentrusgin

## Example

```go
import (
    "github.com/orandin/sentrus"
    sentrusgin "github.com/orandin/sentrus/gin"
)

// Set up Sentry here

// Add hook
logrus.AddHook(sentrus.NewHook(
    []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},
    sentrus.WithCustomCaptureLog(sentrusgin.CaptureLog),
))

app := gin.Default()

app.Use(func(ctx *gin.Context) {
    if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
        hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
    }
    ctx.Next()
})

app.GET("/", func(ctx *gin.Context) {
    logHandler := logrus.WithContext(ctx) // Inject gin.Context into logrus.Entry
    logHandler.Warn("It's a test")

    ctx.Status(http.StatusOK)
})

app.GET("/foo", func(ctx *gin.Context) {
    logHandler := logrus.WithContext(ctx)
    logHandler.WithError(fmt.Errorf("test error")).Warn("It's a test with error")

	ctx.Status(http.StatusOK)
})

app.Run(":3000")
```

