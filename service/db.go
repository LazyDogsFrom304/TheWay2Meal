// Dummy DataBase Provided
// Will mirgrate to redis or mysql later
package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Table map[int]interface{}
type DataBase map[string]Table

// Global instance of database
var MealTable, OrderDoneTable, OrderPendingTable, UserTable = make(Table), make(Table), make(Table), make(Table)
var singleInstanceDB DataBase

// Align with redis api
func GetDefaultDB() DataBase {
	if singleInstanceDB != nil {
		fmt.Printf("singleInstanceDB is %v\n", singleInstanceDB)
		return singleInstanceDB
	}
	return constructDB()
}

func constructDB() DataBase {
	var MealTable, OrderDoneTable, OrderPendingTable, UserTable = make(Table), make(Table), make(Table), make(Table)
	singleInstanceDB = DataBase{
		"meals":         MealTable,
		"ordersDone":    OrderDoneTable,
		"ordersPending": OrderPendingTable,
		"users":         UserTable,
	}
	return singleInstanceDB
}

// simpleSQL formats like:
// `meals:id`
// return obj implements Detachable
func (db DataBase) Get(simpleSQL string) (interface{}, bool) {
	tableName, id, ok := analysisSQL(simpleSQL)
	if !ok {
		return nil, ok
	}
	table, ok := db[tableName]
	if !ok {
		log.Printf("tableName %s is not found\n", tableName)
		return nil, ok
	}
	obj, ok := table[id]
	if !ok && obj != nil {
		log.Printf("index %d is not found\n", id)
		return nil, ok
	}
	return obj, ok
}

// simpleSQL formats like:
// `"users:id", obj`
// pass by value
func (db DataBase) Set(simpleSQL string, obj interface{}) bool {
	tableName, id, ok := analysisSQL(simpleSQL)
	if !ok {
		return ok
	}
	db[tableName][id] = obj
	return true
}

func (db DataBase) DeleteOrder(simpleSQL string) (interface{}, bool) {
	tableName, id, ok := analysisSQL(simpleSQL)
	if !ok {
		return nil, ok
	}
	table, ok := db["ordersPending"]
	if !ok {
		log.Printf("tableName %s is not found\n", tableName)
		return nil, ok
	}
	obj, ok := table[id]
	if !ok && obj != nil {
		log.Printf("index %d is not found\n", id)
		return nil, ok
	}
	delete(table, id)
	return obj, ok
}

func analysisSQL(simpleSQL string) (string, int, bool) {
	keymap := strings.Split(simpleSQL, ":")
	tableName := keymap[0]
	id, ok := strconv.Atoi(keymap[1])
	if ok != nil {
		log.Printf("Error occurs when analysising %s\n, %s", simpleSQL, ok)
		return "", 0, false
	}

	return tableName, id, true
}
