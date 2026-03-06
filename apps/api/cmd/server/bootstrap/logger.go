package bootstrap

import "mytodo/apps/api/pkg/logger"

type Logger = logger.Logger

func NewLogger() (Logger, error) {
	return logger.NewFromEnv()
}
