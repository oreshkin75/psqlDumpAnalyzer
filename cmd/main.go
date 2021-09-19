package main

import (
	"PostgreDumpAnalyzer/modules/process"
	"PostgreDumpAnalyzer/modules/win32"
	"errors"
	"fmt"
)

var (
	dumpFile    = "C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\files\\psql"
	dllFile     = "C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\cmd\\dumpCreator.dll"
	nameProcess = "postgres.exe"
)

func main() {
	sys, err := win32.Loader(dllFile)
	if err != nil {
		panic(err)
	}

	proc, err := win32.FindProc(sys, "DumpProcessImpl")
	if err != nil {
		panic(err)
	}

	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	ids := process.FindProcessByName(processes, nameProcess)
	fmt.Println(ids)

	if ids == nil {
		err = errors.New("There is no such process")
		panic(err)
	}

	for _, p := range ids {
		// TODO сделать нормальное форматирование строки
		err = win32.CallMiniDump(proc, p.ProcessID, dumpFile+fmt.Sprint(p.ProcessID)+".dmp")
		if err != nil {
			panic(err)
		}
	}
}
