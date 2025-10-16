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
	Role      Role
	GameId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role int

const (
	Unknown Role = iota
	Admin
	Default
	Guest
)
