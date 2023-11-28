package cvt

import (
	"encoding/json"
	"fmt"
	"time"
)

// TimeLocation default time location
// you can change this for global location
// if cvt < v0.2.7, Time() and TimeE() use time.Parse() to parse the time string, its default time.UTC
// since cvt >= v0.2.7, add TimeInLocation() and TimeInLocationE() support,
// add this variable to setting default time.Location
var TimeLocation = time.UTC

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
	return TimeInLocationE(val, TimeLocation)
}

// TimeInLocation convert an interface to a time.Time type, with time.Location, with default
func TimeInLocation(v interface{}, loc *time.Location, def ...time.Time) time.Time {
	if v, err := TimeInLocationE(v, loc); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return time.Time{}
}

// TimeInLocationE convert an interface to a time.Time type, with time.Location, with error
func TimeInLocationE(val interface{}, loc *time.Location) (t time.Time, err error) {
	if loc == nil {
		loc = TimeLocation
	}

	// direct type(for improve performance)
	switch vv := val.(type) {
	case nil:
		return
	case time.Time:
		return vv, nil
	case string:
		return parseDate(vv, loc)
	case time.Duration:
		return time.Unix(int64(vv)/1e9, int64(vv)%1e9), nil
	case int, int32, int64, uint, uint32, uint64:
		return time.Unix(Int64(vv), 0), nil
	case json.Number:
		// timestamp
		vvv, err := vv.Int64()
		if err == nil {
			return time.Unix(Int64(vvv), 0), nil
		}
		// time string
		return parseDate(vv.String(), loc)
	}

	// indirect type
	v, _ := Indirect(val)
	switch vv := v.(type) {
	case nil:
		return
	case time.Time:
		return vv, nil
	case string:
		return parseDate(vv, loc)
	case int, int32, int64, uint, uint32, uint64:
		return time.Unix(Int64(vv), 0), nil
	}

	// interface implements
	switch vv := val.(type) {
	case fmt.Stringer:
		return parseDate(vv.String(), loc)
	}

	return t, newErr(val, "time.Time")
}

// TimeFormats all supported time formats
// you can add your custom time format
var TimeFormats = []string{
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
	"20060102",                 // date of int
	"20060102150405",           // datetime of int
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

func parseDate(s string, loc *time.Location) (t time.Time, err error) {
	for _, dateType := range TimeFormats {
		if t, err = time.ParseInLocation(dateType, s, loc); err == nil {
			return
		}
	}

	return t, fmt.Errorf("unable to parse date: %s", s)
}
