package origami

func (ps *PointSet) Fold(fold FoldSpec) *PointSet {
	result := make(PointSet)

	for point := range *ps {
		newX := point.X
		newY := point.Y

		if fold.direction == Horizontal && point.X > fold.foldPoint {
			newX = (fold.foldPoint * 2) - point.X
		}

		if fold.direction == Vertical && point.Y > fold.foldPoint {
			newY = (fold.foldPoint * 2) - point.Y
		}

		result[Point{newX, newY}] = struct{}{}
	}

	return &result
}
