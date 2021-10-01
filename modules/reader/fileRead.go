package reader

import (
	"io"
	"os"
)

func FileRead(path string) ([]byte, int, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, 0, err
	}

	data := make([]byte, 128)
	n, err := file.Read(data)
	if err == io.EOF {
		return nil, n, nil
	} else if err != nil {
		return nil, n, err
	}

	return data, n, nil

	/*data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	os.Stdout.Write(data)
	//fmt.Printf("%X\n", data)
	encodedStr := hex.EncodeToString(data)

	return []byte(encodedStr), nil
	//fmt.Printf("%s\n", encodedStr)
	//fmt.Println(data)*/
}

func BytesRead(bytes uint, file *os.File) ([]byte, int, error) {
	data := make([]byte, bytes)
	n, err := file.Read(data)
	if err == io.EOF {
		return nil, n, nil
	} else if err != nil {
		return nil, n, err
	}
	return data, n, nil
}
