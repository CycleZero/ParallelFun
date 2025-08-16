package biz

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
