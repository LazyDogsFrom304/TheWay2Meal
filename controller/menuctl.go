package controller

import (
	"html/template"
	"theway2meal/models"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func floorSelector(floor int) func(interface{}) bool {
	return func(i interface{}) bool {
		_meal := i.(models.Meal)
		return _meal.Floor == floor
	}
}

func menuHandler(c *gin.Context) {
	_userName := c.MustGet(gin.AuthUserKey).(string)
	_authUser := service.UserService.Select(0, func(i interface{}) bool {
		_user := i.(models.User)
		return _user.Name == _userName
	})[0] //Must exited
	c.Set("auth_user", _authUser)

	_t := template.Must(template.ParseFiles(HTMLPath + "Menu.html"))

	_firstFloor := service.MealService.Select(0, floorSelector(1))
	_secondFloor := service.MealService.Select(0, floorSelector(2))
	_menuItems := [...]interface{}{_userName, _firstFloor, _secondFloor}

	_t.Execute(c.Writer, _menuItems)
}
