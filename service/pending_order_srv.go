package service

import (
	"fmt"
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
func processPendings(obj interface{}, changes ...interface{}) (target interface{}, err error) {
	if len(changes) > 1 {
		err = fmt.Errorf("index out of the range of changes, expect 0 or 1, get %d", len(changes))
		return
	}

	if obj != nil && changes[0] == nil {
		return
	}

	target = changes[0]
	return
}

func (srv *pendingOrderService) GetPendingOrder(orderId int) (order models.Order, err error) {
	obj, err := srv.internalGet(orderId)
	if err != nil {
		order = models.Order{}
		return
	}

	order, ok := obj.(models.Order)
	if !ok {
		err = fmt.Errorf("can't construct %v as models.Order", order)
	}
	return
}

func (srv *pendingOrderService) GenerateUID() int {
	srv.rwmutex.Lock()
	defer srv.rwmutex.Unlock()

	uid := srv.indexNext
	srv.indexNext += 1
	return uid
}
