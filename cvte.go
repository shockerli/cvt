package cvt

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var convErr = errors.New("convert failed")
var formatOutOfLimit = "%w, out of max limit value(%d)"
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

	switch rk {
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
	if v > math.MaxUint64 {
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "uint64"), uint64(math.MaxUint64))
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
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "uint32"), uint32(math.MaxUint32))
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
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "uint16"), uint16(math.MaxUint16))
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
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "uint8"), uint8(math.MaxUint8))
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
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "uint"), ^uint(0))
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

	return 0, convErr
}

// Int64E convert an interface to a int64 type
func Int64E(val interface{}) (int64, error) {
	v, e := convInt64(val)
	if e := catch("int64", val, e); e != nil {
		return 0, e
	}
	if strconv.IntSize == 64 && v > math.MaxInt64 {
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "int64"), math.MaxInt64)
	} else if strconv.IntSize == 32 && v > math.MaxInt32 {
		return 0, fmt.Errorf(formatOutOfLimit, newErr(val, "int64"), math.MaxInt32)
	}

	return v, nil
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

	return 0, convErr
}

// Indirect returns the value with base type
func Indirect(a interface{}) (val interface{}, k reflect.Kind, v reflect.Value) {
	if a == nil {
		return
	}

	t := reflect.TypeOf(a)
	v = reflect.ValueOf(a)
	k = t.Kind()

	switch t.Kind() {
	case reflect.Ptr:
		for v.Kind() == reflect.Ptr && !v.IsNil() {
			v = v.Elem()
		}
		return Indirect(v.Interface())
	case reflect.Bool:
		val = v.Bool()
	case reflect.Int:
		val = int(v.Int())
	case reflect.Int8:
		val = int8(v.Int())
	case reflect.Int16:
		val = int16(v.Int())
	case reflect.Int32:
		val = int32(v.Int())
	case reflect.Int64:
		val = v.Int()
	case reflect.Uint:
		val = uint(v.Uint())
	case reflect.Uint8:
		val = uint8(v.Uint())
	case reflect.Uint16:
		val = uint16(v.Uint())
	case reflect.Uint32:
		val = uint32(v.Uint())
	case reflect.Uint64:
		val = v.Uint()
	case reflect.Uintptr:
		val = uintptr(v.Uint())
	case reflect.Float32:
		val = float32(v.Float())
	case reflect.Float64:
		val = v.Float()
	case reflect.Complex64:
		val = complex64(v.Complex())
	case reflect.Complex128:
		val = v.Complex()
	case reflect.String:
		val = v.String()
	default:
		val = v.Interface()
	}

	return
}

func newErr(val interface{}, t string) error {
	return fmt.Errorf("unable to convert %#v of type %T to %s", val, val, t)
}

func catch(t string, val interface{}, e error) error {
	if e != nil {
		if errors.Is(e, convErr) {
			return newErr(val, t)
		}
		return fmt.Errorf(formatExtend, newErr(val, t), e)
	}
	return nil
}
