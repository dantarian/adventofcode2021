package bingo

import "errors"

func Play(balls []int, boards []*Board) (*Board, error) {
	for _, ball := range balls {
		for _, board := range boards {
			board.mark(ball)
			if board.HasWon {
				return board, nil
			}
		}
	}

	return nil, errors.New("no winning board found")
}

func PlayToLose(balls []int, boards []*Board) (*Board, error) {
	winningBoards := []*Board{}

	for _, ball := range balls {
		for _, board := range boards {
			if !board.HasWon {
				board.mark(ball)
				if board.HasWon {
					winningBoards = append(winningBoards, board)
				}
			}
		}
	}

	if len(winningBoards) == 0 {
		return nil, errors.New("no winning board found")
	}

	return winningBoards[len(winningBoards)-1], nil
}

func (board *Board) mark(ball int) {
	cell, present := board.Cells[ball]

	if !present {
		return
	}

	cell.Marked = true
	board.ColMarkedCounts[cell.Column]++
	board.RowMarkedCounts[cell.Row]++

	if (board.ColMarkedCounts[cell.Column] == 5 || board.RowMarkedCounts[cell.Row] == 5) && !board.HasWon {
		board.HasWon = true
		board.WinningBall = ball
	}
}
