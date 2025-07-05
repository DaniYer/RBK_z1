package parser

import (
	"regexp"
	"strings"
)

// Предобработка текста: сокращения и пунктуация
func CleanText(text string) string {
	reApostropheFix := regexp.MustCompile(`\b(\w+)\s*'\s*(m|re|s|ve|d|ll)\b`)
	text = reApostropheFix.ReplaceAllString(text, `$1@@$2`)
	text = strings.ReplaceAll(text, "@@", `'`)

	reSpaceBeforePunct := regexp.MustCompile(`\s+([.,!?;:])`)
	text = reSpaceBeforePunct.ReplaceAllString(text, `$1`)

	return text
}

// Разделяет строку на слова, команды и знаки препинания
func SplitWithPunctuation(text string) []string {
	re := regexp.MustCompile(`[\w@']+|\([^)]+\)|[.,!?;:]`)
	return re.FindAllString(text, -1)
}

// Собирает текст обратно, соблюдая пробелы и пунктуацию
func JoinWithSpacing(words []string) string {
	var sb strings.Builder
	punctuation := map[string]bool{
		".": true, ",": true, "!": true, "?": true, ";": true, ":": true,
	}

	for i, word := range words {
		if i > 0 && !punctuation[word] {
			sb.WriteString(" ")
		}
		sb.WriteString(word)
	}
	return sb.String()
}
