package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/octopus"
)

func init() {
	rootCmd.AddCommand(octopusCmd)
}

var octopusCmd = &cobra.Command{
	Use:   "octopus <file>",
	Short: "Count flashes",
	Long:  "Count octopus flashes over time",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		swarm, err := octopus.NewSwarm(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing octopuses: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result = swarm.TotalFlashes(100)
		} else {
			result = swarm.SyncedFlashStep()
		}

		fmt.Printf("%v\n", result)
	},
}
