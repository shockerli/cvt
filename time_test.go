package cvt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shockerli/cvt"
	"github.com/stretchr/testify/assert"
)

func TestTime_HasDefault(t *testing.T) {
	loc := time.UTC
	tests := []struct {
		input  interface{}
		def    time.Time
		expect time.Time
	}{
		// supported value, def is not used, def != expect
		{"2018-10-21T23:21:29+0200", time.Date(2010, 4, 23, 11, 11, 11, 0, loc), time.Date(2018, 10, 21, 21, 21, 29, 0, loc)},

		// unsupported value, def == expect
		{"hello world", time.Date(2010, 4, 23, 11, 11, 11, 0, loc), time.Date(2010, 4, 23, 11, 11, 11, 0, loc)},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Time(tt.input, tt.def)
		assert.Equal(t, tt.expect, v.UTC(), msg)
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
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestTimeE(t *testing.T) {
	loc := time.UTC

	tests := []struct {
		input  interface{}
		expect time.Time
		isErr  bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, loc), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, loc), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, loc), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, loc), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, loc), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, loc), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, loc), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, loc), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, loc), false},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, loc), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, loc), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, loc), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, loc), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, loc), false},
		{"2010.03.07", time.Date(2010, 3, 7, 0, 0, 0, 0, loc), false},
		{"2010.03.07 18:08:18", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
		{"2010/03/07", time.Date(2010, 3, 7, 0, 0, 0, 0, loc), false},
		{"2010/03/07 18:08:18", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
		{"2010???03???07???", time.Date(2010, 3, 7, 0, 0, 0, 0, loc), false},
		{"2010???03???07??? 18:08:18", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
		{"2010???03???07??? 18???08???18???", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, loc), false},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, loc), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, loc), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, loc), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, loc), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, loc), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, loc), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, loc), time.Date(2009, 2, 13, 23, 31, 30, 0, loc), false},
		{TestTimeStringer{time.Date(2010, 3, 7, 0, 0, 0, 0, loc)}, time.Date(2010, 3, 7, 0, 0, 0, 0, loc), false},

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
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v.UTC(), v, msg)

		// Non-E test
		v = cvt.Time(tt.input)
		assert.Equal(t, tt.expect, v.UTC(), msg)
	}
}
