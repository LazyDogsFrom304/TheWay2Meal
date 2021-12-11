package models

import (
	"fmt"
)

type Order struct {
	OrderID         int     `json:"orderid" structs:"orderid"`
	OrderTime       string  `json:"ordertime" structs:"ordertime"`
	RequesterName   string  `json:"requestername" structs:"requestername"`
	AcceptorName    string  `json:"acceptorname" structs:"acceptorname"`
	OrderedMealName string  `json:"mealname" structs:"mealname"`
	Price           float64 `json:"price" structs:"price"`
	RequesterId     int     `json:"requesterId" structs:"requesterId"`
	AcceptorId      int     `json:"acceptorId" structs:"acceptorId"`
	OrderedMealId   int     `json:"mealId" structs:"mealId"`
	IsReadyDelete   bool    `json:"isreadydelete" structs:"isreadydelete"`
}

func (order *Order) Done() {
	order.IsReadyDelete = true
}

func (order *Order) String() string {
	return fmt.Sprintf("Order create. No.%d,"+
		" at %s, %s ordered %s,"+
		" price %f,"+
		" %s accepted the request. "+
		"is Done? %t",
		order.OrderID,
		order.OrderTime,
		order.RequesterName,
		order.OrderedMealName,
		order.Price,
		order.AcceptorName,
		order.IsReadyDelete)
}

func (order *Order) Detach() interface{} {
	return *order
}

func (order *Order) DecodeFromMap(m map[string]interface{}) interface{} {
	order.OrderID = int(m["orderid"].(float64))
	order.OrderTime = m["ordertime"].(string)
	order.RequesterName = m["requestername"].(string)
	order.AcceptorName = m["acceptorname"].(string)
	order.OrderedMealName = m["mealname"].(string)
	order.Price = m["price"].(float64)
	order.RequesterId = int(m["requesterId"].(float64))
	order.AcceptorId = int(m["acceptorId"].(float64))
	order.OrderedMealId = int(m["mealId"].(float64))
	order.IsReadyDelete = m["isreadydelete"].(bool)

	return order.Detach()
}

const (
	OrderStatusOK = iota
)
