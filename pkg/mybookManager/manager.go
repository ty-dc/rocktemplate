package mybookManager

import (
	"github.com/spidernet-io/rocktemplate/pkg/mybookManager/types"
	"go.uber.org/zap"
)

type mybookManager struct {
	logger   *zap.Logger
	webhook  *webhookhander
	informer *informerHandler
}

func New(logger *zap.Logger) types.MybookManager {
	return &mybookManager{
		logger: logger.Named("mybookManager"),
	}
}
