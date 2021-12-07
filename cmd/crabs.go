package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/crabs"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(crabsCmd)
}

var crabsCmd = &cobra.Command{
	Use:   "crabs <file>",
	Short: "Align crab submarines",
	Long:  "Calculate the fuel needed to align crab submarines.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.IntCSV(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		var result int
		if !part2 {
			result, err = crabs.StraightAlign(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error calculating fuel: %v\n", err)
				os.Exit(1)
			}
		} else {
			result = crabs.WeightedAlign(data)
		}

		fmt.Printf("Fuel used: %v\n", result)
	},
}
