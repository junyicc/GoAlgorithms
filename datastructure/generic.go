package datastructure

import (
	"fmt"
	"reflect"
)

// Elem type of data structures
type Elem interface{}

func String(e Elem) string {
	if e == nil {
		return ""
	}
	return fmt.Sprint(e)
}

// Less comparison
func Less(e1, e2 Elem) bool {
	if e1 == nil || e2 == nil {
		return false
	}
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		return false
	}
	switch e1.(type) {
	case nil:
		return false
	case int, int8, int16, int32, int64:
		return e1.(int) < e2.(int)
	case uint, uint8, uint16, uint32, uint64:
		return e1.(uint) < e2.(uint)
	case float32, float64:
		return e1.(float64) < e2.(float64)
	case string:
		return e1.(string) < e2.(string)
	default:
		return false
	}
}

// Equal comparison
func Equal(e1, e2 Elem) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if e1 == nil || e2 == nil {
		return false
	}
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		return false
	}
	switch e1.(type) {
	case nil:
		return false
	case int, int8, int16, int32, int64:
		return e1.(int) == e2.(int)
	case uint, uint8, uint16, uint32, uint64:
		return e1.(uint) == e2.(uint)
	case float32, float64:
		return e1.(float64) == e2.(float64)
	case string:
		return e1.(string) == e2.(string)
	default:
		return false
	}
}
