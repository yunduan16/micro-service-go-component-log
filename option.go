package log

import (
    "github.com/sirupsen/logrus"
    "time"
)

// Option option
type Option func(*Options)

var defaultOptions = Options{
    timestampFormat: "2006-01-02 15:04:05",
    logLevel:        logrus.TraceLevel,
    rotationTime:    5*time.Hour,
    maxAge:          5*time.Hour,
    callerDeep:      1,
}

// Options option
type Options struct {
    timestampFormat string
    logLevel        logrus.Level
    fileName        string
    rotationTime    time.Duration
    maxAge          time.Duration
    callerDeep      int
}

// TimestampFormat option
func TimestampFormat(timestampFormat string) Option {
    return func(opts *Options) {
        opts.timestampFormat = timestampFormat
    }
}

// LogLevel option
func LogLevel(logLevel logrus.Level) Option {
    return func(opts *Options) {
        opts.logLevel = logLevel
    }
}

// FileName option
func FileName(fileName string) Option {
    return func(opts *Options) {
        opts.fileName = fileName
    }
}

// RotationTime option
func RotationTime(rotationTime time.Duration) Option {
    return func(opts *Options) {
        opts.rotationTime = rotationTime
    }
}

// MaxAge option
func MaxAge(maxAge time.Duration) Option {
    return func(opts *Options) {
        opts.maxAge = maxAge
    }
}

// CallerDeep option
func CallerDeep(callerDeep int) Option {
    return func(opts *Options) {
        opts.callerDeep = callerDeep
    }
}