package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string
	Password string
}

type Todos struct {
	gorm.Model
	UserId  uint
	title   string
	content string
}
