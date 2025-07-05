package commands

import (
	"testing"
)

func TestApplyCmd(t *testing.T) {
	tests := []struct {
		name     string
		cmd      string
		input    string
		expected string
	}{
		{
			name:     "uppercase",
			cmd:      "up",
			input:    "hello",
			expected: "HELLO",
		},
		{
			name:     "lowercase",
			cmd:      "low",
			input:    "HELLO",
			expected: "hello",
		},
		{
			name:     "capitalize",
			cmd:      "cap",
			input:    "hELLo",
			expected: "Hello",
		},
		{
			name:     "binary to decimal",
			cmd:      "bin",
			input:    "101",
			expected: "5",
		},
		{
			name:     "invalid binary input",
			cmd:      "bin",
			input:    "12",
			expected: "12", // неверный формат — без изменений
		},
		{
			name:     "hexadecimal to decimal",
			cmd:      "hex",
			input:    "2a",
			expected: "42",
		},
		{
			name:     "invalid hex input",
			cmd:      "hex",
			input:    "zz",
			expected: "zz", // неверный hex — оставить как есть
		},
		{
			name:     "unknown command fallback",
			cmd:      "unknown",
			input:    "test",
			expected: "test", // default поведение
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyCmd(tt.cmd, tt.input)
			if got != tt.expected {
				t.Errorf("ApplyCmd(%q, %q) = %q; want %q", tt.cmd, tt.input, got, tt.expected)
			}
		})
	}
}
