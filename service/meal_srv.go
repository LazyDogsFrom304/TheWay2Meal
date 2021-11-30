package service

import (
	"sync"
	"theway2meal/models"
)

// Cache support
type mealService struct {
	service
}

var MealService = &mealService{
	service: service{
		rwmutex:            &sync.RWMutex{},
		tableName:          "meals",
		cache:              make([]interface{}, cacheCap),
		handleBeforeUpdate: updateMealRecords,
	},
}

// changes: [nil]
func updateMealRecords(obj interface{}, changes ...interface{}) interface{} {
	if len(changes) > 0 {
		panic("update Meals record needs no args")
	}
	targetMeal := obj.(models.Meal)
	targetMeal.Update()
	return targetMeal
}

func (srv *mealService) GetMeal(mealId uint32) *models.Meal {
	obj := srv.internalGet(mealId)
	targetMeal, ok := obj.(models.Meal)
	if !ok {
		return nil
	}
	return &targetMeal
}
