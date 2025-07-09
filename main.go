package main

import (
	"RBK_z1/internal/commands"
	"RBK_z1/internal/iohelper"
	"RBK_z1/internal/parser"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

// Проверка, состоит ли слово из латинских букв (для команд и a/an)
func isLatin(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && r > unicode.MaxLatin1 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	raw, err := iohelper.ReadInput("./files/" + args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	clean := parser.CleanText(string(raw))
	words := parser.SplitWithPunctuation(clean)

	for i := 0; i < len(words); i++ {
		word := words[i]

		// 🔤 Обработка "a"/"an" для латинских слов
		if i < len(words)-1 && (word == "a" || word == "an" || word == "A" || word == "An" || word == "AN") {
			next := words[i+1]
			if isLatin(next) && len(next) > 0 {
				firstLetter := []rune(next)[0]
				isVowel := regexp.MustCompile(`(?i)^[aeiou]`).MatchString(string(firstLetter))

				original := word
				if isVowel && word == "a" {
					words[i] = "an"
				} else if !isVowel && word == "an" {
					words[i] = "a"
				}

				if words[i] == "an" && (original == "A" || original == "AN") {
					words[i] = "An"
				} else if words[i] == "a" && (original == "An" || original == "AN") {
					words[i] = "A"
				}
			}
		}

		if !isLatin(word) {
			continue
		}

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

	result := parser.JoinWithSpacing(words)

	if err := iohelper.WriteOutput("./files/"+args[1], result); err != nil {
		fmt.Println("Error:", err)
	}
}
