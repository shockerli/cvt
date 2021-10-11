package cvt

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var errConvFail = errors.New("convert failed")
var errFieldNotFound = errors.New("field not found")
var errUnsupportedTypeNil = errors.New("unsupported type: nil")
var formatOutOfLimitInt = "%w, out of max limit value(%d)"
var formatOutOfLimitFloat = "%w, out of max limit value(%f)"
var formatExtend = "%v, %w"

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

// returns the value with base type
func indirect(a interface{}) (val interface{}, rt reflect.Type, rv reflect.Value) {
	if a == nil {
		return
	}

	rt = reflect.TypeOf(a)
	rv = reflect.ValueOf(a)
	val = rv.Interface()

	switch rt.Kind() {
	case reflect.Ptr: // indirect the base type, if is been referenced many times
		for rv.Kind() == reflect.Ptr {
			// stop indirect until nil, avoid stack overflow
			if rv.IsNil() {
				val = nil
				return
			}
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
	case reflect.Float32:
		val = float32(rv.Float())
	case reflect.Float64:
		val = rv.Float()
	case reflect.String:
		val = rv.String()
	case reflect.Slice:
		// []byte
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			val = rv.Bytes()
		}
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
