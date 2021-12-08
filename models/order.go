package models

import (
	"fmt"
)

type Order struct {
	OrderID         uint32  `json:"orderid" structs:"orderid"`
	OrderTime       string  `json:"ordertime" structs:"ordertime"`
	RequesterName   string  `json:"requestername" structs:"requestername"`
	AcceptorName    string  `json:"acceptorname" structs:"acceptorname"`
	OrderedMealName string  `json:"mealname" structs:"mealname"`
	Price           float64 `json:"price" structs:"price"`
	RequesterId     uint32  `json:"requesterId" structs:"requesterId"`
	AcceptorId      uint32  `json:"acceptorId" structs:"acceptorId"`
	OrderedMealId   uint32  `json:"mealId" structs:"mealId"`
	IsReadyDelete   bool    `json:"isdone" structs:"isdone"`
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

const (
	OrderStatusOK = iota
)
