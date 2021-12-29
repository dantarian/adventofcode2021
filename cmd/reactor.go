package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/reactor"
)

func init() {
	rootCmd.AddCommand(reactorCmd)
}

var reactorCmd = &cobra.Command{
	Use:   "reactor <file>",
	Short: "Reboot reactor",
	Long:  "Calculate activated cubes in reactor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		instructions, err := reactor.ParseInstructions(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
			os.Exit(1)
		}

		result := int64(0)
		if !part2 {
			result = int64(reactor.SmallVolume(instructions))
		} else {
			result = reactor.LargeVolume(instructions)
		}

		fmt.Printf("%v\n", result)
	},
}
