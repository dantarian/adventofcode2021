package beacons

import (
	"fmt"
	"strconv"
	"strings"
)

type Scanner struct {
	id             int
	location       Vector
	visibleBeacons []Vector
}

func (s *Scanner) FindLinkedScanner(others map[int]*Scanner) (bool, *Scanner) {
	found := false
	var scanner *Scanner
	translation := Vector{}
	direction := XPos
	rotation := Zero

scannerSearch:
	for _, other := range others {
		for _, direction = range directions() {
			for _, rotation = range rotations() {
				vectorCounts := make(map[Vector]int)
				for _, vec1 := range s.visibleBeacons {
					for _, vec2 := range other.RotatedBeacons(direction, rotation) {
						diff := vec1.Minus(vec2)
						vectorCounts[diff]++
						if vectorCounts[diff] < 12 {
							continue
						}

						// We have a match!
						found = true
						scanner = other
						translation = diff
						break scannerSearch
					}
				}
			}
		}
	}

	if !found {
		return false, nil
	}

	scanner.Rotate(direction, rotation)
	scanner.Locate(translation.Plus(s.location))
	return found, scanner
}

func (s *Scanner) Locate(location Vector) {
	s.location = location
}

func (s *Scanner) RotatedBeacons(direction Direction, rotation Rotation) []Vector {
	rotatedBeacons := []Vector{}
	for _, beacon := range s.visibleBeacons {
		rotatedBeacons = append(rotatedBeacons, beacon.Rotate(direction, rotation))
	}

	return rotatedBeacons
}

func (s *Scanner) Rotate(direction Direction, rotation Rotation) {
	s.visibleBeacons = s.RotatedBeacons(direction, rotation)
}

func (s *Scanner) TranslatedBeacons() []Vector {
	translatedBeacons := []Vector{}
	for _, beacon := range s.visibleBeacons {
		translatedBeacons = append(translatedBeacons, beacon.Plus(s.location))
	}

	return translatedBeacons
}

func (s *Scanner) String() string {
	return fmt.Sprintf("{ id: %v, location: %v }", s.id, s.location)
}

func (s *Scanner) ManhattanDistance(other *Scanner) int {
	return s.location.ManhattanDistance(other.location)
}

func ParseScanners(data []string) ([]*Scanner, error) {
	scanners := []*Scanner{}

	scanner := &Scanner{}
	for i, line := range data {
		if line == "" {
			if len(scanner.visibleBeacons) > 0 {
				scanners = append(scanners, scanner)
			}
			scanner = &Scanner{}
			continue
		}

		if strings.HasPrefix(line, "--- scanner ") {
			id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(line, "--- scanner "), " ---"))
			if err != nil {
				return nil, err
			}
			scanner = &Scanner{id: id}
			continue
		}

		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			return nil, fmt.Errorf("unexpected line format on line %v: %v", i, line)
		}

		x, err := strconv.Atoi(coords[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(coords[2])
		if err != nil {
			return nil, err
		}

		scanner.visibleBeacons = append(scanner.visibleBeacons, Vector{x, y, z})
	}

	if len(scanner.visibleBeacons) > 0 {
		scanners = append(scanners, scanner)
	}

	return scanners, nil
}

func LinkScanners(list []*Scanner) []*Scanner {
	known := []*Scanner{list[0]}
	unknown := make(map[int]*Scanner)
	for _, scanner := range list[1:] {
		unknown[scanner.id] = scanner
	}

	added := true
	for added {
		added = false
		for _, knownScanner := range known {
			found, scanner := knownScanner.FindLinkedScanner(unknown)
			if !found {
				continue
			}

			known = append(known, scanner)
			delete(unknown, scanner.id)
			added = true
		}
	}

	return known
}
