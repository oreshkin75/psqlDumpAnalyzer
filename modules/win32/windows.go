package win32

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/win32/dll"
	"PostgreDumpAnalyzer/modules/win32/process"
	"errors"
	"fmt"
)

func CreateDump(config *env.Config) ([]string, error) {
	sys, err := dll.Loader(config.DllPath)
	if err != nil {
		return nil, err
	}

	proc, err := dll.FindProc(sys, config.FuncName)
	if err != nil {
		return nil, err
	}

	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	ids := process.FindProcessByName(processes, config.NameProcess)

	if ids == nil {
		err = errors.New("There is no such process")
		return nil, err
	}

	var files []string
	var fileName string
	for _, p := range ids {
		fileName = config.FilesDir + fmt.Sprint(p.ProcessID) + ".dmp"
		err = dll.CallMiniDump(proc, p.ProcessID, fileName)
		files = append(files, fileName)
		if err != nil {
			return files, err
		}
	}
	fileName = config.FilesDir + fmt.Sprint(ids[0].ParentProcessID) + "_ParentProc" + ".dmp"
	err = dll.CallMiniDump(proc, ids[0].ParentProcessID, fileName)
	files = append(files, fileName)
	if err != nil {
		return files, err
	}

	return files, nil
}
