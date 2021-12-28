package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/image"
)

func init() {
	rootCmd.AddCommand(imageCmd)
}

var imageCmd = &cobra.Command{
	Use:   "image <algorithm-file> <image-file>",
	Short: "Enhance images",
	Long:  "Enhance images according to the supplied algorithm.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		algorithmData, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading algorithm file: %v\n", err)
			os.Exit(1)
		}
		algorithm, err := image.NewAlgorithm(algorithmData[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing algorithm: %v\n", err)
			os.Exit(1)
		}

		data, err := file.StringLines(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		image := image.NewImage(data)

		result := 0
		if !part2 {
			result = image.Enhance(algorithm).Enhance(algorithm).LitPixelCount()
		} else {
			for i := 0; i < 50; i++ {
				image = image.Enhance(algorithm)
			}
			result = image.LitPixelCount()
		}

		fmt.Printf("%v\n", result)
	},
}
