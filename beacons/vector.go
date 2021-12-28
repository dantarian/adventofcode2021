package beacons

type Vector struct {
	x, y, z int
}

type Direction int

const (
	XPos Direction = iota
	XNeg
	YPos
	YNeg
	ZPos
	ZNeg
)

func directions() []Direction {
	return []Direction{XPos, XNeg, YPos, YNeg, ZPos, ZNeg}
}

type Rotation int

const (
	Zero Rotation = iota
	Ninety
	OneEighty
	TwoSeventy
)

func rotations() []Rotation {
	return []Rotation{Zero, Ninety, OneEighty, TwoSeventy}
}

func (v Vector) Plus(other Vector) Vector {
	return Vector{v.x + other.x, v.y + other.y, v.z + other.z}
}

func (v Vector) Minus(other Vector) Vector {
	return Vector{v.x - other.x, v.y - other.y, v.z - other.z}
}

func (v Vector) Rotate(direction Direction, rotation Rotation) Vector {
	switch direction {
	case XPos:
		switch rotation {
		case Zero:
			return Vector{v.x, v.y, v.z}
		case Ninety:
			return Vector{v.x, v.z, -v.y}
		case OneEighty:
			return Vector{v.x, -v.y, -v.z}
		case TwoSeventy:
			return Vector{v.x, -v.z, v.y}
		}
	case XNeg:
		switch rotation {
		case Zero:
			return Vector{-v.x, v.y, -v.z}
		case Ninety:
			return Vector{-v.x, -v.z, -v.y}
		case OneEighty:
			return Vector{-v.x, -v.y, v.z}
		case TwoSeventy:
			return Vector{-v.x, v.z, v.y}
		}
	case YPos:
		switch rotation {
		case Zero:
			return Vector{-v.y, v.x, v.z}
		case Ninety:
			return Vector{v.z, v.x, v.y}
		case OneEighty:
			return Vector{v.y, v.x, -v.z}
		case TwoSeventy:
			return Vector{-v.z, v.x, -v.y}
		}
	case YNeg:
		switch rotation {
		case Zero:
			return Vector{v.y, -v.x, v.z}
		case Ninety:
			return Vector{-v.z, -v.x, v.y}
		case OneEighty:
			return Vector{-v.y, -v.x, -v.z}
		case TwoSeventy:
			return Vector{v.z, -v.x, -v.y}
		}
	case ZPos:
		switch rotation {
		case Zero:
			return Vector{-v.z, v.y, v.x}
		case Ninety:
			return Vector{-v.y, -v.z, v.x}
		case OneEighty:
			return Vector{v.z, -v.y, v.x}
		case TwoSeventy:
			return Vector{v.y, v.z, v.x}
		}
	case ZNeg:
		switch rotation {
		case Zero:
			return Vector{v.z, v.y, -v.x}
		case Ninety:
			return Vector{v.y, -v.z, -v.x}
		case OneEighty:
			return Vector{-v.z, -v.y, -v.x}
		case TwoSeventy:
			return Vector{-v.y, v.z, -v.x}
		}
	}

	return Vector{} // This should never happen.
}

func (v Vector) ManhattanDistance(other Vector) int {
	x := v.x - other.x
	y := v.y - other.y
	z := v.z - other.z
	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	if z < 0 {
		z = -z
	}

	return x + y + z
}
