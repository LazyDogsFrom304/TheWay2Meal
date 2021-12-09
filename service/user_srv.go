package service

import (
	"fmt"
	"sync"
	"theway2meal/models"
)

// Todo: Cache support
type userService struct {
	service
}

var UserService = &userService{
	service: service{
		rwmutex:            &sync.RWMutex{},
		tableName:          "users",
		cache:              make([]interface{}, cacheCap),
		handleBeforeUpdate: handleBill,
	},
}

// changes:[float64 payment]
func handleBill(obj interface{}, changes ...interface{}) (interface{}, error) {
	if len(changes) != 1 {
		return nil, fmt.Errorf("index out of the range of changes, expect 1, get %d", len(changes))
	}
	targetUser, ok := obj.(models.User)
	if !ok {
		return nil, fmt.Errorf("can't handleBill %v as models.Meal", obj)
	}
	targetUser.UpdateBalance(changes[0].(float64))
	return targetUser, nil
}

func (srv *userService) GetUser(userId int) (user models.User, err error) {
	obj, err := srv.internalGet(userId)
	if err != nil {
		user = models.User{}
		return
	}

	user, ok := obj.(models.User)
	if !ok {
		err = fmt.Errorf("can't construct %v as models.User", user)
	}
	return
}
