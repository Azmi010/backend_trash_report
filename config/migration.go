package config

import (
	"trash_report/entities"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}