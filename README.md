# Album Counter - YouTube Tracklist Generator

**Album Counter** is a simple yet powerful Go tool designed to assist content creators, music enthusiasts, and archivists in generating timestamped tracklists for YouTube videos.

When uploading full albums or DJ mixes to YouTube, adding timestamps in the description allows viewers to jump to specific tracks. This tool automatically calculates the cumulative start time for each track based on its duration, saving you from manual calculations.

## Features

- **Automatic Calculation**: Input a list of tracks with durations, and get the start time for each.
- **Flexible Input**: Supports `MM:SS`, `H:MM:SS`, and `M:SS` formats.
- **Mixed Formats**: Handles lines with different timestamp formats gracefully.
- **Pipe Support**: works with standard input/output for easy integration into scripts.

## Installation

### From Source

Ensure you have [Go](https://golang.org/dl/) installed.

```bash
git clone https://github.com/arranubels/albumCounter.git
cd albumCounter
go install ./cmd/albumCounter
```

## Usage

Run the tool and paste your tracklist, or pipe a file into it.

### Interactive Mode

```bash
albumCounter
```

Paste your tracklist (Ctrl+D to finish):
```text
1. First Song 3:20
2. Second Song 4:15
3. Third Song 2:45
```

**Output:**
```text
1. First Song 3:20 0:00
2. Second Song 4:15 3:20
3. Third Song 2:45 7:35
```

### File Input

```bash
albumCounter < album.txt
```

### Examples

#### Example 1: Basic Album

**Input (`album.txt`):**
```
01 - Intro 1:30
02 - Hit Single 3:45
03 - Ballad 4:20
```

**Command:**
```bash
albumCounter < album.txt
```

**Output:**
```
01 - Intro 1:30 0:00
02 - Hit Single 3:45 1:30
03 - Ballad 4:20 5:15
```

#### Example 2: Long Mix (Hours)

**Input:**
```
Part 1 45:00
Part 2 30:00
Part 3 50:00
```

**Output:**
```
Part 1 45:00 0:00
Part 2 30:00 45:00
Part 3 50:00 1:15:00
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
