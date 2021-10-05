package logs

import (
	"PostgreDumpAnalyzer/modules/env"
	"errors"
	"log"
	"os"
)

type LogCreator struct {
	config *env.Config
}

type Loggers struct {
	LogInfo      *log.Logger
	LogError     *log.Logger
	LogFileInfo  *os.File
	LogFileError *os.File
}

func New(config *env.Config) *LogCreator {
	return &LogCreator{config: config}
}

func (c *LogCreator) checkConfig() error {
	if c.config.LogInfo == c.config.LogError {
		if c.config.LogInfo == "stdout" && c.config.LogError == "stdout" {
			return nil
		} else if c.config.LogInfo == "stderr" && c.config.LogError == "stderr" {
			return nil
		}
		err := errors.New("error and information log files must be different ")
		return err
	}
	return nil
}

func (c *LogCreator) CreateLogger() (*Loggers, error) {
	var myLog Loggers
	var err error

	err = c.checkConfig()
	if err != nil {
		return nil, err
	}

	if c.config.LogInfo == "stdout" {
		myLog.LogInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	} else if c.config.LogInfo == "stderr" {
		myLog.LogInfo = log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)
	} else {
		myLog.LogFileInfo, err = os.OpenFile(c.config.LogInfo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		myLog.LogInfo = log.New(myLog.LogFileInfo, "INFO\t", log.Ldate|log.Ltime)
	}

	if c.config.LogError == "stdout" {
		myLog.LogError = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	} else if c.config.LogError == "stderr" {
		myLog.LogError = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		myLog.LogFileError, err = os.OpenFile(c.config.LogError, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		myLog.LogError = log.New(myLog.LogFileError, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	}

	return &myLog, nil
}
