package line

import (
	"errors"
	"strconv"
	"strings"
)

func NewLine(description string) (*Line, error) {
	pointStrings := strings.Split(description, " -> ")

	if len(pointStrings) != 2 {
		return nil, errors.New("incorrect number of points found for line")
	}

	start, err := NewPointFromString(pointStrings[0])
	if err != nil {
		return nil, err
	}

	end, err := NewPointFromString(pointStrings[1])
	if err != nil {
		return nil, err
	}

	return &Line{start, end}, nil
}

func NewPointFromString(description string) (*Point, error) {
	scalarStrings := strings.Split(description, ",")

	x, err := strconv.Atoi(scalarStrings[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(scalarStrings[1])
	if err != nil {
		return nil, err
	}

	return &Point{x, y}, nil
}
