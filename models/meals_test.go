package models

import (
	"log"
	"testing"
)

var mealFmt = PaintStringFunc("meal")

func Test_MealSerializable(t *testing.T) {
	for _, meal := range Meals {
		log.Println(mealFmt(meal))
	}
}
