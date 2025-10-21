package biz

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     Role
	GameId   string
}

type Role int

const (
	Unknown Role = iota
	Admin
	Default
	Guest
)
