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
        logger.SetOutput(os.Stdout)
        logger.SetFormatter(&logrus.JSONFormatter{})
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

// ensureInitialized ensures the logger is initialized.
func ensureInitialized() {
    if logger == nil {
        Init()
    }
}
