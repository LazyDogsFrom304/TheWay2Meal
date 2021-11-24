package models

import (
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
