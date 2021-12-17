package octopus

import "strconv"

func NewSwarm(initialPowers []string) (*Swarm, error) {
	swarm := make(Swarm)
	for j, line := range initialPowers {
		for i, value := range line {
			location := Point{i, j}
			power, err := strconv.Atoi(string(value))
			if err != nil {
				return nil, err
			}

			swarm[location] = &Octopus{location, power}
		}
	}
	return &swarm, nil
}
