package service

import (
	"os"
	"testing"
	"theway2meal/models"
)

const test_case = 100

// For test only
func clear() {
	singleInstanceDB = nil
	PendingOrderService.indexNext = 0
	dbPath = "database_test.gdb"
}

func Test_SingleDBWrite(t *testing.T) {
	clear()
	defer os.Remove(dbPath)
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, true, true)
	obj, _ := db.Get("users:1")
	user1, _ := obj.(models.User)
	user1.Balance += 64.0

	db2 := GetDefaultDB()
	db2.Set("users:0", user1)
	t.Log(user1)
	user_val, _ := db2.Get("users:0")
	t.Log(user_val.(models.User))
	if user_val.(models.User).Balance != 74.3 {
		t.Error("DB Write failed.")
	}

}

func Test_SingleDBRead(t *testing.T) {
	clear()
	defer os.Remove(dbPath)
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, true, true)
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

func Test_Presisted(t *testing.T) {
	//test loading data
	clear()
	var err error
	db := GetDefaultDB()
	DBResetDataTemplate(db, true, false, true)

	dbPath := "../data/database.gdb"
	err = db.PersistDataBase(dbPath)
	if err != nil {
		t.Error("Database marshal failed, ", err)
		return
	}
	// defer os.Remove(dbPath)

	db_cp := make(DataBase)
	err = (&db_cp).LoadDataBase(dbPath)
	if err != nil {
		t.Error("Database unmarshal failed, ", err)
		return
	}

	for k, v := range db_cp {
		for _k := range v {
			if db[k][_k] != db_cp[k][_k] {
				t.Errorf("DataBase presisted failed, since %+v differs from %+v", db[k][_k], db_cp[k][_k])
			}
		}
	}
}
