package logging

import (
	"PostgreDumpAnalyzer/internal/configs"
	"log"
	"os"
)

type Creator struct {
	logInfo      *log.Logger
	logError     *log.Logger
	logFileInfo  *os.File
	logFileError *os.File
	config       *configs.Config
}

func New(config *configs.Config) *Creator {
	return &Creator{config: config}
}
