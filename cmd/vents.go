package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/line"
)

func init() {
	rootCmd.AddCommand(ventsCmd)
}

var ventsCmd = &cobra.Command{
	Use:   "vents <file>",
	Short: "Find vent overlaps",
	Long:  "Find where clouds of gas from vents are particularly thick.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		lines := []*line.Line{}
		for _, lineDescription := range data {
			line, err := line.NewLine(lineDescription)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing lines: %v\n", err)
				os.Exit(1)
			}
			lines = append(lines, line)
		}

		result := 0
		if !part2 {
			pointCounts := make(map[line.Point]int)
			for _, l := range lines {
				if !l.IsDiagonal() {
					coveredPoints := l.CoveredPoints()
					for _, point := range coveredPoints {
						_, present := pointCounts[point]
						if !present {
							pointCounts[point] = 0
						}
						pointCounts[point]++

						if pointCounts[point] == 2 {
							result++
						}
					}
				}
			}
		} else {
			pointCounts := make(map[line.Point]int)
			for _, l := range lines {
				coveredPoints := l.CoveredPoints()
				for _, point := range coveredPoints {
					_, present := pointCounts[point]
					if !present {
						pointCounts[point] = 0
					}
					pointCounts[point]++

					if pointCounts[point] == 2 {
						result++
					}
				}
			}
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error following route: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%v\n", result)
	},
}
