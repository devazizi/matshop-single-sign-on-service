package db

import "time"

type User struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name" gorm:"type:varchar(200);uniqueIndex`
	LastName     string `json:"last_name" gorm:"type:varchar(200);uniqueIndex`
	NationalCode string `json:"national_code" gorm:"type:varchar(10);uniqueIndex`
	Email        string `json:"email"  gorm:"type:varchar(250);uniqueIndex`
	CellNumber   string `json:"email"  gorm:"type:varchar(30);uniqueIndex`
	Password     string `gorm:"type:varchar(250);`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type OauthAccessToken struct {
	ID         uint   `json:"id"`
	Token      string `gorm:"type:varchar(200);uniqueIndex`
	UserId     uint
	Audience   string `gorm:"type:varchar(200);uniqueIndex`
	ExpiryDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
