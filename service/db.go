// Dummy DataBase Provided
// Will mirgrate to redis or mysql later
package service

import (
	"log"
	"strconv"
	"strings"
	"theway2meal/models"
)

type Table map[int]interface{}
type DataBase map[string]Table

// Global instance of database
var MealTable, OrderTable, UserTable = make(Table), make(Table), make(Table)
var singleInstanceDB = DataBase{
	"meals":  MealTable,
	"orders": OrderTable,
	"users":  UserTable,
}

// Align with redis api
func GetDefaultDB() DataBase {
	return singleInstanceDB
}

// simpleSQL formats like:
// `meals:id`
// 
func (db DataBase) Get(simpleSQL string) (models.Detachable, bool) {
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
	if !ok {
		log.Printf("index %d is not found\n", id)
		return nil, ok
	}
	return obj, ok
}

// simpleSQL formats like:
// `user:`
func （db DataBase) Set(simpleSQL string) (boll)

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
