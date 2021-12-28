package snailfish

import "fmt"

type Element interface {
	Magnitude() int
	Plus(other Element) Element
	Clone() Element
	explode(level int) explodeResponse
	split() (bool, Element)
	addToLeftmost(value int) bool
	addToRightmost(value int) bool
}

type pair struct {
	left, right Element
}

type value struct {
	value int
}

type explodeResponse struct {
	done                                 bool
	leftUsed, rightUsed, replacementUsed bool
	replacement                          Element
	left, right                          int
}

func (p *pair) String() string {
	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func (v *value) String() string {
	return fmt.Sprint(v.value)
}
