// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spidernet-io/rocktemplate/pkg/debug"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupUtility() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			rootLogger.Sugar().Warnf("got signal=%+v \n", s)
		}
	}()

	// run gops
	d := debug.New(rootLogger)
	if globalConfig.GopsPort != 0 {
		d.RunGops(int(globalConfig.GopsPort))
	}

	if globalConfig.PyroscopeServerAddress != "" {
		d.RunPyroscope(globalConfig.PyroscopeServerAddress, globalConfig.PodName)
	}
}

func DaemonMain() {

	SetupUtility()

	SetupHttpServer()

	RunMetricsServer("agent")

	rootLogger.Info("hello world")
	time.Sleep(time.Hour)
}
