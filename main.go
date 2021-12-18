package main

import (
	"context"
	"rishte/database"
	"rishte/migrations"
	"rishte/models"
	"rishte/repository"
	"time"

	"gorm.io/gorm/logger"
)

func main() {
	db, err := database.GetDatabase()
	if err != nil {
		logger.Default.Error(context.Background(), err.Error())
	}

	err = migrations.RunUserMigrations(db)
	if err != nil {
		logger.Default.Error(context.Background(), err.Error())
	}

	repositories := repository.InitRepositories(db)
	modelUser := models.User{
		FirstName:    "Himanshu",
		LastName:     "Dhankhar",
		ProfileFor:   "self",
		Religion:     "hindu",
		Community:    "Indian",
		Gender:       "male",
		Email:        "dhankhar7924@gmail.com",
		MobileNumber: 9079161380,
		Password:     "Dhankhar7924@",
		DOB: time.Date(1997, time.January,
			11, 21, 34, 01, 0, time.UTC),
		Country: "India",
	}
	repositories.UserRepo.CreateUser(modelUser)

}
