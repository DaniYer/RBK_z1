package main

import (
	"RBK_z1/internal/iohelper"
	"RBK_z1/internal/parser"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"RBK_z1/internal/commands"
)

// Точка входа
func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	// Чтение и очистка входного файла
	raw, err := iohelper.ReadInput("./files/" + args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Предобработка текста: исправление сокращений и удаление лишних пробелов
	clean := parser.CleanText(string(raw))

	// Разбиваем текст на слова, команды и знаки препинания
	words := parser.SplitWithPunctuation(clean)

	// Обработка команд (cap, low, up, bin, hex), включая расширенные (cap, N)
	for i := 0; i < len(words); i++ {
		word := words[i]

		// Расширенные команды: (cmd, N)
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
			// Удаляем команду из списка
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		// Простые команды: (cmd)
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

	// Запись результата в выходной файл
	if err := iohelper.WriteOutput("./files/"+args[1], result); err != nil {
		fmt.Println("Error:", err)
	}
}
