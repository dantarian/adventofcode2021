package octopus

func (s *Swarm) Neighbours(o *Octopus) []*Octopus {
	neighbours := []*Octopus{}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			neighbour, present := (*s)[Point{o.Location.X + dx, o.Location.Y + dy}]
			if !present {
				continue
			}

			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}
