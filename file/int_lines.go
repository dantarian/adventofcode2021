package file

import (
	"os"
	"strconv"
	"strings"
)

func IntLines(filename string) ([]int, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(fileBytes)), "\n")
	values := make([]int, len(lines))
	for i, val := range lines {
		if values[i], err = strconv.Atoi(val); err != nil {
			return nil, err
		}
	}
	return values, nil
}
