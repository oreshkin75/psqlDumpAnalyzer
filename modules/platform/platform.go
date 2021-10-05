package platform

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/logs"
)

type DumpCreator struct {
	logger *logs.Loggers
	config *env.Config
}

func New(logger *logs.Loggers, config *env.Config) *DumpCreator {
	return &DumpCreator{
		logger: logger,
		config: config,
	}
}
