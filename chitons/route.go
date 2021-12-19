package chitons

import (
	"errors"
	"math"
)

func (g *Grid) CheapestRoute(start Point, end Point) (int, error) {
	unvisited := make(Grid)
	unvisitedNeighbours := make(Grid)
	for location := range *g {
		unvisited[location] = math.MaxInt32
	}

	current := start
	unvisited[current] = 0

	for current != end {
		for _, point := range unvisited.neighbours(current) {
			routeCost := unvisited[current] + (*g)[point]
			if routeCost > unvisited[point] {
				continue
			}
			unvisited[point] = routeCost
			unvisitedNeighbours[point] = routeCost
		}

		delete(unvisited, current)
		delete(unvisitedNeighbours, current)

		if len(unvisited) == 0 {
			return 0, errors.New("failed to reach end point")
		}

		var err error
		current, err = unvisitedNeighbours.minimum()
		if err != nil {
			return 0, err
		}
	}

	return unvisited[current], nil
}

func (g *Grid) neighbours(point Point) []Point {
	candidates := []Point{
		{point.X - 1, point.Y},
		{point.X + 1, point.Y},
		{point.X, point.Y - 1},
		{point.X, point.Y + 1},
	}

	neighbours := []Point{}
	for _, candidate := range candidates {
		if _, found := (*g)[candidate]; !found {
			continue
		}
		neighbours = append(neighbours, candidate)
	}

	return neighbours
}

func (g *Grid) minimum() (Point, error) {
	if len(*g) == 0 {
		return Point{0, 0}, errors.New("empty grid")
	}

	var minLocation Point
	minValue := math.MaxInt32
	for location, value := range *g {
		if value >= minValue {
			continue
		}
		minLocation = location
		minValue = value
	}

	return minLocation, nil
}
