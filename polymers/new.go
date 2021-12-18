package polymers

import (
	"fmt"
	"strings"
)

func NewReplacementMap(data []string) (*ReplacementMap, error) {
	repMap := make(ReplacementMap)
	for _, line := range data {
		fields := strings.Split(line, " -> ")
		if len(fields) != 2 || len(fields[0]) != 2 || len(fields[1]) != 1 {
			return nil, fmt.Errorf("unexpected line format: %v", line)
		}

		pairRunes := []rune(fields[0])
		var sb strings.Builder
		sb.WriteRune(pairRunes[0])
		sb.WriteString(fields[1])
		sb.WriteRune(pairRunes[1])
		repMap[fields[0]] = sb.String()
	}

	return &repMap, nil
}

func NewInsertionMap(data []string) (*InsertionMap, error) {
	insMap := make(InsertionMap)
	for _, line := range data {
		fields := strings.Split(line, " -> ")
		if len(fields) != 2 || len(fields[0]) != 2 || len(fields[1]) != 1 {
			return nil, fmt.Errorf("unexpected line format: %v", line)
		}

		insMap[fields[0]] = []rune(fields[1])[0]
	}

	return &insMap, nil
}
