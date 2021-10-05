// +build windows

package platform

import (
	"PostgreDumpAnalyzer/modules/platform/windows"
	"errors"
	"fmt"
)

func (d *DumpCreator) CreateDump() ([]string, error) {
	sys, err := windows.Loader(d.config.DllPath)
	if err != nil {
		return nil, err
	}

	proc, err := windows.FindProc(sys, d.config.FuncName)
	if err != nil {
		return nil, err
	}

	processes, err := windows.Processes()
	if err != nil {
		return nil, err
	}
	ids := windows.FindProcessByName(processes, d.config.NameProcess)

	if ids == nil {
		err = errors.New("There is no such process")
		return nil, err
	}
	d.logger.LogInfo.Print("find processes: ", ids)

	var files []string
	var fileName string
	for _, p := range ids {
		fileName = d.config.FilesDir + fmt.Sprint(p.ProcessID) + ".dmp"
		err = windows.CallMiniDump(proc, p.ProcessID, fileName)
		files = append(files, fileName)
		if err != nil {
			return files, err
		}
	}
	fileName = d.config.FilesDir + fmt.Sprint(ids[0].ParentProcessID) + "_ParentProc" + ".dmp"
	err = windows.CallMiniDump(proc, ids[0].ParentProcessID, fileName)
	files = append(files, fileName)
	if err != nil {
		return files, err
	}

	d.logger.LogInfo.Print("successfully create dump files ", files)
	return files, nil
}
