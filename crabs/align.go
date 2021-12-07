package crabs

import (
	"fmt"
	"math"

	"pencethren.org/aoc2021/stats"
)

func StraightAlign(startPositions []int) (int, error) {
	median, err := stats.Median(startPositions)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, position := range startPositions {
		total += abs(position - median)
	}
	return fuel(startPositions, median, identity), nil
}

func WeightedAlign(startPositions []int) int {
	min, max := limits(startPositions)
	guess := min + max/2

	return search(startPositions, guess, min, max)
}

func search(startPositions []int, guess int, min int, max int) int {
	fmt.Printf("Guess %v Min %v Max %v\n", guess, min, max)
	guessCost := fuel(startPositions, guess, triangle)
	belowCost := fuel(startPositions, guess-1, triangle)
	aboveCost := fuel(startPositions, guess+1, triangle)

	if guessCost > aboveCost {
		return search(startPositions, (guess+max)/2, guess+1, max)
	}

	if guessCost > belowCost {
		return search(startPositions, ((guess+min)/2)+1, min, guess-1)
	}

	return guessCost
}

func fuel(startPositions []int, guess int, cost func(int) int) int {
	total := 0
	for _, position := range startPositions {
		total += cost(abs(position - guess))
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func triangle(x int) int {
	return (int(math.Pow(float64(x), 2)) + x) / 2
}

func identity(x int) int {
	return x
}

func limits(data []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt
	for _, value := range data {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
