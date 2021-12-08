package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/ssd"
)

func init() {
	rootCmd.AddCommand(segmentsCmd)
}

var segmentsCmd = &cobra.Command{
	Use:   "segments <file>",
	Short: "Analyse seven-segment displays",
	Long:  "Analyse and decode seven-segment displays that have been scrambled",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result, err = ssd.CountSimple(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error analysing displays: %v\n", err)
				os.Exit(1)
			}
		} else {
			result, err = ssd.DecodeAndSum(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Not yet implemented.\n")
				os.Exit(1)
			}
		}

		fmt.Printf("%v\n", result)
	},
}
