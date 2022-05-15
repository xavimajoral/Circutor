package handler

import (
	"embed"
	"xorm.io/xorm"
)

type (
	Handler struct {
		DB        *xorm.Engine
		DataFiles embed.FS
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
