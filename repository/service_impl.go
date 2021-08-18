package repository

import "toufu3.jp/Idp/model"

func (r *AllDB) CreateService(s model.IService) error {
	return r.db.Create(*s.(*model.Service)).Error
}

func (r *AllDB) UpdateService(s model.IService) error {
	return r.db.Save(*s.(*model.Service)).Error
}

func (r *AllDB) GetService(uuid string) model.IService {
	var d []model.Service
	var count int64 = 0
	r.db.Select("uuid = ?", uuid).Find(&d).Count(&count)
	if count == 0 {
		return nil
	}
	return &d[0]
}

func (r *AllDB) GetAllServices() []model.IService {
	var d []model.Service
	var result = make([]model.IService, 1)
	r.db.Find(&d)
	for _, v := range d {
		result = append(result, &v)
	}
	return result
}

func (r *AllDB) GetServiceByAdmin(admin string) []model.IService {
	var d []model.Service
	var result = make([]model.IService, 1)
	r.db.Find(&d)
	for _, v := range d {
		for _, serviceAdmin := range v.GetAdminUUID() {
			if admin == serviceAdmin {
				result = append(result, &v)
			}
		}
	}
	return result
}

func (r *AllDB) DeleteService(uuid string) {
	panic("implement me")
}
