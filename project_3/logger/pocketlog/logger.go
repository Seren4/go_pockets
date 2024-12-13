package pocketlog

import "fmt"

// Logger is used to log information
type Logger struct {
	threshold Level
}

// New returns you a logger ready to log at required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof formats and prints a message if the log level is warn or higher
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

// Errorf formats and prints a message if the log level is error
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}
