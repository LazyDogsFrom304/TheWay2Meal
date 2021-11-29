package service

import (
	"fmt"
	"testing"
	"theway2meal/models"
)

func Test_DB(t *testing.T) {
	db := GetDefaultDB()
	for _, user := range models.Users {
		db.Set("users:"+fmt.Sprint(user.UserID), *user)
		tmp := user.Detach().(models.User)
		tmp.UserID += 1
	}
	fmt.Println(db["users"][1].(models.User))
	fmt.Println(models.Users[1])
	user1, _ := db.Get("users:1")
	fmt.Println(user1.(models.User))

}
