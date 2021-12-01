package service

import (
	"sync"
	"theway2meal/models"
)

type doneOrderService struct {
	service
}

var DoneOrderService = &doneOrderService{
	service: service{
		rwmutex:            &sync.RWMutex{},
		tableName:          "ordersDone",
		cache:              make([]interface{}, cacheCap),
		handleBeforeUpdate: appendDoneOrder,
	},
}

// A filter actually
// obj:Order changes:[nil]
func appendDoneOrder(obj interface{}, changes ...interface{}) interface{} {
	if len(changes) != 1 {
		panic("index out of the range of changes")
	}

	return obj
}

func (srv *doneOrderService) GetDoneOrder(orderId uint32) *models.Order {
	obj := srv.internalGet(orderId)
	targetOrder, ok := obj.(models.Order)
	if !ok {
		return nil
	}
	return &targetOrder
}
