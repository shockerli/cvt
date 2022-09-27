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

	_, rv := Indirect(val)

	switch rv.Kind() {
	case reflect.String:
		var length = rv.Len()
		if length > 0 {
			sl = make([]interface{}, length)
			for j, vvv := range rv.String() {
				sl[j] = vvv
			}
		}
	case reflect.Slice, reflect.Array:
		var length = rv.Len()
		if length > 0 {
			sl = make([]interface{}, length)
			for j := 0; j < length; j++ {
				sl[j] = rv.Index(j).Interface()
			}
		}
	case reflect.Map:
		var length = rv.Len()
		if length > 0 {
			sl = make([]interface{}, length)
			for j, key := range sortedMapKeys(rv) {
				sl[j] = rv.MapIndex(key).Interface()
			}
		}
	case reflect.Struct:
		sl = deepStructValues(rv)
	default:
		err = newErr(val, "slice")
	}
	return
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

	if len(list) > 0 {
		var vv int
		sl = make([]int, len(list))
		for j, v := range list {
			vv, err = IntE(v)
			if err != nil {
				return
			}
			sl[j] = vv
		}
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

	if len(list) > 0 {
		var vv int64
		sl = make([]int64, len(list))
		for j, v := range list {
			vv, err = Int64E(v)
			if err != nil {
				return
			}
			sl[j] = vv
		}
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

	if len(list) > 0 {
		var vv float64
		sl = make([]float64, len(list))
		for j, v := range list {
			vv, err = Float64E(v)
			if err != nil {
				return
			}
			sl[j] = vv
		}
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

	if len(list) > 0 {
		var vv string
		sl = make([]string, len(list))
		for j, v := range list {
			vv, err = StringE(v)
			if err != nil {
				return
			}
			sl[j] = vv
		}
	}

	return
}

// ColumnsE return the values from a single column in the input array/slice/map of struct/map
func ColumnsE(val interface{}, field interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rv := Indirect(val)

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		var vv interface{}
		for j := 0; j < rv.Len(); j++ {
			vv, err = FieldE(rv.Index(j).Interface(), field)
			if err != nil {
				return nil, fmt.Errorf("unsupported type: %s", rv.Type().String())
			}
			sl = append(sl, vv)
		}
	case reflect.Map:
		var vv interface{}
		for _, key := range sortedMapKeys(rv) {
			vv, err = FieldE(rv.MapIndex(key).Interface(), field)
			if err != nil {
				return nil, fmt.Errorf("unsupported type: %s", rv.Type().String())
			}
			sl = append(sl, vv)
		}
	default:
		return nil, fmt.Errorf("unsupported type: %s", rv.Type().String())
	}

	return
}

// KeysE return the keys of map, sorted by asc; or fields of struct
func KeysE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rv := Indirect(val)

	switch rv.Kind() {
	case reflect.Map:
		var length = rv.Len()
		if length > 0 {
			sl = make([]interface{}, length)
			for j, key := range sortedMapKeys(rv) {
				sl[j] = key.Interface()
			}
		}
	case reflect.Struct:
		fs := deepStructFields(rv.Type())
		if len(fs) > 0 {
			sl = make([]interface{}, len(fs))
			for j, v := range fs {
				sl[j] = v
			}
		}
	default:
		err = fmt.Errorf("unsupported type: %s", rv.Type().Name())
	}

	return
}
