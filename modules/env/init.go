package env

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DllPath     string
	FilesDir    string
	NameProcess string
	FuncName    string
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvFile(path string) error {
	if err := godotenv.Load(path); err != nil {
		err = errors.New("unable to read config file")
		return err
	}
	return nil
}

func New(path string) (*Config, error) {
	err := getEnvFile(path)
	return &Config{
		DllPath:     getEnv("DLL_PATH", ""),
		FilesDir:    getEnv("DUMP_DIRECTORY", ""),
		NameProcess: getEnv("PROCESS_NAME", ""),
		FuncName:    getEnv("FUNCTION_NAME", ""),
	}, err
}
