package album

import (
	"testing"
)

func TestAddTimestamps(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Basic Case",
			input: ` 1. La la la 3:20
 2. La da de de 1:20
 3. Ba beep bop 3:20`,
			expected: ` 1. La la la 3:20 0:00
 2. La da de de 1:20 3:20
 3. Ba beep bop 3:20 4:40`,
		},
		{
			name: "Longer than hour",
			input: `1. Song 1 30:00
2. Song 2 30:00
3. Song 3 10:00`,
			expected: `1. Song 1 30:00 0:00
2. Song 2 30:00 30:00
3. Song 3 10:00 1:00:00`,
		},
		{
			name:     "Empty Input",
			input:    "",
			expected: "",
		},
		{
			name:     "No Timestamps",
			input:    "Just some text",
			expected: "Just some text",
		},
		{
			name: "Mixed Formats",
			input: `Song A 1:00
Song B 0:00:30
Song C 65:00`,
			// 1:00 (1m) -> Start 0:00. End +60s. Total 60s.
			// 0:00:30 (30s) -> Start 1:00 (60s). End +30s. Total 90s.
			// 65:00 (65m) -> Start 1:30 (90s).
			expected: `Song A 1:00 0:00
Song B 0:00:30 1:00
Song C 65:00 1:30`,
		},
		{
			name: "Multiple Timestamps per Line",
			input: "Start 2:00 End 3:00",
			expected: "Start 2:00 0:00 End 3:00 2:00",
		},
		{
			name: "Partial Match in Numbers",
			input: "Year 2024 is long 123:45",
			// Matches 23:45 inside 123:45.
			expected: "Year 2024 is long 123:45 0:00",
		},
		{
			name: "Parentheses",
			input: "Song (3:20)",
			expected: "Song (3:20 0:00)",
		},
		{
			name: "Malformed seconds",
			input: "Song 2:5",
			expected: "Song 2:5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTimestamps(tt.input)
			if got != tt.expected {
				t.Errorf("AddTimestamps() = %q, want %q", got, tt.expected)
			}
		})
	}
}
