package logging

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	// Only log the warning severity or above.
	env := strings.ToUpper(os.Getenv("ENVIRONMENT"))

	if env == "PRODUCTION" {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.DebugLevel)
	} else {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}
}

// Trace logrus trace log
func Trace(args ...interface{}) {
	logger.Trace(args)
}

// Debug logrus debug log
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Info logrus info log
func Info(args ...interface{}) {
	logger.Info(args)
}

// Warn logrus warn log
func Warn(args ...interface{}) {
	logger.Warn(args)
}

// Error logrus error log
func Error(args ...interface{}) {
	logger.Error(args)
}

// Fatal logrus fatal log
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Panic logrus panic log
func Panic(args ...interface{}) {
	logger.Panic(args)
}
