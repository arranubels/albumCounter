package main

import (
	"fmt"
	"io"
	"os"

	"github.com/arranubels/albumCounter/internal/album"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "albumCounter",
		Short: "Generate timestamped tracklists for YouTube",
		Long: `Album Counter calculates the cumulative start times for a list of tracks.
It accepts input via stdin or file piping and outputs the list with start times appended.`,
		Run: func(cmd *cobra.Command, args []string) {
			input, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(album.AddTimestamps(string(input)))
			fmt.Println("https://github.com/arranubels/albumCounter")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
