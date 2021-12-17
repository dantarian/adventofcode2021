package origami

import (
	"strings"
)

type Point struct {
	X, Y int
}

type PointSet map[Point]struct{}

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

type FoldSpec struct {
	direction Direction
	foldPoint int
}

func (ps *PointSet) String() string {
	maxX, maxY := 0, 0
	for point := range *ps {
		if point.X > maxX {
			maxX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}

	var sb strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			char := ' '
			if _, present := (*ps)[Point{x, y}]; present {
				char = '#'
			}
			sb.WriteRune(char)
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
