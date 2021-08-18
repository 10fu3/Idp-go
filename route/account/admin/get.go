package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/request/account/admin"
)

type fullAccountInfo struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
	model.IAccountAttribute
}

func convertFullInfo(account model.IAccount) fullAccountInfo {
	return fullAccountInfo{
		Uuid:              account.GetUUID(),
		Name:              account.GetName(),
		Mail:              account.GetMail(),
		Avatar:            account.GetAvatar(),
		Bio:               account.GetBio(),
		IAccountAttribute: account.GetAttribute(),
	}
}

func GetHandler(r repository.DB, u model.IAccount, ctx *gin.Context) {
	var req admin.AccountTargetRequest

	if ctx.BindJSON(req) != nil || req.Uuid == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "need_uuid",
		})
		return
	}

	account := r.(repository.IAccountRepository).GetAccount(req.Uuid)

	if account == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "not_exist_account",
		})
		return
	}

	ctx.JSON(http.StatusOK, convertFullInfo(account))

}
