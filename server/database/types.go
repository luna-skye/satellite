package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)




// Setting table entry, containing key value pairs for settings
// Corresponds to the `settings` sqlite table
type Setting struct {
  Key   string `gorm:"primaryKey"`
  Value string
}




// A bookmark entry, containing URL, icon, and name, used in JSON bookmark lists
type Bookmark struct {
  Name string
  Icon string
  Link string
}

// A list of bookmarks, with Value/Scan implementation for JSON in SQLite
type BookmarkList []Bookmark
func (bl BookmarkList) Value() (driver.Value, error) {
  return json.Marshal(bl)
}
func (bl *BookmarkList) Scan(value interface{}) error {
  bytes, ok := value.([]byte)
  if !ok {
    return errors.New("type assertion to []byte failed")
  }
  return json.Unmarshal(bytes, &bl)
}

// A list of categorized bookmarks, with a sort order for category sorting, name, icon, and JSON-based links
// Corresponds to the `bookmark_categories` sqlite table
type BookmarkCategory struct {
  gorm.Model
  SortOrder int
  Name      string
  Icon      string
  Links     BookmarkList `gorm:"type:json"`
}




// Service Settings
type ServiceSettings map[string]interface{}
func (ss ServiceSettings) Value() (driver.Value, error) {
  return json.Marshal(ss)
}
func (ss *ServiceSettings) Scan(value interface{}) error {
  bytes, ok := value.([]byte)
  if !ok {
    return errors.New("type assertion to []byte failed")
  }
  return json.Unmarshal(bytes, &ss)
}

// A configured service, such as 
// Corresponds to the `services` sqlite table
type Service struct {
  gorm.Model
  SortOrder int
  Name      string
  Icon      string
  Type      string
  URL       string
  Settings  ServiceSettings `gorm:"type:json"`
}
