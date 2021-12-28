package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/file"
	"pencethren.org/aoc2021/snailfish"
)

func init() {
	rootCmd.AddCommand(snailfishCmd)
}

var snailfishCmd = &cobra.Command{
	Use:   "snailfish <file>",
	Short: "Do snailfish maths",
	Long:  "Parse and evaluate snailfish numbers",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		numbers := []snailfish.Element{}
		for i, line := range data {
			number, err := snailfish.ParsePair(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing line %v: %v\n", i, err)
				os.Exit(1)
			}
			numbers = append(numbers, number)
		}

		result := 0
		if !part2 {
			var base snailfish.Element
			for _, number := range numbers {
				if base == nil {
					base = number
					continue
				}

				base = base.Plus(number)
			}

			result = base.Magnitude()
		} else {
			maximum := 0
			numbersToAdd := []snailfish.Element{}
			for _, number1 := range numbers {
				for _, number2 := range numbersToAdd {
					xPlusY := number1.Clone().Plus(number2.Clone())
					yPlusX := number2.Clone().Plus(number1.Clone())

					xPlusYMag := xPlusY.Magnitude()
					yPlusXMag := yPlusX.Magnitude()

					if maximum < xPlusYMag {
						maximum = xPlusYMag
					}

					if maximum < yPlusXMag {
						maximum = yPlusXMag
					}
				}

				numbersToAdd = append(numbersToAdd, number1)
			}

			result = maximum
		}

		fmt.Printf("%v\n", result)
	},
}
