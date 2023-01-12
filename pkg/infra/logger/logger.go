package logger

import (
	"log"
	"os"

	"go.opentelemetry.io/otel/trace"
)

type Logger interface {
	// Infof(args ...any)
	// Warnf(args ...any)
	// Debugf(args ...any)
	// ... all the methods for a log
}

type logger struct {
	logger *log.Logger
	tracer trace.Tracer
}

type Option func(l *logger)

func WithTrace(tracer trace.Tracer) Option {
	return func(l *logger) {
		l.tracer = tracer
	}
}

func New(opts ...Option) Logger {
	l := logger{logger: log.New(os.Stdout, "", 0)}

	for _, opt := range opts {
		opt(&l)
	}

	return l
}
