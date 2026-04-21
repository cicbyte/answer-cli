package utils

import (
	"fmt"
	"os"
)

func EnsureDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func InitAppDirs() error {
	config := ConfigInstance

	dirs := []string{
		config.GetAppSeriesDir(),
		config.GetAppDir(),
		config.GetConfigDir(),
		config.GetLogDir(),
	}

	for _, dir := range dirs {
		if err := EnsureDir(dir); err != nil {
			return fmt.Errorf("directory init failed: %v", err)
		}
	}

	return nil
}
