package models

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

var meals = []*Meal{
	{
		Id:    0,
		Name:  "grilled goose",
		Price: 10.3,
	},
	{
		Id:    1,
		Name:  "fried chicken",
		Price: 11.3,
	},
}

func Test_MealSerializable(t *testing.T) {
	for _, meal := range meals {
		fmt.Println(meal)
	}
}

func Test_Order(t *testing.T) {
	totalCase := 10000 //which should be enough
	var wg sync.WaitGroup
	wg.Add(totalCase)

	orderMeal := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			meals[rand.Intn(len(meals))].SyncAddition()
			wg.Done()
		}
	}

	go orderMeal(totalCase / 2)
	go orderMeal(totalCase / 2)
	wg.Wait()

	for _, meal := range meals {
		fmt.Println(meal)
	}

	if meals[0].popularity+meals[1].popularity != totalCase {
		t.Error("Unsynced error ")
	}
}
