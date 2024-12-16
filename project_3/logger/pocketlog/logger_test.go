package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, 1000)
	debugLogger.Debugf("Hello, %s!", "world")
	// Output: [DEBUG] Hello, world!
}

func ExampleLogger_Infof()() {
	 debugLogger := pocketlog.New(pocketlog.LevelInfo, 6)
	 debugLogger.Infof("Hello, %s!", "world")
	// Output: [INFO] Hello,
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
			expected: "[DEBUG] " + debugMsg + "\n" + "[INFO] " + infoMsg + "\n" + "[ERROR] " + errorMsg + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] " + infoMsg + "\n" + "[ERROR] " + errorMsg + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] " + errorMsg + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, 1000, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMsg)
			// OR
			//testedLogger.Logf(pocketlog.LevelDebug, debugMsg)

			//testedLogger.Infof(infoMsg)
			// OR
			testedLogger.Logf(pocketlog.LevelInfo, infoMsg)

			//testedLogger.Errorf(errorMsg)
			// OR
			testedLogger.Logf(pocketlog.LevelError, errorMsg)

			// TODO test loggerLength

			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q got %q", tc.expected, tw.contents)
			}

		})
	}
}

func TestLoggerLength_DebugfInfofErrorf(t *testing.T) {
	type TestCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]TestCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: "[DEBUG] this is a debug\n[INFO] this is an info\n[ERROR] this is an erro\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: "[INFO] this is an info\n[ERROR] this is an erro\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: "[ERROR] this is an erro\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, 15, pocketlog.WithOutput(tw))

			testedLogger.Debugf(debugMsg)
			testedLogger.Logf(pocketlog.LevelInfo, infoMsg)
			testedLogger.Logf(pocketlog.LevelError, errorMsg)
			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q got %q", tc.expected, tw.contents)
			}

		})
	}
}
