package main

import (
	"fmt"
	"io"
	"os"

	"github.com/arranubels/albumCounter/internal/album"
)

//go:generate sh -c "command -v gosubc >/dev/null 2>&1 && gosubc generate || go run github.com/arran4/go-subcommand/cmd/gosubc generate"

// AlbumCounter is a subcommand `albumCounter`
//
// Album Counter calculates the cumulative start times for a list of tracks.
// It accepts input via stdin or file piping and outputs the list with start times appended.
func AlbumCounter() error {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}
	fmt.Println(album.AddTimestamps(string(input)))
	fmt.Println("https://github.com/arranubels/albumCounter")
	return nil
}
