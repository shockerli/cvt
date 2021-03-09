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
