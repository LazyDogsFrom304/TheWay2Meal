package service

import (
	"sync"
)

// Todo: Cache support
type userService struct {
	mutex *sync.Mutex
}

var UserService = &userService{
	&sync.Mutex{},
}

func (srv *userService) 