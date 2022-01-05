package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/alu"
)

func init() {
	rootCmd.AddCommand(aluCmd)
}

var aluCmd = &cobra.Command{
	Use:   "alu",
	Short: "Run MONAD",
	Long:  "Find valid serial numbers.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result := int64(0)
		if !part2 {
			result = alu.MaxValidFinder()
		} else {
			result = alu.MinValidFinder()
		}

		fmt.Printf("%v\n", result)
	},
}
