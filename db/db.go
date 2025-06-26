package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortUrls struct {
	gorm.Model
	longurl  string
	shorturl string
}

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}

func Migrate(db *gorm.DB) (ok bool) {
	err := db.AutoMigrate(&ShortUrls{})
	if err != nil {
		return false
	}
	return true
}
