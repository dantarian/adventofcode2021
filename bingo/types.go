package bingo

type Cell struct {
	Value, Row, Column int
	Marked             bool
}

type Board struct {
	Cells           map[int]*Cell
	RowMarkedCounts [5]int
	ColMarkedCounts [5]int
	HasWon          bool
	WinningBall     int
}
