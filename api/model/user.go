package model

import (
	"time"
)

type UserSignup struct {
	Email    string
	Password string
}

type LoginUser struct {
	ID    uint
	Email string
	Token string
}

type User struct {
	ID        uint   `swaggerignore:"true" xorm:"autoincr pk"`
	Email     string `xorm:"unique"`
	Password  string
	Token     string `xorm:"-"`
	Bookmarks []Bookmark
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
