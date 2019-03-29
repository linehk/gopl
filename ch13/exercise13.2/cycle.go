package cycle

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	t    reflect.Type
	addr unsafe.Pointer
}

func Cycle(v interface{}) bool {
	seen := make(map[comparison]bool)
	return cycle(reflect.ValueOf(v), seen)
}

func cycle(v reflect.Value, seen map[comparison]bool) bool {
	if v.CanAddr() {
		c := comparison{v.Type(), unsafe.Pointer(v.UnsafeAddr())}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cycle(v.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if cycle(v.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if cycle(v.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if cycle(v.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
