package main

import (
	"learngo-pockets/logger/pocketlog"
	"os"
	"time"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, 22, pocketlog.WithOutput(os.Stdout))
	lgr.Debugf("Hello, %s!", "world debug")
	lgr.Infof("Hello, %s!", "world info")
	lgr.Errorf("Hello error, %d %v", 2024, time.Now())
	lgr.Logf(pocketlog.LevelError, "This will be logged, %d %v", 2024, time.Now())
	lgr.Logf(pocketlog.LevelDebug, "This won't be logged, %d %v", 2024, time.Now())
	msg := "This will be truncated, skjhfljsknflkehfenedfhweubdsjkfbuid "
	lgr.Logf(pocketlog.LevelError, "%s %v",msg, time.Now())



}
