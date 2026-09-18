package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var defaultLogger *logrus.Logger

func init() {
	defaultLogger = logrus.New()
	defaultLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	defaultLogger.SetOutput(os.Stdout)
	defaultLogger.SetLevel(logrus.InfoLevel)
}

type Field struct {
	Key   string
	Value any
}

func Err(err error) Field {
	return Field{Key: "error", Value: err}
}

func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Any(key string, value any) Field {
	return Field{Key: key, Value: value}
}

func toLogrusFields(fields []Field) logrus.Fields {
	f := make(logrus.Fields, len(fields))
	for _, field := range fields {
		if err, ok := field.Value.(error); ok && err != nil {
			f[field.Key] = err.Error()
		} else {
			f[field.Key] = field.Value
		}
	}
	return f
}

func Info(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WithFields(toLogrusFields(fields)).Info(msg)
}

func Error(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WithFields(toLogrusFields(fields)).Error(msg)
}

func Panic(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WithFields(toLogrusFields(fields)).Panic(msg)
}

func Debug(ctx context.Context, msg string, fields ...Field) {
	defaultLogger.WithFields(toLogrusFields(fields)).Debug(msg)
}

func Init(isLocal bool) {
	if isLocal {
		defaultLogger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
			DisableColors: true,
		})
	}
}
