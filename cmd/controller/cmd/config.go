// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spidernet-io/rocktemplate/pkg/logger"
	"reflect"
	"strconv"
)

type Config struct {
	EnableMetric           bool
	MetricPort             int32
	HealthPort             int32
	GopsPort               int32
	WebhookPort            int32
	PyroscopeServerAddress string
	ConfigMapPath          string
	TlsCaCertPath          string
	TlsServerCertPath      string
	TlsServerKeyPath       string
}

var globalConfig Config

type _envMapping struct {
	envName      string
	defaultValue string
	p            interface{}
}

var envMapping = []_envMapping{
	{"ENV_ENABLED_METRIC", "false", &globalConfig.EnableMetric},
	{"ENV_METRIC_HTTP_PORT", "", &globalConfig.MetricPort},
	{"ENV_HEALTH_PORT", "", &globalConfig.HealthPort},
	{"ENV_GOPS_LISTEN_PORT", "", &globalConfig.GopsPort},
	{"ENV_WEBHOOK_PORT", "", &globalConfig.WebhookPort},
	{"ENV_PYROSCOPE_PUSH_SERVER_ADDRESS", "", &globalConfig.PyroscopeServerAddress},
}

func init() {

	viper.AutomaticEnv()
	if t := viper.GetString("ENV_LOG_LEVEL"); len(t) > 0 {
		rootLogger = logger.NewStdoutLogger(t)
	} else {
		rootLogger = logger.NewStdoutLogger("")
	}

	logger := rootLogger.Named("config")
	// env built in the image
	if t := viper.GetString("ENV_VERSION"); len(t) > 0 {
		logger.Info("app version " + t)
	}
	if t := viper.GetString("ENV_GIT_COMMIT_VERSION"); len(t) > 0 {
		logger.Info("git commit version " + t)
	}
	if t := viper.GetString("ENV_GIT_COMMIT_TIMESTAMP"); len(t) > 0 {
		logger.Info("git commit timestamp " + t)
	}

	for _, v := range envMapping {
		m := v.defaultValue
		if t := viper.GetString(v.envName); len(t) > 0 {
			m = t
		}
		if len(m) > 0 {
			switch v.p.(type) {
			case *int32:
				if s, err := strconv.ParseInt(m, 10, 64); err == nil {
					v.p = int32(s)
				} else {
					logger.Fatal("failed to parse env value of " + v.envName + " to int32, value=" + m)
				}
			case *string:
				v.p = m
			case *bool:
				if s, err := strconv.ParseBool(m); err == nil {
					v.p = s
				} else {
					logger.Fatal("failed to parse env value of " + v.envName + " to bool, value=" + m)
				}
			default:
				logger.Sugar().Fatal("unsupported type to parse %v, config type=%v ", v.envName, reflect.TypeOf(v.p))
			}
		}

		logger.Info(v.envName + " = " + m)
	}

	// command flags
	globalFlag := rootCmd.PersistentFlags()
	globalFlag.StringVarP(&globalConfig.ConfigMapPath, "config-path", "C", "", "configmap file path")
	globalFlag.StringVarP(&globalConfig.TlsCaCertPath, "tls-ca-cert", "R", "", "ca file path")
	globalFlag.StringVarP(&globalConfig.TlsServerCertPath, "tls-server-cert", "T", "", "server cert file path")
	globalFlag.StringVarP(&globalConfig.TlsServerKeyPath, "tls-server-key", "Y", "", "server key file path")
	if e := viper.BindPFlags(globalFlag); e != nil {
		logger.Sugar().Fatalf("failed to BindPFlags, reason=%v", e)
	}
	printFlag := func() {
		logger.Info("config-path = " + globalConfig.ConfigMapPath)
		logger.Info("tls-ca-cert = " + globalConfig.TlsCaCertPath)
		logger.Info("tls-server-cert = " + globalConfig.TlsServerCertPath)
		logger.Info("tls-server-key = " + globalConfig.TlsServerKeyPath)
	}
	cobra.OnInitialize(printFlag)

}
