package amphipod

type flavour rune

const (
	A    flavour = 'A'
	B    flavour = 'B'
	C    flavour = 'C'
	D    flavour = 'D'
	None flavour = ' '
)

type cell struct {
	id             int
	occupiable     bool
	target         bool
	allowedFlavour flavour
	linkedTo       []*cell
}

type amphipod struct {
	flavour  flavour
	location *cell
}

type state []amphipod

func (s state) String() string {
	representation := make([]rune, len(s)+11)
	for i := range representation {
		representation[i] = ' '
	}
	for _, a := range s {
		representation[a.location.id] = rune(a.flavour)
	}
	return string(representation)
}

func FromString(representation string, cells []*cell) state {
	newState := state{}
	for i, letter := range representation {
		switch letter {
		case 'A':
			newState = append(newState, amphipod{A, cells[i]})
		case 'B':
			newState = append(newState, amphipod{B, cells[i]})
		case 'C':
			newState = append(newState, amphipod{C, cells[i]})
		case 'D':
			newState = append(newState, amphipod{D, cells[i]})
		}
	}
	return newState
}

func (c *cell) link(other *cell) {
	c.linkedTo = append(c.linkedTo, other)
	other.linkedTo = append(other.linkedTo, c)
}

func createCells(expanded bool) []*cell {
	size := 19
	if expanded {
		size += 8
	}
	cells := make([]*cell, size)

	cells[0] = &cell{0, true, false, None, []*cell{}}
	cells[1] = &cell{1, true, false, None, []*cell{}}
	cells[2] = &cell{2, false, false, None, []*cell{}}
	cells[3] = &cell{3, true, false, None, []*cell{}}
	cells[4] = &cell{4, false, false, None, []*cell{}}
	cells[5] = &cell{5, true, false, None, []*cell{}}
	cells[6] = &cell{6, false, false, None, []*cell{}}
	cells[7] = &cell{7, true, false, None, []*cell{}}
	cells[8] = &cell{8, false, false, None, []*cell{}}
	cells[9] = &cell{9, true, false, None, []*cell{}}
	cells[10] = &cell{10, true, false, None, []*cell{}}
	cells[11] = &cell{11, true, true, A, []*cell{}}
	cells[12] = &cell{12, true, true, A, []*cell{}}
	cells[13] = &cell{13, true, true, B, []*cell{}}
	cells[14] = &cell{14, true, true, B, []*cell{}}
	cells[15] = &cell{15, true, true, C, []*cell{}}
	cells[16] = &cell{16, true, true, C, []*cell{}}
	cells[17] = &cell{17, true, true, D, []*cell{}}
	cells[18] = &cell{18, true, true, D, []*cell{}}
	if expanded {
		cells[19] = &cell{19, true, true, A, []*cell{}}
		cells[20] = &cell{20, true, true, A, []*cell{}}
		cells[21] = &cell{21, true, true, B, []*cell{}}
		cells[22] = &cell{22, true, true, B, []*cell{}}
		cells[23] = &cell{23, true, true, C, []*cell{}}
		cells[24] = &cell{24, true, true, C, []*cell{}}
		cells[25] = &cell{25, true, true, D, []*cell{}}
		cells[26] = &cell{26, true, true, D, []*cell{}}
	}

	cells[0].link(cells[1])
	cells[1].link(cells[2])
	cells[2].link(cells[3])
	cells[3].link(cells[4])
	cells[4].link(cells[5])
	cells[5].link(cells[6])
	cells[6].link(cells[7])
	cells[7].link(cells[8])
	cells[8].link(cells[9])
	cells[9].link(cells[10])
	cells[2].link(cells[11])
	cells[11].link(cells[12])
	cells[4].link(cells[13])
	cells[13].link(cells[14])
	cells[6].link(cells[15])
	cells[15].link(cells[16])
	cells[8].link(cells[17])
	cells[17].link(cells[18])

	if expanded {
		cells[12].link(cells[19])
		cells[19].link(cells[20])
		cells[14].link(cells[21])
		cells[21].link(cells[22])
		cells[16].link(cells[23])
		cells[23].link(cells[24])
		cells[18].link(cells[25])
		cells[25].link(cells[26])
	}

	return cells
}

func createInitialState(cells []*cell) state {
	if len(cells) > 19 {
		return FromString("           ADCCBBAADDBDABCC", cells)
	}
	return FromString("           ADCDBBAC", cells)
}
