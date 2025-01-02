package server

import (
	"fmt"
	"log"
	"reflect"

	"strings"

	"luna-skye/satellite/server/database"
	"luna-skye/satellite/server/routes"
	"luna-skye/satellite/server/services"
	"luna-skye/satellite/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)


func StartFiberServer() {
  log.Print("Starting Fiber Server...")

  // sqlite db
  db := database.ConnectDB()
  err := database.AutoMigrateModels(db)
  if err != nil { panic(err) }

  // hbs engine
  engine := handlebars.New("./client/templates", ".hbs")
  engine.AddFunc("eq", func(a, b string) bool { return a == b })
  engine.AddFunc("getIconType", func(input string) string { return strings.Split(input, "_")[0] })
  engine.AddFunc("getIconName", func(input string) string { return strings.Split(input, "_")[1] })
  engine.AddFunc("servicePartial", func(name string) string { return fmt.Sprintf("partials/service/%s", name) })
  engine.AddFunc("getServiceResults", func(service database.Service) fiber.Map {
    //  TODO: implement dynamic results
    s := services.GetResults{}
    serviceMethod := reflect.ValueOf(s).MethodByName(utils.SnakeToPascalCase(service.Type))

    if serviceMethod.IsValid() {
      results := serviceMethod.Call([]reflect.Value{reflect.ValueOf(service)})

      if len(results) > 0 {
        if returnMap, ok := results[0].Interface().(map[string]interface{}); ok {
          log.Println(returnMap)
          return returnMap
        }

        log.Println("Service did not return map")
        return nil
      }

      log.Println("Service returned no results")
      return nil
    }

    log.Printf("Service not found: %s", utils.SnakeToPascalCase(service.Type))
    return nil
  })

  // create fiber app, with views and locals passed to them
  app := fiber.New(fiber.Config{
    Views: engine,
    PassLocalsToViews: true,
  })

  app.Use(func(c *fiber.Ctx) error {
    // disable all cache
    c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")

    // hbs local variables
    c.Locals("currentPath", c.Path())
    c.Locals("bookmarks", database.GetBookmarkCategories(db))
    c.Locals("services", database.GetServices(db))

    return c.Next()
  })

  // routes
  app.Static("/", "./client/static")
  app.Get("/", routes.IndexPage)

  log.Print("Fiber server started!")
  log.Fatal(app.Listen(":3000"))
}


// Make HTTP GET request to a URL endpoint and return the data as unmarshalled JSON
