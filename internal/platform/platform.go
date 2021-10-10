package platform

import (
	"PostgreDumpAnalyzer/internal/configs"
	"log"
)

type Creator struct {
	logInfo  *log.Logger
	logError *log.Logger
	config   *configs.Config
}

func New(logInfo, logError *log.Logger, config *configs.Config) *Creator {
	return &Creator{
		logInfo:  logInfo,
		logError: logError,
		config:   config,
	}
}
