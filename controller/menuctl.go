package controller

import (
	"html/template"
	"strconv"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func menuHandler(c *gin.Context) {
	_userName := c.MustGet(gin.AuthUserKey).(string)
	_authUser := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		return _user.Name == _userName
	})[0] // Must exited

	_t := template.Must(template.ParseFiles(HTMLPath + "Menu.html"))

	_firstFloor := service.MealService.Select(0, floorSelector(1))
	_secondFloor := service.MealService.Select(0, floorSelector(2))
	_menuItems := [...]interface{}{_authUser, _firstFloor, _secondFloor, ACCEPTERID}

	// set cookie
	setCookies(c, map[string]string{
		AUTHKEY: strconv.Itoa(int(_authUser.(models.User).UserID))})

	_t.Execute(c.Writer, _menuItems)
}
