package modules

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/logs"
	"PostgreDumpAnalyzer/modules/platform"
)

type DumpCreator interface {
	CreateDump() ([]string, error)
}

type Modules struct {
	Loggers     *logs.Loggers
	Config      *env.Config
	DumpCreator DumpCreator
}

func New(configPath string) (*Modules, error) {
	configCreate := env.New(configPath)
	config, err := configCreate.SetConfig()
	if err != nil {
		return nil, err
	}
	logCreate := logs.New(config)
	logger, err := logCreate.CreateLogger()
	if err != nil {
		return nil, err
	}

	return &Modules{
		Loggers:     logger,
		Config:      config,
		DumpCreator: platform.New(logger, config),
	}, nil
}
