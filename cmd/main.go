package main

import (
	"PostgreDumpAnalyzer/modules"
	"errors"
	"os"
)

func main() {
	envPath, err := cmdArgs()
	if err != nil {
		panic(err)
	}
	mods, err := modules.New(envPath)
	if err != nil {
		panic(err)
	}
	defer mods.Loggers.LogFileError.Close()
	defer mods.Loggers.LogFileInfo.Close()

	_, err = mods.DumpCreator.CreateDump()
	if err != nil {
		mods.Loggers.LogError.Print(err)
	}

	mods.Loggers.LogInfo.Print("--- the program ended successfully ")
	/*data, err := ioutil.ReadFile(files[0])
	if err != nil {
		logger.LogError.Fatal(err)
		panic(err)
	}
	file, err := os.Create("dump.txt")
	if err != nil {
		logger.LogError.Fatal(err)
		panic(err)
	}
	var cleanData []byte
	for _, ch := range data {
		if ch == 0 {
			continue
		}
		cleanData = append(cleanData, ch)
	}
	dataStr := string(cleanData)
	r := []rune(dataStr)
	var cleanUni []rune
	for _, ch := range r {
		if unicode.IsGraphic(ch) {
			if ch == 0xFFFD {
				ch = 0x000A
			}
			cleanUni = append(cleanUni, ch)
		}
	}
	file.WriteString(string(cleanUni))
	/*var clearData []byte
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
	file.WriteString(clean)*/

}

func cmdArgs() (string, error) {
	args := os.Args
	if len(args) < 2 {
		err := errors.New("the first command line argument must be the path to the environment variable file")
		return "", err
	}

	return args[1], nil
}
