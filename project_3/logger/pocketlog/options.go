package pocketlog

import "io"

// Option defines a functional option to tour logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs.
// (this type of ft can be passed to the New() ft as variadic parameters: a list of zero or more arguments of the same type)
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
	
}
