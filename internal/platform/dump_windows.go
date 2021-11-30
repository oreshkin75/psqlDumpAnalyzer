//go:build windows
// +build windows

package platform

import (
	"PostgreDumpAnalyzer/internal/platform/windows"
	"errors"
	"fmt"
)

func (d *Creator) CreateDump() ([]string, []windows.Process, error) {
	var err error

	if d.processes == nil {
		err = errors.New("There is no such process ")
		d.logError.Print(err, d.config.NameProcess)
		return nil, nil, err
	}

	var files []string
	var fileName string
	for _, p := range d.processes {
		fileName = d.config.FilesDir + fmt.Sprint(p.ProcessID) + ".dmp"
		err = d.ProcessSearcher.CallMiniDump(p.ProcessID, fileName)
		files = append(files, fileName)
		if err != nil {
			return files, d.processes, err
		}
	}

	var processes []windows.Process
	processes = append(processes, d.processes...)
	d.processes = nil
	return files, processes, nil
}

func (d *Creator) FindAllPsqlProcesses() ([]windows.Process, error) {
	var err error
	d.processes, err = d.ProcessSearcher.FindProcessByName(d.config.NameProcess)
	if err != nil {
		return nil, err
	}
	if d.processes == nil {
		err = errors.New("There are no such process ")
		d.logError.Print(err, d.config.NameProcess)
		return nil, err
	}

	return d.processes, nil
}

func (d *Creator) FindQueryProcess() ([]windows.Process, error) {
	newProcesses, err := d.ProcessSearcher.FindProcessByName(d.config.NameProcess)
	if err != nil {
		return nil, err
	}
	if newProcesses == nil {
		err = errors.New("There are no such process. Maybe your close psql server ")
		d.logError.Print(err, d.config.NameProcess)
		return nil, err
	}
	var resultProcesses []windows.Process
	for i, np := range newProcesses {
		for _, p := range d.processes {
			if p.ProcessID == np.ProcessID {
				newProcesses[i] = windows.Process{Exe: "delete"}
			}
		}
	}
	for _, p := range newProcesses {
		if p.Exe != "delete" {
			resultProcesses = append(resultProcesses, p)
		}
	}
	d.processes = nil
	d.processes = append(d.processes, resultProcesses...)
	return d.processes, nil
}
