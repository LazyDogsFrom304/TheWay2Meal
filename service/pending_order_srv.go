package service

import (
	"sync"
	"theway2meal/models"
)

type pendingOrderService struct {
	service
}

var PendingOrderService = &pendingOrderService{
	service: service{
		rwmutex:            &sync.RWMutex{},
		tableName:          "ordersPending",
		cache:              make([]interface{}, cacheCap),
		handleBeforeUpdate: processPendings,
	},
}

// A filter actually
// obj:nil changes:[models.Order]
// obj:Order changes:[nil]
func processPendings(obj interface{}, changes ...interface{}) interface{} {
	if len(changes) > 1 {
		panic("index out of the range of changes")
	}

	if obj != nil {
		return obj
	} else {
		return nil
	}
}

func (srv *pendingOrderService) GetPendingOrder(orderId uint32) *models.Order {
	obj := srv.internalGet(orderId)
	targetOrder, ok := obj.(models.Order)
	if !ok {
		return nil
	}
	return &targetOrder
}
