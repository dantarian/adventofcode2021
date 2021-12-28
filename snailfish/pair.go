package snailfish

func (p *pair) Magnitude() int {
	return 3*p.left.Magnitude() + 2*p.right.Magnitude()
}

func (p *pair) Plus(other Element) Element {
	sum := pair{p, other}

	for reduced := true; reduced; {
		if result := sum.explode(4); result.done {
			reduced = true
			continue
		}

		reduced, _ = sum.split()
	}

	return &sum
}

func (p *pair) Clone() Element {
	return &pair{p.left.Clone(), p.right.Clone()}
}

func (p *pair) explode(level int) explodeResponse {
	if level > 0 {
		if result := p.left.explode(level - 1); result.done {
			if !result.replacementUsed {
				p.left = result.replacement
				result.replacementUsed = true
			}
			if !result.rightUsed {
				result.rightUsed = p.right.addToLeftmost(result.right)
			}
			return result
		}

		if result := p.right.explode(level - 1); result.done {
			if !result.replacementUsed {
				p.right = result.replacement
				result.replacementUsed = true
			}
			if !result.leftUsed {
				result.leftUsed = p.left.addToRightmost(result.left)
			}
			return result
		}

		return explodeResponse{}
	}

	return explodeResponse{done: true, replacement: &value{0}, left: p.left.Magnitude(), right: p.right.Magnitude()}
}

func (p *pair) split() (bool, Element) {
	if done, result := p.left.split(); done {
		if result != nil {
			p.left = result
		}

		return true, nil
	}

	if done, result := p.right.split(); done {
		if result != nil {
			p.right = result
		}

		return true, nil
	}

	return false, nil
}

func (p *pair) addToLeftmost(value int) bool {
	return p.left.addToLeftmost(value)
}

func (p *pair) addToRightmost(value int) bool {
	return p.right.addToRightmost(value)
}
