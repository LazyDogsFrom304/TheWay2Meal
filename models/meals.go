package models

// import {
// 	"reflect"
// }

type Meal struct {
	id            int
	name          string
	price         float32
	isPackingNeed bool
}

//
func (Meal) Get(attr string) interface{} {
	return nil
}
