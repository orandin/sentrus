package sentrus

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

type CaptureLog func(entry *logrus.Entry, defaultHub *sentry.Hub, tags map[string]string) error

type Hook struct {
	hub        *sentry.Hub
	tags       map[string]string
	levels     []logrus.Level
	captureLog CaptureLog
}

type Option func(h *Hook)

func NewHook(levels []logrus.Level, options ...Option) Hook {
	hook := Hook{
		levels:     levels,
		hub:        sentry.CurrentHub(),
		captureLog: DefaultCaptureLog,
	}

	for _, option := range options {
		option(&hook)
	}

	if hook.hub.Client() != nil {
		// Override the in-app detection logic to exclude Logrus and Sentrus.
		// https://github.com/getsentry/sentry-go/issues/124
		// https://github.com/getsentry/sentry-go/issues/136
		sentrusIsNotInAppFrame := new(IsNotInAppFrameIntegration)
		sentrusIsNotInAppFrame.SetupOnce(hook.hub.Client())
	}

	return hook
}

// Options

func WithTags(tags map[string]string) Option {
	return func(h *Hook) {
		h.tags = tags
	}
}

func WithCustomHub(hub *sentry.Hub) Option {
	return func(h *Hook) {
		h.hub = hub
	}
}

func WithCustomCaptureLog(captureLog CaptureLog) Option {
	return func(h *Hook) {
		h.captureLog = captureLog
	}
}

// Implement interface logrus.Hook

func (h Hook) Fire(entry *logrus.Entry) error {
	return h.captureLog(entry, h.hub, h.tags)
}

func (h Hook) Levels() []logrus.Level {
	return h.levels
}

// ---------------

func DefaultCaptureLog(entry *logrus.Entry, hub *sentry.Hub, tags map[string]string) error {
	hub.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(LevelMap[entry.Level])
		scope.SetTags(tags)
		scope.SetExtras(entry.Data)

		if err, ok := GetErrorFromEntry(entry); ok {
			scope.SetExtra("log.message", entry.Message)
			hub.CaptureException(err)
		} else {
			hub.CaptureMessage(entry.Message)
		}
	})

	return nil
}
