package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/route/account"
)

func setupOAuthRouter(r repository.DB, authRouter *gin.RouterGroup) {

}

func Setup(env repository.Env, r repository.DB) {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "heart beat")
	})

	authRouter := router.Group("/oauth2/v1")

	apiRouter := router.Group("/api/v1")

	//TODO
	//OAuth IDToken回りの実装
	setupOAuthRouter(r, authRouter)
	account.SetupRouter(r, apiRouter)

	if err := router.Run(fmt.Sprintf(":%d", env.App_Port)); err != nil {
		fmt.Println(err.Error())
	}

}
