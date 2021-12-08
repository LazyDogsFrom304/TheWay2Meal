package models

import (
	"fmt"
	"time"
)

type Meal struct {
	Id          uint32    `json:"id" structs:"id"`
	Name        string    `json:"name" structs:"name"`
	Price       float64   `json:"price" structs:"price"`
	Floor       int       `json:"floor" structs:"floor"`
	Popularity  uint32    `json:"popularity" structs:"popularity"`
	LastOrdered time.Time `json:"lastordered" structs:"lastordered"`
}

// Increment one by one synchronously
func (meal *Meal) Update() {
	meal.Popularity += 1
	meal.LastOrdered = time.Now()
}

func (meal *Meal) Detach() interface{} {
	return *meal
}

func (meal *Meal) String() string {
	return fmt.Sprintf("Meal Id %d is %s, cost %f yuan, "+
		"ordered for %d times, "+
		"locate at %d Floor"+
		"and the last time when it's ordered is at %s.",
		meal.Id,
		meal.Name,
		meal.Price,
		meal.Popularity,
		meal.Floor,
		meal.LastOrdered.Format("2006-01-02 15:04:05"))
}

const (
	MealStatusOK = iota
)
