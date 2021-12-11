package controller

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes() *gin.Engine {
	ret := gin.Default()
	ret.LoadHTMLGlob(HTMLPath + "*.html")

	customers := ret.Group("/", gin.BasicAuth(getAccounts(AUTHPATH)))
	customers.GET("/menu", menuHandler)
	customers.GET("/menu/:"+MEALKEY, orderPreviewHandler)
	customers.POST("/order", orderApplyHandler)

	users := customers.Group("user", userPremissionInterceotor)
	users.GET(":userid", userHandler)
	users.POST(":userid", userActionsHandler)
	users.GET(":userid/sync", userOrderPresentor)
	users.GET(":userid/pwchange", userPasswordChanging)
	users.POST(":userid/pwchange", onUserPasswordChanged)

	return ret
}
