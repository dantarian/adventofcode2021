package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/probe"
)

func init() {
	rootCmd.AddCommand(probeCmd)
}

var probeCmd = &cobra.Command{
	Use:   "probe <minX> <minY> <maxX> <maxY>",
	Short: "Shoot probes",
	Long:  "Work out how many ways there are to launch a probe into a target area.",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		minX, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[0])
			os.Exit(1)
		}
		minY, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[1])
			os.Exit(1)
		}
		maxX, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[2])
			os.Exit(1)
		}
		maxY, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[3])
			os.Exit(1)
		}

		if minY >= maxY {
			fmt.Fprintln(os.Stderr, "minY must be less than maxY")
			os.Exit(1)
		}
		if minX >= maxX {
			fmt.Fprintln(os.Stderr, "minX must be less than maxX")
			os.Exit(1)
		}

		var result int
		if !part2 {
			result = probe.MaxHeight(minY)
		} else {
			result = probe.Target(minX, minY, maxX, maxY)
		}

		fmt.Printf("%v\n", result)
	},
}
