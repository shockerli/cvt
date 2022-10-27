package cvt_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/shockerli/cvt"
)

var locUTC8 = time.FixedZone("UTC", 8*3600)

func TestTimeInLocation_HasDefault(t *testing.T) {

	var (
		expect = time.Date(2009, 2, 13, 23, 31, 30, 0, locUTC8)
	)

	tests := []struct {
		input  interface{}
		loc    *time.Location
		def    time.Time
		expect time.Time
	}{
		// supported value, def is not used, def != expect
		{"2009-02-13T23:31:30+0800", time.Local, time.Time{}, expect},
		{"2009-02-13T23:31:30", locUTC8, time.Time{}, expect},

		// unsupported value, def == expect
		{"hello world", locUTC8, expect, expect},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.TimeInLocation(tt.input, tt.loc, tt.def)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestTimeInLocation_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		loc    *time.Location
		expect time.Time
	}{
		{"hello world", cvt.TimeLocation, time.Time{}},
		{testing.T{}, cvt.TimeLocation, time.Time{}},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.TimeInLocation(tt.input, tt.loc)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestTimeInLocationE(t *testing.T) {
	tests := []struct {
		input  interface{}
		loc    *time.Location
		expect time.Time
		isErr  bool
	}{
		{nil, nil, time.Time{}, false},

		// errors
		{"2006", nil, time.Time{}, true},
		{"hello world", nil, time.Time{}, true},
		{testing.T{}, nil, time.Time{}, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], expect[%+v], isErr[%v]",
			i, tt.input, tt.expect, tt.isErr,
		)

		v, err := cvt.TimeInLocationE(tt.input, tt.loc)
		if tt.isErr {
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqualTime(t, tt.expect, v, "[WithE] "+msg)

		// Non-E test
		v = cvt.TimeInLocation(tt.input, tt.loc)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestTime_HasDefault(t *testing.T) {
	var expect = time.Date(2009, 2, 13, 23, 31, 30, 0, cvt.TimeLocation)

	tests := []struct {
		input  interface{}
		def    time.Time
		expect time.Time
	}{
		// supported value, def is not used, def != expect
		{"2009-02-13 23:31:30", time.Time{}, expect},

		// unsupported value, def == expect
		{"hello world", expect, expect},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Time(tt.input, tt.def)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestTime_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
	}{
		{"hello world", time.Time{}},
		{testing.T{}, time.Time{}},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Time(tt.input)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestTimeE(t *testing.T) {
	var (
		// UNIX(123456789)
		expect1 = time.Date(2009, 2, 13, 23, 31, 30, 0, cvt.TimeLocation)
		expect2 = time.Date(2009, 2, 13, 0, 0, 0, 0, cvt.TimeLocation)
	)

	tests := []struct {
		input  interface{}
		expect time.Time
		isErr  bool
	}{
		{nil, time.Time{}, false},

		// string
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, cvt.TimeLocation), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, cvt.TimeLocation), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, cvt.TimeLocation), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, cvt.TimeLocation), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, cvt.TimeLocation), false},         // RFC3339 without T or timezone hh:mm colon

		{"20090213", time.Date(2009, 2, 13, 0, 0, 0, 0, cvt.TimeLocation), false},
		{"20090213233130", time.Date(2009, 2, 13, 23, 31, 30, 0, cvt.TimeLocation), false},
		{"Fri Sep 25 13:58:21 2016 -0400", time.Date(2016, 9, 25, 17, 58, 21, 0, cvt.TimeLocation), false},
		{"13 Feb 2009", expect2, false},
		{"2009-02-13", expect2, false},
		{"2009-02-13 23:31:30", expect1, false},
		{"2009-02-13 23:31:30.618", expect1.Add(618000000), false},
		{"2009-02-13 23:31:30.618003", expect1.Add(618003000), false},
		{"2009-02-13 23:31:30.618003001", expect1.Add(618003001), false},
		{"2009-02-13 23:31:30 -0000", expect1, false},
		{"2009-02-13 23:31:30 -00:00", expect1, false},
		{"2009.02.13", expect2, false},
		{"2009.02.13 23:31:30", expect1, false},
		{"2009.02.13 23:31:30.618", expect1.Add(618000000), false},
		{"2009.02.13 23:31:30.618003", expect1.Add(618003000), false},
		{"2009.02.13 23:31:30.618003001", expect1.Add(618003001), false},
		{"2009/02/13", expect2, false},
		{"2009/02/13 23:31:30", expect1, false},
		{"2009/02/13 23:31:30.618", expect1.Add(618000000), false},
		{"2009/02/13 23:31:30.618003", expect1.Add(618003000), false},
		{"2009/02/13 23:31:30.618003001", expect1.Add(618003001), false},
		{"2009年02月13日", expect2, false},
		{"2009年02月13日 23:31:30", expect1, false},
		{"2009年02月13日 23:31:30.618", expect1.Add(618000000), false},
		{"2009年02月13日 23:31:30.618003", expect1.Add(618003000), false},
		{"2009年02月13日 23:31:30.618003001", expect1.Add(618003001), false},
		{"2009年02月13日 23时31分30秒", expect1, false},

		// int
		{1234567890, expect1, false},
		{int64(1234567890), expect1, false},
		{int32(1234567890), expect1, false},
		{uint(1234567890), expect1, false},
		{uint64(1234567890), expect1, false},
		{uint32(1234567890), expect1, false},

		{time.Duration(1234567890*1e9 + 1), expect1.Add(1), false},
		{expect1, expect1, false},
		{time1, expect1, false},
		{&time1, expect1, false},
		{TestTimeStringer{expect1}, expect1, false},
		{pointerIntNil, time.Time{}, false},
		{aliasTypeStringTime1, expect1, false},
		{&aliasTypeStringTime1, expect1, false},
		{aliasTypeIntTime1, expect1, false},
		{&aliasTypeIntTime1, expect1, false},
		{json.Number("1234567890"), expect1, false},
		{json.Number(aliasTypeStringTime1), expect1, false},

		// errors
		{"2006", time.Time{}, true},
		{"hello world", time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], expect[%+v], isErr[%v]",
			i, tt.input, tt.expect, tt.isErr,
		)

		v, err := cvt.TimeE(tt.input)
		if tt.isErr {
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqualTime(t, tt.expect, v, "[WithE] "+msg)

		// Non-E test
		v = cvt.Time(tt.input)
		assertEqualTime(t, tt.expect, v, "[NonE] "+msg)
	}
}
