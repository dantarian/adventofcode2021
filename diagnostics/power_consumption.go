package diagnostics

import "math"

func PowerConsumption(data []string) float64 {
	length := len(data[0])
	counts := make([]int, length);
	for _, val := range data {
		for i, char := range val {
			if string(char) == "1" {
				counts[i]++
			}
		}
	}

	gamma, epsilon := 0.0, 0.0
	for i, val := range counts {
		if val > len(data) / 2 {
			gamma += math.Pow(2, float64(length - i - 1))
		} else {
			epsilon += math.Pow(2, float64(length - i - 1))
		}
	}

	return gamma * epsilon
}