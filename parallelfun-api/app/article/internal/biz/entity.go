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

type VideoPost struct {
	gorm.Model
	Title     string
	Content   string
	AuthorID  uint64
	VideoLink string
}

type Comment struct {
	gorm.Model
	Content   string
	AuthorID  uint64
	ArticleID uint64
}

type BatchSelectParam struct {
	PageNum  int
	PageSize int
	IDs      []uint64
	Title    string
	AuthorID uint64
}

type Author struct {
	ID   uint64
	Name string
}
