package repository

import "toufu3.jp/Idp/model"

type ITempAccountRepository interface {
	CreateTempAccount(account model.ITempAccount)error
	GetTempAccount(uuid string) (model.ITempAccount,error)
	GetAllTempAccounts() ([]model.ITempAccount,error)
	DeleteTempAccount(uuid string)error
	DeleteExpiredTempAccount()error
}

type IAccountRepository interface {
	CreateAccountByTempAccount(account model.ITempAccount)error
	CreateAccount(account model.IAccount) error
	GetAccount(uuid string) model.IAccount
	GetAccountByMail(mail string) model.IAccount
	GetAllAccounts() []model.IAccount
	UpdateAccount(account model.IAccount)error
	DeleteAccount(uuid string)
}