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

	// customers.GET("/usr/:name", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if res, ok := getAccounts(AUTHPATH)[user]; ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "pw": res})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "pw": "not found"})
	// 	}
	// })

	return ret
}
