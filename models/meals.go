package models

import (
	"fmt"
	"time"
)

type Meal struct {
	Id          int     `json:"id" structs:"id"`
	Name        string  `json:"name" structs:"name"`
	Price       float64 `json:"price" structs:"price"`
	Floor       int     `json:"floor" structs:"floor"`
	Popularity  int     `json:"popularity" structs:"popularity"`
	LastOrdered string  `json:"lastordered" structs:"lastordered"`
}

// Increment one by one synchronously
func (meal *Meal) Update() {
	meal.Popularity += 1
	meal.LastOrdered = time.Now().Format(TimeFormat)
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
		meal.LastOrdered)
}

func (meal *Meal) DecodeFromMap(m map[string]interface{}) interface{} {
	meal.Id = int(m["id"].(float64))
	meal.Name = m["name"].(string)
	meal.Price = m["price"].(float64)
	meal.Floor = int(m["floor"].(float64))
	meal.Popularity = int(m["popularity"].(float64))
	meal.LastOrdered = m["lastordered"].(string)

	return meal.Detach()
}

const (
	MealStatusOK = iota
)
