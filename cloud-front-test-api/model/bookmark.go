package model

import "time"

type Bookmark struct {
	ID         uint `xorm:"pk autoincr"`
	BuildingId string
	UserID     uint `xorm:"user_id"`
	//User       User
	CreatedAt time.Time `xorm:"updated"`
	UpdatedAt time.Time `xorm:"created"`
}

type addBookmark struct {
	BuildingId string
}
