package cucumbers

import (
	"strings"
)

func (s *Swarm) MovesUntilNoMoves() int {
	moveMade := true
	i := 0
	for ; moveMade; i++ {
		movedEast := s.move(east)
		movedSouth := s.move(south)
		moveMade = movedEast || movedSouth
	}

	return i
}

func (s *Swarm) move(direction direction) bool {
	movingCucumbers := make(map[point]point)

	for point, cucumber := range s.cucumbers {
		if cucumber.direction != direction {
			continue
		}
		target := s.target(cucumber, point)
		if _, occupied := s.cucumbers[target]; !occupied {
			movingCucumbers[point] = target
		}
	}

	for point, target := range movingCucumbers {
		s.cucumbers[target] = s.cucumbers[point]
		delete(s.cucumbers, point)
	}

	return len(movingCucumbers) > 0
}

func (s *Swarm) target(c Cucumber, p point) point {
	switch c.direction {
	case south:
		return point{p.x, (p.y + 1) % s.height}
	case east:
		return point{(p.x + 1) % s.width, p.y}
	}
	return p
}

func (s *Swarm) String() string {
	var b strings.Builder
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			c, found := s.cucumbers[point{x, y}]
			if !found {
				b.WriteRune('.')
				continue
			}
			switch c.direction {
			case south:
				b.WriteRune('v')
			case east:
				b.WriteRune('>')
			}
		}

		b.WriteRune('\n')
	}
	return b.String()
}
