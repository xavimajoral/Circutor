package model

import "time"

type Site struct {
	ID         uint
	LocationId string
	UserID     uint
	//User       User
	CreatedAt time.Time
	UpdatedAt time.Time
}
