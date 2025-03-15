package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

type Logger struct {
	baseLogger *slog.Logger
	context    context.Context
}

var (
	defaultLogger *Logger
)

// Init initializes both regular and traced loggers
func Init(ctx context.Context) {
	programLevel := new(slog.LevelVar)
	programLevel.Set(slog.LevelDebug)
	w := os.Stderr

	baseLogger := slog.New(tint.NewHandler(w, &tint.Options{
		Level:     programLevel,
		AddSource: true,
	}))

	defaultLogger = &Logger{
		baseLogger: baseLogger,
		context:    context.Background(),
	}

}

func (l *Logger) log(level slog.Level, msg string, args ...interface{}) {
	l.baseLogger.Log(l.context, level, msg, args...)
}

func Info(msg string, args ...interface{}) {
	defaultLogger.log(slog.LevelInfo, msg, args...)
}

func Debug(msg string, args ...interface{}) {
	defaultLogger.log(slog.LevelDebug, msg, args...)
}

func Error(msg string, args ...interface{}) {
	defaultLogger.log(slog.LevelError, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.log(slog.LevelWarn, msg, args...)
}
