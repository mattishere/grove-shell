package files

import (
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	return lines, nil
}

func ReadRCFile() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	path := filepath.Join(homeDir, ".groverc")

	_, err = os.Stat(path)
	if err != nil {
		return nil
	}

	data, err := ReadFile(path)
	if err != nil {
		return nil
	}

	return data
}
