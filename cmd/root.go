package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc2021",
	Short: "Advent of Code 2021 solvers",
	Long:  "Tools for solving Advent of Code 2021 problems.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please supply a subcommand.")
	},
}

var part2 bool

func Execute() {
	rootCmd.PersistentFlags().BoolVar(&part2, "part2", false, "Run the second part of the solution.")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
