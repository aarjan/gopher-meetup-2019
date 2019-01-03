package log

import (
	"io"
	"log"
)

type Option func(*Logger) *Logger

// Logger defines a logger to log the data
type Logger struct {
	log *log.Logger
}

func NewLogger(options ...Option) *Logger {
	l := &Logger{log: &log.Logger{}}
	for _, option := range options {
		option(l)
	}
	return l
}

func (l Logger) Info(message string) {
	l.log.Printf(" [info] %s", message)
}

func SetPrefix(s string) Option {
	return func(l *Logger) *Logger {
		l.log.SetPrefix(s)
		return l
	}
}

func SetOutput(out io.Writer) Option {
	return func(l *Logger) *Logger {
		l.log.SetOutput(out)
		return l
	}
}
