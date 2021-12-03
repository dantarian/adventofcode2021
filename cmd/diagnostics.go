package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/diagnostics"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(diagnosticsCmd)
}

var diagnosticsCmd = &cobra.Command{
	Use:   "diagnostics <file>",
	Short: "Perform diagnostics",
	Long:  "Calculate power consumption or life support rating",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result = int(diagnostics.PowerConsumption(data))
		} else {
			floatResult, err := diagnostics.LifeSupportRating(data)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error calculating life support rating: %v\n", err)
				os.Exit(1)
			}

			result = int(floatResult)
		}

		fmt.Printf("%v\n", result)
	},
}
