package logger

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

// Logrus to level
var (
	// Log level
	logLevel = log.DebugLevel

	// Log level uint mapping with string
	levelToString = map[log.Level]string{
		log.DebugLevel: DebugLevel,
		log.InfoLevel:  InfoLevel,
		log.WarnLevel:  WarnLevel,
		log.ErrorLevel: ErrorLevel,
		log.PanicLevel: CriticalLevel,
	}
)

// SetLevel will log anything that is level or above (seq : debug, info, warn, error, panic).
func SetLevel(level string) {
	switch level {
	case InfoLevel:
		log.SetLevel(log.InfoLevel)
		logLevel = log.InfoLevel

	case WarnLevel:
		log.SetLevel(log.WarnLevel)
		logLevel = log.WarnLevel

	case ErrorLevel:
		log.SetLevel(log.ErrorLevel)
		logLevel = log.ErrorLevel

	case CriticalLevel:
		log.SetLevel(log.PanicLevel)
		logLevel = log.PanicLevel

	default:
		log.SetLevel(log.DebugLevel)
		logLevel = log.DebugLevel
	}

}

// GetLevel return level
func GetLevel() string {
	return levelToString[logLevel]
}

// log with interface
func logobj(lv log.Level, payload interface{}) {

	log.WithFields(
		log.Fields{
			"severity": strings.ToUpper(levelToString[lv]),
			"payload":  payload,
		},
	).Log(lv)

}

// logf with string
func logf(lv log.Level, format string, args ...interface{}) {

	log.WithFields(
		log.Fields{
			"severity": strings.ToUpper(levelToString[lv]),
			"payload":  fmt.Sprintf(format, args...),
		},
	).Log(lv)

}

// Debug logging with debug severity
func Debug(payload interface{}) {
	logobj(log.DebugLevel, payload)
}

// Debugf logging with debug severity
func Debugf(format string, args ...interface{}) {
	logf(log.DebugLevel, format, args...)
}

// Info logging with info severity
func Info(payload interface{}) {
	logobj(log.InfoLevel, payload)
}

// Infof logging with info severity
func Infof(format string, args ...interface{}) {
	logf(log.InfoLevel, format, args...)
}

// Warn logging with warning severity
func Warn(payload interface{}) {
	logobj(log.WarnLevel, payload)
}

// Warnf logging with warning severity
func Warnf(format string, args ...interface{}) {
	logf(log.WarnLevel, format, args...)
}

// Error logging with error severity
func Error(payload interface{}) {
	logobj(log.ErrorLevel, payload)
}

// Errorf logging with error severity
func Errorf(format string, args ...interface{}) {
	logf(log.ErrorLevel, format, args...)
}

// Critical logging with critical severity
func Critical(payload interface{}) {
	logobj(log.PanicLevel, payload)
}

// Criticalf logging with critical severity
func Criticalf(format string, args ...interface{}) {
	logf(log.PanicLevel, format, args...)
}
