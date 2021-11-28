package cvt_test

import (
	"fmt"
	"github.com/shockerli/cvt"
	"testing"
)

func TestStringMapE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]interface{}
		isErr  bool
	}{
		// JSON String
		{`{"name":"cvt","age":3.21}`, map[string]interface{}{"name": "cvt", "age": 3.21}, false},
		{`{"name":"cvt","tag":"convert"}`, map[string]interface{}{"name": "cvt", "tag": "convert"}, false},
		{`{"name":"cvt","build":true}`, map[string]interface{}{"name": "cvt", "build": true}, false},

		// Map
		{map[string]interface{}{}, map[string]interface{}{}, false},
		{map[string]interface{}{"name": "cvt", "age": 3.21}, map[string]interface{}{"name": "cvt", "age": 3.21}, false},
		{map[interface{}]interface{}{"name": "cvt", "age": 3.21}, map[string]interface{}{"name": "cvt", "age": 3.21}, false},
		{map[interface{}]interface{}{111: "cvt", "222": 3.21}, map[string]interface{}{"111": "cvt", "222": 3.21}, false},

		// Struct
		{struct {
			Name string
			Age  int
		}{"cvt", 3}, map[string]interface{}{"Name": "cvt", "Age": 3}, false},
		{&struct {
			Name string
			Age  int
		}{"cvt", 3}, map[string]interface{}{"Name": "cvt", "Age": 3}, false},
		{struct {
			A1 string
			TestStructC
		}{"a1", TestStructC{"c1"}}, map[string]interface{}{"A1": "a1", "C1": "c1"}, false},
		{struct {
			A1 string
			TestStructC
			C1 string
		}{"a1", TestStructC{"c1-1"}, "c1-2"}, map[string]interface{}{"A1": "a1", "C1": "c1-2"}, false},
		{struct {
			A1 string
			*TestStructC
			C1 string
		}{"a1", &TestStructC{"c1-1"}, "c1-2"}, map[string]interface{}{"A1": "a1", "C1": "c1-2"}, false},
		{struct {
			C1 string
			*TestStructC
			A1 string
		}{"c1-1", &TestStructC{"c1-2"}, "a1"}, map[string]interface{}{"A1": "a1", "C1": "c1-1"}, false},
		{struct {
			AliasTypeInt8
		}{5}, map[string]interface{}{"AliasTypeInt8": AliasTypeInt8(5)}, false},
		{struct {
			*AliasTypeInt
		}{&aliasTypeInt0}, map[string]interface{}{"AliasTypeInt": aliasTypeInt0}, false},
		{struct {
			*AliasTypeInt
		}{}, map[string]interface{}{}, false},
		{struct {
			*TestStructC
		}{}, map[string]interface{}{}, false},

		// errors
		{nil, nil, true},
		{"", nil, true},
		{"hello", nil, true},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf(
			"i = %d, input[%+v], expect[%+v], isErr[%v]",
			i, tt.input, tt.expect, tt.isErr,
		)

		v, err := cvt.StringMapE(tt.input)
		if tt.isErr {
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
	}
}
