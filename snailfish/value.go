package snailfish

func (v *value) Magnitude() int {
	return v.value
}

func (v *value) Plus(other Element) Element {
	sum := pair{v, other}

	for reduced := true; reduced; {
		if result := sum.explode(4); result.done {
			reduced = true
			continue
		}

		reduced, _ = sum.split()
	}

	return &sum
}

func (v *value) Clone() Element {
	return &value{v.value}
}

func (v *value) explode(level int) explodeResponse {
	return explodeResponse{}
}

func (v *value) split() (bool, Element) {
	if v.value < 10 {
		return false, nil
	}

	left := v.value / 2
	right := v.value - left

	replacement := &pair{&value{left}, &value{right}}
	return true, replacement
}

func (v *value) addToLeftmost(value int) bool {
	v.value += value
	return true
}

func (v *value) addToRightmost(value int) bool {
	v.value += value
	return true
}
