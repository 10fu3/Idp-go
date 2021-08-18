package service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"toufu3.jp/Idp/model"
)

type CreateRequest struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Redirect    string `json:"redirect_uri"`
	Description string `json:"description"`
	OpenID      bool   `json:"has_openid"`
	Profile     bool   `json:"has_profile"`
	Mail        bool   `json:"has_mail"`
}

func (c *CreateRequest) Convert(createdBy model.IAccount) (model.Service, error) {

	if c.Name == "" {
		return model.Service{}, fmt.Errorf("need_%s", "name")
	}
	if c.Redirect == "" {
		return model.Service{}, fmt.Errorf("need_%s", "redirect")
	}
	if c.Description == "" {
		return model.Service{}, fmt.Errorf("need_%s", "description")
	}

	return model.Service{
		Model:       gorm.Model{},
		Uuid:        uuid.NewString(),
		Name:        c.Name,
		Secret:      uuid.NewString(),
		Redirect:    c.Redirect,
		Avatar:      c.Avatar,
		Description: c.Description,
		AdminUUID:   []string{createdBy.GetUUID()},
		ServiceScope: model.ServiceScope{
			Model:   gorm.Model{},
			OpenID:  c.OpenID,
			Profile: c.Profile,
			Mail:    c.Mail,
		},
	}, nil
}
