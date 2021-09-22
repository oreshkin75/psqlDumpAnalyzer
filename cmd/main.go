package main

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/reader"
	"PostgreDumpAnalyzer/modules/win32"
	"encoding/hex"
	"fmt"
)

func main() {

	/*http.Handle("/", http.FileServer(http.Dir("frontend/build/web")))
	http.ListenAndServe(":8080", nil)*/

	config, err := env.New("C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\cmd\\.env")
	if err != nil {
		panic(err)
	}

	err = win32.CreateDump(config)
	if err != nil {
		panic(err)
	}

	for {
		data, _, err := reader.FileRead("C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\files\\5940.dmp")
		if err != nil {
			panic(err)
		}
		if data == nil {
			break
		}
		fmt.Println(hex.EncodeToString(data))
	}
}
