package cvt_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/shockerli/cvt"
	"github.com/stretchr/testify/assert"
)

func TestFloat64_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    float64
		expect float64
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
		{float32(8.31), 1, 8.31},
		{float64(8.31), 1, 8.31},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8.01},
		{int(-8), 1, -8},
		{int8(-8), 1, -8},
		{int16(-8), 1, -8},
		{int32(-8), 1, -8},
		{int64(-8), 1, -8},
		{float32(-8.31), 1, -8.31},
		{float64(-8.31), 1, -8.31},
		{int64(math.MaxInt64), 1, float64(math.MaxInt64)},
		{uint64(math.MaxUint64), 1, float64(math.MaxUint64)},
		{"-8", 1, -8},
		{"-8.01", 1, -8.01},
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
		{aliasTypeString_8d15, 2, 8.15},
		{&aliasTypeString_8d15, 2, 8.15},
		{aliasTypeString_8d15_minus, 1, -8.15},
		{&aliasTypeString_8d15_minus, 1, -8.15},

		// unsupported value, def == expect
		{"10a", 1.11, 1.11},
		{"a10a", 1.11, 1.11},
		{"8.01a", 1.11, 1.11},
		{"8.01 ", 1.11, 1.11},
		{"hello", 1.11, 1.11},
		{testing.T{}, 1.11, 1.11},
		{&testing.T{}, 1.11, 1.11},
		{[]int{}, 1.11, 1.11},
		{[]string{}, 1.11, 1.11},
		{[...]string{}, 1.11, 1.11},
		{map[int]string{}, 1.11, 1.11},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Float64(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFloat64_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
	}{
		{"8.01a", 0},
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Float64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFloat32_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    float32
		expect float32
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
		{float32(8.31), 1, 8.31},
		{float64(8.31), 1, 8.31},
		{"8", 2, 8},
		{"8.00", 2, 8},
		{"8.01", 2, 8.01},
		{int(-8), 1, -8},
		{int8(-8), 1, -8},
		{int16(-8), 1, -8},
		{int32(-8), 1, -8},
		{int64(-8), 1, -8},
		{float32(-8.31), 1, -8.31},
		{float64(-8.31), 1, -8.31},
		{int(math.MaxInt32), 1, float32(math.MaxInt32)},
		{"-8", 1, -8},
		{"-8.01", 1, -8.01},
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
		{aliasTypeString_8d15, 2, 8.15},
		{&aliasTypeString_8d15, 2, 8.15},
		{aliasTypeString_8d15_minus, 1, -8.15},
		{&aliasTypeString_8d15_minus, 1, -8.15},

		// unsupported value, def == expect
		{"10a", 1.11, 1.11},
		{"a10a", 1.11, 1.11},
		{"8.01a", 1.11, 1.11},
		{"8.01 ", 1.11, 1.11},
		{"hello", 1.11, 1.11},
		{testing.T{}, 1.11, 1.11},
		{&testing.T{}, 1.11, 1.11},
		{[]int{}, 1.11, 1.11},
		{[]string{}, 1.11, 1.11},
		{[...]string{}, 1.11, 1.11},
		{map[int]string{}, 1.11, 1.11},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Float32(tt.input, tt.def)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFloat32_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
	}{
		{"8.01a", 0},
		{testing.T{}, 0},
		{&testing.T{}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 0},
		{[]string{}, 0},
		{[]string{"a", "b", "c"}, 0},
		{[...]string{}, 0},
		{map[int]string{}, 0},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Float32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFloat64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float64
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
		{float32(8.31), float64(8.31), false},
		{float64(8.31), float64(8.31), false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8.31, false},
		{float32(-8.3190), -8.3190, false},
		{float64(-8.31), -8.31, false},
		{"-8", -8, false},
		{"-8.01", -8.01, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8.01, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8.01, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8.01, false},
		{int64(math.MaxInt64), float64(math.MaxInt64), false},
		{uint64(math.MaxUint64), float64(math.MaxUint64), false},
		{nil, 0, false},
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8.15, false},
		{&aliasTypeString_8d15, 8.15, false},
		{aliasTypeString_8d15_minus, -8.15, false},
		{&aliasTypeString_8d15_minus, -8.15, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"hello", 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Float64E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Float64(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFloat32E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect float32
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
		{float32(8.31), float32(8.31), false},
		{float64(8.31), float32(8.31), false},
		{true, 1, false},
		{false, 0, false},
		{int(-8), -8, false},
		{int8(-8), -8, false},
		{int16(-8), -8, false},
		{int32(-8), -8, false},
		{int64(-8), -8, false},
		{float32(-8.31), -8.31, false},
		{float64(-8.31), -8.31, false},
		{"-8", -8, false},
		{"-8.01", -8.01, false},
		{"8", 8, false},
		{"8.00", 8, false},
		{"8.01", 8.01, false},
		{[]byte("-8"), -8, false},
		{[]byte("-8.01"), -8.01, false},
		{[]byte("8"), 8, false},
		{[]byte("8.00"), 8, false},
		{[]byte("8.01"), 8.01, false},
		{int(math.MaxInt32), float32(math.MaxInt32), false},
		{nil, 0, false},
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8.15, false},
		{&aliasTypeString_8d15, 8.15, false},
		{aliasTypeString_8d15_minus, -8.15, false},
		{&aliasTypeString_8d15_minus, -8.15, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"hello", 0, true},
		{float64(math.MaxFloat64), 0, true},
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.Float32E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Float32(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}
