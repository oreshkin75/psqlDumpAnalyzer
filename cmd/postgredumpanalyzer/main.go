package main

import (
	"PostgreDumpAnalyzer/internal/configs"
	"PostgreDumpAnalyzer/internal/dumps"
	"PostgreDumpAnalyzer/internal/logging"
	"errors"
	"fmt"
	"io"
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

	/*psql := psql.New(logInfo, logError, config)
	_, err = psql.Connect()
	if err != nil {
		panic(err)
	}
	err = psql.Select("select * from person")
	if err != nil {
		panic(err)
	}

	platform := platform.New(logInfo, logError, config)
	files, processes, err := platform.CreateDump()
	if err != nil {
		logError.Panic(err)
	}
	logInfo.Print("Find processes ", processes)
	logInfo.Print("Create dump files ", files)*/

	dumps := dumps.New(logInfo, logError)
	err = dumps.OpenDumpFile("C:\\Users\\vkont\\GolandProjects\\psqlDumpAnalyzer\\assets\\dumps39176.dmp")
	if err != nil {
		panic(err)
	}
	defer dumps.CloseDumpFile()
	var allData []byte
	for {
		data, _, err := dumps.Read(1000)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		allData = append(allData, data...)
	}

	dataWithoutNulls := dumps.DeleteNulls(allData)
	clearData := dumps.DeleteUnprintableCharacters(dataWithoutNulls)
	fmt.Println(clearData)

	f, err := os.Create("cleardump.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(clearData)
	if err != nil {
		panic(err)
	}
}

func cmdArgs() (string, error) {
	args := os.Args
	if len(args) < 2 {
		err := errors.New("the first command line argument must be the path to the environment variable file")
		return "", err
	}

	return args[1], nil
}
