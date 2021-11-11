package sentrus

import (
	"strings"

	"github.com/getsentry/sentry-go"
)

type IsNotInAppFrameIntegration struct{}

func (sentrusIntegration *IsNotInAppFrameIntegration) Name() string {
	return "SentrusIsNotInAppFrame"
}

func (sentrusIntegration *IsNotInAppFrameIntegration) SetupOnce(client *sentry.Client) {
	client.AddEventProcessor(sentrusIntegration.processor)
}

func (sentrusIntegration *IsNotInAppFrameIntegration) processor(event *sentry.Event, _ *sentry.EventHint) *sentry.Event {

	for i := range event.Threads {
		exception := &event.Threads[i]
		stacktrace := exception.Stacktrace

		if stacktrace != nil {
			filterFrame(stacktrace.Frames)
		}
	}

	for i := range event.Exception {
		exception := &event.Exception[i]
		stacktrace := exception.Stacktrace

		if stacktrace != nil {
			filterFrame(stacktrace.Frames)
		}
	}

	return event
}

func filterFrame(frames []sentry.Frame) {
	for j := range frames {
		frame := &frames[j]
		if frame.Module == "github.com/sirupsen/logrus" || strings.HasPrefix(frame.Module, "github.com/orandin/sentrus") {
			frame.InApp = false
		}
	}
}
