package controller

import (
	"fmt"
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

	t := template.Must(template.ParseFiles(HTMLPath + "Menu.html"))

	_firstFloor := service.MealService.Select(0, floorSelector(1))
	_secondFloor := service.MealService.Select(0, floorSelector(2))
	fmt.Println(_firstFloor...)
	fmt.Println(_secondFloor...)
	_floorList := [...]interface{}{_firstFloor, _secondFloor}

	t.Execute(c.Writer, _floorList)
}
