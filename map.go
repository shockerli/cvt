package cvt

import (
	"encoding/json"
	"reflect"
)

// IntMapE convert an interface to `map[int]interface{}`
// * Support JSON string of map
// * Support any `map` type
func IntMapE(val interface{}) (m map[int]interface{}, err error) {
	m = make(map[int]interface{})
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	// direct type(for improve performance)
	switch v := val.(type) {
	case map[int]interface{}:
		return v, nil
	}

	// indirect type
	_, rv := Indirect(val)
	switch rv.Kind() {
	case reflect.Map:
		var idx int
		for _, key := range rv.MapKeys() {
			idx, err = IntE(key.Interface())
			if err != nil {
				return
			}
			m[idx] = rv.MapIndex(key).Interface()
		}
	case reflect.Slice:
		// []byte
		// Example: []byte(`{1:"bob",2:18}`)
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			err = json.Unmarshal(rv.Bytes(), &m)
		}
	case reflect.String:
		// JSON string of map
		// Example: `{1:"bob",2:18}`
		err = json.Unmarshal([]byte(rv.String()), &m)
	}

	return
}

// StringMapE convert an interface to `map[string]interface{}`
// * Support JSON string of map
// * Support any `map` type
// * Support any `struct` type
func StringMapE(val interface{}) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	if val == nil {
		return nil, errUnsupportedTypeNil
	}

	// direct type(for improve performance)
	switch v := val.(type) {
	case map[string]interface{}:
		return v, nil
	case []byte:
		err = json.Unmarshal(v, &m)
		return
	case string:
		err = json.Unmarshal([]byte(v), &m)
		return
	}

	// indirect type
	_, rv := Indirect(val)
	switch rv.Kind() {
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			m[String(key.Interface())] = rv.MapIndex(key).Interface()
		}
	case reflect.Struct:
		m = struct2map(rv)
	case reflect.Slice:
		// []byte, JSON
		// Example: []byte(`{"name":"bob","age":18}`)
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			err = json.Unmarshal(rv.Bytes(), &m)
		}
	case reflect.String:
		// JSON string of map
		// Example: `{"name":"bob","age":18}`
		err = json.Unmarshal([]byte(rv.String()), &m)
	}

	return
}

func struct2map(rv reflect.Value) map[string]interface{} {
	var m = make(map[string]interface{})
	if !rv.IsValid() {
		return m
	}

	for j := 0; j < rv.NumField(); j++ {
		f := rv.Type().Field(j)
		t := ptrType(f.Type)
		vv := ptrValue(rv.Field(j))
		if f.Anonymous && t.Kind() == reflect.Struct {
			for k, v := range struct2map(vv) {
				// anonymous subfield has a low priority
				if _, ok := m[k]; !ok {
					m[k] = v
				}
			}
		} else if vv.IsValid() && vv.CanInterface() {
			m[f.Name] = vv.Interface()
		}
	}
	return m
}
