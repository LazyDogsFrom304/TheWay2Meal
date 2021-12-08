package service

import (
	"math"
	"math/rand"
	"sync"
	"testing"
	"theway2meal/models"
)

func Test_UserGet(t *testing.T) {
	clear()
	db := GetDefaultDB()
	DB_loadTestingData(db, true, true, true)
	user2 := UserService.GetUser(2)
	if user2.Name != models.Users[2].Name {
		t.Error("UserService Get test failed")
	}
}

func Test_UserHandle(t *testing.T) {
	clear()
	db := GetDefaultDB()
	DB_loadTestingData(db, true, true, true)
	var wg sync.WaitGroup
	wg.Add(test_case)

	appendOrder := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			meal := models.Meals[rand.Intn(len(models.Meals))]
			UserService.Update(uint32(rand.Intn(len(models.Users))),
				meal.Price)
			UserService.Update(uint32(rand.Intn(len(models.Users))),
				-meal.Price)
			wg.Done()
		}
	}

	N := 4
	for i := 0; i < N; i++ {
		go appendOrder(int(test_case) / N)
	}
	wg.Wait()

	var sum float64
	for _, user := range UserService.SelectAll() {
		sum += user.(models.User).Balance
	}

	if math.Abs(sum) > 1e4 {
		t.Errorf("Bill checks failed: sum is %f", sum)
	}

}
