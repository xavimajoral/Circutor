package model

import (
	"time"
)

type UserSignup struct {
	Email    string
	Password string
}

type User struct {
	ID        uint `swaggerignore:"true"`
	Email     string
	Password  string
	Token     string
	Sites     []Site
	CreatedAt time.Time
	UpdatedAt time.Time
}
