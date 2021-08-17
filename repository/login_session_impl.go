package repository

import (
	"github.com/google/uuid"
	"time"
	"toufu3.jp/Idp/model"
)

func (r *AllDB) CreateLoginSession(account model.IAccount) (string,error) {
	session := model.LoginSession{
		Uuid:      uuid.NewString(),
		ExpiredAt: time.Now().AddDate(0,0,30).Unix(),
		AccountId: account.GetUUID(),
	}

	if err := r.db.Create(session).Error;err != err{
		return "",err
	}
	return session.Uuid,nil
}

func (r *AllDB) GetLoginSession(session string) model.ILoginSession {
	var savedSession model.LoginSession
	if s := r.db.Select("uuid = ?",session).Find(&savedSession);s.Error != nil{
		return nil
	}

	if savedSession.GetExpiredAt() < time.Now().Unix(){
		return nil
	}

	return &savedSession
}

func (r *AllDB) HasLoginSession(session string, account model.IAccount) bool {

	savedSession := r.GetLoginSession(session)

	if savedSession == nil{
		return false
	}

	if savedSession.GetAccountID() == account.GetUUID(){
		if savedSession.GetExpiredAt() >= time.Now().Unix(){
			return true
		}
		r.DeleteLoginSession(session)
		return false
	}
	return false
}

func (r *AllDB) DeleteLoginSession(session string) {
	r.db.Unscoped().Delete(&model.LoginSession{},"uuid = ?",session)
}

func (r *AllDB) DeleteLoginSessionByAccountID(uuid string) {
	r.db.Unscoped().Delete(&model.LoginSession{},"account_id = ?",uuid)
}

func (r *AllDB) UpdateLoginSession(session string,account model.IAccount) bool {
	if s := r.GetLoginSession(session); s != nil{

		if s.GetAccountID() == account.GetUUID(){
			update := model.LoginSession{
				Uuid:      s.GetUUID(),
				ExpiredAt: time.Now().AddDate(0, 1, 0).Unix(),
				AccountId: s.GetAccountID(),
			}
			if r.db.Save(update).Error != nil{
				return false
			}
			return true
		}

	}
	return false
}
