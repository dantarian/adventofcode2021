package smoke

type point2d struct {
	x, y int
}

func BasinSize(heights [][]int, startPoint Point) int {
	maxY := len(heights) - 1
	maxX := len(heights[0]) - 1
	candidates := func(point point2d) []point2d {
		res := []point2d{}
		if point.x > 0 {
			res = append(res, point2d{point.x - 1, point.y})
		}

		if point.x < maxX {
			res = append(res, point2d{point.x + 1, point.y})
		}

		if point.y > 0 {
			res = append(res, point2d{point.x, point.y - 1})
		}

		if point.y < maxY {
			res = append(res, point2d{point.x, point.y + 1})
		}
		return res
	}

	startPoint2d := point2d{startPoint.X, startPoint.Y}

	includedPoints := map[point2d]struct{}{startPoint2d: {}}
	excludedPoints := map[point2d]struct{}{}
	pointsToCheck := map[point2d]struct{}{startPoint2d: {}}

	for len(pointsToCheck) > 0 {
		for point := range pointsToCheck {
			for _, candidatePoint := range candidates(point) {
				if _, pres := includedPoints[candidatePoint]; pres {
					continue
				}

				if _, pres := excludedPoints[candidatePoint]; pres {
					continue
				}

				if heights[candidatePoint.y][candidatePoint.x] == 9 {
					excludedPoints[candidatePoint] = struct{}{}
					continue
				}

				if heights[candidatePoint.y][candidatePoint.x] <= heights[point.y][point.x] {
					// Assume no nested basins.
					continue
				}

				includedPoints[candidatePoint] = struct{}{}

				if _, pres := pointsToCheck[candidatePoint]; !pres {
					pointsToCheck[candidatePoint] = struct{}{}
				}
			}
			delete(pointsToCheck, point)
		}

	}

	return len(includedPoints)
}
