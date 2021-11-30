package service

import (
	"fmt"
	"testing"
	"theway2meal/models"
)

func clear() {
	// Dangerous
	singleInstanceDB = nil
}

func Test_DBRead(t *testing.T) {
	clear()
	db := GetDefaultDB()
	for _, user := range models.Users {
		db.Set("users:"+fmt.Sprint(user.UserID), *user)
		tmp := user.Detach().(models.User)
		tmp.UserID += 1
	}
	t.Log(db["users"][1].(models.User))
	t.Log(models.Users[1])
	user1, _ := db.Get("users:1")
	t.Log(user1.(models.User))
	if db["users"][1].(models.User).UserID != models.Users[1].UserID ||
		models.Users[1].UserID != user1.(models.User).UserID {
		t.Error("Pointer passed during reading database.")
	}

}

func Test_DBWrite(t *testing.T) {
	clear()
	db := GetDefaultDB()
	for _, user := range models.Users {
		db.Set("users:"+fmt.Sprint(user.UserID), *user)
	}
	obj, _ := db.Get("users:1")
	user1, _ := obj.(models.User)
	user1.Balance += 64.0
	db.Set("users:0", user1)
	t.Log(user1)
	user_val, _ := db.Get("users:0")
	t.Log(user_val.(models.User))
	if user_val.(models.User).Balance != 64.0 {
		t.Error("DB Write failed.")
	}

}

func Test_Assert(t *testing.T) {
	var userI interface{} = models.User{}
	if user, ok := userI.(models.Meal); !ok {
		fmt.Printf("user is %#v\n", user) // Zero value of models.Meal
	}

}
