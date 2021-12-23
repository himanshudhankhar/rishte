package main

import (
	"context"
	"os"
	"rishte/database"
	"rishte/handlers"
	"rishte/migrations"
	"rishte/models"
	"rishte/repository"
	"time"

	"github.com/gin-gonic/gin"
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

	repos := repository.InitRepositories(db)

	testUser := models.User{
		FirstName:    "himanshu",
		LastName:     "dhankhar",
		ProfileFor:   "self",
		Religion:     "hindu",
		Community:    "jat",
		Gender:       "male",
		Email:        "dhankhar7924@gmail.com",
		MobileNumber: 9079161380,
		Password:     "Dhankhar7924@",
		Country:      "India",
		DOB: time.Date(1997, time.January, 10, 0,0,0,0, time.UTC),
	}
	repos.UserRepo.CreateUser(testUser)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	handler := handlers.NewHandler(*repos)

	router := gin.Default()
	router.POST("/login", handler.Login)

	router.Run(":" + port)
}
