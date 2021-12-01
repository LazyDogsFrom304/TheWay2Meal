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

	_db := GetDefaultDB()
	obj, ok := _db.Get(srv.tableName + fmt.Sprint(id))
	if !ok {
		return nil
	}
	return obj
}

func (srv *service) Update(id uint32, changes ...interface{}) (interface{}, error) {
	srv.rwmutex.Lock()
	defer srv.rwmutex.Unlock()

	_db := GetDefaultDB()
	obj, _ := _db.Get(srv.tableName + fmt.Sprint(id))

	obj = srv.handleBeforeUpdate(obj, changes)

	old, ok := _db.Set(srv.tableName+fmt.Sprint(id), obj)
	if !ok {
		return nil, constructError(srv.tableName,
			fmt.Sprintf("can't apply modification business id %d", id))
	}
	return old, nil
}

func (srv *service) Select(maxLen int, filter func(interface{}) bool) []interface{} {
	srv.rwmutex.RLock()
	defer srv.rwmutex.RUnlock()

	_db := GetDefaultDB()
	_targetTable := _db[srv.tableName]
	if maxLen == 0 {
		maxLen = len(_targetTable)
	}
	candidates := make([]interface{}, 0, maxLen)

	_counter := 0
	for _, cand := range _targetTable {
		if _counter >= maxLen {
			break
		}
		if filter(cand) {
			candidates = append(candidates, cand)
		}
		_counter += 1
	}

	return candidates
}

func (srv *service) SelectAll(filter func(interface{}) bool) []interface{} {
	return srv.Select(0, filter)
}
