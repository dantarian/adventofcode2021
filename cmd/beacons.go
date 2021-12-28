package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/beacons"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(beaconsCmd)
}

var beaconsCmd = &cobra.Command{
	Use:   "beacons <file>",
	Short: "Analyse scanners and beacons",
	Long:  "Map an area using scanners and beacons.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		scanners, err := beacons.ParseScanners(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
			os.Exit(1)
		}
		mappedScanners := beacons.LinkScanners(scanners)

		result := 0
		if !part2 {
			beacons := make(map[beacons.Vector]struct{})
			for _, scanner := range mappedScanners {
				for _, beacon := range scanner.TranslatedBeacons() {
					beacons[beacon] = struct{}{}
				}
			}

			result = len(beacons)
		} else {
			maxDistance := 0
			for i, scanner := range mappedScanners {
				for _, other := range mappedScanners[i+1:] {
					distance := scanner.ManhattanDistance(other)
					if distance > maxDistance {
						maxDistance = distance
					}
				}
			}
			result = maxDistance
		}

		fmt.Printf("%v\n", result)
	},
}
