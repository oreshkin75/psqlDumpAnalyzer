package logging

import (
	"log"
	"os"
)

/*func (c *LogCreator) checkConfig() error {
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
}*/

func (c *Creator) CreateLoggerInfo() (*log.Logger, error) {
	var err error
	if c.config.LogInfo == "stdout" {
		c.logInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	} else if c.config.LogInfo == "stderr" {
		c.logInfo = log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)
	} else {
		c.logFileInfo, err = os.OpenFile(c.config.LogInfo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			c.logInfo = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
			return c.logInfo, err
		}
		c.logInfo = log.New(c.logFileInfo, "INFO\t", log.Ldate|log.Ltime)
	}
	return c.logInfo, nil
}

func (c *Creator) CreateLoggerError() (*log.Logger, error) {
	var err error
	if c.config.LogError == "stdout" {
		c.logError = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	} else if c.config.LogError == "stderr" {
		c.logError = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		c.logFileError, err = os.OpenFile(c.config.LogError, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			c.logError = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
			return c.logError, err
		}
		c.logError = log.New(c.logFileError, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	}

	return c.logError, nil
}
