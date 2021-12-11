package syntax

func Complete(line string) (int, bool) {
	var brackets stack

	for _, character := range line {
		switch character {
		case '(', '[', '{', '<':
			brackets.push(character)
			continue
		case ')', ']', '}', '>':
			stackBracket, exists := brackets.pop()
			if !exists {
				return 0, false
			}

			if match(stackBracket, character) {
				continue
			}

			return 0, false
		}
		return 0, false
	}

	// Whatever's left we can score.
	score := 0
	for character, found := brackets.pop(); !(brackets.isEmpty() && !found); character, found = brackets.pop() {
		charVal := 0
		switch character {
		case '(':
			charVal = 1
		case '[':
			charVal = 2
		case '{':
			charVal = 3
		case '<':
			charVal = 4
		}
		score = (score * 5) + charVal
	}
	return score, true
}
