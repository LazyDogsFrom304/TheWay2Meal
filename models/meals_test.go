package models

import (
	"log"
	"math/rand"
	"sync"
	"testing"
)

var mealFmt = PaintStringFunc("meal")

func Test_MealSerializable(t *testing.T) {
	for _, meal := range Meals {
		log.Println(mealFmt(meal))
	}
}

// TODO: Mirgrate to service
func Test_Order(t *testing.T) {
	totalCase := 10000 //which should be enough
	var wg sync.WaitGroup
	wg.Add(totalCase)

	orderMeal := func(dealNums int) {
		for i := 0; i < dealNums; i++ {
			Meals[rand.Intn(len(Meals))].Update()
			wg.Done()
		}
	}

	go orderMeal(totalCase / 2)
	go orderMeal(totalCase / 2)
	wg.Wait()

	if Meals[0].Popularity+Meals[1].Popularity != uint32(totalCase) {
		t.Error("Unsynced error ")
	}
}
