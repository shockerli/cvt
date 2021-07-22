package cvt_test

import (
	"errors"
	"fmt"
	"html/template"
	"math/big"
	"testing"
	"time"

	"github.com/shockerli/cvt"
	"github.com/stretchr/testify/assert"
)

func TestString_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    string
		expect string
	}{
		// supported value, def is not used, def != expect
		{uint64(8), "xxx", "8"},
		{float32(8.31), "xxx", "8.31"},
		{float64(-8.31), "xxx", "-8.31"},
		{true, "xxx", "true"},
		{int64(-8), "xxx", "-8"},
		{[]byte("8.01"), "xxx", "8.01"},
		{[]rune("我❤️中国"), "xxx", "我❤️中国"},
		{nil, "xxx", ""},
		{aliasTypeInt_0, "xxx", "0"},
		{&aliasTypeString_8d15_minus, "xxx", "-8.15"},
		{aliasTypeBool_true, "xxx", "true"},
		{errors.New("errors"), "xxx", "errors"},
		{time.Friday, "xxx", "Friday"},
		{big.NewInt(123), "xxx", "123"},
		{TestMarshalJSON{}, "xxx", "MarshalJSON"},
		{&TestMarshalJSON{}, "xxx", "MarshalJSON"},

		// unsupported value, def == expect
		{testing.T{}, "xxx", "xxx"},
		{&testing.T{}, "xxx", "xxx"},
		{[]int{}, "xxx", "xxx"},
		{[]string{}, "xxx", "xxx"},
		{[...]string{}, "xxx", "xxx"},
		{map[int]string{}, "xxx", "xxx"},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.String(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestString_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
	}{
		{testing.T{}, ""},
		{&testing.T{}, ""},
		{[]int{}, ""},
		{[]string{}, ""},
		{[...]string{}, ""},
		{map[int]string{}, ""},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.String(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestStringE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
		isErr  bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{true, "true", false},
		{false, "false", false},
		{int(-8), "-8", false},
		{int8(-8), "-8", false},
		{int16(-8), "-8", false},
		{int32(-8), "-8", false},
		{int64(-8), "-8", false},
		{float32(-8.31), "-8.31", false},
		{float64(-8.31), "-8.31", false},
		{[]byte("-8"), "-8", false},
		{[]byte("-8.01"), "-8.01", false},
		{[]byte("8"), "8", false},
		{[]byte("8.00"), "8.00", false},
		{[]byte("8.01"), "8.01", false},
		{[]rune("我❤️中国"), "我❤️中国", false},
		{nil, "", false},
		{aliasTypeInt_0, "0", false},
		{&aliasTypeInt_0, "0", false},
		{aliasTypeInt_1, "1", false},
		{&aliasTypeInt_1, "1", false},
		{aliasTypeString_0, "0", false},
		{&aliasTypeString_0, "0", false},
		{aliasTypeString_1, "1", false},
		{&aliasTypeString_1, "1", false},
		{aliasTypeString_8d15, "8.15", false},
		{&aliasTypeString_8d15, "8.15", false},
		{aliasTypeString_8d15_minus, "-8.15", false},
		{&aliasTypeString_8d15_minus, "-8.15", false},
		{aliasTypeBool_true, "true", false},
		{&aliasTypeBool_true, "true", false},
		{aliasTypeBool_false, "false", false},
		{&aliasTypeBool_false, "false", false},
		{errors.New("errors"), "errors", false},
		{time.Friday, "Friday", false},
		{big.NewInt(123), "123", false},
		{TestMarshalJSON{}, "MarshalJSON", false},
		{&TestMarshalJSON{}, "MarshalJSON", false},
		{template.URL("http://host.foo"), "http://host.foo", false},

		// errors
		{testing.T{}, "", true},
		{&testing.T{}, "", true},
		{[]int{}, "", true},
		{[]string{}, "", true},
		{[...]string{}, "", true},
		{map[int]string{}, "", true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.StringE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.String(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}
