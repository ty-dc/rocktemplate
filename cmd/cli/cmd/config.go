package cmd

import (
	"github.com/spf13/viper"
	"github.com/spidernet-io/template/pkg/logger"
)

func init() {

	viper.AutomaticEnv()
	if t := viper.GetString("ENV_LOG_LEVEL"); len(t) > 0 {
		rootLogger = logger.NewStdoutLogger(t)
	} else {
		rootLogger = logger.NewStdoutLogger("")
	}

	logger := rootLogger.Named("config")
	if t := viper.GetString("ENV_VERSION"); len(t) > 0 {
		logger.Info("version " + t)
	}
	if t := viper.GetString("ENV_GIT_COMMIT_VERSION"); len(t) > 0 {
		logger.Info("git commit version " + t)
	}
	if t := viper.GetString("ENV_GIT_COMMIT_TIMESTAMP"); len(t) > 0 {
		logger.Info("git commit timestamp " + t)
	}

}
