package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/model/request/admin"
	"toufu3.jp/Idp/repository"
)

func DeleteHandler(r repository.DB, u model.IAccount,ctx *gin.Context)  {

	var req admin.AccountTargetRequest

	if ctx.BindJSON(&req) != nil || req.Uuid == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":"need_uuid",
		})
		return
	}

	r.(repository.IAccountRepository).DeleteAccount(req.Uuid)
	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}
