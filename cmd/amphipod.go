package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/amphipod"
)

func init() {
	rootCmd.AddCommand(amphipodCmd)
}

var amphipodCmd = &cobra.Command{
	Use:   "amphipod",
	Short: "Arrange rooms",
	Long:  "Assign amphipods to their rooms.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result := 0
		var err error
		if !part2 {
			result, err = amphipod.Solve(false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error solving puzzle: %v\n", err)
				os.Exit(1)
			}
		} else {
			result, err = amphipod.Solve(true)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error solving puzzle: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("%v\n", result)
	},
}
