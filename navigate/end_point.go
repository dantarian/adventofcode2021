package navigate

import (
	"strconv"
	"strings"
)

func EndPoint(data []string) (int, error) {
	x, y, err := followRoute(data)

	if err != nil {
		return 0, err
	}

	return x * y, nil
}

func EndPointWithAim(data []string) (int, error) {
	x, y, err := followRouteWithAim(data)

	if err != nil {
		return 0, err
	}

	return x * y, nil
}

func followRoute(data []string) (int, int, error) {
	x, y := 0, 0

	for _, row := range data {
		fields := strings.Fields(row)
		command := fields[0]
		value, err := strconv.Atoi(fields[1])

		if err != nil {
			return 0, 0, err
		}

		switch command {
		case "forward":
			x += value
		case "up":
			y -= value
		case "down":
			y += value
		}
	}

	return x, y, nil
}

func followRouteWithAim(data []string) (int, int, error) {
	x, y, aim := 0, 0, 0

	for _, row := range data {
		fields := strings.Fields(row)
		command := fields[0]
		value, err := strconv.Atoi(fields[1])

		if err != nil {
			return 0, 0, err
		}

		switch command {
		case "forward":
			x += value
			y += value * aim
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	return x, y, nil
}
