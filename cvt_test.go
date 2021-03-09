package cvt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shockerli/cvt"
)

func TestBool_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    bool
		expect bool
	}{
		// supported value, def is not used, def != expect
		{0, true, false},
		{float64(0.00), true, false},
		{int(0.00), true, false},
		{int64(0.00), true, false},
		{uint(0.00), true, false},
		{uint64(0.00), true, false},
		{uint8(0.00), true, false},
		{nil, true, false},
		{"false", true, false},
		{"FALSE", true, false},
		{"False", true, false},
		{"f", true, false},
		{"F", true, false},
		{false, true, false},
		{"off", true, false},
		{"Off", true, false},
		{[]byte("Off"), true, false},
		{aliasTypeInt_0, true, false},
		{&aliasTypeInt_0, true, false},
		{aliasTypeString_0, true, false},
		{&aliasTypeString_0, true, false},
		{aliasTypeString_off, true, false},
		{&aliasTypeString_off, true, false},

		{[]int{}, true, false},
		{[]string{}, true, false},
		{[...]string{}, true, false},
		{map[int]int{}, true, false},
		{map[string]string{}, true, false},

		{"true", false, true},
		{"TRUE", false, true},
		{"True", false, true},
		{"t", false, true},
		{"T", false, true},
		{1, false, true},
		{true, false, true},
		{-1, false, true},
		{"on", false, true},
		{"On", false, true},
		{0.01, false, true},
		{aliasTypeInt_1, false, true},
		{&aliasTypeInt_1, false, true},
		{aliasTypeString_1, false, true},
		{&aliasTypeString_1, false, true},
		{aliasTypeString_on, false, true},
		{&aliasTypeString_on, false, true},

		{[]int{1, 2, 3}, false, true},
		{[]string{"a", "b", "c"}, false, true},
		{[...]string{"a", "b", "c"}, false, true},
		{map[int]int{1: 111, 2: 222}, false, true},
		{map[string]string{"a": "aaa"}, false, true},

		// unsupported value, def == expect
		{"hello", true, true},
		{"hello", false, false},
		{testing.T{}, true, true},
		{testing.T{}, false, false},
		{&testing.T{}, true, true},
		{&testing.T{}, false, false},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Bool(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestBool_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
	}{
		{testing.T{}, false},
		{&testing.T{}, false},
		{[]int{}, false},
		{[]int{1, 2, 3}, true},
		{[]string{}, false},
		{[]string{"a", "b", "c"}, true},
		{[...]string{}, false},
		{map[int]string{}, false},
		{aliasTypeString_8d15_minus, true},
		{&aliasTypeString_8d15_minus, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Bool(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint64_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    uint64
		expect uint64
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{true, 2, 1},
		{false, 2, 0},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},

		// unsupported value, def == expect
		{int(-8), 1, 1},
		{int8(-8), 1, 1},
		{int16(-8), 1, 1},
		{int32(-8), 1, 1},
		{int64(-8), 1, 1},
		{float32(-8.31), 1, 1},
		{float64(-8.31), 1, 1},
		{"-8", 1, 1},
		{"-8.01", 1, 1},
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
		{aliasTypeString_8d15_minus, 1, 1},
		{&aliasTypeString_8d15_minus, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Uint64(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint64_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{aliasTypeString_8d15_minus, 0},
		{&aliasTypeString_8d15_minus, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Uint64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint32_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    uint32
		expect uint32
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{true, 2, 1},
		{false, 2, 0},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},

		// unsupported value, def == expect
		{int(-8), 1, 1},
		{int8(-8), 1, 1},
		{int16(-8), 1, 1},
		{int32(-8), 1, 1},
		{int64(-8), 1, 1},
		{float32(-8.31), 1, 1},
		{float64(-8.31), 1, 1},
		{"-8", 1, 1},
		{"-8.01", 1, 1},
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
		{aliasTypeString_8d15_minus, 1, 1},
		{&aliasTypeString_8d15_minus, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Uint32(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint32_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{aliasTypeString_8d15_minus, 0},
		{&aliasTypeString_8d15_minus, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Uint32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint16_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    uint16
		expect uint16
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{true, 2, 1},
		{false, 2, 0},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},

		// unsupported value, def == expect
		{int(-8), 1, 1},
		{int8(-8), 1, 1},
		{int16(-8), 1, 1},
		{int32(-8), 1, 1},
		{int64(-8), 1, 1},
		{float32(-8.31), 1, 1},
		{float64(-8.31), 1, 1},
		{"-8", 1, 1},
		{"-8.01", 1, 1},
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
		{aliasTypeString_8d15_minus, 1, 1},
		{&aliasTypeString_8d15_minus, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Uint16(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint16_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{aliasTypeString_8d15_minus, 0},
		{&aliasTypeString_8d15_minus, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Uint16(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint8_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    uint8
		expect uint8
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{true, 2, 1},
		{false, 2, 0},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},

		// unsupported value, def == expect
		{int(-8), 1, 1},
		{int8(-8), 1, 1},
		{int16(-8), 1, 1},
		{int32(-8), 1, 1},
		{int64(-8), 1, 1},
		{float32(-8.31), 1, 1},
		{float64(-8.31), 1, 1},
		{"-8", 1, 1},
		{"-8.01", 1, 1},
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
		{aliasTypeString_8d15_minus, 1, 1},
		{&aliasTypeString_8d15_minus, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Uint8(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint8_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{aliasTypeString_8d15_minus, 0},
		{&aliasTypeString_8d15_minus, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Uint8(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    uint
		expect uint
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{true, 2, 1},
		{false, 2, 0},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},

		// unsupported value, def == expect
		{int(-8), 1, 1},
		{int8(-8), 1, 1},
		{int16(-8), 1, 1},
		{int32(-8), 1, 1},
		{int64(-8), 1, 1},
		{float32(-8.31), 1, 1},
		{float64(-8.31), 1, 1},
		{"-8", 1, 1},
		{"-8.01", 1, 1},
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
		{aliasTypeString_8d15_minus, 1, 1},
		{&aliasTypeString_8d15_minus, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Uint(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{aliasTypeString_8d15_minus, 0},
		{&aliasTypeString_8d15_minus, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Uint(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt64_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    int64
		expect int64
	}{
		// supported value, def is not used, def != expect
		{int(8), 1, 8},
		{int8(8), 1, 8},
		{int16(8), 1, 8},
		{int32(8), 1, 8},
		{int64(8), 1, 8},
		{uint(8), 1, 8},
		{uint8(8), 1, 8},
		{uint16(8), 1, 8},
		{uint32(8), 1, 8},
		{uint64(8), 1, 8},
		{float32(8.31), 1, 8},
		{float64(8.31), 1, 8},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8},
		{int(-8), 1, -8},
		{int8(-8), 1, -8},
		{int16(-8), 1, -8},
		{int32(-8), 1, -8},
		{int64(-8), 1, -8},
		{float32(-8.31), 1, -8},
		{float64(-8.31), 1, -8},
		{"-8", 1, -8},
		{"-8.01", 1, -8},
		{true, 2, 1},
		{false, 2, 0},
		{nil, 2, 0},
		{aliasTypeInt_0, 2, 0},
		{&aliasTypeInt_0, 2, 0},
		{aliasTypeInt_1, 2, 1},
		{&aliasTypeInt_1, 2, 1},
		{aliasTypeString_0, 2, 0},
		{&aliasTypeString_0, 2, 0},
		{aliasTypeString_1, 2, 1},
		{&aliasTypeString_1, 2, 1},
		{aliasTypeString_8d15, 2, 8},
		{&aliasTypeString_8d15, 2, 8},
		{aliasTypeString_8d15_minus, 1, -8},
		{&aliasTypeString_8d15_minus, 1, -8},

		// unsupported value, def == expect
		{"10a", 1, 1},
		{"a10a", 1, 1},
		{"8.01a", 1, 1},
		{"8.01 ", 1, 1},
		{"hello", 1, 1},
		{testing.T{}, 1, 1},
		{&testing.T{}, 1, 1},
		{[]int{}, 1, 1},
		{[]string{}, 1, 1},
		{[...]string{}, 1, 1},
		{map[int]string{}, 1, 1},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Int64(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt64_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int64
	}{
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
		{"4873546382743564386435354655456575456754356765546554643456", 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Int64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}
