package models

import "gorm.io/gorm"

// User the user of our application
type User struct {
	gorm.Model

	Email string `gorm:"type:varchar(100);unique_index"`
}
