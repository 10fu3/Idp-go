package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/request/service"
)

func CreateHandler(r repository.DB, u model.IAccount, ctx *gin.Context) {
	var req service.CreateRequest
	if ctx.BindJSON(&req) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "convert_error",
		})
		return
	}

	s, err := req.Convert(u)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if r.(repository.IServiceRepository).CreateService(&s) != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "internal_error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
