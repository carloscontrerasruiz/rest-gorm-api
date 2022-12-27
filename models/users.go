package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Firstname string `gorm:"not null;default:null" json:"first_name"`
	Lastname  string `gorm:"not null;default:null" json:"lastname"`
	Email     string `gorm:"not null;default:null;unique" json:"email"`
	Tasks     []Task `json:"tasks"`
}
