package parser

import (
	"regexp"
	"strings"
	"unicode"
)

func CleanText(text string) string {
	text = regexp.MustCompile(`\s+'`).ReplaceAllString(text, "'")
	text = regexp.MustCompile(`'\s+`).ReplaceAllString(text, "'")
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = regexp.MustCompile(`\s+([.,!?;:])`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`,([^\s])`).ReplaceAllString(text, ", $1")
	return strings.TrimSpace(text)
}

func SplitWithPunctuation(text string) []string {
	var tokens []string
	var current strings.Builder

	for _, r := range text {
		if unicode.IsSpace(r) {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		} else if isPunctuation(r) {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(r))
		} else {
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

func isPunctuation(r rune) bool {
	punct := []rune{'.', ',', ';', ':', '!', '?', '"', '\'', '“', '”', '‘', '’', '(', ')'}
	for _, p := range punct {
		if r == p {
			return true
		}
	}
	return false
}

func JoinWithSpacing(words []string) string {
	var sb strings.Builder
	for i := 0; i < len(words); i++ {
		word := words[i]
		if i > 0 && needsSpace(words[i-1], word) {
			sb.WriteRune(' ')
		}
		sb.WriteString(word)
	}
	return sb.String()
}

func needsSpace(prev, curr string) bool {
	if isPunctuationRune([]rune(curr)[0]) {
		return false
	}
	if isPunctuationRune([]rune(prev)[len([]rune(prev))-1]) {
		return true
	}
	return true
}

func isPunctuationRune(r rune) bool {
	return strings.ContainsRune(".,!?;:\")('“”‘’", r)
}
