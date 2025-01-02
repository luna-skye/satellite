package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


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
