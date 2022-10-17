// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"github.com/google/gops/agent"
	"github.com/pyroscope-io/client/pyroscope"
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
	if globalConfig.GopsPort != 0 {
		address := fmt.Sprintf("127.0.0.1:%d", globalConfig.GopsPort)
		op := agent.Options{
			ShutdownCleanup: true,
			Addr:            address,
		}
		if err := agent.Listen(op); err != nil {
			rootLogger.Sugar().Fatalf("gops failed to listen on port %s, reason=%v", address, err)
		}
		rootLogger.Sugar().Infof("gops is listening on %s ", address)
		defer agent.Close()
	}

	if globalConfig.PyroscopeServerAddress != "" {
		// push mode ,  push to pyroscope server
		rootLogger.Sugar().Infof("pyroscope works in push mode, server %s ", globalConfig.PyroscopeServerAddress)
		node, e := os.Hostname()
		if e != nil || len(node) == 0 {
			rootLogger.Sugar().Fatalf("failed to get hostname, reason=%v", e)
		}
		_, e = pyroscope.Start(pyroscope.Config{
			ApplicationName: BinName,
			ServerAddress:   globalConfig.PyroscopeServerAddress,
			Logger:          pyroscope.StandardLogger,
			Tags:            map[string]string{"node": node},
			ProfileTypes: []pyroscope.ProfileType{
				pyroscope.ProfileCPU,
				pyroscope.ProfileAllocObjects,
				pyroscope.ProfileAllocSpace,
				pyroscope.ProfileInuseObjects,
				pyroscope.ProfileInuseSpace,
			},
		})
		if e != nil {
			rootLogger.Sugar().Fatalf("failed to setup pyroscope, reason=%v", e)
		}
	}
}

func DaemonMain() {

	rootLogger.Sugar().Infof("config: %+v", globalConfig)

	SetupUtility()

	SetupHttpServer()

	rootLogger.Info("hello world")
	time.Sleep(time.Hour)
}
