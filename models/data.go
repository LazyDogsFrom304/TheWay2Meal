package models

import (
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

var Meals = []Meal{
	{
		Id:          0,
		Name:        "烧鹅饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          1,
		Name:        "鸡腿饭",
		Price:       11.3,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          2,
		Name:        "鱼香肉丝",
		Price:       12.3,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
}

var Users = []User{
	{
		UserID:  0,
		Name:    "gs",
		Balance: 0.0,
	},
	{
		UserID:  1,
		Name:    "jzh",
		Balance: 0.0,
	},
	{
		UserID:  2,
		Name:    "zcm",
		Balance: 0.0,
	},
}

var Orders = []Order{}

var console_color = map[string][2]int{
	"order": {31, 40}, //red
	"meal":  {32, 40}, //green
	"user":  {33, 40}, //yellow
}

func PaintStringFunc(obj string) func(fmt.Stringer) string {
	return func(s fmt.Stringer) string {
		return fmt.Sprintf("\033[1;%d;%dm%s\033[0m", console_color[obj][0],
			console_color[obj][1], s)
	}
}
