package logger

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type LoggerWrapper struct {
	logger *logrus.Logger
}

func GetLogLevel() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	logLevel := os.Getenv("LOGLEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	return logLevel
}

func GetLogger() *LoggerWrapper {
	logLevel := GetLogLevel()
	logger := logrus.New()

	logger.SetOutput(os.Stdout)

	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.Warn("Invalid log level, defaulting to info")
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	return &LoggerWrapper{logger: logger}
}

func (l *LoggerWrapper) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
func (l *LoggerWrapper) Debug(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}
func (l *LoggerWrapper) Warn(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}
func (l *LoggerWrapper) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}
func (l *LoggerWrapper) Fatal(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
	os.Exit(1)
}
