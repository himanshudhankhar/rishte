package repository

import (
	userrepo "rishte/repository/user"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo *userrepo.UserRepo
}

// InitRepositories should be called in main.go
func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := userrepo.NewUserRepo(db)
	return &Repositories{UserRepo: userRepo}
}
