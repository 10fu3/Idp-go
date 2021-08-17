package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"toufu3.jp/Idp/model"
	"toufu3.jp/Idp/repository"
)

func LoginFilter(r repository.DB,success func(r repository.DB,u model.IAccount,ctx *gin.Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authN := ctx.GetHeader("Authorization")
		if authN == ""{
			authN = ctx.GetHeader("authorization")
		}

		split := strings.Split(authN," ")

		if len(split) != 2{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"status":"error",
				"message":"bad_authorization_header",
			})
		}

		account := r.(repository.IAccountRepository).GetAccount(split[1])

		if account == nil{
			ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{
				"status":"error",
				"message":"bad_header",
			})
		}

		success(r,account,ctx)
	}
}

func AdminFilter(r repository.DB,success func(r repository.DB,u model.IAccount,ctx *gin.Context)) gin.HandlerFunc  {
	return LoginFilter(r, func(r repository.DB, u model.IAccount, ctx *gin.Context) {
		if u.GetAttribute().Admin{
			success(r,u,ctx)
		}else{
			ctx.AbortWithStatusJSON(http.StatusForbidden,gin.H{
				"status":"error",
				"message":"you_don't_have_permission",
			})
		}
	})
}
