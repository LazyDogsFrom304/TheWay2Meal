package service

import (
	"errors"
	"fmt"
	"sync"
)

type capError struct {
	e   error
	msg string
}

const cacheCap = 10

func constructError(errorPos, info string) capError {
	return capError{errors.New(errorPos), info}
}

func (ce capError) Error() string {
	return ce.e.Error() + " " + ce.msg
}

type service struct {
	rwmutex   *sync.RWMutex
	tableName string
	cache     []interface{}

	// As a constraint, users need to implement
	// Service's own handle function.
	handleBeforeUpdate func(interface{}, ...interface{}) interface{}
}

func (srv *service) internalGet(id uint32) interface{} {
	srv.rwmutex.RLock()
	defer srv.rwmutex.RUnlock()

	db := GetDefaultDB()
	obj, ok := db.Get(srv.tableName + fmt.Sprint(id))
	if !ok {
		return nil
	}
	return obj
}

func (src *service) Update(id uint32, changes ...interface{}) (interface{}, error) {
	src.rwmutex.Lock()
	defer src.rwmutex.Unlock()

	db := GetDefaultDB()
	obj, _ := db.Get(src.tableName + fmt.Sprint(id))

	obj = src.handleBeforeUpdate(obj, changes)

	old, ok := db.Set(src.tableName+fmt.Sprint(id), obj)
	if !ok {
		return nil, constructError(src.tableName,
			fmt.Sprintf("can't apply modification business id %d", id))
	}
	return old, nil
}
