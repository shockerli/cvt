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
	"time"
)

var errConvFail = errors.New("convert failed")
var errFieldNotFound = errors.New("field not found")
var errUnsupportedTypeNil = errors.New("unsupported type: nil")
var formatOutOfLimitInt = "%w, out of max limit value(%d)"
var formatOutOfLimitFloat = "%w, out of max limit value(%f)"
var formatExtend = "%v, %w"

// BoolE convert an interface to a bool type
func BoolE(val interface{}) (bool, error) {
	v, rk, rv := indirect(val)

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
	v, _, rv := indirect(val)

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
	v, _, rv := indirect(val)

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
	v, _, rv := indirect(val)

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
	v, _, rv := indirect(val)

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

// TimeE convert an interface to a time.Time type
func TimeE(val interface{}) (t time.Time, err error) {
	v, _, _ := indirect(val)

	// source type
	switch vv := v.(type) {
	case time.Time:
		return vv, nil
	case string:
		return parseDate(vv)
	case int, int32, int64, uint, uint32, uint64:
		return time.Unix(Int64(vv), 0), nil
	}

	// interface implements
	switch vv := val.(type) {
	case fmt.Stringer:
		return parseDate(vv.String())
	}

	return time.Time{}, newErr(val, "time.Time")
}

func parseDate(s string) (t time.Time, err error) {
	fs := []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		"2006.01.02",
		"2006.01.02 15:04:05",
		"2006/01/02",
		"2006/01/02 15:04:05",
		"2006年01月02日",
		"2006年01月02日 15:04:05",
		"2006年01月02日 15时04分05秒",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, dateType := range fs {
		if t, err = time.Parse(dateType, s); err == nil {
			return
		}
	}

	return t, fmt.Errorf("unable to parse date: %s", s)
}

// SliceE convert an interface to a []interface{} type
func SliceE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rt, rv := indirect(val)

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
		for _, key := range sortedMapKeys(rv) {
			sl = append(sl, rv.MapIndex(key).Interface())
		}
		return
	case reflect.Struct:
		sl = deepStructValues(rt, rv)
		return
	}

	return nil, newErr(val, "slice")
}

// SliceIntE convert an interface to a []int type
func SliceIntE(val interface{}) (sl []int, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	for _, v := range list {
		vv, err := IntE(v)
		if err != nil {
			return nil, err
		}
		sl = append(sl, vv)
	}

	return
}

// SliceInt64E convert an interface to a []int64 type
func SliceInt64E(val interface{}) (sl []int64, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	for _, v := range list {
		vv, err := Int64E(v)
		if err != nil {
			return nil, err
		}
		sl = append(sl, vv)
	}

	return
}

// SliceFloat64E convert an interface to a []float64 type
func SliceFloat64E(val interface{}) (sl []float64, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	for _, v := range list {
		vv, err := Float64E(v)
		if err != nil {
			return nil, err
		}
		sl = append(sl, vv)
	}

	return
}

// SliceStringE convert an interface to a []string type
func SliceStringE(val interface{}) (sl []string, err error) {
	list, err := SliceE(val)
	if err != nil {
		return
	}

	for _, v := range list {
		vv, err := StringE(v)
		if err != nil {
			return nil, err
		}
		sl = append(sl, vv)
	}

	return
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

// return the map keys, which sorted by asc
func sortedMapKeys(v reflect.Value) (s []reflect.Value) {
	s = v.MapKeys()
	sort.Slice(s, func(i, j int) bool {
		return strings.Compare(String(s[i].Interface()), String(s[j].Interface())) < 0
	})
	return
}

// FieldE return the field value from map/struct, ignore the filed type
func FieldE(val interface{}, field interface{}) (interface{}, error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	sf := String(field) // match with the String of field, so field can be any type
	_, rt, rv := indirect(val)

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

// ColumnsE return the values from a single column in the input array/slice/map of struct/map
func ColumnsE(val interface{}, field interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rt, rv := indirect(val)

	switch rt.Kind() {
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

	return nil, fmt.Errorf("unsupported type: %s", rt.Name())
}

// KeysE return the keys of map, sorted by asc
func KeysE(val interface{}) (sl []interface{}, err error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	_, rt, rv := indirect(val)

	switch rt.Kind() {
	case reflect.Map:
		for _, key := range sortedMapKeys(rv) {
			sl = append(sl, key.Interface())
		}
		return
	}

	return nil, fmt.Errorf("unsupported type: %s", rt.Name())
}

// returns the value with base type
func indirect(a interface{}) (val interface{}, rt reflect.Type, rv reflect.Value) {
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
		return indirect(rv.Interface())
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
		if errors.Is(e, errConvFail) || errors.Is(e, errFieldNotFound) || errors.Is(e, errUnsupportedTypeNil) {
			return newErr(val, t)
		}
		return fmt.Errorf(formatExtend, newErr(val, t), e)
	}
	return nil
}
