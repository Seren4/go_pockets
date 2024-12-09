package pocketlog

// Level represents an available logging level
type Level byte


const (
	// LevelDebug represents the lowest level, mostly used for debugging
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable
	LevelInfo
	// LevelWarn represents a logging level that contains information deemed valuable and warnings
	LevelWarn
	// LevelError represents the highest level, only to be used to trace error
	LevelError
)