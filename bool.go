package cvt

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Bool convert an interface to a bool type, with default value
func Bool(v interface{}, def ...bool) bool {
	if v, err := BoolE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return false
}

// BoolP convert and store in a new bool value, and returns a pointer to it
func BoolP(v interface{}, def ...bool) *bool {
	i := Bool(v, def...)
	return &i
}

// BoolE convert an interface to a bool type
func BoolE(val interface{}) (bool, error) {
	// direct type(for improve performance)
	switch vv := val.(type) {
	case nil:
		return false, nil
	case bool:
		return vv, nil
	case
		float32, float64:
		return Float64(vv) != 0, nil
	case
		time.Duration,
		int, int8, int16, int32, int64:
		return Int64(vv) != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		return Uint64(vv) != 0, nil
	case []byte:
		return str2bool(string(vv))
	case string:
		return str2bool(vv)
	case json.Number:
		vvv, err := vv.Float64()
		if err != nil {
			return false, newErr(val, "bool")
		}
		return vvv != 0, nil
	}

	// indirect type
	v, rv := Indirect(val)

	switch vv := v.(type) {
	case nil:
		return false, nil
	case bool:
		return vv, nil
	case int, int8, int16, int32, int64:
		return rv.Int() != 0, nil
	case uint, uint8, uint16, uint32, uint64:
		return rv.Uint() != 0, nil
	case float32, float64:
		return rv.Float() != 0, nil
	case []byte:
		return str2bool(string(vv))
	case string:
		return str2bool(vv)
	}

	switch rv.Kind() {
	// by elem length
	case reflect.Array, reflect.Slice, reflect.Map:
		return rv.Len() > 0, nil
	}

	return false, newErr(val, "bool")
}

// returns the boolean value represented by the string
func str2bool(str string) (bool, error) {
	if val, err := strconv.ParseBool(str); err == nil {
		return val, nil
	} else if val, err := strconv.ParseFloat(str, 64); err == nil {
		return val != 0, nil
	}

	switch strings.ToLower(strings.TrimSpace(str)) {
	case "on", "yes", "y":
		return true, nil
	case "off", "no", "n":
		return false, nil
	default:
		return false, newErr(str, "bool")
	}
}
