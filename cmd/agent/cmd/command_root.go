// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime/debug"
)

var BinName = filepath.Base(os.Args[0])
var rootLogger *zap.Logger

// rootCmd represents the base command.
var rootCmd = &cobra.Command{
	Use:   BinName,
	Short: "short description",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if e := recover(); nil != e {
				rootLogger.Sugar().Errorf("Panic details: %v", e)
				debug.PrintStack()
				os.Exit(1)
			}
		}()
		DaemonMain()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootLogger.Fatal(err.Error())
	}
}
