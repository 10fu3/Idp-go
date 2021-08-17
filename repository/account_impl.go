package repository

import (
	"fmt"
	"toufu3.jp/Idp/model"
)

func (r *AllDB) CreateAccountByTempAccount(account model.ITempAccount) error {

	if r.GetAccountByMail(account.GetMail()) != nil{
		return fmt.Errorf("exist mail address")
	}

	m := account.ConvertAccount()
	return r.db.Create(m).Error
}

func (r *AllDB) CreateAccount(account model.IAccount) error {
	if r.GetAccountByMail(account.GetMail()) != nil{
		return fmt.Errorf("exist mail address")
	}
	return r.db.Create(account).Error
}

func (r *AllDB) GetAccount(uuid string) model.IAccount {
	var a model.Account
	r.db.Select("uuid = ?",uuid).Find(&a)
	return &a
}

func (r *AllDB) GetAccountByMail(mail string) model.IAccount {
	var a model.Account
	r.db.Select("mail = ?",mail).Find(&a)
	return &a
}

func (r *AllDB) GetAllAccounts() []model.IAccount {
	var a []model.Account
	r.db.Find(&a)
	var b = make([]model.IAccount,5)

	for _, v := range a {
		b = append(b, &v)
	}

	return b
}

func (r *AllDB) UpdateAccount(account model.IAccount) error {
	return r.db.Save(*account.(*model.Account)).Error
}

func (r *AllDB) DeleteAccount(uuid string) {
	r.db.Unscoped().Delete(model.Account{},"uuid = ?",uuid)
}
