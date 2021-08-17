package me

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/model/request"
	"toufu3.jp/Idp/repository"
)

func UpdateHandler(r repository.DB, u model.IAccount,ctx *gin.Context){

	var json request.AccountUpdateRequest

	ctx.BindJSON(&json)

	err := r.(repository.IAccountRepository).UpdateAccount(json.Convert(u))

	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"message":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
	})

}
