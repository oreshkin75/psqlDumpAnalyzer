//go:build windows
// +build windows

package platform

import (
	windows2 "PostgreDumpAnalyzer/internal/platform/windows"
	"errors"
	"fmt"
)

func (d *Creator) CreateDump() ([]string, []windows2.Process, error) {
	sys, err := windows2.Loader(d.config.DllPath)
	if err != nil {
		return nil, nil, err
	}

	proc, err := windows2.FindProc(sys, d.config.FuncName)
	if err != nil {
		return nil, nil, err
	}

	processes, err := windows2.Processes()
	if err != nil {
		return nil, nil, err
	}
	ids := windows2.FindProcessByName(processes, d.config.NameProcess)

	if ids == nil {
		err = errors.New("There is no such process")
		return nil, nil, err
	}

	var files []string
	var fileName string
	for _, p := range ids {
		fileName = d.config.FilesDir + fmt.Sprint(p.ProcessID) + ".dmp"
		err = windows2.CallMiniDump(proc, p.ProcessID, fileName)
		files = append(files, fileName)
		if err != nil {
			return files, ids, err
		}
	}

	return files, ids, nil
}
