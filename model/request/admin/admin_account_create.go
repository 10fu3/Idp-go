package admin

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"toufu3.jp/Idp/model"
)

type AccountCreateRequest struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	Bio string `json:"bio"`
}

func (a *AccountCreateRequest) Convert() model.IAccount {
	return &model.Account{
		Model:            gorm.Model{},
		Uuid:             uuid.NewString(),
		Name:             a.Name,
		Mail:             a.Mail,
		Password:         a.Password,
		Avatar:           a.Avatar,
		Bio:              a.Bio,
		IAccountAttribute: model.IAccountAttribute{
			Model:  gorm.Model{},
			Admin:  false,
			Active: true,
		},
	}
}
