package octopus

func (s *Swarm) TotalFlashes(steps int) int {
	total := 0

	for i := 0; i < steps; i++ {
		total += s.step()
	}

	return total
}

func (s *Swarm) SyncedFlashStep() int {
	var i int
	for i = 0; s.step() < len(*s); i++ {
		continue
	}

	return i + 1
}

func (s *Swarm) step() int {
	flashed := make(map[Point]*Octopus)
	newlyFlashed := make(map[Point]*Octopus)

	// Increment powers and mark initial flashes
	for _, octopus := range *s {
		octopus.Power++
		if octopus.Power <= 9 {
			continue
		}
		flashed[octopus.Location] = octopus
		newlyFlashed[octopus.Location] = octopus
	}

	// Handle cascades
	for len(newlyFlashed) > 0 {
		lastFlashed := newlyFlashed
		newlyFlashed = make(map[Point]*Octopus)
		for _, octopus := range lastFlashed {
			for _, neighbour := range s.Neighbours(octopus) {
				neighbour.Power++
				_, alreadyFlashed := flashed[neighbour.Location]
				if alreadyFlashed || neighbour.Power <= 9 {
					continue
				}
				flashed[neighbour.Location] = neighbour
				newlyFlashed[neighbour.Location] = neighbour
			}
		}
	}

	// Reset flashed octopuses to 0
	for _, octopus := range flashed {
		octopus.Power = 0
	}

	return len(flashed)
}
