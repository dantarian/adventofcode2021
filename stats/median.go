package stats

import (
	"errors"
	"sort"
)

func Median(values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("no values to find median of")
	}

	sort.Ints(values)

	if len(values)%2 == 0 {
		return (values[len(values)/2-1] + values[len(values)/2]) / 2, nil
	}

	return values[len(values)/2], nil
}
