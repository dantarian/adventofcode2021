package caves

import (
	"fmt"
	"regexp"
	"strings"
)

func NewNetwork(data []string) (Network, error) {
	network := make(Network)
	var err error

	network["start"], err = newCave("start")
	if err != nil {
		return nil, err
	}

	network["end"], err = newCave("end")
	if err != nil {
		return nil, err
	}

	for _, line := range data {
		caveNames := strings.Split(line, "-")
		if len(caveNames) != 2 {
			return nil, fmt.Errorf("failed to parse line: %v", line)
		}

		cave1, cave1Present := network[caveNames[0]]
		cave2, cave2Present := network[caveNames[1]]

		if !cave1Present {
			cave1, err = newCave(caveNames[0])
			if err != nil {
				return nil, err
			}
			network[caveNames[0]] = cave1
		}
		if !cave2Present {
			cave2, err = newCave(caveNames[1])
			if err != nil {
				return nil, err
			}
			network[caveNames[1]] = cave2
		}

		cave1.LinkedCaves = append(cave1.LinkedCaves, cave2)
		cave2.LinkedCaves = append(cave2.LinkedCaves, cave1)
	}

	return network, nil
}

func newCave(name string) (*Cave, error) {
	big, err := regexp.MatchString("[A-Z]+", name)
	if err != nil {
		return nil, err
	}

	return &Cave{name, big, []*Cave{}}, nil
}

func newPath(base *Path, cave *Cave) *Path {
	route := []*Cave{}
	used := make(map[string]int)
	if base != nil {
		route = append(route, base.Route...)
		for k, v := range base.Used {
			used[k] = v
		}
	}

	route = append(route, cave)
	if !cave.Big {
		used[cave.Name] = used[cave.Name] + 1
	}

	return &Path{route, used}
}
