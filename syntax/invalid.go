package syntax

import (
	"errors"
	"fmt"
)

// Returns 0 if the string is valid or incomplete, or a score if an invalid
// character is encountered.
func Invalid(source string) (int, error) {
	var brackets stack

	for _, character := range source {
		switch character {
		case '(', '[', '{', '<':
			brackets.push(character)
			continue
		case ')', ']', '}', '>':
			stackBracket, exists := brackets.pop()
			if !exists {
				return 0, errors.New("closing bracket with no opening bracket")
			}

			if match(stackBracket, character) {
				continue
			}

			return score(character)
		}
		return 0, fmt.Errorf("unexpected character: %v", string(character))
	}

	return 0, nil
}

func match(opening rune, closing rune) bool {
	if opening == '(' && closing == ')' {
		return true
	}

	if opening == '[' && closing == ']' {
		return true
	}

	if opening == '{' && closing == '}' {
		return true
	}

	if opening == '<' && closing == '>' {
		return true
	}

	return false
}

func score(r rune) (int, error) {
	switch r {
	case ')':
		return 3, nil
	case ']':
		return 57, nil
	case '}':
		return 1197, nil
	case '>':
		return 25137, nil
	}

	return 0, fmt.Errorf("unexpected character: %v", string(r))
}
