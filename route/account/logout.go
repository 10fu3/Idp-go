package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"toufu3.jp/Idp/repository"
)

func LogoutHandler(r repository.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authN := ctx.GetHeader("Authorization")
		if authN == ""{
			authN = ctx.GetHeader("authorization")
		}

		split := strings.Split(authN," ")

		if len(split) != 2{
			ctx.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
				"status":"error",
				"message":"bad_header",
			})
			return
		}
		r.(repository.ILoginSessionStore).DeleteLoginSession(split[1])
		ctx.JSON(http.StatusOK,gin.H{
			"status":"success",
		})
	}
}
