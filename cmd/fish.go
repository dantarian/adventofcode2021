package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/fish"
)

func init() {
	rootCmd.AddCommand(fishCmd)
}

var fishCmd = &cobra.Command{
	Use:   "fish <file>",
	Short: "Model fish populations",
	Long:  "Calculate the numbers of lanternfish based on initial population.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.IntCSV(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			fishNumbers := fish.Model(80)
			for _, fish := range data {
				result += fishNumbers[fish]
			}
		} else {
			fishNumbers := fish.Model(256)
			for _, fish := range data {
				result += fishNumbers[fish]
			}
		}

		fmt.Printf("Number of fish: %v\n", result)
	},
}
