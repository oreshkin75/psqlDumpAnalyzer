package ui

import (
	"PostgreDumpAnalyzer/internal/configs"
	"PostgreDumpAnalyzer/internal/dumps"
	"PostgreDumpAnalyzer/internal/platform"
	"PostgreDumpAnalyzer/internal/platform/windows"
	"PostgreDumpAnalyzer/internal/psql"
	"database/sql"
	"log"
)

type DumpReader interface {
	OpenDumpFile(path string) error
	Read(bytesToRead int) ([]byte, int, error)
	CloseDumpFile()
	DeleteNulls(data []byte) []byte
	DeleteUnprintableCharacters(data []byte) string
}

type PlatformWorker interface {
	CreateDump() ([]string, []windows.Process, error)
	FindAllPsqlProcesses() ([]windows.Process, error)
	FindQueryProcess() ([]windows.Process, error)
}

type PsqlWorker interface {
	Connect() (*sql.DB, error)
	Select(query string) ([]string, error)
	Insert(query string) error
	Update(query string) error
	Delete(query string) error
}

type Creator struct {
	logInfo        *log.Logger
	logError       *log.Logger
	config         *configs.Config
	columns        []string
	dumpFiles      []string
	psqlWorker     PsqlWorker
	platformWorker PlatformWorker
	dumpReader     DumpReader
}

func New(logInfo, logError *log.Logger, config *configs.Config) *Creator {
	return &Creator{
		logInfo:        logInfo,
		logError:       logError,
		config:         config,
		psqlWorker:     psql.New(logInfo, logError, config),
		platformWorker: platform.New(logInfo, logError, config),
		dumpReader:     dumps.New(logInfo, logError),
	}
}
