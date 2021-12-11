package syntax

type stack []rune

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() (rune, bool) {
	if s.isEmpty() {
		return ' ', false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
