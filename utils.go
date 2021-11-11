package sentrus

import (
	"github.com/sirupsen/logrus"
)

func GetErrorFromEntry(entry *logrus.Entry) (error, bool) {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	return err, ok
}
