# Album Counter

A simple Go tool that calculates the cumulative start times for a list of tracks with their durations. This is particularly useful for creating timestamped tracklists for YouTube videos or other audio compilations.

## Usage

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Running the tool

You can run the tool directly using `go run`:

```bash
go run replace.go
```

Once the program is running, paste your tracklist into the terminal. The tool expects lines containing timestamps in the format `MM:SS`, `H:MM:SS`, or `M:SS`.

Example input:
```
 1. La la la 3:20
 2. La da de de 1:20
 3. Ba beep bop 3:20
```

After pasting the input, signal EOF (Ctrl+D on Linux/macOS, Ctrl+Z on Windows) to process the input.

### Piping Input

You can also pipe input from a file:

```bash
go run replace.go < input.txt
```

## Output

The tool appends the calculated start time to the end of each line:

```
 1. La la la 3:20 0:00
 2. La da de de 1:20 3:20
 3. Ba beep bop 3:20 4:40
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
