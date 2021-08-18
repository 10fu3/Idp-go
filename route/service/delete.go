package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/request/service"
)

func DeleteHandler(r repository.DB, u model.IAccount, ctx *gin.Context) {
	var req service.DeleteRequest
	if ctx.BindJSON(&req) == nil || req.Uuid == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "need_uuid",
		})
		return
	}

	notFound := func() {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "not_found",
		})
	}

	found := func() {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}

	target := r.(repository.IServiceRepository).GetService(req.Uuid)

	if target == nil {
		notFound()
		return
	}

	for _, v := range target.GetAdminUUID() {
		if v == u.GetUUID() {
			found()
			return
		}
	}
	notFound()

}
