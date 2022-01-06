package cucumbers

type point struct {
	x, y int
}

type direction int

const (
	east direction = iota
	south
)

type Cucumber struct {
	direction direction
}

type Swarm struct {
	cucumbers     map[point]Cucumber
	height, width int
}

func NewSwarm(data []string) Swarm {
	swarm := Swarm{make(map[point]Cucumber), 0, 0}
	for y, line := range data {
		for x, char := range line {
			switch char {
			case 'v':
				swarm.cucumbers[point{x, y}] = Cucumber{south}
			case '>':
				swarm.cucumbers[point{x, y}] = Cucumber{east}
			}
			swarm.width = x + 1
		}
		swarm.height = y + 1
	}

	return swarm
}
