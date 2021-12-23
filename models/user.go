package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	ProfileFor   string
	Religion     string
	Community    string
	Gender       string
	Email        string `gorm:"unique"`
	MobileNumber int64  `gorm:"unique"`
	Password     string
	DOB          time.Time
	Country      string
}

type UserDetailedProfile struct {
	UserID            uint `gorm:"primaryKey"`
	City              string
	MartialStatus     string
	Diet              string
	Height            int
	Age               int
	Religion          string
	Caste             string
	CasteNoBar        bool
	Qualification     string
	CollegeName       string
	WorkAt            string
	Job               string
	YearlyIncomeLower int
	YearlyIncomeUpper int
	AboutUser         string
	ImageUrls         string
	IDProofUrls       string
	ProfileVerified   bool
}
