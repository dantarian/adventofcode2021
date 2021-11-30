package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(day1Cmd)
}

var day1Cmd = &cobra.Command{
	Use:   "day1 <file>",
	Short: "Solve it",
	Long:  "Solve the puzzle.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.IntLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		if !part2 {
			fmt.Printf("No solution here yet.\n")
		} else {
			fmt.Printf("No solution here yet.\n")
		}
	},
}
