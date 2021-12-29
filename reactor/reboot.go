package reactor

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Instruction struct {
	state                              bool
	xMin, xMax, yMin, yMax, zMin, zMax int
}

func ParseInstructions(data []string) ([]Instruction, error) {
	instructions := []Instruction{}
	for i, line := range data {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("unexpected format on line %v: %v", i, line)
		}

		state := fields[0] == "on"

		coords := strings.Split(fields[1], ",")
		if len(coords) != 3 {
			return nil, fmt.Errorf("unexpected format on line %v: %v", i, line)
		}

		xCoords := strings.Split(strings.TrimPrefix(coords[0], "x="), "..")
		if len(xCoords) != 2 {
			return nil, fmt.Errorf("unexpected format on line %v: %v", i, line)
		}
		yCoords := strings.Split(strings.TrimPrefix(coords[1], "y="), "..")
		if len(yCoords) != 2 {
			return nil, fmt.Errorf("unexpected format on line %v: %v", i, line)
		}
		zCoords := strings.Split(strings.TrimPrefix(coords[2], "z="), "..")
		if len(zCoords) != 2 {
			return nil, fmt.Errorf("unexpected format on line %v: %v", i, line)
		}

		xMin, err := strconv.Atoi(xCoords[0])
		if err != nil {
			return nil, err
		}
		xMax, err := strconv.Atoi(xCoords[1])
		if err != nil {
			return nil, err
		}
		yMin, err := strconv.Atoi(yCoords[0])
		if err != nil {
			return nil, err
		}
		yMax, err := strconv.Atoi(yCoords[1])
		if err != nil {
			return nil, err
		}
		zMin, err := strconv.Atoi(zCoords[0])
		if err != nil {
			return nil, err
		}
		zMax, err := strconv.Atoi(zCoords[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{state, xMin, xMax, yMin, yMax, zMin, zMax})
	}

	return instructions, nil
}

func (i Instruction) Contains(x, y, z int) bool {
	return x >= i.xMin && x <= i.xMax && y >= i.yMin && y <= i.yMax && z >= i.zMin && z <= i.zMax
}

func SmallVolume(instructions []Instruction) int {
	count := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				for i := len(instructions) - 1; i >= 0; i-- {
					if instructions[i].Contains(x, y, z) {
						if instructions[i].state {
							count++
						}
						break
					}
				}
			}
		}
	}
	return count
}

func LargeVolume(instructions []Instruction) int64 {
	xEvents := []int{}
	for _, i := range instructions {
		xEvents = append(xEvents, i.xMin, i.xMax+1)
	}

	subvolumes := make(chan int64, len(xEvents)-1)

	sort.Ints(xEvents)

	lastEvent := xEvents[0]
	total := int64(0)
	workers := 0
	for _, event := range xEvents {
		length := event - lastEvent
		lastEvent = event
		if length > 0 {
			workers++
			go calculateSubVolume(length, event-1, instructions, subvolumes)
		}
	}

	for i := 0; i < workers; i++ {
		total += <-subvolumes
	}

	return total
}

func calculateSubVolume(length int, x int, instructions []Instruction, response chan<- int64) {
	area := calculateArea(x, instructions)
	volume := int64(length) * area
	response <- volume
}

func calculateArea(x int, instructions []Instruction) int64 {
	yEvents := []int{}
	for _, i := range instructions {
		if x < i.xMin || x > i.xMax {
			continue
		}
		yEvents = append(yEvents, i.yMin, i.yMax+1)
	}

	if len(yEvents) == 0 {
		return 0
	}

	sort.Ints(yEvents)
	line := int64(0)
	lastEvent := yEvents[0]
	total := int64(0)
	for _, event := range yEvents {
		if line > 0 {
			area := line * int64(event-lastEvent)
			total += area
		}
		line = calculateLine(x, event, instructions)
		lastEvent = event
	}

	return total
}

func calculateLine(x int, y int, instructions []Instruction) int64 {
	zEvents := []int{}
	for _, i := range instructions {
		if x < i.xMin || x > i.xMax || y < i.yMin || y > i.yMax {
			continue
		}
		zEvents = append(zEvents, i.zMin, i.zMax+1)
	}

	if len(zEvents) == 0 {
		return 0
	}

	sort.Ints(zEvents)
	total := int64(0)
	lastEvent := zEvents[0]
	for _, event := range zEvents {
		length := event - lastEvent
		if length > 0 {
			for i := len(instructions) - 1; i >= 0; i-- {
				if instructions[i].Contains(x, y, lastEvent) {
					if instructions[i].state {
						total += int64(length)
					}
					break
				}
			}
		}
		lastEvent = event
	}

	return total
}
