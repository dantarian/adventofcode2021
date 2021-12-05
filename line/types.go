package line

type Point struct {
	X, Y int
}

type Line struct {
	Start *Point
	End   *Point
}
