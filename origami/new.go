package origami

import (
	"fmt"
	"strconv"
	"strings"
)

func NewPointSet(data []string) (*PointSet, error) {
	points := make(PointSet)
	for _, line := range data {
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			return nil, fmt.Errorf("unparseable line: %v", line)
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}

		points[Point{x, y}] = struct{}{}
	}

	return &points, nil
}

func NewFolds(data []string) ([]FoldSpec, error) {
	folds := []FoldSpec{}
	for _, line := range data {
		info := strings.Split(line, "=")
		if len(info) != 2 {
			return nil, fmt.Errorf("unparseable line: %v", line)
		}

		direction := Horizontal
		if info[0] == "y" {
			direction = Vertical
		}

		foldPoint, err := strconv.Atoi(info[1])
		if err != nil {
			return nil, err
		}

		folds = append(folds, FoldSpec{direction, foldPoint})
	}

	return folds, nil
}
