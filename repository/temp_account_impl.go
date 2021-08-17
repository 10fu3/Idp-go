package repository

import (
	"time"
	"toufu3.jp/Idp/model"
)

func (r *AllDB) CreateTempAccount(account model.ITempAccount) error {
	return r.db.Create(account).Error
}

func (r *AllDB) GetTempAccount(uuid string) (model.ITempAccount,error) {
	var u model.ITempAccount
	if e := r.db.Where("uuid = ?",uuid).First(&u).Error; e!=nil{
		return nil,e
	}
	return u,nil
}

func (r *AllDB) GetAllTempAccounts() ([]model.ITempAccount, error) {
	var u []model.ITempAccount
	if e := r.db.Find(&u).Error;e != nil{
		return nil,e
	}
	return u,nil
}

func (r *AllDB) DeleteTempAccount(uuid string) error {
	return r.db.Delete(model.Account{}, "uuid = ?",uuid).Error
}

func (r *AllDB) DeleteExpiredTempAccount() error {
	return r.db.Delete(model.Account{},"expired_at > ?",time.Now().Unix()).Error
}
