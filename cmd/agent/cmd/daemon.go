package cmd

import "time"

func DaemonMain() {
	rootLogger.Info("hello world")
	time.Sleep(time.Hour)
}
