package model

import "gorm.io/gorm"

type IService interface {
	GetName() string
	GetUUID() string
	GetSecret() string
	GetRedirect() string
	GetAvatar() string
	GetDescription() string
	GetAdminUUID() []string
	GetScope() IServiceScope
}
type Service struct {
	gorm.Model
	Uuid string
	Name string
	Secret string
	Redirect string
	Avatar string
	Description string
	AdminUUID []string
	ServiceScope
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) GetUUID() string {
	return s.Uuid
}

func (s *Service) GetSecret() string {
	return s.Secret
}

func (s *Service) GetRedirect() string {
	return s.Redirect
}

func (s *Service) GetAvatar() string {
	return s.Avatar
}

func (s *Service) GetDescription() string {
	return s.Description
}

func (s *Service) GetAdminUUID() []string {
	return s.AdminUUID
}

func (s *Service) GetScope() IServiceScope {
	return &s.ServiceScope
}

type IServiceScope interface{
	HasOpenID() bool
	HasProfile() bool
	HasMail() bool
}

type ServiceScope struct {
	gorm.Model
	OpenID bool
	Profile bool
	Mail bool
}

func (s *ServiceScope) HasOpenID() bool {
	return s.OpenID
}

func (s *ServiceScope) HasProfile() bool {
	return s.Profile
}

func (s *ServiceScope) HasMail() bool {
	return s.Mail
}



