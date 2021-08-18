package service

import (
	"github.com/gin-gonic/gin"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/route"
)

func SetupRouter(r repository.DB, apiRouter *gin.RouterGroup) {
	serviceRouter := apiRouter.Group("/service")

	serviceRouter.GET("", route.LoginFilter(r, GetAllHandler))
	serviceRouter.POST("", route.LoginFilter(r, CreateHandler))
	serviceRouter.PUT("/:uuid", route.LoginFilter(r, UpdateHandler))
	serviceRouter.DELETE("/:uuid", route.LoginFilter(r, DeleteHandler))
}
