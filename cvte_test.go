package cvt_test

import (
	"errors"
	"fmt"
	"html/template"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shockerli/cvt"
)

// [test data]

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

type TestMarshalJSON struct{}

func (TestMarshalJSON) MarshalJSON() ([]byte, error) {
	return []byte("MarshalJSON"), nil
}

type TestStructA struct {
	A1 int
	TestStructB
	A2 string
	DD TestStructD
}

type TestStructB struct {
	TestStructC
	B1 int
}

type TestStructC struct {
	C1 string
}

func (c TestStructC) String() string {
	return c.C1
}

type TestStructD struct {
	D1 int
}

type TestStructE struct {
	D1 int
	DD *TestStructD
}

type TestTimeStringer struct {
	time time.Time
}

func (t TestTimeStringer) String() string {
	return t.time.String()
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cvt.Bool(aliasTypeString_0, true)
	}
}

// [function tests]

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

func TestSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []interface{}
		isErr  bool
	}{
		{"hello", []interface{}{'h', 'e', 'l', 'l', 'o'}, false},
		{[]byte("hey"), []interface{}{byte('h'), byte('e'), byte('y')}, false},
		{[]rune("我爱中国"), []interface{}{'我', '爱', '中', '国'}, false},
		{[]int{}, nil, false},
		{[]int{1, 2, 3}, []interface{}{1, 2, 3}, false},
		{[]string{}, nil, false},
		{[]string{"a", "b", "c"}, []interface{}{"a", "b", "c"}, false},
		{[]interface{}{1, "a", -1, nil}, []interface{}{1, "a", -1, nil}, false},
		{[...]string{}, nil, false},
		{[...]string{"a", "b", "c"}, []interface{}{"a", "b", "c"}, false},
		{map[int]string{}, nil, false},
		{map[int]string{1: "111", 2: "222"}, []interface{}{"111", "222"}, false},
		{map[int]TestStructC{}, nil, false},
		{map[int]TestStructC{1: {"c1"}, 2: {"c2"}}, []interface{}{TestStructC{"c1"}, TestStructC{"c2"}}, false},
		// map key convert to string, and sorted by key asc
		{map[interface{}]string{
			"k":  "k",
			1:    "1",
			0:    "0",
			"b":  "b",
			-1:   "-1",
			"3c": "3c",
			-0.1: "-0.1",
		}, []interface{}{"-0.1", "-1", "0", "1", "3c", "b", "k"}, false},

		{testing.T{}, nil, false},
		{&testing.T{}, nil, false},
		{TestStructA{}, []interface{}{0, "", 0, "", TestStructD{0}}, false},
		{&TestStructB{}, []interface{}{"", 0}, false},
		{&TestStructE{}, []interface{}{0, (*TestStructD)(nil)}, false},

		// errors
		{int(123), nil, true},
		{uint16(123), nil, true},
		{float64(12.3), nil, true},
		{func() {}, nil, true},
		{nil, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.SliceE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)

		// Non-E test with no default value:
		v = cvt.Slice(tt.input)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestSliceIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []int
		isErr  bool
	}{
		{[]int{}, nil, false},
		{[]int{1, 2, 3}, []int{1, 2, 3}, false},
		{[]string{}, nil, false},
		{[]interface{}{1, "-1", -1, nil}, []int{1, -1, -1, 0}, false},
		{[...]string{}, nil, false},
		{[...]string{"1", "2", "3"}, []int{1, 2, 3}, false},

		// sorted by key asc
		{map[int]string{}, nil, false},
		{map[int]string{2: "222", 1: "111"}, []int{111, 222}, false},
		{map[int]TestStructC{}, nil, false},
		{map[interface{}]string{
			1:    "1",
			0:    "0",
			-1:   "-1",
			-0.1: "-0.1",
		}, []int{0, -1, 0, 1}, false},

		{testing.T{}, nil, false},
		{&testing.T{}, nil, false},

		// errors
		{int(123), nil, true},
		{uint16(123), nil, true},
		{float64(12.3), nil, true},
		{func() {}, nil, true},
		{nil, nil, true},
		{[]string{"a", "b", "c"}, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.SliceIntE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestSliceInt64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []int64
		isErr  bool
	}{
		{[]int{}, nil, false},
		{[]int{1, 2, 3}, []int64{1, 2, 3}, false},
		{[]string{}, nil, false},
		{[]interface{}{1, "-1", -1, nil}, []int64{1, -1, -1, 0}, false},
		{[...]string{}, nil, false},
		{[...]string{"1", "2", "3"}, []int64{1, 2, 3}, false},

		// sorted by key asc
		{map[int]string{}, nil, false},
		{map[int]string{2: "222", 1: "111"}, []int64{111, 222}, false},
		{map[int]TestStructC{}, nil, false},
		{map[interface{}]string{
			1:    "1",
			0:    "0",
			-1:   "-1",
			-0.1: "-0.1",
		}, []int64{0, -1, 0, 1}, false},

		{testing.T{}, nil, false},
		{&testing.T{}, nil, false},

		// errors
		{int(123), nil, true},
		{uint16(123), nil, true},
		{float64(12.3), nil, true},
		{func() {}, nil, true},
		{nil, nil, true},
		{[]string{"a", "b", "c"}, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.SliceInt64E(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestFieldE(t *testing.T) {
	tests := []struct {
		input  interface{}
		field  interface{}
		expect interface{}
		isErr  bool
	}{
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "D1", 1, false},
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "DD", &TestStructD{D1: 2}, false},
		{TestStructB{B1: 1, TestStructC: TestStructC{C1: "c1"}}, "C1", "c1", false},
		{map[int]interface{}{123: "112233"}, "123", "112233", false},
		{map[int]interface{}{123: "112233"}, 123, "112233", false},
		{map[string]interface{}{"123": "112233"}, 123, "112233", false},
		{map[string]interface{}{"c": "ccc"}, TestStructC{C1: "c"}, "ccc", false},

		// errors
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "", nil, true},
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "Age", nil, true},
		{int(123), "Name", nil, true},
		{uint16(123), "Name", nil, true},
		{float64(12.3), "Name", nil, true},
		{func() {}, "Name", nil, true},
		{nil, "Name", nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], field[%s], expect[%+v], isErr[%v]",
			i, tt.input, tt.field, tt.expect, tt.isErr,
		)

		v, err := cvt.FieldE(tt.input, tt.field)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)
	}
}

func TestColumnsE(t *testing.T) {
	tests := []struct {
		input  interface{}
		field  interface{}
		expect interface{}
		isErr  bool
	}{
		{[]interface{}{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}}, "D1", []interface{}{1}, false},
		{[]TestStructE{{D1: 1}, {D1: 2}}, "D1", []interface{}{1, 2}, false},
		{[]TestStructE{{DD: &TestStructD{}}, {D1: 2}}, "DD", []interface{}{&TestStructD{}, (*TestStructD)(nil)}, false},
		{[]interface{}{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}}, "DD", []interface{}{&TestStructD{D1: 2}}, false},
		{[]map[string]interface{}{{"1": 111, "DDD": "D1"}, {"2": 222, "DDD": "D2"}, {"DDD": nil}}, "DDD", []interface{}{"D1", "D2", nil}, false},
		{map[int]map[string]interface{}{1: {"1": 111, "DDD": "D1"}, 2: {"2": 222, "DDD": "D2"}, 3: {"DDD": nil}}, "DDD", []interface{}{"D1", "D2", nil}, false},
		{map[int]TestStructD{1: {11}, 2: {22}}, "D1", []interface{}{11, 22}, false},

		// errors
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "", nil, true},
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "Age", nil, true},
		{int(123), "Name", nil, true},
		{uint16(123), "Name", nil, true},
		{float64(12.3), "Name", nil, true},
		{"Name", "Name", nil, true},
		{func() {}, "Name", nil, true},
		{nil, "Name", nil, true},
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "D1", nil, true},
		{TestStructE{D1: 1, DD: &TestStructD{D1: 2}}, "DD", nil, true},
		{TestStructB{B1: 1, TestStructC: TestStructC{C1: "c1"}}, "C1", nil, true},
		{map[int]interface{}{123: "112233"}, "123", nil, true},
		{map[int]interface{}{123: "112233"}, 123, nil, true},
		{map[string]interface{}{"123": "112233"}, 123, nil, true},
		{map[string]interface{}{"c": "ccc"}, TestStructC{C1: "c"}, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], field[%s], expect[%+v], isErr[%v]",
			i, tt.input, tt.field, tt.expect, tt.isErr,
		)

		v, err := cvt.ColumnsE(tt.input, tt.field)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
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
		{"2010年03月07日", time.Date(2010, 3, 7, 0, 0, 0, 0, loc), false},
		{"2010年03月07日 18:08:18", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
		{"2010年03月07日 18时08分18秒", time.Date(2010, 3, 7, 18, 8, 18, 0, loc), false},
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

func TestKeysE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []interface{}
		isErr  bool
	}{
		{map[int]map[string]interface{}{1: {"1": 111, "DDD": 12.3}, 2: {"2": 222, "DDD": "321"}, 3: {"DDD": nil}}, []interface{}{1, 2, 3}, false},
		{map[string]interface{}{"A": 1, "2": 2}, []interface{}{"2", "A"}, false},
		{map[float64]TestStructD{1: {11}, 2: {22}}, []interface{}{float64(1), float64(2)}, false},
		{map[interface{}]interface{}{1: 1, 2.2: 2.22, "A": "A"}, []interface{}{1, 2.2, "A"}, false},

		// errors
		{nil, nil, true},
		{"Name", nil, true},
		{testing.T{}, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], expect[%+v], isErr[%v]",
			i, tt.input, tt.expect, tt.isErr,
		)

		v, err := cvt.KeysE(tt.input)
		if tt.isErr {
			assert.Error(t, err, msg)
			continue
		}

		assert.NoError(t, err, msg)
		assert.Equal(t, tt.expect, v, msg)
	}
}
