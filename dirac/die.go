package dirac

type DeterministicDie struct {
	rolls int
}

func (d *DeterministicDie) Roll() int {
	result := (d.rolls % 100) + 1
	d.rolls++

	return result
}
