package service

import (
	"fmt"
	"strconv"
	"sync"
)

const cacheCap = 10

type service struct {
	rwmutex   *sync.RWMutex
	tableName string
	cache     []interface{}

	// As a constraint, users need to implement
	// Service's own handle function.
	handleBeforeUpdate func(interface{}, ...interface{}) (interface{}, error)
}

func (srv *service) internalGet(id int) (obj interface{}, err error) {
	srv.rwmutex.RLock()
	defer srv.rwmutex.RUnlock()

	_db := GetDefaultDB()
	obj, err = _db.Get(srv.tableName + ":" + strconv.Itoa(id))
	return
}

func (srv *service) Update(id int, changes ...interface{}) (old interface{}, err error) {
	srv.rwmutex.Lock()
	defer srv.rwmutex.Unlock()

	_db := GetDefaultDB()
	obj, err := _db.Get(srv.tableName + ":" + strconv.Itoa(id))
	if err != nil {
		return
	}

	obj, err = srv.handleBeforeUpdate(obj, changes...)
	if err != nil {
		return
	}

	old, err = _db.Set(srv.tableName+":"+strconv.Itoa(id), obj)
	if err != nil {
		err = fmt.Errorf("can't apply modification business id %d, due to %s", id, err.Error())
	}
	return
}

func (srv *service) Select(maxLen int, filter func(interface{}) bool) []interface{} {
	srv.rwmutex.RLock()

	_db := GetDefaultDB()
	defer func() {
		srv.rwmutex.RUnlock()
		_db.PersistDataBase(dbPath)
	}()
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

func (srv *service) SelectAll() []interface{} {
	filter := func(interface{}) bool {
		return true
	}
	return srv.Select(0, filter)
}
