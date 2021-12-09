package controller

import (
	"fmt"
	"html/template"
	"log"
	"strconv"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func menuHandler(c *gin.Context) {
	_userName := c.MustGet(gin.AuthUserKey).(string)
	_authUsers := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		return _user.Name == _userName
	})

	if len(_authUsers) != 1 {
		log.Println(fmt.Errorf("fail to find user %s in database", _userName))
		return
	}

	_authUser := _authUsers[0] // Must exited

	_t := template.Must(template.ParseFiles(HTMLPath + "Menu.html"))

	_firstFloor := service.MealService.Select(0, floorSelector(1))
	_secondFloor := service.MealService.Select(0, floorSelector(2))
	_menuItems := [...]interface{}{_authUser, _firstFloor, _secondFloor, ACCEPTERID}

	// set cookie
	setCookies(c, map[string]string{
		AUTHKEY: strconv.Itoa(_authUser.(models.User).UserID)})

	_t.Execute(c.Writer, _menuItems)
}
