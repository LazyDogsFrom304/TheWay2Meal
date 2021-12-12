// Dummy DataBase Provided
// Will mirgrate to redis or mysql later
package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"theway2meal/models"
)

type Table map[int]interface{}
type DataBase map[string]Table

// Global instance of database
var MealTable, OrderDoneTable, OrderPendingTable, UserTable = make(Table), make(Table), make(Table), make(Table)
var singleInstanceDB DataBase
var dbPath = "data/database.gdb"

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

func (db DataBase) PersistDataBase(path string) (err error) {

	data, err := json.Marshal(db)
	if err != nil {
		log.Printf("json marshal failed, err: %v\n", err)
		return
	}

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		log.Printf("Write failed, err: %v\n", err)
	}
	return
}

func (db DataBase) LoadDataBase(path string) (err error) {

	f, err := os.Open(path)
	if err != nil {
		log.Printf("open %s fail, %v\n", path, err)
		return
	}
	defer f.Close()

	data_r, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("read %s fail, %v\n", path, err)
		return
	}

	err = json.Unmarshal([]byte(data_r), &db)
	if err != nil {
		log.Println("Unmarshal failed, ", err)
		return
	}

	// refactor map 2 interface{}
	refactorFunc := func(tableName string, instance models.Detachable) {
		for k, v := range db[tableName] {
			db[tableName][k] = instance.DecodeFromMap(v.(map[string]interface{}))
		}
	}

	refactorFunc("meals", &models.Meal{})
	refactorFunc("ordersDone", &models.Order{})
	refactorFunc("ordersPending", &models.Order{})
	refactorFunc("users", &models.User{})

	return
}

func DBResetDataTemplate(db DataBase, u, o, m bool) {
	if u {
		for _, user := range models.Users {
			db.Set("users:"+strconv.Itoa(user.UserID), user)
		}
	}

	// Test only
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

func OrdersReset() {
	mux := &sync.Mutex{}

	mux.Lock()
	defer mux.Unlock()
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, false, false)
}

func DataBasePrepare(r bool) (err error) {
	db := GetDefaultDB()
	if r {
		PendingOrderService.indexNext = 0
		DBResetDataTemplate(singleInstanceDB, true, false, true)
	} else {
		err = db.LoadDataBase(dbPath)
	}
	return
}
