package commands

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Выполняет команду преобразования над словом
func ApplyCmd(cmd, word string) string {
	switch cmd {
	case "cap":
		return Cap(word)
	case "low":
		return Low(word)
	case "up":
		return Up(word)
	case "bin":
		return Bin(word)
	case "hex":
		return Hex(word)
	}
	return word
}

// --- Функции регистров и преобразований

func Cap(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

func Up(word string) string {
	return strings.ToUpper(word)
}

func Low(word string) string {
	return strings.ToLower(word)
}

func Hex(word string) string {
	val, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return fmt.Sprintf("%d", val)
}

func Bin(word string) string {
	val, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return fmt.Sprintf("%d", val)
}
