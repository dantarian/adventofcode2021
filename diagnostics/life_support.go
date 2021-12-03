package diagnostics

import (
	"errors"
	"math"
)

func LifeSupportRating(data []string) (float64, error) {
	oxygenGeneratorRatingString, err := filterMostCommon(data, 0)
	if err != nil {
		return 0.0, err
	}

  co2ScrubberRatingString, err := filterLeastCommon(data, 0)
	if err != nil {
		return 0.0, err
	}

	oxygenGeneratorRating := binStringToFloat64(oxygenGeneratorRatingString)
	co2ScrubberRating := binStringToFloat64(co2ScrubberRatingString)
	
	return oxygenGeneratorRating * co2ScrubberRating, nil
}

func filterMostCommon(data []string, position int) (string, error) {
	if len(data) == 1 {
		return data[0], nil
	}

	if position >= len(data[0]) {
		return "", errors.New("no unique value found")
	}

	numberOfOnes := 0
	for _, val := range data {
		if string(val[position]) == "1" {
			numberOfOnes++
		}
	}

	mostCommon := "0"
	if float64(numberOfOnes) >= float64(len(data)) / 2.0 {
		mostCommon = "1"
	}

	filteredData := []string{}
	for _, val := range data {
		if string(val[position]) == mostCommon {
			filteredData = append(filteredData, val)
		}
	}

	return filterMostCommon(filteredData, position + 1)
}

func filterLeastCommon(data []string, position int) (string, error) {
	if len(data) == 1 {
		return data[0], nil
	}

	if position >= len(data[0]) {
		return "", errors.New("no unique value found")
	}

	numberOfOnes := 0
	for _, val := range data {
		if string(val[position]) == "1" {
			numberOfOnes++
		}
	}

	leastCommon := "0"
	if float64(numberOfOnes) < float64(len(data)) / 2.0 {
		leastCommon = "1"
	}

	filteredData := []string{}
	for _, val := range data {
		if string(val[position]) == leastCommon {
			filteredData = append(filteredData, val)
		}
	}

	return filterLeastCommon(filteredData, position + 1)
}

func binStringToFloat64(binstr string) float64 {
	result := 0.0
	length := len(binstr)
	for i, val := range binstr {
		if string(val) == "1" {
			result += math.Pow(2, float64(length - i - 1))
		}
	}
	return result
}