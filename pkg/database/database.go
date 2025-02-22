package database

import (
	"blog/config"
	"blog/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.DbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Set table prefix
	model.TablePrefix = config.AppConfig.TablePrefix

	// Auto migrate models
	err = DB.AutoMigrate(
		&model.Post{},
		&model.Category{},
		&model.Tag{},
		&model.Comment{},
		&model.User{},
		&model.Setting{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
