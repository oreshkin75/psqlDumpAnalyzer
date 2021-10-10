package dumps

import (
	"log"
	"os"
)

type Creator struct {
	loggerInfo  *log.Logger
	loggerError *log.Logger
	file        *os.File
}

func New(loggerInfo, loggerError *log.Logger) *Creator {
	return &Creator{
		loggerInfo:  loggerInfo,
		loggerError: loggerError,
	}
}
