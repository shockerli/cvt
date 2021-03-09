package cvt_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shockerli/cvt"
)

// alias type: bool
type AliasTypeBool bool

// alias type: int
type AliasTypeInt int

var (
	aliasTypeBool_true  AliasTypeBool = true
	aliasTypeBool_false AliasTypeBool = false
)

var (
	aliasTypeInt_0 AliasTypeInt = 0
	aliasTypeInt_1 AliasTypeInt = 1
)

// alias type: string
type AliasTypeString string

var (
	aliasTypeString_0          AliasTypeString = "0"
	aliasTypeString_1          AliasTypeString = "1"
	aliasTypeString_8d15       AliasTypeString = "8.15"
	aliasTypeString_8d15_minus AliasTypeString = "-8.15"
	aliasTypeString_on         AliasTypeString = "on"
	aliasTypeString_off        AliasTypeString = "off"
)

func TestBoolE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
		isErr  bool
	}{
		// false/scale
		{0, false, false},
		{float64(0.00), false, false},
		{int(0.00), false, false},
		{int64(0.00), false, false},
		{uint(0.00), false, false},
		{uint64(0.00), false, false},
		{uint8(0.00), false, false},
		{nil, false, false},
		{false, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{"off", false, false},
		{"Off", false, false},
		{"0", false, false},
		{"0.00", false, false},
		{[]byte("false"), false, false},
		{[]byte("Off"), false, false},
		{aliasTypeInt_0, false, false},
		{&aliasTypeInt_0, false, false},
		{aliasTypeString_0, false, false},
		{&aliasTypeString_0, false, false},
		{aliasTypeString_off, false, false},
		{&aliasTypeString_off, false, false},
		{aliasTypeBool_false, false, false},
		{&aliasTypeBool_false, false, false},

		// false/slice/array/map
		{[]int{}, false, false},
		{[]string{}, false, false},
		{[...]string{}, false, false},
		{map[int]int{}, false, false},
		{map[string]string{}, false, false},

		// true/scale
		{true, true, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{-1, true, false},
		{"on", true, false},
		{"On", true, false},
		{0.01, true, false},
		{"0.01", true, false},
		{[]byte("true"), true, false},
		{aliasTypeInt_1, true, false},
		{&aliasTypeInt_1, true, false},
		{aliasTypeString_1, true, false},
		{&aliasTypeString_1, true, false},
		{aliasTypeString_on, true, false},
		{&aliasTypeString_on, true, false},
		{aliasTypeBool_true, true, false},
		{&aliasTypeBool_true, true, false},

		// true/slice/array/map
		{[]int{1, 2, 3}, true, false},
		{[]string{"a", "b", "c"}, true, false},
		{[...]string{"a", "b", "c"}, true, false},
		{map[int]int{1: 111, 2: 222}, true, false},
		{map[string]string{"a": "aaa"}, true, false},

		// errors
		{"hello", false, true},
		{testing.T{}, false, true},
		{&testing.T{}, false, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%+v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.BoolE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Bool(tt.input)
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString_8d15_minus, 0, true},
		{&aliasTypeString_8d15_minus, 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString_8d15_minus, 0, true},
		{&aliasTypeString_8d15_minus, 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString_8d15_minus, 0, true},
		{&aliasTypeString_8d15_minus, 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString_8d15_minus, 0, true},
		{&aliasTypeString_8d15_minus, 0, true},
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
		{uint64(math.MaxUint64), uint(math.MaxUint64), false},
		{uint32(math.MaxUint32), uint(math.MaxUint32), false},
		{nil, 0, false},
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
		{aliasTypeString_8d15_minus, 0, true},
		{&aliasTypeString_8d15_minus, 0, true},
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
	tests := []struct {
		input  interface{}
		expect int64
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
		{uint32(math.MaxUint32), int64(math.MaxUint32), false},
		{nil, 0, false},
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeString_8d15_minus, -8, false},
		{&aliasTypeString_8d15_minus, -8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeString_8d15_minus, -8, false},
		{&aliasTypeString_8d15_minus, -8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeString_8d15_minus, -8, false},
		{&aliasTypeString_8d15_minus, -8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

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
		{testing.T{}, 0, true},
		{&testing.T{}, 0, true},
		{[]int{}, 0, true},
		{[]string{}, 0, true},
		{[...]string{}, 0, true},
		{map[int]string{}, 0, true},
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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeString_8d15_minus, -8, false},
		{&aliasTypeString_8d15_minus, -8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

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
		{aliasTypeInt_0, 0, false},
		{&aliasTypeInt_0, 0, false},
		{aliasTypeInt_1, 1, false},
		{&aliasTypeInt_1, 1, false},
		{aliasTypeString_0, 0, false},
		{&aliasTypeString_0, 0, false},
		{aliasTypeString_1, 1, false},
		{&aliasTypeString_1, 1, false},
		{aliasTypeString_8d15, 8, false},
		{&aliasTypeString_8d15, 8, false},
		{aliasTypeString_8d15_minus, -8, false},
		{&aliasTypeString_8d15_minus, -8, false},
		{aliasTypeBool_true, 1, false},
		{&aliasTypeBool_true, 1, false},
		{aliasTypeBool_false, 0, false},
		{&aliasTypeBool_false, 0, false},

		// errors
		{"10a", 0, true},
		{"a10a", 0, true},
		{"8.01a", 0, true},
		{"8.01 ", 0, true},
		{"4873546382743564386435354655456575456754356765546554643456", 0, true},
		{float64(4873546382743564386435354655456575456754356765546554643456), 0, true},
		{uint64(math.MaxUint64), 0, true},
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
