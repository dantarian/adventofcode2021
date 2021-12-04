package bingo

import (
	"strconv"
	"strings"
)

func NewBoard(lines []string) (*Board, error) {
	board := new(Board)
	board.Cells = make(map[int]*Cell)
	board.WinningBall = -1

	for rowNum, row := range lines {
		cells := strings.Fields(row)
		for colNum, cell := range cells {
			value, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}

			board.Cells[value] = &Cell{value, rowNum, colNum, false}
		}
	}

	return board, nil
}
