package win32

import "C"
import (
	"fmt"
	"syscall"
)

func Loader(dllName string) (*syscall.DLL, error) {
	dumpDll, err := syscall.LoadDLL(dllName)
	if err != nil {
		return nil, err
	}

	return dumpDll, nil
}

func FindProc(dll *syscall.DLL, funcName string) (*syscall.Proc, error) {
	miniDump, err := dll.FindProc(funcName)
	if err != nil {
		return nil, err
	}

	return miniDump, err
}

func CallMiniDump(proc *syscall.Proc, processId uint32) error {
	retCode, _, err := proc.Call(uintptr(processId))
	if err != nil {
		fmt.Println("return code: ", retCode)
		fmt.Println("syscall error: ", syscall.GetLastError())
		return err
	}
	return nil
}
