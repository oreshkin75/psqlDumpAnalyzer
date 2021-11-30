package windows

import (
	"golang.org/x/sys/windows"
	"strings"
	"syscall"
	"unsafe"
)

func (c *Creator) FindProcessByName(name string) ([]Process, error) {
	handle, err := windows.CreateToolhelp32Snapshot(Th32csSnapprocess, 0)
	if err != nil {
		c.logError.Print(err)
		return nil, err
	}
	defer windows.CloseHandle(handle)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	// get the first process
	err = windows.Process32First(handle, &entry)
	if err != nil {
		c.logError.Print(err)
		return nil, err
	}

	results := make([]Process, 0, 50)
	for {
		results = append(results, newWindowsProcess(&entry))

		err = windows.Process32Next(handle, &entry)
		if err != nil {
			// windows sends ERROR_NO_MORE_FILES on last process
			if err == syscall.ERROR_NO_MORE_FILES {
				break
			}
			c.logError.Print(err)
			return nil, err
		}
	}

	var findProcesses []Process
	for _, p := range results {
		if strings.ToLower(p.Exe) == strings.ToLower(name) {
			findProcesses = append(findProcesses, p)
		}
	}
	return findProcesses, nil
}

func newWindowsProcess(e *windows.ProcessEntry32) Process {
	// Find when the string ends for decoding
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return Process{
		ProcessID:       e.ProcessID,
		ParentProcessID: e.ParentProcessID,
		Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
	}
}
