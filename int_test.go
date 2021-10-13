package cvt_test

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/shockerli/cvt"
	"github.com/stretchr/testify/assert"
)

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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},

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
		{aliasTypeString8d15Minus, 1, 1},
		{&aliasTypeString8d15Minus, 1, 1},
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
		{aliasTypeString8d15Minus, 0},
		{&aliasTypeString8d15Minus, 0},
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},

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
		{aliasTypeString8d15Minus, 1, 1},
		{&aliasTypeString8d15Minus, 1, 1},
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
		{aliasTypeString8d15Minus, 0},
		{&aliasTypeString8d15Minus, 0},
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},

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
		{aliasTypeString8d15Minus, 1, 1},
		{&aliasTypeString8d15Minus, 1, 1},
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
		{aliasTypeString8d15Minus, 0},
		{&aliasTypeString8d15Minus, 0},
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},

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
		{aliasTypeString8d15Minus, 1, 1},
		{&aliasTypeString8d15Minus, 1, 1},
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
		{aliasTypeString8d15Minus, 0},
		{&aliasTypeString8d15Minus, 0},
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},

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
		{aliasTypeString8d15Minus, 1, 1},
		{&aliasTypeString8d15Minus, 1, 1},
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
		{aliasTypeString8d15Minus, 0},
		{&aliasTypeString8d15Minus, 0},
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},
		{aliasTypeString8d15Minus, 1, -8},
		{&aliasTypeString8d15Minus, 1, -8},

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

func TestInt32_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    int32
		expect int32
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},
		{aliasTypeString8d15Minus, 1, -8},
		{&aliasTypeString8d15Minus, 1, -8},

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

		v := cvt.Int32(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt32_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
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

		v := cvt.Int32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt16_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    int16
		expect int16
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},
		{aliasTypeString8d15Minus, 1, -8},
		{&aliasTypeString8d15Minus, 1, -8},

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

		v := cvt.Int16(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt16_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
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

		v := cvt.Int16(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt8_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    int8
		expect int8
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},
		{aliasTypeString8d15Minus, 1, -8},
		{&aliasTypeString8d15Minus, 1, -8},

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

		v := cvt.Int8(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt8_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
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

		v := cvt.Int8(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    int
		expect int
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
		{aliasTypeInt0, 2, 0},
		{&aliasTypeInt0, 2, 0},
		{aliasTypeInt1, 2, 1},
		{&aliasTypeInt1, 2, 1},
		{aliasTypeString0, 2, 0},
		{&aliasTypeString0, 2, 0},
		{aliasTypeString1, 2, 1},
		{&aliasTypeString1, 2, 1},
		{aliasTypeString8d15, 2, 8},
		{&aliasTypeString8d15, 2, 8},
		{aliasTypeString8d15Minus, 1, -8},
		{&aliasTypeString8d15Minus, 1, -8},

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

		v := cvt.Int(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
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

		v := cvt.Int(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint64
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{int64(1487354638276643554), 1487354638276643554, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{uint64(1487354638276643554), 1487354638276643554, false},
		{uint(math.MaxUint32), uint64(math.MaxUint32), false},
		{uint32(math.MaxUint32), uint64(math.MaxUint32), false},
		{uint64(math.MaxUint64), uint64(math.MaxUint64), false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"-8.01", 0, true},
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString8d15Minus, 0, true},
		{&aliasTypeString8d15Minus, 0, true},
		{AliasTypeInt(-1), 0, true},
		{AliasTypeInt8(-1), 0, true},
		{AliasTypeInt16(-1), 0, true},
		{AliasTypeInt32(-1), 0, true},
		{AliasTypeInt64(-1), 0, true},
		{AliasTypeFloat32(-1.01), 0, true},
		{AliasTypeFloat64(-1.01), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%+v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Uint64E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Uint64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint32
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{uint(math.MaxUint32), uint32(math.MaxUint32), false},
		{uint32(math.MaxUint32), uint32(math.MaxUint32), false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"-8.01", 0, true},
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{int64(1487354638276643554), 0, true},
		{uint64(1487354638276643554), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString8d15Minus, 0, true},
		{&aliasTypeString8d15Minus, 0, true},
		{AliasTypeInt(-1), 0, true},
		{AliasTypeInt8(-1), 0, true},
		{AliasTypeInt16(-1), 0, true},
		{AliasTypeInt32(-1), 0, true},
		{AliasTypeInt64(-1), 0, true},
		{AliasTypeFloat32(-1.01), 0, true},
		{AliasTypeFloat64(-1.01), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%+v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Uint32E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Uint32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint16
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"-8.01", 0, true},
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{int64(1487354638276643554), 0, true},
		{uint64(1487354638276643554), 0, true},
		{uint(math.MaxUint32), 0, true},
		{uint32(math.MaxUint32), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString8d15Minus, 0, true},
		{&aliasTypeString8d15Minus, 0, true},
		{AliasTypeInt(-1), 0, true},
		{AliasTypeInt8(-1), 0, true},
		{AliasTypeInt16(-1), 0, true},
		{AliasTypeInt32(-1), 0, true},
		{AliasTypeInt64(-1), 0, true},
		{AliasTypeFloat32(-1.01), 0, true},
		{AliasTypeFloat64(-1.01), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%+v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Uint16E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Uint16(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUint8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint8
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"-8.01", 0, true},
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{int64(1487354638276643554), 0, true},
		{uint64(1487354638276643554), 0, true},
		{uint(math.MaxUint32), 0, true},
		{uint32(math.MaxUint32), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString8d15Minus, 0, true},
		{&aliasTypeString8d15Minus, 0, true},
		{AliasTypeInt(-1), 0, true},
		{AliasTypeInt8(-1), 0, true},
		{AliasTypeInt16(-1), 0, true},
		{AliasTypeInt32(-1), 0, true},
		{AliasTypeInt64(-1), 0, true},
		{AliasTypeFloat32(-1.01), 0, true},
		{AliasTypeFloat64(-1.01), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%+v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Uint8E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Uint8(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestUintE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect uint
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{uint64(math.MaxUint32), uint(math.MaxUint32), false},
		{uint32(math.MaxUint32), uint(math.MaxUint32), false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{int(-8), 0, true},
		{int8(-8), 0, true},
		{int16(-8), 0, true},
		{int32(-8), 0, true},
		{int64(-8), 0, true},
		{float32(-8.31), 0, true},
		{float64(-8.31), 0, true},
		{"-8", 0, true},
		{"-8.01", 0, true},
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString8d15Minus, 0, true},
		{&aliasTypeString8d15Minus, 0, true},
		{AliasTypeInt(-1), 0, true},
		{AliasTypeInt8(-1), 0, true},
		{AliasTypeInt16(-1), 0, true},
		{AliasTypeInt32(-1), 0, true},
		{AliasTypeInt64(-1), 0, true},
		{AliasTypeFloat32(-1.01), 0, true},
		{AliasTypeFloat64(-1.01), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.UintE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Uint(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt64E(t *testing.T) {
	type T struct {
		input  interface{}
		expect int64
		isErr  bool
	}
	tests := []T{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8, false},
		{float64(-8.31), -8, false},
		{"-8", -8, false},
		{"-8.01", -8, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{uint32(math.MaxUint32), int64(math.MaxUint32), false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat32(-1.01), -1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{AliasTypeFloat64(-1.01), -1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeString8d15Minus, -8, false},
		{&aliasTypeString8d15Minus, -8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{json.Number("hello"), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Int64E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Int64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int32
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8, false},
		{float64(-8.31), -8, false},
		{"-8", -8, false},
		{"-8.01", -8, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{math.MaxInt32, int32(math.MaxInt32), false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat32(-1.01), -1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{AliasTypeFloat64(-1.01), -1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeString8d15Minus, -8, false},
		{&aliasTypeString8d15Minus, -8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{int64(math.MaxInt64), 0, true},
		{uint32(math.MaxUint32), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{json.Number("hello"), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Int32E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Int32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt16E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int16
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8, false},
		{float64(-8.31), -8, false},
		{"-8", -8, false},
		{"-8.01", -8, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{math.MaxInt16, int16(math.MaxInt16), false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat32(-1.01), -1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{AliasTypeFloat64(-1.01), -1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeString8d15Minus, -8, false},
		{&aliasTypeString8d15Minus, -8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{uint32(math.MaxUint32), 0, true},
		{int64(math.MaxInt64), 0, true},
		{int32(math.MaxInt32), 0, true},
		{uint16(math.MaxUint16), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{json.Number("hello"), 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Int16E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Int16(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestInt8E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int8
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8, false},
		{float64(-8.31), -8, false},
		{"-8", -8, false},
		{"-8.01", -8, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{int8(math.MaxInt8), math.MaxInt8, false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat32(-1.01), -1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{AliasTypeFloat64(-1.01), -1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeString8d15Minus, -8, false},
		{&aliasTypeString8d15Minus, -8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{uint32(math.MaxUint32), 0, true},
		{int64(math.MaxInt64), 0, true},
		{int32(math.MaxInt32), 0, true},
		{int16(math.MaxInt16), 0, true},
		{uint8(math.MaxUint8), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Int8E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Int8(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect int
		isErr  bool
	}{
		{int(8), 8, false},
		{int8(8), 8, false},
		{int16(8), 8, false},
		{int32(8), 8, false},
		{int64(8), 8, false},
		{uint(8), 8, false},
		{uint8(8), 8, false},
		{uint16(8), 8, false},
		{uint32(8), 8, false},
		{uint64(8), 8, false},
		{float32(8.31), 8, false},
		{float64(8.31), 8, false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8, false},
		{float64(-8.31), -8, false},
		{"-8", -8, false},
		{"-8.01", -8, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8, false},
		{int(math.MaxInt32), int(math.MaxInt32), false},
		{nil, 0, false},
		{aliasTypeInt0, 0, false},
		{&aliasTypeInt0, 0, false},
		{aliasTypeInt1, 1, false},
		{&aliasTypeInt1, 1, false},
		{AliasTypeInt8(1), 1, false},
		{AliasTypeInt16(1), 1, false},
		{AliasTypeInt32(1), 1, false},
		{AliasTypeInt64(1), 1, false},
		{aliasTypeUint0, 0, false},
		{&aliasTypeUint0, 0, false},
		{aliasTypeUint1, 1, false},
		{&aliasTypeUint1, 1, false},
		{AliasTypeUint8(1), 1, false},
		{AliasTypeUint16(1), 1, false},
		{AliasTypeUint32(1), 1, false},
		{AliasTypeUint64(1), 1, false},
		{AliasTypeFloat32(1), 1, false},
		{AliasTypeFloat32(0), 0, false},
		{AliasTypeFloat32(1.01), 1, false},
		{AliasTypeFloat32(-1.01), -1, false},
		{AliasTypeFloat64(0), 0, false},
		{AliasTypeFloat64(1), 1, false},
		{AliasTypeFloat64(1.01), 1, false},
		{AliasTypeFloat64(-1.01), -1, false},
		{aliasTypeString0, 0, false},
		{&aliasTypeString0, 0, false},
		{aliasTypeString1, 1, false},
		{&aliasTypeString1, 1, false},
		{aliasTypeString8d15, 8, false},
		{&aliasTypeString8d15, 8, false},
		{aliasTypeString8d15Minus, -8, false},
		{&aliasTypeString8d15Minus, -8, false},
		{AliasTypeBytes("0"), 0, false},
		{AliasTypeBytes("10.98"), 10, false},
		{aliasTypeBool4True, 1, false},
		{&aliasTypeBool4True, 1, false},
		{aliasTypeBool4False, 0, false},
		{&aliasTypeBool4False, 0, false},
		{json.Number("1"), 1, false},
		{pointerInterNil, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
		{"hello", 0, true},
		{[]byte("hello"), 0, true},
		{json.Number("hello"), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.IntE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Int(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}
