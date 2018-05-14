package datastructure

// Elem type of data structures
type Elem interface{}

// Less comparison
func Less(e1, e2 Elem) bool {
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
