package windows

import "C"
import (
	"fmt"
	"syscall"
	"unsafe"
)

func (c *Creator) CallMiniDump(processId uint32, path string) error {
	dumpDll, err := syscall.LoadDLL(c.config.DllPath)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	miniDump, err := dumpDll.FindProc(c.config.FuncName)
	if err != nil {
		c.logError.Print(err)
		return err
	}

	retCode, _, err := miniDump.Call(uintptr(processId), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))))
	if retCode != 1 {
		fmt.Println("return code: ", retCode)
		fmt.Println("syscall error: ", syscall.GetLastError())
		c.logError.Print(err)
		return err
	}
	return nil
}
