package win32

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/win32/dll"
	"PostgreDumpAnalyzer/modules/win32/process"
	"errors"
	"fmt"
)

func CreateDump(config *env.Config) error {
	sys, err := dll.Loader(config.DllPath)
	if err != nil {
		return err
	}

	proc, err := dll.FindProc(sys, config.FuncName)
	if err != nil {
		return err
	}

	processes, err := process.Processes()
	if err != nil {
		return err
	}

	ids := process.FindProcessByName(processes, config.NameProcess)

	if ids == nil {
		err = errors.New("There is no such process")
		return err
	}

	for _, p := range ids {
		err = dll.CallMiniDump(proc, p.ProcessID, config.FilesDir+fmt.Sprint(p.ProcessID)+".dmp")
		if err != nil {
			return err
		}
	}
	err = dll.CallMiniDump(proc, ids[0].ParentProcessID, config.FilesDir+fmt.Sprint(ids[0].ParentProcessID)+"_ParentProc"+".dmp")
	if err != nil {
		return err
	}

	return nil
}
