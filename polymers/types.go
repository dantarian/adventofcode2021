package polymers

type ReplacementMap map[string]string

type InsertionMap map[string]rune

type RuneCountMap map[rune]int64

type RuneCountAddress struct {
	rune1, rune2 rune
	steps        int
}

type RuneCountResultsMap map[RuneCountAddress]RuneCountMap
