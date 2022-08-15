package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Emailgorm string `gorm:"not null;unique_index"`
	Tasks     []Task
}