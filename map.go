package cvt

import (
	"encoding/json"
	"reflect"
)

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
	_, rv := indirect(val)
	switch rv.Kind() {
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			m[String(key.Interface())] = rv.MapIndex(key).Interface()
		}
	case reflect.Struct:
		m = struct2map(rv)
	case reflect.Slice:
		// []byte
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			err = json.Unmarshal(rv.Bytes(), &m)
		}
	case reflect.String:
		// JSON string of map
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
				// anonymous sub-field has a low priority
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
