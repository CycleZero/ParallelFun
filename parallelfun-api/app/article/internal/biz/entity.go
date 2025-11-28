package biz

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string
	Content    string
	AuthorID   uint64
	MediaInfos []MediaInfo `gorm:"-"`
}

type MediaInfo struct {
	gorm.Model
	OssId     string
	Type      string
	Link      string `gorm:"-"`
	ArticleID uint64
}

type Comment struct {
	gorm.Model
	Content   string
	AuthorID  uint64
	ArticleID uint64
	Likes     int
	Dislikes  int
}

type VideoPost struct {
	gorm.Model
	Title     string
	Content   string
	AuthorID  uint64
	VideoLink string
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
