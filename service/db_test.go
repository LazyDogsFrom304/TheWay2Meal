package service

import (
	"fmt"
	"testing"
	"theway2meal/models"
)

const test_case = 10000

// For test only
func clear() {
	singleInstanceDB = nil
}

func db_loadTestingData(db DataBase) {
	for _, user := range models.Users {
		db.Set("users:"+fmt.Sprint(user.UserID), *user)
	}
	for _, order := range models.Orders[:1] {
		db.Set("ordersDone:"+fmt.Sprint(order.OrderID), *order)
	}
	for _, order := range models.Orders[1:] {
		db.Set("ordersPending:"+fmt.Sprint(order.OrderID), *order)
	}
	for _, meal := range models.Meals {
		db.Set("meals:"+fmt.Sprint(meal.Id), *meal)
	}
}

func Test_SingleDBWrite(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	obj, _ := db.Get("users:1")
	user1, _ := obj.(models.User)
	user1.Balance += 64.0

	db2 := GetDefaultDB()
	db2.Set("users:0", user1)
	t.Log(user1)
	user_val, _ := db2.Get("users:0")
	t.Log(user_val.(models.User))
	if user_val.(models.User).Balance != 64.0 {
		t.Error("DB Write failed.")
	}

}

func Test_SingleDBRead(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	t.Log(db["users"][1].(models.User))
	t.Log(models.Users[1])
	user1, _ := db.Get("users:1")
	t.Log(user1.(models.User))
	db2 := GetDefaultDB()
	if db2["users"][1].(models.User).UserID != models.Users[1].UserID ||
		models.Users[1].UserID != user1.(models.User).UserID {
		t.Error("Pointer passed during reading database.")
	}
}

func Test_Assert(t *testing.T) {
	var userI interface{} = models.User{}
	if user, ok := userI.(models.Meal); !ok {
		t.Logf("user is %#v\n", user) // Zero value of models.Meal
	}
}
