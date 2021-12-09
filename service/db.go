// Dummy DataBase Provided
// Will mirgrate to redis or mysql later
package service

import (
	"fmt"
	"strconv"
	"strings"
	"theway2meal/models"
)

type Table map[int]interface{}
type DataBase map[string]Table

// Global instance of database
var MealTable, OrderDoneTable, OrderPendingTable, UserTable = make(Table), make(Table), make(Table), make(Table)
var singleInstanceDB DataBase

// Align with redis api
func GetDefaultDB() DataBase {
	if singleInstanceDB != nil {
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
func (db DataBase) Get(simpleSQL string) (obj interface{}, err error) {
	tableName, id, err := analysisSQL(simpleSQL)
	if err != nil {
		return
	}
	table, ok := db[tableName]
	if !ok {
		err = fmt.Errorf("dbGET: tableName %s is not found", tableName)
		return
	}
	obj = table[id]
	return
}

// simpleSQL formats like:
// `"users:id", obj`
// pass by value
func (db DataBase) Set(simpleSQL string, obj interface{}) (old interface{}, err error) {

	tableName, id, err := analysisSQL(simpleSQL)
	if err != nil {
		return
	}
	old = db[tableName][id]

	if obj != nil {
		db[tableName][id] = obj
	} else {
		delete(db[tableName], id)
	}
	return
}

func analysisSQL(simpleSQL string) (string, int, error) {
	keymap := strings.Split(simpleSQL, ":")
	tableName := keymap[0]
	id, ok := strconv.Atoi(keymap[1])
	if ok != nil {
		return "", 0, fmt.Errorf("error occurs when analysising %s\n, %s", simpleSQL, ok)
	}

	return tableName, id, nil
}

// Test only
func DB_loadTestingData(db DataBase, u, o, m bool) {
	if u {
		for _, user := range models.Users {
			db.Set("users:"+strconv.Itoa(user.UserID), user)
		}
	}

	if o {
		index := 0
		for _, order := range models.Orders[:1] {
			db.Set("ordersDone:"+strconv.Itoa(order.OrderID), order)
			index++
		}
		for _, order := range models.Orders[1:] {
			db.Set("ordersPending:"+strconv.Itoa(order.OrderID), order)
			index++
		}
		PendingOrderService.indexNext = index
	}

	if m {
		for _, meal := range models.Meals {
			db.Set("meals:"+strconv.Itoa(meal.Id), meal)
		}
	}

}
