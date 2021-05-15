package logger

import (
	"alfred.sh/common/logger"
	"go.uber.org/zap"
)

var (
	// Log TODO: Just a quick fix, now always import logger from "alfred.sh/common/logger"
	Log *zap.Logger = logger.Log
)

