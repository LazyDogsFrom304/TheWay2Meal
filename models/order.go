package models

import (
	"fmt"
	"time"
)

type Order struct {
	OrderID         uint64    `json:"orderid" structs:"orderid"`
	OrderTime       time.Time `json:"ordertime" structs:"ordertime"`
	RequesterName   string    `json:"requestername" structs:"requestername"`
	AcceptorName    string    `json:"acceptorname" structs:"acceptorname"`
	OrderedMealName string    `json:"mealname" structs:"mealname"`
	RequesterId     uint64    `json:"requesterId" structs:"requesterId"`
	AcceptorId      uint64    `json:"acceptorId" structs:"acceptorId"`
	OrderedMealId   uint64    `json:"mealId" structs:"mealId"`
	IsDone          bool      `json:"isdone" structs:"isdone"`
}

func (order *Order) Done() {
	order.IsDone = true
}

func (order *Order) String() string {
	return fmt.Sprintf("Order create. No.%d,"+
		" at %s, %s ordered %s,"+
		" %s accepted the request. "+
		"is Done? %t",
		order.OrderID,
		order.OrderTime.Format("2006-01-02 15:04:05"),
		order.RequesterName,
		order.OrderedMealName,
		order.AcceptorName,
		order.IsDone)
}

func (order *Order) Detach() interface{} {
	return *order
}

const (
	OrderStatusOK = iota
)

type Operator interface {
	Apply() bool
}

// func (order *Order) Apply() bool {
// 	// cache
// 	// or db

// 	return true
// }
