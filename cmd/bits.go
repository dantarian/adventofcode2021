package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pencethren.org/aoc2021/bits"
	"pencethren.org/aoc2021/file"
)

func init() {
	rootCmd.AddCommand(bitsCmd)
}

var bitsCmd = &cobra.Command{
	Use:   "bits <file>",
	Short: "Handle BITS transmissions",
	Long:  "Handle BITS transmissions.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := file.StringLines(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
			os.Exit(1)
		}

		packet, err := bits.ParsePacket(data[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing data: %v\n", err)
			os.Exit(1)
		}

		result := uint64(0)
		if !part2 {
			result = packet.VersionSum()
		} else {
			result = packet.Value()
		}

		fmt.Printf("%v\n", result)
	},
}
