package handler

import (
	"xorm.io/xorm"
)

type (
	Handler struct {
		DB *xorm.Engine
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
