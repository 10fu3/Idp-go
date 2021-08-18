package admin

import "toufu3.jp/Idp/model"

type AccountUpdateRequest struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
}

func (a *AccountUpdateRequest) Convert(source model.IAccount) model.IAccount {

	m := model.Account{}

	if a.Name != "" {
		m.Name = source.GetName()
	}
	if a.Password != "" {
		m.Password = source.GetPassword()
	}
	if a.Mail != "" {
		m.Mail = source.GetMail()
	}
	if a.Bio != "" {
		m.Bio = source.GetBio()
	}
	if a.Avatar != "" {
		m.Avatar = source.GetAvatar()
	}

	m.Model = source.GetModel()
	m.IAccountAttribute = source.GetAttribute()

	return &m
}
