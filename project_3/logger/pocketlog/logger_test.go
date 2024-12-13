package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s!", "world")
	// Output: Hello, world!
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

const (
	debugMsg = "this is a debug message"
	infoMsg  = "this is an info message"
	errorMsg = "this is an error message"
)

func TestLogger_DebugfInfofErrorf(t *testing.T) {
	type TestCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]TestCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugMsg + "\n" + infoMsg + "\n" + errorMsg + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoMsg + "\n" + errorMsg + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: errorMsg + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testedLogger.Debugf(debugMsg)
			testedLogger.Infof(infoMsg)
			testedLogger.Errorf(errorMsg)

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q got %q", tc.expected, tw.contents)
			}

		})
	}
}
