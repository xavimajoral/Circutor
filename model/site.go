package model

import "time"

type Site struct {
	ID         uint `xorm:"pk autoincr"`
	LocationId string
	UserID     uint `xorm:"user_id"`
	//User       User
	CreatedAt time.Time
	UpdatedAt time.Time
}
