package cvt_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"math/big"
	"testing"
	"time"

	"github.com/shockerli/cvt"
)

func TestString_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    string
		expect string
	}{
		// supported value, def is not used, def != expect
		{"hello", "world", "hello"},
		{uint64(8), "xxx", "8"},
		{float32(8.31), "xxx", "8.31"},
		{float64(-8.31), "xxx", "-8.31"},
		{true, "xxx", "true"},
		{int64(-8), "xxx", "-8"},
		{[]byte("8.01"), "xxx", "8.01"},
		{[]rune("我❤️中国"), "xxx", "我❤️中国"},
		{nil, "xxx", ""},
		{aliasTypeInt0, "xxx", "0"},
		{&aliasTypeString8d15Minus, "xxx", "-8.15"},
		{aliasTypeBool4True, "xxx", "true"},
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
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
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
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestStringP(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect string
	}{
		{"123", "123"},
		{123, "123"},
		{123.01, "123.01"},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.StringP(tt.input)
		assertEqual(t, tt.expect, *v, "[NonE] "+msg)
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
		{"hello world!", "hello world!", false},
		{[]byte("-8"), "-8", false},
		{[]byte("-8.01"), "-8.01", false},
		{[]byte("8"), "8", false},
		{[]byte("8.00"), "8.00", false},
		{[]byte("8.01"), "8.01", false},
		{[]rune("我❤️中国"), "我❤️中国", false},
		{nil, "", false},
		{pointerInterNil, "", false},
		{AliasTypeBytesNil, "", false},
		{&AliasTypeBytesNil, "", false},
		{aliasTypeInt0, "0", false},
		{&aliasTypeInt0, "0", false},
		{aliasTypeInt1, "1", false},
		{&aliasTypeInt1, "1", false},
		{aliasTypeString0, "0", false},
		{&aliasTypeString0, "0", false},
		{aliasTypeString1, "1", false},
		{&aliasTypeString1, "1", false},
		{aliasTypeString8d15, "8.15", false},
		{&aliasTypeString8d15, "8.15", false},
		{aliasTypeString8d15Minus, "-8.15", false},
		{&aliasTypeString8d15Minus, "-8.15", false},
		{aliasTypeBool4True, "true", false},
		{&aliasTypeBool4True, "true", false},
		{aliasTypeBool4False, "false", false},
		{&aliasTypeBool4False, "false", false},
		{AliasTypeBytes("hello"), "hello", false},
		{&pointerRunes, "中国", false},
		{AliasTypeUint(12), "12", false},
		{AliasTypeUint8(12), "12", false},
		{AliasTypeUint16(12), "12", false},
		{AliasTypeUint32(12), "12", false},
		{AliasTypeUint64(12), "12", false},
		{AliasTypeInt(-12), "-12", false},
		{AliasTypeInt8(-12), "-12", false},
		{AliasTypeInt16(-12), "-12", false},
		{AliasTypeInt32(-12), "-12", false},
		{AliasTypeInt64(-12), "-12", false},
		{AliasTypeFloat32(-12.34), "-12.34", false},
		{AliasTypeFloat64(12.34), "12.34", false},
		{errors.New("errors"), "errors", false},
		{time.Friday, "Friday", false},
		{big.NewInt(123), "123", false},
		{TestMarshalJSON{}, "MarshalJSON", false},
		{&TestMarshalJSON{}, "MarshalJSON", false},
		{template.URL("https://host.foo"), "https://host.foo", false},
		{template.HTML("<html></html>"), "<html></html>", false},
		{json.Number("12.34"), "12.34", false},
		{pointerInterNil, "", false},
		{&pointerInterNil, "", false},
		{pointerIntNil, "", false},
		{&pointerIntNil, "", false},
		{(*AliasTypeInt)(nil), "", false},
		{(*PointerTypeInt)(nil), "", false},

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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)

		// Non-E test
		v = cvt.String(tt.input)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
	}
}
