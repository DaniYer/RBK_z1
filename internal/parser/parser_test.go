package parser_test

import (
	"RBK_z1/internal/commands"
	"RBK_z1/internal/parser"
	"regexp"
	"strconv"
	"testing"
)

func runPipeline(input string) string {
	clean := parser.CleanText(input)
	words := parser.SplitWithPunctuation(clean)

	for i := 0; i < len(words); i++ {
		word := words[i]

		if m := regexp.MustCompile(`^\((cap|low|up),\s*(\d+)\)$`).FindStringSubmatch(word); m != nil {
			cmd := m[1]
			count, _ := strconv.Atoi(m[2])
			start := i - count
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				words[j] = commands.ApplyCmd(cmd, words[j])
			}
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if m := regexp.MustCompile(`^\((cap|low|up|bin|hex)\)$`).FindStringSubmatch(word); m != nil {
			cmd := m[1]
			if i > 0 {
				words[i-1] = commands.ApplyCmd(cmd, words[i-1])
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return parser.JoinWithSpacing(words)
}

func TestCleanText(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"I 'm ready .", "I'm ready."},
		{"we 've been there , right ?", "we've been there, right?"},
	}
	for _, test := range tests {
		got := parser.CleanText(test.input)
		if got != test.output {
			t.Errorf("CleanText(%q) = %q; want %q", test.input, got, test.output)
		}
	}
}

func TestSplitWithPunctuation(t *testing.T) {
	input := "Let's go, now! (cap) Here."
	expected := []string{"Let's", "go", ",", "now", "!", "(cap)", "Here", "."}
	result := parser.SplitWithPunctuation(input)
	if len(result) != len(expected) {
		t.Fatalf("Wrong number of tokens: got %d, want %d", len(result), len(expected))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Token %d: got %q, want %q", i, result[i], expected[i])
		}
	}
}

func TestJoinWithSpacing(t *testing.T) {
	input := []string{"Hello", ",", "world", "!"}
	want := "Hello, world!"
	got := parser.JoinWithSpacing(input)
	if got != want {
		t.Errorf("JoinWithSpacing = %q; want %q", got, want)
	}
}

func TestFullTextProcessing(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			"classic sample.txt",
			"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			"binary and hex",
			"2a (hex) 101 (bin)",
			"42 5",
		},
		{
			"cap and low mix",
			"she is AMAZING (low, 1) (cap, 2)",
			"She Is amazing",
		},
		{
			"redundant commands",
			"wow (up) (up) (cap)",
			"WOW",
		},
		{
			"punctuation logic",
			"Hello , world ! (up) . (low) IT IS (cap, 3)",
			"Hello, world! WORLD. it is It Is World",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := runPipeline(tc.input)
			if got != tc.output {
				t.Errorf("\nGot:   %q\nWant: %q", got, tc.output)
			}
		})
	}
}
