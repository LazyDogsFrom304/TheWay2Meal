package models

import (
	"fmt"
)

type User struct {
	UserID  uint32  `json:"userid" structs:"userid"`
	Name    string  `json:"name" structs:"name"`
	Balance float64 `json:"balance" structs:"balance"`
}

// it's ok to only store balance for each user
// excluding the ownership of every bet.
func (user *User) UpdateBalance(pay float64) {
	user.Balance += pay
}

func (user *User) Detach() interface{} {
	return *user
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
