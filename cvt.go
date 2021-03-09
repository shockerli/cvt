package cvt

// Bool convert an interface to a bool type, with default value
func Bool(v interface{}, def ...bool) bool {
	if v, err := BoolE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return false
}

// Uint64 convert an interface to a uint64 type, with default value
func Uint64(v interface{}, def ...uint64) uint64 {
	if v, err := Uint64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint32 convert an interface to a uint32 type, with default value
func Uint32(v interface{}, def ...uint32) uint32 {
	if v, err := Uint32E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint16 convert an interface to a uint16 type, with default value
func Uint16(v interface{}, def ...uint16) uint16 {
	if v, err := Uint16E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint8 convert an interface to a uint8 type, with default value
func Uint8(v interface{}, def ...uint8) uint8 {
	if v, err := Uint8E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Uint convert an interface to a uint type, with default value
func Uint(v interface{}, def ...uint) uint {
	if v, err := UintE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int64 convert an interface to a int64 type, with default value
func Int64(v interface{}, def ...int64) int64 {
	if v, err := Int64E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int32 convert an interface to a int32 type, with default value
func Int32(v interface{}, def ...int32) int32 {
	if v, err := Int32E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int16 convert an interface to a int16 type, with default value
func Int16(v interface{}, def ...int16) int16 {
	if v, err := Int16E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int8 convert an interface to a int8 type, with default value
func Int8(v interface{}, def ...int8) int8 {
	if v, err := Int8E(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// Int convert an interface to a int type, with default value
func Int(v interface{}, def ...int) int {
	if v, err := IntE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

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
