package model

type ILoginSession interface {
	GetUUID() string
	GetExpiredAt() int64
	GetAccountID() string
}

type LoginSession struct {
	Uuid string `gorm:"column:uuid"`
	ExpiredAt int64 `gorm:"column:expired_at"`
	AccountId string `gorm:"column:account_id"`
}

func (l *LoginSession) GetUUID() string {
	return l.Uuid
}

func (l *LoginSession) GetExpiredAt() int64 {
	return l.ExpiredAt
}

func (l *LoginSession) GetAccountID() string {
	return l.AccountId
}
