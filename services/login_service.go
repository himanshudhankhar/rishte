package services

import (
	// "errors"
	"context"
	"errors"
	"rishte/domain"
	"rishte/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/logger"
)

type LoginService interface {
	LoginByEmail(email string, password string) (*domain.Login, error)
	LoginByPhone(phone string, password string) (*domain.Login, error)
}

type loginService struct {
	userRepo repository.Repositories
}

func NewLoginService(repo repository.Repositories) loginService {
	return loginService{
		userRepo: repo,
	}
}

func (l *loginService) LoginByEmail(email string, password string) (*domain.Login, error) {
	User := l.userRepo.UserRepo.GetExistingUserByEmail(email)

	if User.FirstName == "" && User.LastName == "" {
		return nil, errors.New("email not found")
	}

	passwordHash := []byte(User.Password)
	err := bcrypt.CompareHashAndPassword(passwordHash, []byte(password))

	if err == nil {
		return &domain.Login{
			Username: email,
			Password: password,
		}, nil
	}

	if err != nil {
		logger.Default.Error(context.Background(), err.Error())
		return nil, errors.New("password invalid")
	}

	return nil, err
}

func (l *loginService) LoginByPhone(phone string, password string) (*domain.Login, error) {
	User := l.userRepo.UserRepo.GetExistingUserByPhone(phone)

	if User.FirstName == "" && User.LastName == "" {
		return nil, errors.New("phone not found")
	}

	passwordHash := []byte(User.Password)
	err := bcrypt.CompareHashAndPassword(passwordHash, []byte(password))

	if err == nil {
		return &domain.Login{
			Username: User.Email,
			Password: password,
		}, nil
	}

	if err != nil {
		logger.Default.Error(context.Background(), err.Error())
		return nil, errors.New("password invalid")
	}

	return nil, err
}
