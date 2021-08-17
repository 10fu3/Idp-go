package model

import (
	"gorm.io/gorm"
	"time"
)

type ITempAccount interface {
	GetUUID() string
	GetName() string
	GetMail() string
	GetPassword() string
	GetExpiredAt() int64
	ConvertAccount() IAccount
}

type TempAccount struct {
	gorm.Model
	Uuid      string
	Name      string `form:"name"`
	Mail      string `form:"mail"`
	Password  string `form:"password"`
	ExpiredAt int64
}

func (t *TempAccount) GetUUID() string {
	return t.Uuid
}

func (t *TempAccount) GetName() string {
	return t.Name
}

func (t *TempAccount) GetMail() string {
	return t.Mail
}

func (t *TempAccount) GetPassword() string {
	return t.Password
}

func (t *TempAccount) GetExpiredAt() int64 {
	return t.ExpiredAt
}

func (t *TempAccount) ConvertAccount() IAccount {
	return &Account{
		Model: gorm.Model{
			ID:        t.ID,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: t.DeletedAt,
		},
		Uuid:               t.Uuid,
		Name:             t.Name,
		Mail:             t.Mail,
		Password:         t.Password,
		Avatar:           "https://i.imgur.com/R6tktJ6.jpg", //人の顔
		Bio:              "",
		AccountAttribute: AccountAttribute{
			Model:  gorm.Model{},
			Admin:  false,
			Active: true,
		},
	}
}

type Account struct {
	gorm.Model
	Uuid        string `json:"uuid" gorm:"column:uuid"`
	Name      string `json:"name" gorm:"column:name"`
	Mail      string `json:"mail" gorm:"column:mail"`
	Password  string `json:"password" gorm:"column:password"`
	Avatar    string `json:"avatar" gorm:"column:avatar"`
	Bio       string `json:"bio" gorm:"column:bio"`
	AccountAttribute
}

func (a *Account) GetModel() gorm.Model {
	return a.Model
}

func (a *Account) GetMail() string {
	return a.Mail
}

func (a *Account) GetUUID() string {
	return a.Uuid
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) GetPassword() string {
	return a.Password
}

func (a *Account) GetAvatar() string {
	return a.Avatar
}

func (a *Account) GetBio() string {
	return a.Bio
}

func (a *Account) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *Account) GetAttribute() IAccountAttribute {
	return &a.AccountAttribute
}

type AccountAttribute struct {
	gorm.Model
	Admin  bool `json:"is_admin" gorm:"column:is_admin"`
	Active bool `json:"is_active" gorm:"column:is_active"`
}

func (a *AccountAttribute) IsActive() bool {
	return a.Active
}

func (a *AccountAttribute) IsAdmin() bool {
	return a.Active
}
