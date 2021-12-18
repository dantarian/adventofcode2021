package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/polymers"
)

func init() {
	rootCmd.AddCommand(polymerCmd)
}

var polymerCmd = &cobra.Command{
	Use:   "polymer <template> <file>",
	Short: "Build polymers",
	Long:  "Build and analyse polymer chains.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		template := args[0]
		data, err := file.StringLines(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		insertionMap, err := polymers.NewInsertionMap(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing map: %v\n", err)
			os.Exit(1)
		}

		result := int64(0)
		var runeCounts polymers.RuneCountMap
		if !part2 {
			runeCounts = insertionMap.GetRuneCountsForPolymer(template, 10)
		} else {
			runeCounts = insertionMap.GetRuneCountsForPolymer(template, 40)
		}

		minRuneCount := int64(math.MaxInt64)
		maxRuneCount := int64(0)
		for _, count := range runeCounts {
			if count < minRuneCount {
				minRuneCount = count
			}

			if count > maxRuneCount {
				maxRuneCount = count
			}
		}

		result = maxRuneCount - minRuneCount

		fmt.Printf("%v\n", result)
	},
}
