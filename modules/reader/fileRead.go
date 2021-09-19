package reader

import (
	"encoding/hex"
	"io/ioutil"
)

func FileRead(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	encodedStr := hex.EncodeToString(data)

	return []byte(encodedStr), nil
	//fmt.Printf("%s\n", encodedStr)
	//fmt.Println(data)
}
