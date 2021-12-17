package octopus

type Point struct {
	X, Y int
}

type Octopus struct {
	Location Point
	Power    int
}

type Swarm map[Point]*Octopus
