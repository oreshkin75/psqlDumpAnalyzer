package dumps

import (
	"os"
)

// OpenDumpFile Открывает файл с сохранение его дескриптора для дальнейшего чтения
func (f *Creator) OpenDumpFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	f.file = file
	return nil
}

// Read Читает дамп файл в байтовом представлении
// Проверка на io.EOF осуществляется пользователем
func (f *Creator) Read(bytesToRead int) ([]byte, int, error) {
	data := make([]byte, bytesToRead)
	n, err := f.file.Read(data)
	if err != nil {
		return nil, n, err
	}
	return data, n, nil
}

// CloseDumpFile Закрывает дескриптор файла
func (f *Creator) CloseDumpFile() {
	f.file.Close()
}
