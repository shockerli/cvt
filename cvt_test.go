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
