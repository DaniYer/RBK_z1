package iohelper

import (
	"fmt"
	"os"
)

// ReadInput читает содержимое файла и возвращает его как строку
func ReadInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return string(data), nil
}

// WriteOutput записывает строку в указанный файл
func WriteOutput(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", path, err)
	}
	return nil
}
