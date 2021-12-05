package line

func (line *Line) IsDiagonal() bool {
	if (line.Start.X == line.End.X) || (line.Start.Y == line.End.Y) {
		return false
	}

	return true
}
