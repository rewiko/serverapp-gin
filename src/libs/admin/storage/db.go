package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rewiko/gin-app/libs/admin/model"
)

// InitDB creates and migrates the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "storage.db")
	if err != nil {
		return nil, err
	}

	//db.LogMode(true)
	db.AutoMigrate(&model.User{}, &model.Chocolate{})

	return db, nil
}
