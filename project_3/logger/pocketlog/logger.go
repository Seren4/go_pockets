package pocketlog

// Logger is used to log information
type Logger struct{
	threshold Level
}

// New returns you a logger; ready to log at hte required threshold
func New(threshold Level) *Logger{
	return &Logger{
		threshold: threshold,
	}
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any){
}

// Infof formats and prints a message if the log level is info or higher
func (l *Logger) Infof(format string, args ...any){
}

// Errorf formats and prints a message if the log level is error
func (l *Logger) Errorf(format string, args ...any){
}