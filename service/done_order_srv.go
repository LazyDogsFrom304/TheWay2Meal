package service

import (
	"fmt"
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
func appendDoneOrder(obj interface{}, changes ...interface{}) (interface{}, error) {
	if len(changes) != 1 {
		return nil, fmt.Errorf("index out of the range of changes, expect 1, get %d", len(changes))
	}

	return changes[0], nil
}

func (srv *doneOrderService) GetDoneOrder(orderId int) (order models.Order, err error) {
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
