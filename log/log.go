package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type log struct {
	logger *logrus.Logger
}

func Init() Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	return log{
		logger: logger,
	}
}

// TODO: переписать логгер

func (l log) Info(text string) {
	l.logger.Println("INFO: ", text)
}

func (l log) Debug(text string) {
	l.logger.Println("DEBUG: ", text)
}

func (l log) Error(text string) {
	l.logger.Println("ERROR: ", text)
}

func (l log) Fatal(text string) {
	l.logger.Fatal("FATAL: ", text)
}

func (l log) Infof(text string, args ...interface{}) {
	str := fmt.Sprintf(text, args...)
	l.Info(str)
}

func (l log) Debugf(text string, args ...interface{}) {
	str := fmt.Sprintf(text, args...)
	l.Debug(str)
}

func (l log) Errorf(text string, args ...interface{}) {
	str := fmt.Sprintf(text, args...)
	l.Error(str)
}

func (l log) Fatalf(text string, args ...interface{}) {
	str := fmt.Sprintf(text, args...)
	l.Fatal(str)
}
