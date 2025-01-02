package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Converts a snake_case string to a PascalCase string
func SnakeToPascalCase(snakeCase string) string {
  words := strings.Split(snakeCase, "_")
  for i, word := range words {
    words[i] = cases.Title(language.English).String(word)
  }
  return strings.Join(words, "")
}
