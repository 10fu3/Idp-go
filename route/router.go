package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/route/account"
	"toufu3.jp/Idp/route/account/admin"
	"toufu3.jp/Idp/route/account/me"
	"toufu3.jp/Idp/route/entry"
)

func setupOAuthRouter(r repository.DB, authRouter *gin.RouterGroup) {

}

func setupAccountAdminRouter(r repository.DB,accountRouter *gin.RouterGroup) {
	adminRouter := accountRouter.Group("/admin")
	{
		adminRouter.GET("/all", AdminFilter(r,admin.GetAllHandler))
		adminRouter.GET("/:uuid", AdminFilter(r,admin.GetHandler))
		adminRouter.POST("/:uuid", AdminFilter(r,admin.CreateHandler))
		adminRouter.PUT("/:uuid", AdminFilter(r,admin.UpdateHandler))
		adminRouter.DELETE("/:uud", AdminFilter(r,admin.DeleteHandler))
	}
}

func setupLoginLogoutRouter(r repository.DB,accountRouter *gin.RouterGroup) {
	accountRouter.POST("/login", account.LoginHandler(r))

	//need login
	accountRouter.POST("/logout", account.LogoutHandler(r))
}

func setupRegisterRouter(r repository.DB,accountRouter *gin.RouterGroup) {
	//C self first
	accountRouter.POST("/entry", entry.Handler(r))
	//C self confirmed mail
	accountRouter.POST("/entry/:uuid",entry.ConfirmHandler(r))
}

func setupAccountManagementRouter(r repository.DB,accountRouter *gin.RouterGroup){
	managementRouter := accountRouter.Group("/me")

	//R me / need login
	managementRouter.GET("", LoginFilter(r,me.GetHandler))
	//U me / need login
	managementRouter.PUT("", LoginFilter(r,me.UpdateHandler))
	//D me / need login
	managementRouter.DELETE("", LoginFilter(r,me.DeleteHandler))
}

func setupAccountRouter(r repository.DB,apiRouter *gin.RouterGroup) {
	accountRouter := apiRouter.Group("/account")
	setupRegisterRouter(r,accountRouter)
	setupLoginLogoutRouter(r,accountRouter)
	setupAccountManagementRouter(r,accountRouter)
	setupAccountAdminRouter(r,accountRouter)
}


func Setup(env repository.Env, r repository.DB) {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"heart beat")
	})

	authRouter := router.Group("/oauth2/v1")

	apiRouter := router.Group("/api/v1")

	//TODO
	//OAuth IDToken回りの実装
	setupOAuthRouter(r,authRouter)
	setupAccountRouter(r,apiRouter)

	if err := router.Run(fmt.Sprintf(":%d",env.App_Port));err != nil{
		fmt.Println(err.Error())
	}

}
