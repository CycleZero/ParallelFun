package biz

import (
	"gorm.io/gorm"
)

type Server struct {
	gorm.Model
	Name    string
	Address string
	Port    uint
	Status  int
	OwnerId uint

	Avatar      string
	Cover       string
	Description string
	Tags        []string
}

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
