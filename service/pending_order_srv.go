package service

import (
	"sync"
	"theway2meal/models"
)

type pendingOrderService struct {
	service
	indexNext int
}

var PendingOrderService = &pendingOrderService{
	service: service{
		rwmutex:            &sync.RWMutex{},
		tableName:          "ordersPending",
		cache:              make([]interface{}, cacheCap),
		handleBeforeUpdate: processPendings,
	},
	indexNext: 0,
}

// A filter actually
// obj:nil changes:[models.Order]
// obj:Order changes:[nil]
func processPendings(obj interface{}, changes ...interface{}) interface{} {
	if len(changes) > 1 {
		panic("index out of the range of changes")
	}

	if obj != nil {
		return nil
	} else {
		return changes[0]
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

func (srv *pendingOrderService) GenerateUID() uint32 {
	srv.rwmutex.Lock()
	defer srv.rwmutex.Unlock()

	uid := srv.indexNext
	srv.indexNext += 1
	return uint32(uid)
}
