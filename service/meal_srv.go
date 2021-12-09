package service

import (
	"fmt"
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
func updateMealRecords(obj interface{}, changes ...interface{}) (interface{}, error) {
	if len(changes) > 0 {
		return nil, fmt.Errorf("index out of the range of changes, expect 0, get %d", len(changes))
	}

	targetMeal, ok := obj.(models.Meal)
	if !ok {
		return nil, fmt.Errorf("can't updateMealRecords %v as models.Meal", obj)
	}
	targetMeal.Update()
	return targetMeal, nil
}

func (srv *mealService) GetMeal(mealId int) (meal models.Meal, err error) {
	obj, err := srv.internalGet(mealId)
	if err != nil {
		meal = models.Meal{}
		return
	}

	meal, ok := obj.(models.Meal)
	if !ok {
		err = fmt.Errorf("can't construct %v as models.Meal", obj)
	}
	return
}
