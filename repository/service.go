package repository

import "toufu3.jp/Idp/model"

type IServiceRepository interface {
	CreateService(s model.IService)
	UpdateService(s model.IService)
	GetService(uuid string) model.IService
	GetAllServices() []model.IService
	DeleteService(uuid string)
}
