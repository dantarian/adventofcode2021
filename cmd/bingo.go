package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/bingo"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(bingoCmd)
}

var bingoCmd = &cobra.Command{
	Use:   "bingo <file>",
	Short: "Play bingo",
	Long:  "Find the winning bingo card",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		dieOnError("Error loading file: %v\n", err)

		balls, err := parseBalls(data[0])
		dieOnError("Error parsing balls data: %v\n", err)

		numBoards := len(data[1:]) / 6
		boards := make([]*bingo.Board, numBoards)
		for i := 0; i < numBoards; i++ {
			boards[i], err = bingo.NewBoard(data[i*6+2 : i*6+7])
			dieOnError("Error parsing boards data: %v\n", err)
		}

		result := 0
		if !part2 {
			winningBoard, err := bingo.Play(balls, boards)
			dieOnError("Error playing bingo: %v\n", err)
			result = winningBoard.Score()
		} else {
			lastWinningBoard, err := bingo.PlayToLose(balls, boards)
			dieOnError("Error playing bingo: %v\n", err)
			result = lastWinningBoard.Score()
		}

		fmt.Printf("%v\n", result)
	},
}

func parseBalls(ballsStr string) ([]int, error) {
	ballStrValues := strings.Split(ballsStr, ",")
	balls := make([]int, len(ballStrValues))
	var err error

	for i, val := range ballStrValues {
		if balls[i], err = strconv.Atoi(val); err != nil {
			return nil, err
		}
	}

	return balls, nil
}

func dieOnError(message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, message, err)
		os.Exit(1)
	}
}
