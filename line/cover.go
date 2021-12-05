package line

func (line *Line) CoveredPoints() []Point {
	xDirection := direction(line.Start.X, line.End.X)
	yDirection := direction(line.Start.Y, line.End.Y)

	xLength := abs(line.Start.X-line.End.X) + 1
	yLength := abs(line.Start.Y-line.End.Y) + 1
	numPoints := max(xLength, yLength)

	points := []Point{}
	for i := 0; i < numPoints; i++ {
		points = append(points, Point{line.Start.X + (i * xDirection), line.Start.Y + (i * yDirection)})
	}

	return points
}

func direction(start int, end int) int {
	if start > end {
		return -1
	}

	if start < end {
		return 1
	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
