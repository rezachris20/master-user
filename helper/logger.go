package helper

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// TOPIC for setting topic of log
	TOPIC = "master-service-log"
	// LogTag default log tag
	LogTag = "master-service"
)

func LogContext(c string, s string) *log.Entry {
	return log.WithFields(log.Fields{
		"##time":  time.Now().Format(time.RFC3339Nano),
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

func Log(level log.Level, message string, context string, scope string) {

	entry := LogContext(context, scope)
	switch level {
	case log.DebugLevel:
		entry.Debug(message)
	case log.InfoLevel:
		entry.Info(message)
	case log.WarnLevel:
		entry.Warn(message)
	case log.ErrorLevel:
		entry.Error(message)
	case log.FatalLevel:
		entry.Fatal(message)
	case log.PanicLevel:
		entry.Panic(message)
	}
}
