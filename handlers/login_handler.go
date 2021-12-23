package handlers

import (
	"context"
	"fmt"
	"rishte/domain"
	"rishte/repository"
	"rishte/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

type Handler struct {
	UserRepo repository.Repositories
}

func NewHandler(repos repository.Repositories) Handler {
	return Handler{
		UserRepo: repos,
	}
}

func (h Handler) Login(c *gin.Context) {
	var login domain.Login
	err := c.BindJSON(&login)
	if err != nil {
		logger.Default.Error(context.Background(), err.Error())
		c.JSON(400, gin.H{"error": "username or password not present"})
		return
	}
	if login.Username == "" {
		logger.Default.Error(context.Background(), "username empty")
		c.JSON(400, gin.H{"error": "username not present"})
		return
	}

	if login.Password == "" {
		logger.Default.Error(context.Background(), "password empty")
		c.JSON(400, gin.H{"error": "password not present"})
		return
	}

	loginService := services.NewLoginService(h.UserRepo)

	user, err := loginService.LoginByEmail(login.Username, login.Password)
	if err == nil && user.Username != "" && user.Password != "" {
		c.JSON(200, gin.H{"success": true, "username": user.Username})
		logger.Default.Info(context.Background(), "user found", user.Username)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	if err.Error() == "password invalid" {
		c.JSON(401, gin.H{"success": false, "error": "Password Invalid"})
		logger.Default.Info(context.Background(), "password invalid", login.Username)
		return
	}

	user, err = loginService.LoginByPhone(login.Username, login.Password)
	if err == nil && user.Username != "" && user.Password != "" {
		c.JSON(200, gin.H{"success": true, "username": user.Username})
		logger.Default.Info(context.Background(), "user found", user.Username)
		return
	}

	if err.Error() == "password invalid" {
		c.JSON(401, gin.H{"success": false, "error": "Password Invalid"})
		logger.Default.Info(context.Background(), "password invalid", login.Username)
		return
	}

	if err != nil || user == nil || user.Username == "" || user.Password != "" {
		c.JSON(401, gin.H{"success": false, "error": "User Not Found"})
		logger.Default.Info(context.Background(), "user not found", login.Username)
		return
	}
}
