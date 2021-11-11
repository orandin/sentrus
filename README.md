# Sentrus

`Sentrus` provides a [Sentry](https://sentry.io) hook for the logrus package, using the latest Sentry SDK (`sentry-go`).
Extensible, it can be integrated with a variety of popular frameworks and packages in the Go ecosystem.

## Usage

```go
import(
    "github.com/getsentry/sentry-go"
    "github.com/sirupsen/logrus"
    "github.com/orandin/sentrus"
)

// Sentry init
if err := sentry.Init(sentry.ClientOptions{Dsn: "YOUR_DSN"}); err != nil {
    logger.WithError(err).Fatal("Cannot initiate Sentry")
}
defer sentry.Flush(2 * time.Second)

// Add Sentrus hook
loggrus.AddHook(sentrus.NewHook(
    []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel},

    // Optional: add tags to add in each Sentry event
    sentrus.WithTags(map[string]string{"foo": "bar"}),

    // Optional: set custom CaptureLog function
    sentrus.WithCustomCaptureLog(sentrus.DefaultCaptureLog),
))

logrus.Info("foo") // Not sent to Sentry
logrus.Warn("bar") // Sent to Sentry

// With an error
err := fmt.Errorf("awesome error")

logrus.WithError(err).Info("foo with an error") // Not sent to Sentry
logrus.WithError(err).Warn("bar with an error") // Sent to Sentry
```

### Integration with a framework

Like `Sentry`, `Sentrus` can be integrated with a variety of popular frameworks and packages in the Go ecosystem. For more detailed 
information, checkout the guides:

- [net/http](http/README.md)
- [fasthttp](fasthttp/README.md)
- [gin](gin/README.md)

#### An integration with a framework or package is missing?

`Sentrus` is **extensible**. You can develop custom `CaptureLog` functions to format and push the Sentry events. Don't 
hesitate to look the code of already existing integrations and feel free to contribute.

### Stack traces improvement

By default, `Sentrus` and `Logrus` will be considered by Sentry SDK as part of your application in stack traces. The
consequence is to have stack traces less readable from Sentry UI. 

At this time, `sentry-go` doesn't [make isInAppFrame customizable by clients](https://github.com/getsentry/sentry-go/issues/136). 
So, a Sentry integration (`sentrus.IsNotInAppFrameIntegration`) has been created as a workaround, and it is activated 
automatically when the hook is initialized. Each event will be modified to no longer consider `Sentrus` and `Logrus` as
part of your application.

## License

Licensed under [MIT License](https://opensource.org/licenses/MIT), see [`LICENSE`](LICENSE).
