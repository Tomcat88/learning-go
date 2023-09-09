package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walkValue(val.MapIndex(k))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		res := val.Call(nil)
		for _, r := range res {
			walkValue(r)
		}
	case reflect.String:
		fn(val.String())
	}
}
