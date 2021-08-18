package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
)

func GetAllHandler(r repository.DB, u model.IAccount, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r.(repository.IServiceRepository).GetServiceByAdmin(u.GetUUID()))
}
