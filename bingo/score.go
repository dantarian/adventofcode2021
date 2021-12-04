package bingo

func (board *Board) Score() int {
	sum := 0
	for _, cell := range board.Cells {
		if cell.Marked {
			continue
		}

		sum += cell.Value
	}
	return sum * board.WinningBall
}
