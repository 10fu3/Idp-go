package me

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
)

func GetHandler(r repository.DB, u model.IAccount,ctx *gin.Context){
	ctx.JSON(http.StatusOK,u)
}
