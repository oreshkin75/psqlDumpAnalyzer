package psql

import (
	"PostgreDumpAnalyzer/internal/configs"
	"database/sql"
	"log"
)

type Creator struct {
	logInfo  *log.Logger
	logError *log.Logger
	config   *configs.Config
	db       *sql.DB
}

func New(logInfo, logError *log.Logger, config *configs.Config) *Creator {
	return &Creator{
		logInfo:  logInfo,
		logError: logError,
		config:   config,
	}
}
