package domain

type User struct {
	FirstName    string
	LastName     string
	ProfileFor   string
	Religion     string
	Community    string
	Gender       string
	Email        string `gorm:"unique"`
	MobileNumber int64  `gorm:"unique"`
	Password     string
	DOB          string
	Country      string
}
