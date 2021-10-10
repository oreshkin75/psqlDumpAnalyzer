package main

import (
	"PostgreDumpAnalyzer/internal/configs"
	"PostgreDumpAnalyzer/internal/logging"
	platform2 "PostgreDumpAnalyzer/internal/platform"
	"errors"
	"fmt"
	"os"
)

func main() {
	envPath, err := cmdArgs()
	if err != nil {
		panic(err)
	}
	env := configs.New(envPath)
	config, err := env.GetConfig()
	logger := logging.New(config)
	logInfo, err := logger.CreateLoggerInfo()
	if err != nil {
		fmt.Println("can't open log file. Info logs will be write in stdout")
	}
	logError, err := logger.CreateLoggerError()
	if err != nil {
		fmt.Println("can't open log file. Error logs will be write in stdout")
	}
	platform := platform2.New(logInfo, logError, config)
	files, processes, err := platform.CreateDump()
	if err != nil {
		logError.Panic(err)
	}
	logInfo.Print("Find processes ", processes)
	logInfo.Print("Create dump files ", files)

}

func cmdArgs() (string, error) {
	args := os.Args
	if len(args) < 2 {
		err := errors.New("the first command line argument must be the path to the environment variable file")
		return "", err
	}

	return args[1], nil
}
