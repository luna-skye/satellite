package services

import (
	"github.com/gofiber/fiber/v2"
	"luna-skye/satellite/utils"
	"luna-skye/satellite/server/database"
)


type GetResults struct {}

func (s GetResults) Radarr(service database.Service) fiber.Map {
  return fiber.Map{
    "status": utils.GET("http://192.168.1.238:7878/api/v3/system/status?apiKey=29fb7c6ffede430b905db2461b315c54"),
  }
}

func (s GetResults) Sonarr(service database.Service) fiber.Map {
  return fiber.Map{
    "status": utils.GET(""),
  }
}




//  TODO: rewrite with lambdas

// type ServiceResult func() fiber.Map
//
// functions := map[string]ServiceResult{
//   "radarr": func(service database.Service) {
//     return fiber.Map{}
//   }
// }
