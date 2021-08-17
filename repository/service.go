package repository

type IServiceRepository interface {
	CreateService()
	UpdateService()
	GetService(uuid string)
	DeleteService(uuid string)
}