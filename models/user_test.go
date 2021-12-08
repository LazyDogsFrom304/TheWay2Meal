package models

import (
	"log"
	"testing"
)

var userFmt = PaintStringFunc("user")

func Test_UserSerializable(t *testing.T) {
	for _, user := range Users {
		log.Println(userFmt(user))
	}
}
