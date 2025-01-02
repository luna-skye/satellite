package utils

import (
	"io"
	"log"
	"strings"
	"net/http"
	"encoding/json"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gofiber/fiber/v2"
)


// Converts a snake_case string to a PascalCase string
func SnakeToPascalCase(snakeCase string) string {
  words := strings.Split(snakeCase, "_")
  for i, word := range words {
    words[i] = cases.Title(language.English).String(word)
  }
  return strings.Join(words, "")
}


// Make GET request to a URL and return the resulting JSON
func GET(url string) fiber.Map {
  resp, err := http.Get(url)
  if err != nil { log.Fatalln(err) }
  defer resp.Body.Close()

  body, err := io.ReadAll(resp.Body)
  if err != nil { log.Fatalln(err) }

  data := fiber.Map{}
  err = json.Unmarshal(body, &data)
  if err != nil { log.Fatalln(err) }

  return data
}
