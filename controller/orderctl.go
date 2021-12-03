package controller

import (
	"fmt"
	"html/template"
	"strconv"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func orderPreviewHandler(c *gin.Context) {
	_t := template.Must(template.ParseFiles(HTMLPath + "Order.html"))

	_userName := c.MustGet(gin.AuthUserKey).(string)

	// try to get cookie
	_authCookie, _ := c.Request.Cookie(AUTHKEY)
	fmt.Println(_authCookie.Value)
	if _authCookie == nil {
		fmt.Println("Authkey is not stored in cookie.")
		return
	}

	_mealId, e := strconv.Atoi(c.Params.ByName("id"))
	if e != nil {
		fmt.Printf("url error when precessing mealId.")
		return
	}
	_mealInfo := service.MealService.GetMeal(uint32(_mealId))
	_candiAccepter := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		_thisUserID, _ := strconv.Atoi(_authCookie.Value)
		return _user.UserID != uint32(_thisUserID)
	})

	_orderInfo := [...]interface{}{_userName, _mealInfo, _candiAccepter}

	_t.Execute(c.Writer, _orderInfo)
}

func orderApplyHandler(c *gin.Context) {

}
