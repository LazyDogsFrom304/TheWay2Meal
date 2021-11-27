package models

import (
	"fmt"
	"sync"
	"time"
)

type Meal struct {
	Id          int
	Name        string
	Price       float32
	popularity  int
	lastOrdered time.Time
	lock        sync.Mutex
}

// Increment one by one synchronously
func (meal *Meal) SyncAddition() {
	meal.lock.Lock()
	meal.popularity += 1
	meal.lastOrdered = time.Now()
	meal.lock.Unlock()
}

func (meal *Meal) String() string {
	return fmt.Sprintf("Meal Id %d is %s, cost %f yuan, "+
		"ordered for %d times, "+
		"and the last time when it's ordered is at %s.",
		meal.Id,
		meal.Name,
		meal.Price,
		meal.popularity,
		meal.lastOrdered.Format("2000-01-01 01:01:01"))
}
