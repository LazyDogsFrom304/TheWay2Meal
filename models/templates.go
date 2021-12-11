package models

// obj.Detach() equals *obj
type Detachable interface {
	Detach() interface{}
	DecodeFromMap(map[string]interface{}) interface{}
}
