package main

import (
	"PostgreDumpAnalyzer/modules/process"
	"PostgreDumpAnalyzer/modules/win32"
)

func main() {
	/*data, err := ioutil.ReadFile("C:\\Users\\vkont\\CLionProjects\\test\\cmake-build-debug\\test_20210912_143007.dmp")
	  if err != nil {
	      panic(err)
	  }
	  encodedStr := hex.EncodeToString(data)

	  fmt.Printf("%s\n", encodedStr)
	  //fmt.Println(data)*/

	sys, err := win32.Loader("C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\cmd\\dumpCreator.dll")
	if err != nil {
		panic(err)
	}

	proc, err := win32.FindProc(sys, "DumpProcessImpl")
	if err != nil {
		panic(err)
	}

	id, err := process.ProcessID("pg_ctl.exe")
	if err != nil {
		panic(err)
	}

	err = win32.CallMiniDump(proc, id)
	if err != nil {
		panic(err)
	}
}
