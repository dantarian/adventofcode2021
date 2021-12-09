package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/smoke"
)

func init() {
	rootCmd.AddCommand(smokeCmd)
}

var smokeCmd = &cobra.Command{
	Use:   "smoke <file>",
	Short: "Identify smoky locations",
	Long:  "Find the dangerous points on a landscape full of smoke",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		parsedData := [][]int{}
		for _, line := range data {
			parsedLine := []int{}
			for _, value := range line {
				intVal, _ := strconv.Atoi(string(value))
				parsedLine = append(parsedLine, intVal)
			}
			parsedData = append(parsedData, parsedLine)
		}

		result := 0
		lowPoints := smoke.LowPoints(parsedData)
		if !part2 {
			for _, point := range lowPoints {
				result += point.Z + 1
			}
		} else {
			basinSizes := []int{}
			for _, point := range lowPoints {
				basinSizes = append(basinSizes, smoke.BasinSize(parsedData, point))
			}
			sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
			result = basinSizes[0] * basinSizes[1] * basinSizes[2]
		}

		fmt.Printf("%v\n", result)
	},
}
