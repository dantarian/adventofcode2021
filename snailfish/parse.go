package snailfish

import (
	"errors"
	"strconv"
)

func ParsePair(data string) (Element, error) {
	result, _, err := parse([]rune(data))

	if err != nil {
		return nil, err
	}
	return result, nil
}

func parse(chars []rune) (Element, int, error) {
	if len(chars) == 0 {
		return nil, 0, errors.New("out of characters")
	}

	if chars[0] == '[' {
		left, leftLength, err := parse(chars[1:])
		if err != nil {
			return nil, 0, err
		}

		right, rightLength, err := parse(chars[leftLength+2:])
		if err != nil {
			return nil, 0, err
		}

		return &pair{left, right}, leftLength + rightLength + 3, nil
	}

	val, err := strconv.Atoi(string(chars[0]))
	if err != nil {
		return nil, 0, err
	}

	return &value{val}, 1, nil
}
