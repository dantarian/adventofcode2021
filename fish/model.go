package fish

import "fmt"

type cohort struct {
	count, timeToSpawning int
}

func (c *cohort) String() string {
	return fmt.Sprintf("%v", *c)
}

func Model(days int) []int {
	fishesPerDay := []int{1}
	cohorts := []*cohort{{1, 0}}
	fishes := 1

	for i := 0; i < days; i++ {
		newSpawns := 0
		for _, cohort := range cohorts {
			cohort.timeToSpawning--
			if cohort.timeToSpawning < 0 {
				cohort.timeToSpawning = 6
				newSpawns += cohort.count
			}
		}

		if newSpawns > 0 {
			cohorts = append(cohorts, &cohort{newSpawns, 8})
			fishes += newSpawns
		}

		fishesPerDay = append(fishesPerDay, fishes)
	}

	response := []int{}
	responseLength := 7
	if days < 7 {
		responseLength = days
	}
	for i := 0; i < responseLength; i++ {
		response = append(response, fishesPerDay[len(fishesPerDay)-(i+1)])
	}

	return response
}
