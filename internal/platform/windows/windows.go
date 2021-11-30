package windows

import (
	"PostgreDumpAnalyzer/internal/configs"
	"log"
)

const Th32csSnapprocess = 0x00000002

type Process struct {
	ProcessID       uint32
	ParentProcessID uint32
	Exe             string
}

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
