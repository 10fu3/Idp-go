package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/request/service"
)

func UpdateHandler(r repository.DB, u model.IAccount, ctx *gin.Context) {
	var req service.UpdateRequest
	if ctx.BindJSON(&req) != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "bad_request",
		})
		return
	}

	uuid := ctx.Param("uuid")

	if uuid == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "need_uuid",
		})
		return
	}

	baseService := r.(repository.IServiceRepository).GetService(uuid)

	if baseService == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "not_found",
		})
		return
	}

	converted := req.Convert(baseService)

	if err := r.(repository.IServiceRepository).UpdateService(&converted); err != nil {
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
