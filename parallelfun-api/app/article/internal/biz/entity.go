package biz

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string
	Content  string
	AuthorID uint64
}

type Author struct {
	ID   uint64
	Name string
}
