package file

import (
	"os"
	"strconv"
	"strings"
)

func Int64CSV(filename string) ([]int64, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	strValues := strings.Split(strings.TrimSpace(string(fileBytes)), ",")
	values := make([]int64, len(strValues))
	for i, val := range strValues {
		if values[i], err = strconv.ParseInt(val, 10, 64); err != nil {
			return nil, err
		}
	}
	return values, nil
}
