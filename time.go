package cvt

import (
	"encoding/json"
	"fmt"
	"time"
)

// Time convert an interface to a time.Time type, with default value
func Time(v interface{}, def ...time.Time) time.Time {
	if v, err := TimeE(v); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return time.Time{}
}

// TimeE convert an interface to a time.Time type
func TimeE(val interface{}) (t time.Time, err error) {
	// direct type(for improve performance)
	switch vv := val.(type) {
	case nil:
		return
	case time.Time:
		return vv, nil
	case string:
		return parseDate(vv)
	case time.Duration,
		int, int32, int64, uint, uint32, uint64:
		return time.Unix(Int64(vv), 0), nil
	case json.Number:
		// timestamp
		vvv, err := vv.Int64()
		if err == nil {
			return time.Unix(Int64(vvv), 0), nil
		}
		// time string
		return parseDate(vv.String())
	}

	// indirect type
	v, _ := Indirect(val)
	switch vv := v.(type) {
	case nil:
		return
	case time.Time:
		return vv, nil
	case string:
		return parseDate(vv)
	case time.Duration,
		int, int32, int64, uint, uint32, uint64:
		return time.Unix(Int64(vv), 0), nil
	}

	// interface implements
	switch vv := val.(type) {
	case fmt.Stringer:
		return parseDate(vv.String())
	}

	return t, newErr(val, "time.Time")
}

func parseDate(s string) (t time.Time, err error) {
	fs := []string{
		time.RFC3339,
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		"2006-01-02T15:04:05",           // ISO8601 without timezone
		"Mon Jan 2 15:04:05 2006 -0700", // Git log date
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05.000000",
		"2006-01-02 15:04:05.000000000",
		"2006.01.02",
		"2006.01.02 15:04:05",
		"2006.01.02 15:04:05.000",
		"2006.01.02 15:04:05.000000",
		"2006.01.02 15:04:05.000000000",
		"2006/01/02",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04:05.000",
		"2006/01/02 15:04:05.000000",
		"2006/01/02 15:04:05.000000000",
		"2006年01月02日",
		"2006年01月02日 15:04:05",
		"2006年01月02日 15:04:05.000",
		"2006年01月02日 15:04:05.000000",
		"2006年01月02日 15:04:05.000000000",
		"2006年01月02日 15时04分05秒",
	}

	for _, dateType := range fs {
		if t, err = time.Parse(dateType, s); err == nil {
			return
		}
	}

	return t, fmt.Errorf("unable to parse date: %s", s)
}
