package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/stats"
	"pencethren.org/aoc2021/syntax"
)

func init() {
	rootCmd.AddCommand(syntaxCmd)
}

var syntaxCmd = &cobra.Command{
	Use:   "syntax <file>",
	Short: "Check navigation syntax",
	Long:  "Generate syntax-based scores",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		result := 0
		if !part2 {
			for _, line := range data {
				score, err := syntax.Invalid(line)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error validating syntax: %v\n", err)
					os.Exit(1)
				}
				result += score
			}
		} else {
			scores := []int{}
			for _, line := range data {
				score, valid := syntax.Complete(line)
				if !valid {
					continue
				}
				scores = append(scores, score)
			}

			result, err = stats.Median(scores)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error completing data: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("%v\n", result)
	},
}
