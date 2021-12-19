package chitons

import "strconv"

func NewGrid(data []string) (*Grid, error) {
	grid := make(Grid)

	for y, line := range data {
		for x, r := range line {
			risk, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			grid[Point{x, y}] = risk
		}
	}

	return &grid, nil
}
