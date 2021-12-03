package controller

import (
	"fmt"
	"html/template"
	"strconv"
	"theway2meal/service"

	"github.com/gin-gonic/gin"
)

func orderPreviewHandler(c *gin.Context) {
	_userName := c.MustGet(gin.AuthUserKey).(string)
	_mealId, e := strconv.Atoi(c.Params.ByName("id"))
	if e != nil {
		fmt.Printf("url error when precessing mealId.")
		return
	}
	_mealInfo := service.MealService.GetMeal(uint32(_mealId))

	_t := template.Must(template.ParseFiles(HTMLPath + "Order.html"))

	_orderInfo := [...]interface{}{_userName, _mealInfo}

	_t.Execute(c.Writer, _orderInfo)
}
