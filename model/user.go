package model

import (
	"time"
)

type UserSignup struct {
	Email    string
	Password string
}

type User struct {
	ID        uint   `swaggerignore:"true" xorm:"autoincr pk"`
	Email     string `xorm:"unique"`
	Password  string
	Token     string
	Sites     []Site
	CreatedAt time.Time
	UpdatedAt time.Time
}
