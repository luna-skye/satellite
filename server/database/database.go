package database

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)




// Connect to SQLite Database
func ConnectDB() *gorm.DB {
  log.Print("Opening SQLite Satellite DB...")
  db, err := gorm.Open(sqlite.Open("satellite.db"), &gorm.Config{})
  if err != nil { panic("Failed to connect database") }

  log.Print("Database connected!")
  return db
}


func AutoMigrateModels(db *gorm.DB) error {
  if err := db.AutoMigrate(&BookmarkCategory{}); err != nil { return err }
  if err := db.AutoMigrate(&Service{}); err != nil { return err }
  return nil
}




// Initialize global database settings
func InitializeSettings(db *gorm.DB) error {
  defaultSettings := []Setting{
    { Key: "theme",          Value: "hydrogen" },
    { Key: "lang",           Value: "en"       },
    { Key: "showWeather",    Value: "true"     },
    { Key: "weatherZipCode", Value: ""         },
    { Key: "weatherApiKey",  Value: ""         },
    { Key: "showBookmarks",  Value: "true"     },
    { Key: "showServices",   Value: "true"     },
    { Key: "passwordHash",   Value: ""         },
  }

  return db.Transaction(func(tx *gorm.DB) error {
    if err := db.AutoMigrate(&Setting{}); err != nil { return err }
    for _, setting := range defaultSettings {
      if err := tx.FirstOrCreate(&setting, Setting{Key: setting.Key}).Error; err != nil { return err }
    }
    return nil
  })
}




// Returns all bookmarks from the SQLite Database
func GetBookmarkCategories(db *gorm.DB) []BookmarkCategory {
  bookmarks := []BookmarkCategory{}
  db.Order("sort_order ASC").Find(&bookmarks)
  return bookmarks
}

// Create and Insert New Bookmark Category
func CreateBookmarkCategory(db *gorm.DB, bmCategory BookmarkCategory) *gorm.DB {
  log.Print("Creating new bookmark category...")
  result := db.Create(&bmCategory)
  if result.Error != nil { panic("Failed to insert bookmark category") }
  log.Print("Rows affected: ", result.RowsAffected)

  return db
}




// Returns all service configuration from the SQLite Database
func GetServices(db *gorm.DB) []Service {
  services := []Service{}
  db.Order("sort_order ASC").Find(&services)
  return services
}

// Create a new service configuration and insert it into the database
func CreateServices(db *gorm.DB, service Service) *gorm.DB {
  log.Print("Creating new bookmark category...")
  result := db.Create(&service)
  if result.Error != nil { panic("Failed to insert service") }
  log.Print("Rows affected: ", result.RowsAffected)

  return db
} 

