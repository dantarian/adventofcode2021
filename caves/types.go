package caves

type Cave struct {
	Name        string
	Big         bool
	LinkedCaves []*Cave
}

type Network map[string]*Cave

type Path struct {
	Route []*Cave
	Used  map[string]int
}
