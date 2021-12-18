package migrations

import (
	"log"
	"rishte/models"

	"gorm.io/gorm"
)

func RunUserMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error while Migrating User model", err.Error())
		return err
	}

	err = db.AutoMigrate(&models.UserDetailedProfile{})
	if err != nil {
		log.Fatal("Error while Migrating UserDetailed Profile model", err.Error())
		return err
	}

	return nil
}
