package me

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
)

func DeleteHandler(r repository.DB, u model.IAccount,ctx *gin.Context){
	r.(repository.IAccountRepository).DeleteAccount(u.GetUUID())

	r.(repository.ILoginSessionStore).DeleteLoginSessionByAccountID(u.GetUUID())

	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}
