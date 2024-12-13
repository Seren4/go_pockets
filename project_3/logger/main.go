package main

import (
	"learngo-pockets/logger/pocketlog"
	"os"
	"time"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))
	lgr.Debugf("Hello, %s!", "world debug")
	lgr.Infof("Hello, %s!", "world info")
	lgr.Errorf("Hello error, %d %v", 2024, time.Now())

}
