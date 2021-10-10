package configs

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
	LogInfo     string
	LogError    string
}

// getEnv Получение переменной окружения
func (c *Creator) getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// getEnvFile Чтение файла переменных окружения
func (c *Creator) getEnvFile() error {
	if err := godotenv.Load(c.path); err != nil {
		err = errors.New("unable to read config file")
		return err
	}
	return nil
}

// GetConfig Получение конфигурации
func (c *Creator) GetConfig() (*Config, error) {
	err := c.getEnvFile()
	return &Config{
		DllPath:     c.getEnv("DLL_PATH", ""),
		FilesDir:    c.getEnv("DUMP_DIRECTORY", ""),
		NameProcess: c.getEnv("PROCESS_NAME", ""),
		FuncName:    c.getEnv("FUNCTION_NAME", ""),
		LogInfo:     c.getEnv("LOG_INFO", "stdout"),
		LogError:    c.getEnv("LOG_ERROR", "stderr"),
	}, err
}
