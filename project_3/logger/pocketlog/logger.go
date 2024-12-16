package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
	loggerLength int
}

// New returns you a logger ready to log at required threshold and required message length.
// The default output is Stdout.
func New(threshold Level, loggerLength int, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout, loggerLength: loggerLength}

	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}

// logf prints the message to the output.
func (l *Logger) logf(lvl Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	if l.loggerLength < len(message) {
        message = message[:l.loggerLength]
    }

	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)

}

// Logf formats and prints a mesh if the log level is hugh enough
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}
	l.logf(lvl, format, args...)
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.Logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.Logf(LevelInfo, format, args...)
}

// Infof formats and prints a message if the log level is warn or higher
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	l.Logf(LevelWarn, format, args...)
}

// Errorf formats and prints a message if the log level is error
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.Logf(LevelError, format, args...)
}
