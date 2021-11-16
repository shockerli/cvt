package cvt

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// String convert an interface to a string type, with default value
func String(v interface{}, def ...string) string {
	if v, err := StringE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

// StringE convert an interface to a string type
func StringE(val interface{}) (string, error) {
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

	// direct type
	switch vv := val.(type) {
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
		return strconv.FormatUint(Uint64(vv), 10), nil
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(Int64(vv), 10), nil
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(vv), 'f', -1, 32), nil
	}

	// indirect type
	v, rv := indirect(val)
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
