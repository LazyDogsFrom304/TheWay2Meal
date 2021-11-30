package service

import (
	"errors"
	"fmt"
	"sync"
	"theway2meal/models"
)

// Todo: Cache support
type userService struct {
	rwmutex *sync.RWMutex
}

var UserService = &userService{
	&sync.RWMutex{},
}

func (srv *userService) GetUser(userId uint32) *models.User {
	srv.rwmutex.RLock()
	defer srv.rwmutex.RUnlock()

	db := GetDefaultDB()
	obj, ok := db.Get("users" + fmt.Sprint(userId))
	if !ok {
		return nil
	}
	targetUser, ok := obj.(models.User)
	if !ok {
		return nil
	}
	return &targetUser
}

func (src *userService) RequestUser(userId uint32, bill *models.Meal) error {
	src.rwmutex.Lock()
	defer src.rwmutex.Unlock()

	db := GetDefaultDB()
	obj, _ := db.Get("users" + fmt.Sprint(userId))
	targetUser := obj.(models.User) //it's ok if target is a empty/new User
	targetUser.Request(bill)
	if ok := db.Set("users"+fmt.Sprint(userId), targetUser); !ok {

		return capError{errors.New("db error: "), fmt.Sprintf("Can't apply %s's request.\n", targetUser.Name)}
	}
	return nil
}
