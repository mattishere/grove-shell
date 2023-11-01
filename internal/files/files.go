package files

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]string, error){
    file, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(file), "\n")

    return lines, nil
}
