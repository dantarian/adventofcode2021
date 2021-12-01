package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/sonar"
)

func init() {
	rootCmd.AddCommand(sonarCmd)
}

var sonarCmd = &cobra.Command{
	Use:   "sonar <file>",
	Short: "Interpret sonar data",
	Long:  "Count increasing sonar readings.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.IntLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result = sonar.CountIncreasing(data)
		} else {
			result = sonar.CountIncreasingWindowed(data, 3)
		}

		fmt.Printf("Total increasing values: %v\n", result)
	},
}
