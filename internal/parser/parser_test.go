package parser

import (
	"RBK_z1/internal/commands"
	"regexp"
	"strconv"
)

func runPipeline(input string) string {
	clean := CleanText(input)
	words := SplitWithPunctuation(clean)

	type wordWrapper struct {
		text string
		cmds []string
	}

	// Оборачиваем слова
	wrapped := make([]wordWrapper, 0, len(words))
	for _, word := range words {
		wrapped = append(wrapped, wordWrapper{text: word})
	}

	// Обрабатываем команды
	for i := 0; i < len(wrapped); i++ {
		word := wrapped[i].text

		// (cmd, N) — массовые
		if m := regexp.MustCompile(`^\((cap|low|up),\s*(\d+)\)$`).FindStringSubmatch(word); m != nil {
			cmd := m[1]
			count, _ := strconv.Atoi(m[2])
			start := i - count
			if start < 0 {
				start = 0
			}
			for j := start; j < i; j++ {
				wrapped[j].cmds = append(wrapped[j].cmds, cmd)
			}
			// удаляем саму команду
			wrapped = append(wrapped[:i], wrapped[i+1:]...)
			i--
			continue
		}

		// одиночные команды
		if m := regexp.MustCompile(`^\((cap|low|up|bin|hex)\)$`).FindStringSubmatch(word); m != nil {
			cmd := m[1]
			if i > 0 {
				wrapped[i-1].cmds = append(wrapped[i-1].cmds, cmd)
			}
			wrapped = append(wrapped[:i], wrapped[i+1:]...)
			i--
		}
	}

	// Приоритет команд
	priority := map[string]int{
		"low": 1,
		"cap": 2,
		"up":  3,
		"bin": 4,
		"hex": 4,
	}

	// Применяем команды в порядке приоритета
	for i := range wrapped {
		if len(wrapped[i].cmds) == 0 {
			continue
		}
		// сортируем команды по приоритету
		cmds := wrapped[i].cmds
		for j := 0; j < len(cmds)-1; j++ {
			for k := j + 1; k < len(cmds); k++ {
				if priority[cmds[j]] > priority[cmds[k]] {
					cmds[j], cmds[k] = cmds[k], cmds[j]
				}
			}
		}
		// применяем по порядку
		for _, cmd := range cmds {
			wrapped[i].text = commands.ApplyCmd(cmd, wrapped[i].text)
		}
	}

	// Собираем результат
	final := make([]string, 0, len(wrapped))
	for _, w := range wrapped {
		final = append(final, w.text)
	}

	return JoinWithSpacing(final)
}


