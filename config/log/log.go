package log

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

// Init initializes the logger.
func Init() {
	once.Do(func() {
		logger = logrus.New()

		// Output to stdout instead of the default stderr
		logger.SetOutput(os.Stdout)

		// JSON format logs, suitable for production environments
		logger.SetFormatter(&logrus.JSONFormatter{})

		// Log level set to Info by default; configurable if needed
		logger.SetLevel(logrus.InfoLevel)
	})
}

// Info logs an info message.
func Info(args ...interface{}) {
	ensureInitialized()
	logger.Info(args...)
}

// Infof logs an info message with formatting.
func Infof(format string, args ...interface{}) {
	ensureInitialized()
	logger.Infof(format, args...)
}

// Error logs an error message.
func Error(args ...interface{}) {
	ensureInitialized()
	logger.Error(args...)
}

// Errorf logs an error message with formatting.
func Errorf(format string, args ...interface{}) {
	ensureInitialized()
	logger.Errorf(format, args...)
}

// Fatal logs a fatal message and exits.
func Fatal(args ...interface{}) {
	ensureInitialized()
	logger.Fatal(args...)
}

// Fatalf logs a fatal message with formatting and exits.
func Fatalf(format string, args ...interface{}) {
	ensureInitialized()
	logger.Fatalf(format, args...)
}

// Debug logs a debug message, typically useful for development.
func Debug(args ...interface{}) {
	ensureInitialized()
	logger.Debug(args...)
}

// Debugf logs a debug message with formatting.
func Debugf(format string, args ...interface{}) {
	ensureInitialized()
	logger.Debugf(format, args...)
}

// ensureInitialized ensures the logger is initialized.
func ensureInitialized() {
	if logger == nil {
		Init()
	}
}
