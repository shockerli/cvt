package cvt

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var errConvFail = errors.New("convert failed")
var errFieldNotFound = errors.New("field not found")
var formatOutOfLimitInt = "%w, out of max limit value(%d)"
var formatOutOfLimitFloat = "%w, out of max limit value(%f)"
var formatExtend = "%v, %w"

// BoolE convert an interface to a bool type
func BoolE(val interface{}) (bool, error) {
	v, rk, rv := Indirect(val)

	switch vv := v.(type) {
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
	case nil:
		return false, nil
	}

	switch rk.Kind() {
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
	case "on":
		return true, nil
	case "off":
		return false, nil
	default:
		return false, newErr(str, "bool")
	}
}

// Uint64E convert an interface to a uint64 type
func Uint64E(val interface{}) (uint64, error) {
	v, e := convUint64(val)
	if e := catch("uint64", val, e); e != nil {
		return 0, e
	}

	return v, nil
}

// Uint32E convert an interface to a uint32 type
func Uint32E(val interface{}) (uint32, error) {
	v, e := convUint64(val)
	if e := catch("uint32", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxUint32 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "uint32"), uint32(math.MaxUint32))
	}

	return uint32(v), nil
}

// Uint16E convert an interface to a uint16 type
func Uint16E(val interface{}) (uint16, error) {
	v, e := convUint64(val)
	if e := catch("uint16", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxUint16 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "uint16"), uint16(math.MaxUint16))
	}

	return uint16(v), nil
}

// Uint8E convert an interface to a uint8 type
func Uint8E(val interface{}) (uint8, error) {
	v, e := convUint64(val)
	if e := catch("uint8", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxUint8 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "uint8"), uint8(math.MaxUint8))
	}

	return uint8(v), nil
}

// UintE convert an interface to a uint type
func UintE(val interface{}) (uint, error) {
	v, e := convUint64(val)
	if e := catch("uint", val, e); e != nil {
		return 0, e
	}
	if v > uint64(^uint(0)) {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "uint"), ^uint(0))
	}

	return uint(v), nil
}

func convUint64(val interface{}) (uint64, error) {
	v, _, rv := Indirect(val)

	switch vv := v.(type) {
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		vvv, err := strconv.ParseFloat(vv, 64)
		if err == nil && vvv >= 0 && vvv <= math.MaxUint64 {
			return uint64(math.Trunc(vvv)), nil
		}
	case []byte:
		vvv, err := strconv.ParseFloat(string(vv), 64)
		if err == nil && vvv >= 0 && vvv <= math.MaxUint64 {
			return uint64(math.Trunc(vvv)), nil
		}
	case uint, uint8, uint16, uint32, uint64:
		return rv.Uint(), nil
	case int, int8, int16, int32, int64:
		if rv.Int() >= 0 {
			return uint64(rv.Int()), nil
		}
	case float32, float64:
		if rv.Float() >= 0 && rv.Float() <= math.MaxUint64 {
			return uint64(math.Trunc(rv.Float())), nil
		}
	}

	return 0, errConvFail
}

// Int64E convert an interface to a int64 type
func Int64E(val interface{}) (int64, error) {
	v, e := convInt64(val)
	if e := catch("int64", val, e); e != nil {
		return 0, e
	}

	return v, nil
}

// Int32E convert an interface to a int32 type
func Int32E(val interface{}) (int32, error) {
	v, e := convInt64(val)
	if e := catch("int32", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxInt32 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "int32"), int32(math.MaxInt32))
	}

	return int32(v), nil
}

// Int16E convert an interface to a int16 type
func Int16E(val interface{}) (int16, error) {
	v, e := convInt64(val)
	if e := catch("int16", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxInt16 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "int16"), int16(math.MaxInt16))
	}

	return int16(v), nil
}

// Int8E convert an interface to a int8 type
func Int8E(val interface{}) (int8, error) {
	v, e := convInt64(val)
	if e := catch("int8", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxInt8 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "int8"), int8(math.MaxInt8))
	}

	return int8(v), nil
}

// IntE convert an interface to a int type
func IntE(val interface{}) (int, error) {
	v, e := convInt64(val)
	if e := catch("int", val, e); e != nil {
		return 0, e
	}
	if strconv.IntSize == 32 && v > math.MaxInt32 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "int"), int32(math.MaxInt32))
	}

	return int(v), nil
}

func convInt64(val interface{}) (int64, error) {
	v, _, rv := Indirect(val)

	switch vv := v.(type) {
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		vvv, err := strconv.ParseFloat(vv, 64)
		if err == nil && vvv <= math.MaxInt64 {
			return int64(math.Trunc(vvv)), nil
		}
	case []byte:
		vvv, err := strconv.ParseFloat(string(vv), 64)
		if err == nil && vvv <= math.MaxInt64 {
			return int64(math.Trunc(vvv)), nil
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		if rv.Uint() <= math.MaxInt64 {
			return int64(rv.Uint()), nil
		}
	case int, int8, int16, int32, int64:
		return rv.Int(), nil
	case float32, float64:
		if rv.Float() <= math.MaxInt64 {
			return int64(math.Trunc(rv.Float())), nil
		}
	}

	return 0, errConvFail
}

// Float64E convert an interface to a float64 type
func Float64E(val interface{}) (float64, error) {
	v, _, rv := Indirect(val)

	switch vv := v.(type) {
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		vvv, err := strconv.ParseFloat(vv, 64)
		if err == nil {
			return vvv, nil
		}
	case []byte:
		vvv, err := strconv.ParseFloat(string(vv), 64)
		if err == nil {
			return vvv, nil
		}
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return float64(rv.Uint()), nil
	case int, int8, int16, int32, int64:
		return float64(rv.Int()), nil
	case float32:
		// use fmt to fix float32 -> float64 precision loss
		// eg: cvt.Float64E(float32(8.31))
		vvv, err := strconv.ParseFloat(fmt.Sprintf("%f", vv), 64)
		if err == nil {
			return vvv, nil
		}
	case float64:
		return vv, nil
	}

	return 0, errConvFail
}

// Float32E convert an interface to a float32 type
func Float32E(val interface{}) (float32, error) {
	v, e := Float64E(val)
	if e := catch("float32", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxFloat32 {
		return 0, fmt.Errorf(formatOutOfLimitFloat, newErr(val, "float32"), float32(math.MaxFloat32))
	}

	return float32(v), nil
}

// StringE convert an interface to a string type
func StringE(val interface{}) (string, error) {
	v, _, rv := Indirect(val)

	// interface implements
	switch vv := val.(type) {
	case fmt.Stringer:
		return vv.String(), nil
	case error:
		return vv.Error(), nil
	case json.Marshaler:
		vvv, e := vv.MarshalJSON()
		if e == nil {
			return string(vvv), nil
		}
	}

	// source type
	switch vv := v.(type) {
	case nil:
		return "", nil
	case bool:
		return strconv.FormatBool(vv), nil
	case string:
		return vv, nil
	case []byte:
		return string(vv), nil
	case []rune:
		return string(vv), nil
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return strconv.FormatUint(rv.Uint(), 10), nil
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(rv.Int(), 10), nil
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(vv), 'f', -1, 32), nil
	}

	return "", newErr(val, "string")
}

// SliceE convert an interface to a []interface{} type
func SliceE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return
	}

	_, rt, rv := Indirect(val)

	switch rt.Kind() {
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
		var keys = rv.MapKeys()
		// sorted by keys
		sort.Slice(keys, func(i, j int) bool {
			return strings.Compare(String(keys[i].Interface()), String(keys[j].Interface())) < 0
		})
		for _, key := range keys {
			sl = append(sl, rv.MapIndex(key).Interface())
		}
		return
	case reflect.Struct:
		sl = deepStructValues(rt, rv)
		return
	}

	return nil, newErr(val, "slice")
}

// return the values of struct fields, and deep find the embedded fields
func deepStructValues(rt reflect.Type, rv reflect.Value) (sl []interface{}) {
	for j := 0; j < rv.NumField(); j++ {
		if rt.Field(j).Anonymous {
			sl = append(sl, deepStructValues(rt.Field(j).Type, rv.Field(j))...)
		} else if rv.Field(j).CanInterface() {
			sl = append(sl, rv.Field(j).Interface())
		}
	}
	return
}

// FieldE return the field value from map/struct, ignore the filed type
func FieldE(val interface{}, field interface{}) (interface{}, error) {
	sf := String(field) // match with the String of field, so field can be any type
	_, rt, rv := Indirect(val)

	switch rt.Kind() {
	case reflect.Map: // key of map
		for _, key := range rv.MapKeys() {
			if String(key.Interface()) == sf {
				return rv.MapIndex(key).Interface(), nil
			}
		}
	case reflect.Struct: // field of struct
		vv := rv.FieldByName(sf)
		if vv.IsValid() {
			return vv.Interface(), nil
		}
	}

	return nil, fmt.Errorf("%w(%s)", errFieldNotFound, sf)
}

// Indirect returns the value with base type
func Indirect(a interface{}) (val interface{}, rt reflect.Type, rv reflect.Value) {
	if a == nil {
		return
	}

	rt = reflect.TypeOf(a)
	rv = reflect.ValueOf(a)

	switch rt.Kind() {
	case reflect.Ptr: // indirect the base type, if is be referenced many times
		for rv.Kind() == reflect.Ptr && !rv.IsNil() {
			rv = rv.Elem()
		}
		return Indirect(rv.Interface())
	case reflect.Bool:
		val = rv.Bool()
	case reflect.Int:
		val = int(rv.Int())
	case reflect.Int8:
		val = int8(rv.Int())
	case reflect.Int16:
		val = int16(rv.Int())
	case reflect.Int32:
		val = int32(rv.Int())
	case reflect.Int64:
		val = rv.Int()
	case reflect.Uint:
		val = uint(rv.Uint())
	case reflect.Uint8:
		val = uint8(rv.Uint())
	case reflect.Uint16:
		val = uint16(rv.Uint())
	case reflect.Uint32:
		val = uint32(rv.Uint())
	case reflect.Uint64:
		val = rv.Uint()
	case reflect.Uintptr:
		val = uintptr(rv.Uint())
	case reflect.Float32:
		val = float32(rv.Float())
	case reflect.Float64:
		val = rv.Float()
	case reflect.Complex64:
		val = complex64(rv.Complex())
	case reflect.Complex128:
		val = rv.Complex()
	case reflect.String:
		val = rv.String()
	default:
		val = rv.Interface()
	}

	return
}

func newErr(val interface{}, t string) error {
	return fmt.Errorf("unable to convert %#v of type %T to %s", val, val, t)
}

// catching an error and return a new
func catch(t string, val interface{}, e error) error {
	if e != nil {
		if errors.Is(e, errConvFail) {
			return newErr(val, t)
		}
		return fmt.Errorf(formatExtend, newErr(val, t), e)
	}
	return nil
}
