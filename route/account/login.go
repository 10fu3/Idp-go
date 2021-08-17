package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/model/request"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/util"
)

func LoginHandler(r repository.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.LoginRequest
		ctx.BindJSON(req)
		account := r.(repository.IAccountRepository).GetAccountByMail(req.Mail)
		if account == nil{
			ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
				"status":"error",
				"message":"bad_params",
			})
			return
		}
		if util.PasswordVerify(account.GetPassword(),req.Password) == nil{

			session,err := r.(repository.ILoginSessionStore).CreateLoginSession(account)

			if err != nil{
				ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
					"status":"error",
				})
				return
			}

			ctx.JSON(http.StatusOK,gin.H{
				"session" : session,
			})
		}
	}
}
