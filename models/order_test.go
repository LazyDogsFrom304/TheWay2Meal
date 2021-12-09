package models

import (
	"log"
	"testing"
)

var orderFmt = PaintStringFunc("order")

func Test_OrderSerializable(t *testing.T) {
	for _, order := range Orders {
		log.Println(orderFmt(&order))
	}
}
