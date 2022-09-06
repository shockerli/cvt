package cvt

import (
	"fmt"
	"reflect"
)

// Slice convert an interface to a []interface{} type, with default value
func Slice(v interface{}, def ...[]interface{}) []interface{} {
	if v, err := SliceE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// SliceE convert an interface to a []interface{} type
func SliceE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return sl, errUnsupportedTypeNil
	}

	_, rv := indirect(val)

	switch rv.Kind() {
	case reflect.String:
		for _, vvv := range rv.String() {
			sl = append(sl, vvv)
		}
		return
	case reflect.Slice, reflect.Array:
		for j := 0; j < rv.Len(); j++ {
			sl = append(sl, rv.Index(j).Interface())
		}
		return
	case reflect.Map:
		for _, key := range sortedMapKeys(rv) {
			sl = append(sl, rv.MapIndex(key).Interface())
		}
		return
	case reflect.Struct:
		sl = deepStructValues(rv)
		return
	}

	return sl, newErr(val, "slice")
}

// SliceInt convert an interface to a []int type, with default value
func SliceInt(v interface{}, def ...[]int) []int {
	if v, err := SliceIntE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// SliceIntE convert an interface to a []int type
func SliceIntE(val interface{}) (sl []int, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	var vv int
	for _, v := range list {
		vv, err = IntE(v)
		if err != nil {
			return
		}
		sl = append(sl, vv)
	}

	return
}

// SliceInt64 convert an interface to a []int64 type, with default value
func SliceInt64(v interface{}, def ...[]int64) []int64 {
	if v, err := SliceInt64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// SliceInt64E convert an interface to a []int64 type
func SliceInt64E(val interface{}) (sl []int64, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	var vv int64
	for _, v := range list {
		vv, err = Int64E(v)
		if err != nil {
			return
		}
		sl = append(sl, vv)
	}

	return
}

// SliceFloat64 convert an interface to a []float64 type, with default value
func SliceFloat64(v interface{}, def ...[]float64) []float64 {
	if v, err := SliceFloat64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// SliceFloat64E convert an interface to a []float64 type
func SliceFloat64E(val interface{}) (sl []float64, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	var vv float64
	for _, v := range list {
		vv, err = Float64E(v)
		if err != nil {
			return
		}
		sl = append(sl, vv)
	}

	return
}

// SliceString convert an interface to a []string type, with default value
func SliceString(v interface{}, def ...[]string) []string {
	if v, err := SliceStringE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// SliceStringE convert an interface to a []string type
func SliceStringE(val interface{}) (sl []string, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	var vv string
	for _, v := range list {
		vv, err = StringE(v)
		if err != nil {
			return
		}
		sl = append(sl, vv)
	}

	return
}

// ColumnsE return the values from a single column in the input array/slice/map of struct/map
func ColumnsE(val interface{}, field interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rv := indirect(val)

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		for j := 0; j < rv.Len(); j++ {
			vv, e := FieldE(rv.Index(j).Interface(), field)
			if e == nil {
				sl = append(sl, vv)
			}
		}
	case reflect.Map:
		for _, key := range sortedMapKeys(rv) {
			vv, e := FieldE(rv.MapIndex(key).Interface(), field)
			if e == nil {
				sl = append(sl, vv)
			}
		}
	}

	// non valid field value, means error
	if len(sl) > 0 {
		return
	}

	return nil, fmt.Errorf("unsupported type: %s", rv.Type().Name())
}

// KeysE return the keys of map, sorted by asc; or fields of struct
func KeysE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rv := indirect(val)

	switch rv.Kind() {
	case reflect.Map:
		for _, key := range sortedMapKeys(rv) {
			sl = append(sl, key.Interface())
		}
		return
	case reflect.Struct:
		for _, v := range deepStructFields(rv.Type()) {
			sl = append(sl, v)
		}
		return
	}

	return nil, fmt.Errorf("unsupported type: %s", rv.Type().Name())
}
