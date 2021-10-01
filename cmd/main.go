package main

import (
	"PostgreDumpAnalyzer/modules/env"
	"PostgreDumpAnalyzer/modules/win32"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {

	/*http.Handle("/", http.FileServer(http.Dir("frontend/build/web")))
	http.ListenAndServe(":8080", nil)*/

	config, err := env.New("C:\\Users\\vkont\\GolandProjects\\PostgreDumpAnalyzer\\cmd\\.env")
	if err != nil {
		panic(err)
	}

	files, err := win32.CreateDump(config)
	if err != nil {
		panic(err)
	}

	fmt.Println(files)

	/*file, err := os.Open(files[0])
	defer file.Close()
	if err != nil {
		panic(err)
	}
	for {
		data, _, err := reader.BytesRead(128, file)
		if err != nil {
			panic(err)
		}
		if data == nil {
			break
		}
		//fmt.Println(hex.EncodeToString(data))
	}*/
	data, err := ioutil.ReadFile(files[0])
	if err != nil {
		panic(err)
	}
	file, err := os.Create("dump.txt")
	if err != nil {
		panic(err)
	}

	var clearData []byte
	for _, ch := range data {
		if ch == 0 {
			continue
		}
		clearData = append(clearData, ch)
	}
	defer file.Close()
	myString := string(clearData)
	clean := strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, myString)
	fmt.Printf("%q\n", clean)
	file.WriteString(clean)

}
