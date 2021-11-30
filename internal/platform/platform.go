package platform

import (
	"PostgreDumpAnalyzer/internal/configs"
	"PostgreDumpAnalyzer/internal/platform/windows"
	"log"
)

type ProcessSearcher interface {
	FindProcessByName(name string) ([]windows.Process, error)
	CallMiniDump(processId uint32, path string) error
}

type Creator struct {
	logInfo         *log.Logger
	logError        *log.Logger
	config          *configs.Config
	processes       []windows.Process
	ProcessSearcher ProcessSearcher
}

func New(logInfo, logError *log.Logger, config *configs.Config) *Creator {
	return &Creator{
		logInfo:         logInfo,
		logError:        logError,
		config:          config,
		processes:       nil,
		ProcessSearcher: windows.New(logInfo, logError, config),
	}
}
