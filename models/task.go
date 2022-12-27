package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title      string `gorm:"type:varchar(100);not null;unique_index"`
	Descripton string
	Done       bool `gorm:"default:false"`
	UserID     uint `json:"user_id"`
}
