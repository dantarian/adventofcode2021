package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/caves"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(cavesCmd)
}

var cavesCmd = &cobra.Command{
	Use:   "caves <file>",
	Short: "Plot routes through caves",
	Long:  "Find numbers of routes through a cave network.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		network, err := caves.NewNetwork(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse cave data: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result, err = network.CountPaths(1)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error counting paths: %v\n", err)
				os.Exit(1)
			}
		} else {
			result, err = network.CountPaths(2)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error counting paths: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("%v\n", result)
	},
}
