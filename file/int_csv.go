package file

import (
	"os"
	"strconv"
	"strings"
)

func IntCSV(filename string) ([]int, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	strValues := strings.Split(strings.TrimSpace(string(fileBytes)), ",")
	values := make([]int, len(strValues))
	for i, val := range strValues {
		if values[i], err = strconv.Atoi(val); err != nil {
			return nil, err
		}
	}
	return values, nil
}
