package service

import "toufu3.jp/Idp/model"

type UpdateRequest struct {
	Name        *string `json:"name"`
	Avatar      *string `json:"avatar"`
	Redirect    *string `json:"redirect_uri"`
	Description *string `json:"description"`
	OpenID      *bool   `json:"has_openid"`
	Profile     *bool   `json:"has_profile"`
	Mail        *bool   `json:"has_mail"`
}

func (u *UpdateRequest) Convert(base model.IService) model.Service {
	m := *base.(*model.Service)

	if u.Name != nil {
		m.Name = *u.Name
	}
	if u.Avatar != nil {
		m.Avatar = *u.Avatar
	}
	if u.Redirect != nil {
		m.Redirect = *u.Redirect
	}
	if u.Description != nil {
		m.Description = *u.Description
	}
	if u.OpenID != nil {
		m.OpenID = *u.OpenID
	}
	if u.Profile != nil {
		m.Profile = *u.Profile
	}
	if u.Mail != nil {
		m.Mail = *u.Mail
	}
	return m
}
