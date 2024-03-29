package cvt

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Uint64 convert an interface to an uint64 type, with default value
func Uint64(v interface{}, def ...uint64) uint64 {
	if v, err := Uint64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint64P convert and store in a new uint64 value, and returns a pointer to it
func Uint64P(v interface{}, def ...uint64) *uint64 {
	i := Uint64(v, def...)
	return &i
}

// Uint64E convert an interface to an uint64 type
func Uint64E(val interface{}) (uint64, error) {
	v, e := convUint64(val)
	if e := catch("uint64", val, e); e != nil {
		return 0, e
	}

	return v, nil
}

// Uint32 convert an interface to an uint32 type, with default value
func Uint32(v interface{}, def ...uint32) uint32 {
	if v, err := Uint32E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint32P convert and store in a new uint32 value, and returns a pointer to it
func Uint32P(v interface{}, def ...uint32) *uint32 {
	i := Uint32(v, def...)
	return &i
}

// Uint32E convert an interface to an uint32 type
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

// Uint16 convert an interface to an uint16 type, with default value
func Uint16(v interface{}, def ...uint16) uint16 {
	if v, err := Uint16E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint16P convert and store in a new uint16 value, and returns a pointer to it
func Uint16P(v interface{}, def ...uint16) *uint16 {
	i := Uint16(v, def...)
	return &i
}

// Uint16E convert an interface to an uint16 type
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

// Uint8 convert an interface to an uint8 type, with default value
func Uint8(v interface{}, def ...uint8) uint8 {
	if v, err := Uint8E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint8P convert and store in a new uint8 value, and returns a pointer to it
func Uint8P(v interface{}, def ...uint8) *uint8 {
	i := Uint8(v, def...)
	return &i
}

// Uint8E convert an interface to an uint8 type
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

// Uint convert an interface to an uint type, with default value
func Uint(v interface{}, def ...uint) uint {
	if v, err := UintE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// UintP convert and store in a new uint value, and returns a pointer to it
func UintP(v interface{}, def ...uint) *uint {
	i := Uint(v, def...)
	return &i
}

// UintE convert an interface to an uint type
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

// Int64 convert an interface to an int64 type, with default value
func Int64(v interface{}, def ...int64) int64 {
	if v, err := Int64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int64P convert and store in a new int64 value, and returns a pointer to it
func Int64P(v interface{}, def ...int64) *int64 {
	i := Int64(v, def...)
	return &i
}

// Int64E convert an interface to an int64 type
func Int64E(val interface{}) (int64, error) {
	v, e := convInt64(val)
	if e := catch("int64", val, e); e != nil {
		return 0, e
	}

	return v, nil
}

// Int32 convert an interface to an int32 type, with default value
func Int32(v interface{}, def ...int32) int32 {
	if v, err := Int32E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int32P convert and store in a new int32 value, and returns a pointer to it
func Int32P(v interface{}, def ...int32) *int32 {
	i := Int32(v, def...)
	return &i
}

// Int32E convert an interface to an int32 type
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

// Int16 convert an interface to an int16 type, with default value
func Int16(v interface{}, def ...int16) int16 {
	if v, err := Int16E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int16P convert and store in a new int16 value, and returns a pointer to it
func Int16P(v interface{}, def ...int16) *int16 {
	i := Int16(v, def...)
	return &i
}

// Int16E convert an interface to an int16 type
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

// Int8 convert an interface to an int8 type, with default value
func Int8(v interface{}, def ...int8) int8 {
	if v, err := Int8E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int8P convert and store in a new int8 value, and returns a pointer to it
func Int8P(v interface{}, def ...int8) *int8 {
	i := Int8(v, def...)
	return &i
}

// Int8E convert an interface to an int8 type
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

// Int convert an interface to an int type, with default value
func Int(v interface{}, def ...int) int {
	if v, err := IntE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// IntP convert and store in a new int value, and returns a pointer to it
func IntP(v interface{}, def ...int) *int {
	i := Int(v, def...)
	return &i
}

// IntE convert an interface to an int type
func IntE(val interface{}) (int, error) {
	v, e := convInt64(val)
	if e := catch("int", val, e); e != nil {
		return 0, e
	}
	// 32bit system
	if strconv.IntSize == 32 && v > math.MaxInt32 {
		return 0, fmt.Errorf(formatOutOfLimitInt, newErr(val, "int"), int32(math.MaxInt32))
	}

	return int(v), nil
}

// convert any value to uint64
func convUint64(val interface{}) (uint64, error) {
	// direct type(for improve performance)
	switch vv := val.(type) {
	case int:
		if vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case int64:
		if vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case int32:
		if vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case int16:
		if vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case int8:
		if vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case uint:
		return uint64(vv), nil
	case uint64:
		return vv, nil
	case uint32:
		return uint64(vv), nil
	case uint16:
		return uint64(vv), nil
	case uint8:
		return uint64(vv), nil
	case float64:
		if vv > math.MaxUint64 || vv < 0 {
			return 0, errConvFail
		}
		return uint64(math.Trunc(vv)), nil
	case float32:
		if vv > math.MaxUint64 || vv < 0 {
			return 0, errConvFail
		}
		return uint64(vv), nil
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		return str2uint64(vv)
	case []byte:
		return str2uint64(string(vv))
	}

	// indirect type
	v, rv := Indirect(val)
	switch vv := v.(type) {
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		return str2uint64(vv)
	case []byte:
		return str2uint64(string(vv))
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

// convert any value to int64
func convInt64(val interface{}) (int64, error) {
	// direct type(for improve performance)
	switch vv := val.(type) {
	case int:
		return int64(vv), nil
	case int64:
		return vv, nil
	case int32:
		return int64(vv), nil
	case int16:
		return int64(vv), nil
	case int8:
		return int64(vv), nil
	case uint:
		if strconv.IntSize == 32 && vv > math.MaxInt32 {
			return 0, errConvFail
		}
		return int64(vv), nil
	case uint64:
		if vv > math.MaxInt64 {
			return 0, errConvFail
		}
		return int64(vv), nil
	case uint32:
		return int64(vv), nil
	case uint16:
		return int64(vv), nil
	case uint8:
		return int64(vv), nil
	case float64:
		if vv > math.MaxInt64 {
			return 0, errConvFail
		}
		return int64(math.Trunc(vv)), nil
	case float32:
		return int64(vv), nil
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		return str2int64(vv)
	case []byte:
		return str2int64(string(vv))
	}

	// indirect type
	v, rv := Indirect(val)
	switch vv := v.(type) {
	case nil:
		return 0, nil
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	case string:
		return str2int64(vv)
	case []byte:
		return str2int64(string(vv))
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

// convert an int or float string to int64
//
//	"12" => 12
//	"12.01" => 12
//	"-12" => -12
//	"-12.01" => -12
func str2int64(s string) (i int64, err error) {
	// ensure can be converted to float
	_, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}

	// trim the decimal part
	if i := strings.Index(s, "."); i >= 0 {
		s = s[:i]
	}
	i, err = strconv.ParseInt(s, 10, 64)
	return
}

// convert an int or float string to uint64
//
//	"12" => 12
//	"12.01" => 12
func str2uint64(s string) (i uint64, err error) {
	// ensure can be converted to float
	_, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}

	// trim the decimal part
	if i := strings.Index(s, "."); i >= 0 {
		s = s[:i]
	}
	i, err = strconv.ParseUint(s, 10, 64)
	return
}
