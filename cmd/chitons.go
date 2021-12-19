package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/chitons"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(chitonsCmd)
}

var chitonsCmd = &cobra.Command{
	Use:   "chitons <file>",
	Short: "Calculate route",
	Long:  "Find the route that has the lowest risk of disturbing chitons.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		grid, err := chitons.NewGrid(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			maxCoord := len(data) - 1
			result, err = grid.CheapestRoute(chitons.Point{X: 0, Y: 0}, chitons.Point{X: maxCoord, Y: maxCoord})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error finding route: %v\n", err)
				os.Exit(1)
			}
		} else {
			factor := 5
			grid.Expand(factor)
			maxCoord := len(data)*5 - 1
			result, err = grid.CheapestRoute(chitons.Point{X: 0, Y: 0}, chitons.Point{X: maxCoord, Y: maxCoord})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error finding route: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("%v\n", result)
	},
}
