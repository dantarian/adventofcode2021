package chitons

func (g *Grid) Expand(factor int) {
	maxCoord := 0
	for location := range *g {
		if location.X < maxCoord {
			continue
		}
		maxCoord = location.X
	}

	baseSize := maxCoord + 1

	for location, value := range *g {
		for i := 0; i < factor; i++ {
			for j := 0; j < factor; j++ {
				if i == 0 && j == 0 {
					continue
				}

				newValue := value + i + j
				for ; newValue > 9; newValue -= 9 {
				}

				(*g)[Point{(i * baseSize) + location.X, (j * baseSize) + location.Y}] = newValue
			}
		}
	}
}
