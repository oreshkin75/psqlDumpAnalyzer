package main

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/win32"
)

func main() {

	config := env.New("C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\cmd\\.env")
	err := win32.CreateDump(config)
	if err != nil {
		panic(err)
	}
}
