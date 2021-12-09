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
	DB_loadTestingData(db, true, true, true)
	meal0, e := MealService.GetMeal(0)
	if e != nil || meal0.Name != models.Meals[0].Name {
		t.Errorf("MealService Get test failed, %v", e)
	}
}

func Test_MealUpdate(t *testing.T) {
	clear()
	db := GetDefaultDB()
	DB_loadTestingData(db, true, true, true)
	var wg sync.WaitGroup
	wg.Add(test_case)

	appendOrder := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			MealService.Update(rand.Intn(len(models.Meals)))
			wg.Done()
		}
	}

	N := 4
	for i := 0; i < N; i++ {
		go appendOrder(test_case / N)
	}
	wg.Wait()

	var sum int
	for _, user := range MealService.SelectAll() {
		sum += user.(models.Meal).Popularity
	}

	if sum != test_case {
		t.Errorf("Bill numbers failed: sum is %d", sum)
	}

}
