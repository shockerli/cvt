package cvt

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Float64 convert an interface to a float64 type, with default value
func Float64(v interface{}, def ...float64) float64 {
	if v, err := Float64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Float64P convert and store in a new float64 value, and returns a pointer to it
func Float64P(v interface{}, def ...float64) *float64 {
	i := Float64(v, def...)
	return &i
}

// Float64E convert an interface to a float64 type
func Float64E(val interface{}) (float64, error) {
	v, e := convFloat64E(val)
	if e := catch("float64", val, e); e != nil {
		return 0, e
	}
	return v, nil
}

func convFloat64E(val interface{}) (float64, error) {
	// direct type(for improve performance)
	switch vv := val.(type) {
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
		return 0, errConvFail
	case []byte:
		vvv, err := strconv.ParseFloat(string(vv), 64)
		if err == nil {
			return vvv, nil
		}
		return 0, errConvFail
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return float64(Uint64(vv)), nil
	case int, int8, int16, int32, int64:
		return float64(Int(vv)), nil
	case float32:
		// use fmt to fix float32 -> float64 precision loss
		// eg: cvt.Float64E(float32(8.31))
		return strconv.ParseFloat(fmt.Sprintf("%f", vv), 64)
	case float64:
		return vv, nil
	case json.Number:
		return vv.Float64()
	case time.Duration:
		return float64(vv), nil
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
		return strconv.ParseFloat(vv, 64)
	case []byte:
		return strconv.ParseFloat(string(vv), 64)
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return float64(rv.Uint()), nil
	case int, int8, int16, int32, int64:
		return float64(rv.Int()), nil
	case float32:
		// use fmt to fix float32 -> float64 precision loss
		// eg: cvt.Float64E(float32(8.31))
		return strconv.ParseFloat(fmt.Sprintf("%f", vv), 64)
	case float64:
		return vv, nil
	}

	return 0, errConvFail
}

// Float32 convert an interface to a float32 type, with default value
func Float32(v interface{}, def ...float32) float32 {
	if v, err := Float32E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Float32P convert and store in a new float32 value, and returns a pointer to it
func Float32P(v interface{}, def ...float32) *float32 {
	i := Float32(v, def...)
	return &i
}

// Float32E convert an interface to a float32 type
func Float32E(val interface{}) (float32, error) {
	v, e := convFloat64E(val)
	if e := catch("float32", val, e); e != nil {
		return 0, e
	}
	if v > math.MaxFloat32 {
		return 0, fmt.Errorf(formatOutOfLimitFloat, newErr(val, "float32"), float32(math.MaxFloat32))
	}

	return float32(v), nil
}
