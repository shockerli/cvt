package cvt_test

import (
	"fmt"
	"testing"

	"github.com/shockerli/cvt"
)

func TestSlice_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		def    []interface{}
		expect []interface{}
	}{
		// supported value, def is not used, def != expect
		{[]int{1, 2, 3}, []interface{}{"a", "b"}, []interface{}{1, 2, 3}},
		{testing.T{}, []interface{}{1, 2, 3}, nil},

		// unsupported value, def == expect
		{int(123), []interface{}{"hello"}, []interface{}{"hello"}},
		{uint16(123), nil, nil},
		{func() {}, []interface{}{}, []interface{}{}},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Slice(tt.input, tt.def)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestSlice_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []interface{}
	}{
		{int(123), nil},
		{uint16(123), nil},
		{func() {}, nil},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Slice(tt.input)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)

		// Non-E test
		v = cvt.Slice(tt.input)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
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
		{[]float64{1, 2, 3.3, 4.8}, []int64{1, 2, 3, 4}, false},
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
	}
}

func TestSliceFloat64E(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []float64
		isErr  bool
	}{
		{[]int{}, nil, false},
		{[]int{1, 2, 3}, []float64{1, 2, 3}, false},
		{[]float64{1, 2, 3}, []float64{1, 2, 3}, false},
		{[]float64{1.1, 2.2, 3}, []float64{1.1, 2.2, 3}, false},
		{[]string{}, nil, false},
		{[]interface{}{1, "-1.1", -1.7, nil}, []float64{1, -1.1, -1.7, 0}, false},
		{[...]string{}, nil, false},
		{[...]string{"1.01", "2.22", "3.30", "-1"}, []float64{1.01, 2.22, 3.3, -1}, false},

		// sorted by key asc
		{map[int]string{}, nil, false},
		{map[int]string{2: "222", 1: "11.1"}, []float64{11.1, 222}, false},
		{map[int]TestStructC{}, nil, false},
		{map[interface{}]string{
			1:    "1",
			0.9:  "0.9",
			-1:   "-1",
			-0.1: "-0.1",
		}, []float64{-0.1, -1, 0.9, 1}, false},

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

		v, err := cvt.SliceFloat64E(tt.input)
		if tt.isErr {
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
	}
}

func TestSliceStringE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []string
		isErr  bool
	}{
		{[]int{}, nil, false},
		{[]int{1, 2, 3}, []string{"1", "2", "3"}, false},
		{[]float64{1, 2, 3}, []string{"1", "2", "3"}, false},
		{[]float64{1.1, 2.2, 3.0}, []string{"1.1", "2.2", "3"}, false},
		{[]string{}, nil, false},
		{[]interface{}{1, "-1.1", -1.7, nil}, []string{"1", "-1.1", "-1.7", ""}, false},
		{[...]string{}, nil, false},
		{[...]string{"1.01", "2.22", "3.30", "-1"}, []string{"1.01", "2.22", "3.30", "-1"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false},

		// sorted by key asc
		{map[int]string{}, nil, false},
		{map[int]string{2: "222", 1: "11.1"}, []string{"11.1", "222"}, false},
		{map[int]TestStructC{1: {"C12"}}, []string{"C12"}, false},

		{testing.T{}, nil, false},
		{&testing.T{}, nil, false},

		// errors
		{int(123), nil, true},
		{uint16(123), nil, true},
		{float64(12.3), nil, true},
		{func() {}, nil, true},
		{nil, nil, true},
		{[]interface{}{testing.T{}}, nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v], isErr[%v]", i, tt.input, tt.expect, tt.isErr)

		v, err := cvt.SliceStringE(tt.input)
		if tt.isErr {
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
	}
}
