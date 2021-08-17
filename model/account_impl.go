package model

import (
	"gorm.io/gorm"
	"time"
)

type IAccount interface {
	GetUUID() string
	GetMail() string
	GetName() string
	GetPassword() string
	GetAvatar() string
	GetBio() string
	GetModel() gorm.Model
	GetCreatedAt() time.Time
	GetAttribute() IAccountAttribute
}

type IAccountAttribute interface {
	IsActive() bool
	IsAdmin() bool
}
