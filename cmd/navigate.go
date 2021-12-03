package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/navigate"
)

func init() {
	rootCmd.AddCommand(navigateCmd)
}

var navigateCmd = &cobra.Command{
	Use:   "navigate <file>",
	Short: "Follow navigation plan",
	Long:  "Calculate destination from route plan",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			result, err = navigate.EndPoint(data)
		} else {
			result, err = navigate.EndPointWithAim(data)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error following route: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%v\n", result)
	},
}
