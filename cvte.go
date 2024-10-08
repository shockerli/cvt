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

// Len return size of string, slice, array or map
func Len(v interface{}) int {
	if v == nil {
		return 0
	}

	switch vv := v.(type) {
	case string:
		return len(vv)
	case []bool:
		return len(vv)
	case []byte:
		return len(vv)
	case []rune:
		return len(vv)
	case []string:
		return len(vv)
	case []int:
		return len(vv)
	case []int8:
		return len(vv)
	case []int16:
		return len(vv)
	case []int64:
		return len(vv)
	case []uint:
		return len(vv)
	case []uint16:
		return len(vv)
	case []uint32:
		return len(vv)
	case []uint64:
		return len(vv)
	case []float32:
		return len(vv)
	case []float64:
		return len(vv)
	case []interface{}:
		return len(vv)
	case map[string]interface{}:
		return len(vv)
	case map[string]string:
		return len(vv)
	case map[string]int:
		return len(vv)
	}

	_, rv := Indirect(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return rv.Len()
	}

	return -1
}

// IsEmpty checks value for empty state
func IsEmpty(v interface{}) bool {
	switch vv := v.(type) {
	case nil:
		return true
	case bool:
		return !vv
	case int:
		return vv == 0
	case int8:
		return vv == 0
	case int16:
		return vv == 0
	case int32:
		return vv == 0
	case int64:
		return vv == 0
	case uint:
		return vv == 0
	case uint8:
		return vv == 0
	case uint16:
		return vv == 0
	case uint32:
		return vv == 0
	case uint64:
		return vv == 0
	case float32:
		return vv == 0
	case float64:
		return vv == 0
	case string:
		return vv == ""
	case []int:
		return len(vv) == 0
	case []int8:
		return len(vv) == 0
	case []int16:
		return len(vv) == 0
	case []int32:
		return len(vv) == 0
	case []int64:
		return len(vv) == 0
	case []uint:
		return len(vv) == 0
	case []uint8:
		return len(vv) == 0
	case []uint16:
		return len(vv) == 0
	case []uint32:
		return len(vv) == 0
	case []uint64:
		return len(vv) == 0
	case []float32:
		return len(vv) == 0
	case []float64:
		return len(vv) == 0
	case []bool:
		return len(vv) == 0
	case []interface{}:
		return len(vv) == 0
	case []string:
		return len(vv) == 0
	}

	_, rv := Indirect(v)
	return rv.IsZero()
}

// Field return the field value from map/struct, with default value
func Field(v interface{}, field interface{}, def ...interface{}) interface{} {
	if v, err := FieldE(v, field); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return nil
}

// FieldE return the field value from map/struct, ignore the field type
func FieldE(val interface{}, field interface{}) (interface{}, error) {
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	sf := String(field) // match with the String of field, so field can be any type
	_, rv := Indirect(val)

	switch rv.Kind() {
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

// Typeof returns a string containing the name of the type of `val`.
func Typeof(val interface{}) string {
	return fmt.Sprintf("%T", val)
}

// return the values of struct fields, and deep find the embedded fields
func deepStructValues(rv reflect.Value) (sl []interface{}) {
	for j := 0; j < rv.NumField(); j++ {
		if rv.Type().Field(j).Anonymous {
			sl = append(sl, deepStructValues(rv.Field(j))...)
		} else if rv.Field(j).CanInterface() {
			sl = append(sl, rv.Field(j).Interface())
		}
	}
	return
}

// return the name of struct fields, and deep find the embedded fields
func deepStructFields(rt reflect.Type) (sl []string) {
	rt = ptrType(rt)

	type field struct {
		level int8
		index int
	}
	var exists = make(map[string]field)

	fn := func(v string, level int8) {
		ff, ok := exists[v]
		if ok && ff.level <= level {
			return
		} else if ok && ff.level > level {
			sl = append(sl[:ff.index], sl[ff.index+1:]...)
		}
		sl = append(sl, v)
		exists[v] = field{level, len(sl) - 1}
	}

	// sort by field definition order, include embed field
	for j := 0; j < rt.NumField(); j++ {
		f := rt.Field(j)
		t := ptrType(f.Type)
		// embed struct, include pointer struct
		if f.Anonymous && t.Kind() == reflect.Struct {
			for _, v := range deepStructFields(t) {
				fn(v, 1)
			}
		} else { // single field, include pointer field
			fn(f.Name, 0)
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

func ptrType(rt reflect.Type) reflect.Type {
	if rt.Kind() == reflect.Ptr {
		for rt.Kind() == reflect.Ptr {
			rt = rt.Elem()
		}
	}
	return rt
}

func ptrValue(rv reflect.Value) reflect.Value {
	if rv.Kind() == reflect.Ptr {
		for rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
	}
	return rv
}

// Indirect returns the value with base type
func Indirect(a interface{}) (val interface{}, rv reflect.Value) {
	if a == nil {
		return
	}

	rv = reflect.ValueOf(a)
	val = rv.Interface()

	switch rv.Kind() {
	case reflect.Ptr: // indirect the base type, if has been referenced many times
		for rv.Kind() == reflect.Ptr {
			// stop indirect until nil, avoid stack overflow
			if rv.IsNil() {
				val = nil
				return
			}
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
