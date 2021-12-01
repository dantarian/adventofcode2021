package sonar

func CountIncreasing(data []int) int {
	return CountIncreasingWindowed(data, 1)
}

func CountIncreasingWindowed(data []int, windowSize int) int {
	last := 1000000
	count := 0
	window := 0
	for i, value := range data {
		window += value
		if i < windowSize {
			last = window
		} else {
			window -= data[i-windowSize]
			if window > last {
				count++
			}
			last = window
		}
	}

	return count
}
