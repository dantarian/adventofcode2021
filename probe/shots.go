package probe

import (
	"math"
)

type stepMap map[int][]int
type pair struct {
	first, second int
}

func MaxHeight(minY int) int {
	return triangle(-minY - 1)
}

// Assumptions:
// minX > 0
// maxY < 0
// minX < maxX
// minY < maxY
func Target(minX, minY, maxX, maxY int) int {
	validShots := make(map[pair]struct{})
	minDY := minY
	maxDY := -minY - 1
	maxDX := maxX
	maxSteps := 2 * (-minY)
	validDYForStep := make(stepMap)

	// Find which steps are in the target range for various dy
	for i := minDY; i <= maxDY; i++ {
		for step, y, dy := 0, 0, i; step < maxSteps; step, dy = step+1, dy-1 {
			y += dy
			if y > maxY {
				continue
			}

			if y < minY {
				break
			}

			validDYForStep[step] = append(validDYForStep[step], i)
		}
	}

	// Precalculate triangular numbers up to triangle(maxX)
	t := make(map[int]int)
	for i := 0; i <= maxDX; i++ {
		t[i] = triangle(i)
	}

	// Find which steps are in the target range for various dx
	for i := 0; i <= maxDX; i++ {
		if t[i] < minX {
			continue
		}

		for step, x, dx := 0, 0, i; step < maxSteps; step, dx = step+1, towardsZero(dx) {
			x += dx
			if x > maxX {
				break
			}

			if x < minX {
				continue
			}

			for _, dy := range validDYForStep[step] {
				validShots[pair{i, dy}] = struct{}{}
			}
		}
	}

	return len(validShots)
}

func triangle(x int) int {
	return (int(math.Pow(float64(x), 2)) + x) / 2
}

func towardsZero(x int) int {
	if x == 0 {
		return 0
	}

	return x - int(math.Copysign(1, float64(x)))
}
