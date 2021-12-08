package ssd

import (
	"errors"
	"strings"
)

func CountSimple(data []string) (int, error) {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return 0, errors.New("unexpected line format")
		}

		digits := strings.Fields(parts[1])
		for _, digit := range digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				total++
			}
		}
	}

	return total, nil
}