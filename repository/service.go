package repository

import "toufu3.jp/Idp/model"

type IServiceRepository interface {
	CreateService(s model.IService) error
	UpdateService(s model.IService) error
	GetService(uuid string) model.IService
	GetAllServices() []model.IService
	GetServiceByAdmin(admin string) []model.IService
	DeleteService(uuid string)
}
