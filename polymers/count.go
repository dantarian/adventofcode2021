package polymers

func (im *InsertionMap) GetRuneCountsForPolymer(input string, steps int) RuneCountMap {
	runeCountResultsMap := make(RuneCountResultsMap)
	totalRuneCounts := make(RuneCountMap)
	runes := []rune(input)
	for i, r1 := range runes {
		if i == len(runes)-1 {
			totalRuneCounts[r1]++
			break
		}
		r2 := runes[i+1]

		runeCounts := im.getRuneCountForPair(r1, r2, steps, &runeCountResultsMap)
		totalRuneCounts.merge(&runeCounts)
		totalRuneCounts[r2]-- // Prevent double-counting
	}

	return totalRuneCounts
}

func (im *InsertionMap) getRuneCountForPair(r1, r2 rune, steps int, runeCountResultsMap *RuneCountResultsMap) RuneCountMap {
	// Do we already know the answer?
	result, found := (*runeCountResultsMap)[RuneCountAddress{r1, r2, steps}]
	if found {
		return result
	}

	// No, so calculate:
	pair := string([]rune{r1, r2})
	r3, found := (*im)[pair]
	if !found || steps == 0 {
		result := make(RuneCountMap)
		result[r1]++
		result[r2]++
		(*runeCountResultsMap)[RuneCountAddress{r1, r2, steps}] = result
		return result
	}

	result = make(RuneCountMap)
	rcm1 := im.getRuneCountForPair(r1, r3, steps-1, runeCountResultsMap)
	rcm2 := im.getRuneCountForPair(r3, r2, steps-1, runeCountResultsMap)
	result.merge(&rcm1)
	result.merge(&rcm2)
	result[r3]-- // Prevent double-counting
	(*runeCountResultsMap)[RuneCountAddress{r1, r2, steps}] = result
	return result
}

func (rcm *RuneCountMap) merge(other *RuneCountMap) {
	for r, count := range *other {
		(*rcm)[r] += count
	}
}
