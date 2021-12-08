package ssd

import (
	"errors"
	"math"
	"sort"
	"strings"
)

func DecodeAndSum(data []string) (int, error) {
	total := 0
	for _, line := range data {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			return 0, errors.New("unexpected line format")
		}

		translate := createTranslator(parts[0])
		value := translate(parts[1])
		total += value
	}

	return total, nil
}

func createTranslator(encoded string) func(string) int {
	digits := sortDigitCharacters(strings.Fields(encoded))

	var one, two, three, four, five, six, seven, eight, nine, zero string

	// Identify simple digits
	for _, digit := range digits {
		switch len(digit) {
		case 2:
			one = digit
		case 3:
			seven = digit
		case 4:
			four = digit
		case 7:
			eight = digit
		}
	}

  // Find 6, which is the only six-segment number that doesn't have one of the segments of 1; 0,
	// which of the other two is the only one missing a segment of 4; and 9, which is the other.
	for _, digit := range digits {
		if len(digit) != 6 {
			continue
		}
		
		if !(strings.ContainsRune(digit, []rune(one)[0]) && strings.ContainsRune(digit, []rune(one)[1])) {
			six = digit
			continue
		}

		notFound := []rune{}
		for _, letter := range four {
			if !strings.ContainsRune(digit, letter) {
				notFound = append(notFound, letter)
			}
		}

		if len(notFound) == 1 {
			zero = digit
			continue
		}

		nine = digit
	}

	// Find 3, which is the only five-segment number having both segments of 1; 5, which lacks only
	// one segment of 6; and 2, which is the other.
	for _, digit := range digits {
		if len(digit) != 5 {
			continue
		}

		if strings.ContainsRune(digit, []rune(one)[0]) && strings.ContainsRune(digit, []rune(one)[1]) {
			three = digit
			continue
		}

		notFound := []rune{}
		for _, letter := range six {
			if !strings.ContainsRune(digit, letter) {
				notFound = append(notFound, letter)
			}
		}

		if len(notFound) == 1 {
			five = digit
			continue
		}

		two = digit
	}

	decodedDigits := map[string]int{one: 1, two: 2, three: 3, four: 4, five: 5, six: 6, seven: 7, eight: 8, nine: 9, zero: 0}

	return func(str string) int {
		digits := sortDigitCharacters(strings.Fields(str))
		value := 0
		for i, digit := range digits {
			value += decodedDigits[digit] * int(math.Pow10(len(digits) - (1 + i)))
		}
		return value
	}

}

func sortDigitCharacters(digits []string) []string {
	sortedDigits := make([]string, len(digits))
	for i, digit := range digits {
		runes := []rune(digit)
		sort.Slice(runes, func(a, b int) bool {
			return runes[a] < runes[b]
		})
		sortedDigits[i] = string(runes)
	}
	return sortedDigits
}
