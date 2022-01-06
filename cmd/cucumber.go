package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/cucumbers"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(cucumberCmd)
}

var cucumberCmd = &cobra.Command{
	Use:   "cucumber <file>",
	Short: "Move cucumbers",
	Long:  "Track sea cucumber movements",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		swarm := cucumbers.NewSwarm(data)

		result := 0
		if !part2 {
			result = swarm.MovesUntilNoMoves()
		} else {
			fmt.Fprintln(os.Stderr, "Not yet implemented.")
			os.Exit(1)
		}

		fmt.Printf("%v\n", result)
	},
}
