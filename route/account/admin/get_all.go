package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
)

type simpleAccountInfo struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	CreatedAt int64 `json:"created_at"`
}

func convertSimpleInfo(a model.IAccount) simpleAccountInfo {
	return simpleAccountInfo{
		Uuid:      a.GetUUID(),
		Name:      a.GetName(),
		Avatar:    a.GetAvatar(),
		CreatedAt: a.GetCreatedAt().Unix(),
	}
}

func GetAllHandler(r repository.DB, u model.IAccount,ctx *gin.Context) {

	all := r.(repository.IAccountRepository).GetAllAccounts()

	result := make([]simpleAccountInfo,5)

	for _,v := range all{
		result = append(result,convertSimpleInfo(v))
	}

	ctx.JSON(http.StatusOK,result)
}
