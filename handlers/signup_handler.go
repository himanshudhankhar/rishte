package handlers

import (
	"context"
	"errors"
	"net/mail"
	"regexp"
	"rishte/domain"
	"rishte/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func (h Handler) SignUp(c *gin.Context) {
	var signup domain.User
	err := c.BindJSON(&signup)

	if err != nil {
		logger.Default.Error(context.Background(), "err while signup bindjson", err.Error())
		c.JSON(500, gin.H{"success": false, "error": "user details are invalid" + err.Error()})
		return
	}

	if signup.FirstName == "" {
		logger.Default.Error(context.Background(), "err first name is empty")
		c.JSON(400, gin.H{"success": false, "error": "first name cannot be empty"})
		return
	}

	if signup.LastName == "" {
		logger.Default.Error(context.Background(), "err last name is empty")
		c.JSON(400, gin.H{"success": false, "error": "last name cannot be empty"})
		return
	}

	if signup.ProfileFor == "" {
		logger.Default.Error(context.Background(), "err profile created for is empty")
		c.JSON(400, gin.H{"success": false, "error": "profile created for cannot be empty"})
		return
	}

	if signup.Community == "" {
		logger.Default.Error(context.Background(), "err community is empty")
		c.JSON(400, gin.H{"success": false, "error": "community cannot be empty"})
		return
	}

	if signup.Country == "" {
		logger.Default.Error(context.Background(), "err country is empty")
		c.JSON(400, gin.H{"success": false, "error": "country cannot be empty"})
		return
	}

	if signup.Religion == "" {
		logger.Default.Error(context.Background(), "err religion is empty")
		c.JSON(400, gin.H{"success": false, "error": "religion cannot be empty"})
		return
	}

	if signup.Gender == "" {
		logger.Default.Error(context.Background(), "err gender is empty")
		c.JSON(400, gin.H{"success": false, "error": "gender cannot be empty"})
		return
	}

	email, validMail := h.validMailAddress(signup.Email)
	if email == "" || !validMail {
		logger.Default.Error(context.Background(), "err email is not valid")
		c.JSON(400, gin.H{"success": false, "error": "email is invalid"})
		return
	} else {
		signup.Email = email
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if signup.MobileNumber == 0 || !re.MatchString(strconv.FormatInt(signup.MobileNumber, 10)) {
		logger.Default.Error(context.Background(), "err mobile number is not valid")
		c.JSON(400, gin.H{"success": false, "error": "mobile number is invalid"})
		return
	}

	if signup.Password == "" {
		logger.Default.Error(context.Background(), "err password is not present")
		c.JSON(400, gin.H{"success": false, "error": "password cannot be empty"})
		return
	}

	dob := signup.DOB
	dobValid := h.parseDate(strings.Split(dob, "-"))

	if !dobValid {
		logger.Default.Error(context.Background(), "err dob is invalid")
		c.JSON(400, gin.H{"success": false, "error": "dob is not valid"})
		return
	}

	parsedUser, err := h.parseUser(signup)

	if err != nil {
		logger.Default.Error(context.Background(), "err while parsing user")
		c.JSON(500, gin.H{"success": false, "error": "Error while parsing User"})
		return
	}

	user, err := h.UserRepo.UserRepo.CreateUser(parsedUser)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			logger.Default.Error(c.Request.Context(), "err received while creating user email already in use", err.Error())
			c.JSON(409, gin.H{"success": false, "error": "Email already in use"})
			return
		}else {
		logger.Default.Error(c.Request.Context(), "err received while creating user ", err.Error())
		c.JSON(409, gin.H{"success": false, "error": "Error while creating User" + err.Error()})
		return
		}
	}

	c.JSON(200, gin.H{"success": true, "error": "", "userID": user.ID})
	logger.Default.Warn(c.Request.Context(), "user", user)
}

func (h Handler) validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func (h Handler) parseDate(dates []string) bool {
	if len(dates) != 3 {
		return false
	}
	if len(dates[0]) != 4 {
		return false
	}
	if len(dates[1]) < 1 || len(dates[1]) > 2 {
		return h.validateMonth(dates[1])
	}
	if len(dates[2]) < 1 || len(dates[2]) > 2 {
		return h.validateDay(dates[2])
	}

	_, err := time.Parse("2006-01-02", dates[0]+"-"+dates[1]+"-"+dates[2])

	return err == nil
}

func (h Handler) validateMonth(month string) bool {
	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return false
	}
	if monthInt < 1 || monthInt > 12 {
		return false
	}
	return true
}

func (h Handler) validateDay(month string) bool {
	dayInt, err := strconv.Atoi(month)
	if err != nil {
		return false
	}
	if dayInt < 1 || dayInt > 31 {
		return false
	}
	return true
}

func (h Handler) parseUser(user domain.User) (models.User, error) {
	var modelUser models.User
	modelUser.FirstName = user.FirstName
	modelUser.LastName = user.LastName
	modelUser.Email = user.Email
	modelUser.Gender = user.Gender
	modelUser.Community = user.Community
	modelUser.Country = user.Country
	modelUser.MobileNumber = user.MobileNumber
	modelUser.Password = user.Password
	modelUser.ProfileFor = user.ProfileFor
	modelUser.Religion = user.Religion
	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		logger.Default.Error(context.Background(), "error while parsing dob", err.Error())
		return modelUser, errors.New("dob time parsing err")
	}
	modelUser.DOB = dob
	return modelUser, nil
}
