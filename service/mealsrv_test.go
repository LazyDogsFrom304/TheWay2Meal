package service

import (
	"math/rand"
	"sync"
	"testing"
	"theway2meal/models"
)

func Test_MealGet(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	meal0 := MealService.GetMeal(0)
	if meal0.Name != models.Meals[0].Name {
		t.Error("MealService Get test failed")
	}
}

func Test_MealUpdate(t *testing.T) {
	clear()
	db := GetDefaultDB()
	db_loadTestingData(db)
	var wg sync.WaitGroup
	wg.Add(test_case)

	appendOrder := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			MealService.Update(uint32(rand.Intn(len(models.Meals))))
			wg.Done()
		}
	}

	N := 4
	for i := 0; i < N; i++ {
		go appendOrder(int(test_case) / N)
	}
	wg.Wait()

	var sum uint32
	for _, user := range MealService.SelectAll() {
		sum += user.(models.Meal).Popularity
	}

	if sum != test_case {
		t.Errorf("Bill numbers failed: sum is %d", sum)
	}

}
