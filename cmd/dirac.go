package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/dirac"
)

func init() {
	rootCmd.AddCommand(diracCmd)
}

var diracCmd = &cobra.Command{
	Use:   "dirac <p1Start> <p2Start>",
	Short: "Play Dirac Dice",
	Long:  "Play Dirac Dice.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		p1Start, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[0])
			os.Exit(1)
		}
		p2Start, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Not a number: %v\n", args[1])
			os.Exit(1)
		}

		var result int64
		if !part2 {
			result = int64(dirac.PlayDeterministic(p1Start, p2Start))
		} else {
			result = dirac.PlayQuantum(p1Start, p2Start)
		}

		fmt.Printf("%v\n", result)
	},
}
