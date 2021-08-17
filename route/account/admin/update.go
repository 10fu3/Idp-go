package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/model/request/admin"
	"toufu3.jp/Idp/repository"
)

func UpdateHandler(r repository.DB, u model.IAccount,ctx *gin.Context)  {

	var req admin.AccountUpdateRequest

	bindErr := ctx.BindJSON(&req)

	if bindErr != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":"bad_params",
		})
		return
	}

	if req.Uuid == ""{
		ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":"need_uuid",
		})
		return
	}

	account := r.(repository.IAccountRepository).GetAccount(req.Uuid)

	if account == nil{
		ctx.AbortWithStatusJSON(http.StatusNotFound,gin.H{
			"status":"error",
			"message":"not_exist_account",
		})
		return
	}

	if account.GetUUID() == u.GetUUID(){
		ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{
			"status":"error",
			"message":"use_/me",
		})
		return
	}

	if r.(repository.IAccountRepository).UpdateAccount(req.Convert(account)) != nil{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"message":"internal_error",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
	return

}
