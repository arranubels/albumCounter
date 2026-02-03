package main

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
