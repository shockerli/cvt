package cvt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shockerli/cvt"
)

// alias type: int
type AliasTypeInt int

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
		// true/scale
		{0, false, false},
		{float64(0.00), false, false},
		{int(0.00), false, false},
		{int64(0.00), false, false},
		{uint(0.00), false, false},
		{uint64(0.00), false, false},
		{uint8(0.00), false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},
		{"off", false, false},
		{"Off", false, false},
		{"0", false, false},
		{"0.00", false, false},
		{[]byte("Off"), false, false},
		{aliasTypeInt_0, false, false},
		{&aliasTypeInt_0, false, false},
		{aliasTypeString_0, false, false},
		{&aliasTypeString_0, false, false},
		{aliasTypeString_off, false, false},
		{&aliasTypeString_off, false, false},

		// false/slice/array/map
		{[]int{}, false, false},
		{[]string{}, false, false},
		{[...]string{}, false, false},
		{map[int]int{}, false, false},
		{map[string]string{}, false, false},

		// true/scale
		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{true, true, false},
		{-1, true, false},
		{"on", true, false},
		{"On", true, false},
		{0.01, true, false},
		{"0.01", true, false},
		{aliasTypeInt_1, true, false},
		{&aliasTypeInt_1, true, false},
		{aliasTypeString_1, true, false},
		{&aliasTypeString_1, true, false},
		{aliasTypeString_on, true, false},
		{&aliasTypeString_on, true, false},

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
