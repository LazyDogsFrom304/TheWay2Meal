package models

import (
	"fmt"
	"sync"
)

type User struct {
	UserID  uint64  `json:"userid" structs:"userid"`
	Name    string  `json:"name" structs:"name"`
	Balance float64 `json:"balance" structs:"balance"`
	lock    sync.Mutex
}

// it's ok to only store balance for each user
// excluding the ownership of every bet.
func (user *User) SyncAddBalance(pay float64) {
	user.lock.Lock()
	user.Balance += pay
	user.lock.Unlock()
}

func (user *User) Request(meal *Meal) {
	user.SyncAddBalance(float64(meal.Price))
}

func (user *User) Accept(meal *Meal) {

}

func (user *User) String() string {
	return fmt.Sprintf("User Id %d is %s, whose bill is at %f yuan.",
		user.UserID,
		user.Name,
		user.Balance)
}

const (
	UserStatusOK = iota
)
