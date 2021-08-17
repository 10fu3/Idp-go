package repository

import "toufu3.jp/Idp/model"

type ILoginSessionStore interface {
	CreateLoginSession(account model.IAccount) (string,error)
	GetLoginSession(session string) model.ILoginSession
	HasLoginSession(session string,account model.IAccount) bool
	DeleteLoginSession(session string)
	DeleteLoginSessionByAccountID(uuid string)
	UpdateLoginSession(session string,account model.IAccount) bool
}