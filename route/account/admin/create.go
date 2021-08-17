package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/model/request/admin"
	"toufu3.jp/Idp/repository"
)

func CreateHandler(r repository.DB, u model.IAccount,ctx *gin.Context)  {

	var req admin.AccountCreateRequest

	ctx.BindJSON(req)

	err := r.(repository.IAccountRepository).CreateAccount(req.Convert())

	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
	})

}
