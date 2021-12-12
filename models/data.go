package models

import (
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

var Meals = []Meal{
	{
		Id:          0,
		Name:        "红烧牛蛙饭",
		Price:       15.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          1,
		Name:        "香菇滑鸡饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          2,
		Name:        "松滋鸡煲",
		Price:       12.5,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          3,
		Name:        "香辣鸡煲",
		Price:       12.5,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          4,
		Name:        "农家小炒肉饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          5,
		Name:        "酱香鸭饭",
		Price:       12.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          6,
		Name:        "红烧肉煲",
		Price:       12.5,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          7,
		Name:        "鱼香茄子煲",
		Price:       9.5,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          8,
		Name:        "（双拼）烤肠",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          9,
		Name:        "（双拼）小炒肉",
		Price:       6.2,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          10,
		Name:        "（双拼）培根",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          11,
		Name:        "（双拼）辣子鸡",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          12,
		Name:        "（双拼）鸡排",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          13,
		Name:        "（双拼）荷包蛋",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          14,
		Name:        "（双拼）肉片",
		Price:       5.7,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          15,
		Name:        "鸡腿卤汁饭",
		Price:       11.3,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          16,
		Name:        "鸭腿卤汁饭",
		Price:       10.3,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          17,
		Name:        "8元套餐",
		Price:       8.3,
		Floor:       2,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          18,
		Name:        "鸡片饭",
		Price:       9.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          19,
		Name:        "猪肉饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          20,
		Name:        "鸡肉饭",
		Price:       12.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          21,
		Name:        "烤肉饭",
		Price:       9.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          22,
		Name:        "卤鸭腿饭",
		Price:       8.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          23,
		Name:        "腌菜肉沫饭",
		Price:       9.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          24,
		Name:        "干椒辣子鸡饭",
		Price:       2.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          25,
		Name:        "咖喱鸡饭",
		Price:       9.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          26,
		Name:        "金牌烧鹅腿饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          27,
		Name:        "土豆烧肉饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          28,
		Name:        "孜然肉片饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          29,
		Name:        "香辣牛肉丝饭",
		Price:       12.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          30,
		Name:        "回锅肉饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          31,
		Name:        "金牌黑鸭王饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          32,
		Name:        "蜜汁叉烧饭",
		Price:       13.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          33,
		Name:        "金牌烧鹅饭",
		Price:       12.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          34,
		Name:        "黄金鸡排饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          35,
		Name:        "葱油鸡饭",
		Price:       10.3,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          36,
		Name:        "热卤菜：猪肉鸡肉西兰花胡萝卜",
		Price:       14.1,
		Floor:       1,
		LastOrdered: time.Now().Format(TimeFormat),
	},
	{
		Id:          37,
		Name:        "可口可乐",
		Price:       2.8,
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
