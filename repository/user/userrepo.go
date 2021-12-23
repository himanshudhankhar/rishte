package userrepo

import (
	"context"
	"rishte/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetExistingUserByEmail(email string) models.User {
	var user models.User
	repo.db.Where("email = ?", email).First(&user)
	return user
}

func (repo *UserRepo) GetExistingUserByPhone(phone string) models.User {
	var user models.User
	repo.db.Where("mobile_number = ?", phone).First(&user)
	return user
}

func (repo *UserRepo) CreateUser(user models.User) (models.User, error) {

	hashedPass, _ := repo.hashPassword(user.Password)
	user.Password = hashedPass

	result := repo.db.Create(&user)

	if result.Error != nil {
		logger.Default.Error(context.Background(), result.Error.Error())
		return user, result.Error
	}

	logger.Default.Info(context.Background(), "Inserted a user with ID:", user.ID)
	return user, nil
}

func (repo *UserRepo) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
