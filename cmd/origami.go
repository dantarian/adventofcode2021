package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/origami"
)

func init() {
	rootCmd.AddCommand(origamiCmd)
}

var origamiCmd = &cobra.Command{
	Use:   "origami <points-file> <folds-file>",
	Short: "Fold paper",
	Long:  "Analyse arrangements of dots on a sheet after folding",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		pointsData, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading points file: %v\n", err)
			os.Exit(1)
		}

		foldsData, err := file.StringLines(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading folds file: %v\n", err)
			os.Exit(1)
		}

		points, err := origami.NewPointSet(pointsData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing points file: %v\n", err)
			os.Exit(1)
		}

		folds, err := origami.NewFolds(foldsData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing folds file: %v\n", err)
			os.Exit(1)
		}

		if !part2 {
			points = points.Fold(folds[0])
			fmt.Printf("%v\n", len(*points))
		} else {
			for _, fold := range folds {
				points = points.Fold(fold)
			}
			fmt.Println(points.String())
		}

	},
}
