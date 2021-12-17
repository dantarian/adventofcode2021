package caves

import (
	"errors"
	"fmt"
)

func (n *Network) CountPaths(allowedSmallCaveVisits int) (int, error) {
	start, found := (*n)["start"]
	if !found {
		return 0, errors.New("network contains no start cave")
	}

	basePath := newPath(nil, start)

	paths := basePath.follow(allowedSmallCaveVisits)

	// Deduplicate paths
	pathSet := make(map[string]struct{})
	for _, path := range paths {
		pathSet[path.String()] = struct{}{}
	}

	return len(pathSet), nil
}

func (p *Path) follow(allowedSmallCaveVisits int) []*Path {
	paths := []*Path{}
	baseCave := p.Route[len(p.Route)-1]
	if allowedSmallCaveVisits > 1 {
		for _, visits := range p.Used {
			if visits > 1 {
				allowedSmallCaveVisits = 1
				break
			}
		}
	}

	for _, cave := range baseCave.LinkedCaves {
		if cave.Name == "start" {
			continue
		}

		if cave.Name == "end" {
			paths = append(paths, newPath(p, cave))
			continue
		}

		count, caveUsed := p.Used[cave.Name]
		if caveUsed && count >= allowedSmallCaveVisits {
			continue
		}

		paths = append(paths, newPath(p, cave).follow(allowedSmallCaveVisits)...)
	}

	return paths
}

func (p *Path) String() string {
	return fmt.Sprintf("%v", p.Route)
}

func (c *Cave) String() string {
	return c.Name
}
