package account

import (
	"github.com/gin-gonic/gin"
	"toufu3.jp/Idp/repository"
	"toufu3.jp/Idp/route"
	"toufu3.jp/Idp/route/account/admin"
	"toufu3.jp/Idp/route/account/me"
	"toufu3.jp/Idp/route/entry"
)

func setupAccountAdminRouter(r repository.DB, accountRouter *gin.RouterGroup) {
	adminRouter := accountRouter.Group("/admin")
	{
		adminRouter.GET("/all", route.AdminFilter(r, admin.GetAllHandler))
		adminRouter.GET("/:uuid", route.AdminFilter(r, admin.GetHandler))
		adminRouter.POST("/:uuid", route.AdminFilter(r, admin.CreateHandler))
		adminRouter.PUT("/:uuid", route.AdminFilter(r, admin.UpdateHandler))
		adminRouter.DELETE("/:uud", route.AdminFilter(r, admin.DeleteHandler))
	}
}

func setupLoginLogoutRouter(r repository.DB, accountRouter *gin.RouterGroup) {
	accountRouter.POST("/login", LoginHandler(r))

	//need login
	accountRouter.POST("/logout", LogoutHandler(r))
}

func setupRegisterRouter(r repository.DB, accountRouter *gin.RouterGroup) {
	//C self first
	accountRouter.POST("/entry", entry.Handler(r))
	//C self confirmed mail
	accountRouter.POST("/entry/:uuid", entry.ConfirmHandler(r))
}

func setupAccountManagementRouter(r repository.DB, accountRouter *gin.RouterGroup) {
	managementRouter := accountRouter.Group("/me")

	//R me / need login
	managementRouter.GET("", route.LoginFilter(r, me.GetHandler))
	//U me / need login
	managementRouter.PUT("", route.LoginFilter(r, me.UpdateHandler))
	//D me / need login
	managementRouter.DELETE("", route.LoginFilter(r, me.DeleteHandler))
}

func SetupRouter(r repository.DB, apiRouter *gin.RouterGroup) {
	accountRouter := apiRouter.Group("/account")
	setupRegisterRouter(r, accountRouter)
	setupLoginLogoutRouter(r, accountRouter)
	setupAccountManagementRouter(r, accountRouter)
	setupAccountAdminRouter(r, accountRouter)
}
