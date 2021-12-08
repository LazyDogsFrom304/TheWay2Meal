package service

import (
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
func handleBill(obj interface{}, changes ...interface{}) interface{} {
	if len(changes) != 1 {
		panic("index out of the range of changes, expect 1.")
	}
	targetUser := obj.(models.User) //it's ok if target is a empty/new User
	targetUser.UpdateBalance(changes[0].(float64))
	return targetUser
}

func (srv *userService) GetUser(userId uint32) *models.User {
	obj := srv.internalGet(userId)
	targetUser, ok := obj.(models.User)
	if !ok {
		return nil
	}
	return &targetUser
}
