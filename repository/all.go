package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Env struct {
	DB_Host           string `required:"true"`
	DB_Port           int    `required:"true"`
	DB_User           string `required:"true"`
	DB_Password       string `required:"true"`
	DB_Name           string `required:"true"`
	App_Port          int    `required:"true"`
	App_Client_Secret string `required:"true"`
	App_Client_ID     string `required:"true"`
}

type DbSettings struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (s *DbSettings) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Host, s.Port, s.DbName)
}

func (e *Env) ConvertDBConf() *DbSettings {
	return &DbSettings{
		Host:     e.DB_Host,
		Port:     e.DB_Port,
		User:     e.DB_User,
		Password: e.DB_Password,
		DbName:   e.DB_Name,
	}
}

type AllDB struct {
	db *gorm.DB
}

type DB interface {
	IAccountRepository
	ITempAccountRepository
}

func ConnectDB(dsn string) (DB,error) {
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil,err
	}
	return &AllDB{db: db},nil
}