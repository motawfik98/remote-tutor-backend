package models

import (
	"gorm.io/gorm"
)

// User struct to store the user data
type User struct {
	gorm.Model
	Username     string `gorm:"unique" json:"username"`
	Password     string `json:"password"`
	FullName     string `json:"fullName"`
	Activated    bool   `gorm:"default:0" json:"activated"`
	Admin        bool   `gorm:"default:0" json:"admin"`
	Year         int    `json:"year"`
	PhoneNumber  string `json:"phoneNumber"`
	ParentNumber string `json:"parentNumber"`
}
