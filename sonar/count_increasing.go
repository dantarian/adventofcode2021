package sonar

func CountIncreasing(data []int) int {
	return CountIncreasingWindowed(data, 1)
}

func CountIncreasingWindowed(data []int, windowSize int) int {
	previous := 1000000
	count := 0
	window := 0
	for i, value := range data {
		window += value

		if i >= windowSize {
			window -= data[i-windowSize]
			if window > previous {
				count++
			}
		}

		previous = window
	}

	return count
}
