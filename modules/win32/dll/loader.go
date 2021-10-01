package dll

import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
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

func CallMiniDump(proc *syscall.Proc, processId uint32, path string) error {
	retCode, _, err := proc.Call(uintptr(processId), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))))
	if retCode != 1 {
		fmt.Println("return code: ", retCode)
		fmt.Println("syscall error: ", syscall.GetLastError())
		fmt.Println(err)
		return err
	}
	return nil
}
