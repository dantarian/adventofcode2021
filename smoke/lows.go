package smoke

type Point struct {
	X, Y, Z int
}

func LowPoints(heights [][]int) []Point {
	rows := len(heights)
	cols := len(heights[0])
	points := []Point{}
	for j, line := range heights {
		for i, cell := range line {
			if i > 0 && cell >= line[i-1] {
				continue
			}

			if i < cols-1 && cell >= line[i+1] {
				continue
			}

			if j > 0 && cell >= heights[j-1][i] {
				continue
			}

			if j < rows-1 && cell >= heights[j+1][i] {
				continue
			}

			points = append(points, Point{i, j, cell})
		}
	}
	return points
}
