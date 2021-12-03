package file

import (
	"os"
	"strings"
)

func StringLines(filename string) ([]string, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	return lines, nil
}
