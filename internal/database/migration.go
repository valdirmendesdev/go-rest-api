package database

import (
	"github.com/jinzhu/gorm"
	"github.com/valdirmendesdev/go-rest-api/internal/services"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&services.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
