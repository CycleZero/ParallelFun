package biz

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Name    string
	Address string
	Port    uint
	Status  int
	OwnerId uint
}
