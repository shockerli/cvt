package cvt_test

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/shockerli/cvt"
)

// [test data]

// indirect type
type (
	AliasTypeBool      bool
	AliasTypeInt       int
	PointerTypeInt     *AliasTypeInt
	AliasTypeInt8      int8
	AliasTypeInt16     int16
	AliasTypeInt32     int32
	AliasTypeInt64     int64
	AliasTypeUint      uint
	AliasTypeUint8     uint8
	AliasTypeUint16    uint16
	AliasTypeUint32    uint32
	AliasTypeUint64    uint64
	AliasTypeFloat32   float32
	AliasTypeFloat64   float64
	AliasTypeString    string
	AliasTypeBytes     []byte
	AliasTypeInterface interface{}
)

var (
	aliasTypeBool4True  AliasTypeBool = true
	aliasTypeBool4False AliasTypeBool = false

	aliasTypeInt0 AliasTypeInt = 0
	aliasTypeInt1 AliasTypeInt = 1

	aliasTypeUint0 AliasTypeUint = 0
	aliasTypeUint1 AliasTypeUint = 1

	aliasTypeFloat0    AliasTypeFloat64 = 0
	aliasTypeFloat1    AliasTypeFloat64 = 1
	aliasTypeFloat8d15 AliasTypeFloat64 = 8.15

	aliasTypeString0                    AliasTypeString = "0"
	aliasTypeString1                    AliasTypeString = "1"
	aliasTypeString8d15                 AliasTypeString = "8.15"
	aliasTypeString8d15Minus            AliasTypeString = "-8.15"
	aliasTypeStringOn                   AliasTypeString = "on"
	aliasTypeStringOff                  AliasTypeString = "off"
	aliasTypeStringLosePrecisionInt64   AliasTypeString = "7138826985640367621"
	aliasTypeStringLosePrecisionFloat64 AliasTypeString = "7138826985640367621.18"

	pointerRunes       = []rune("中国")
	pointerInterNil    *AliasTypeInterface
	pointerIntNil      *AliasTypeInt
	aliasTypeBytesNil  AliasTypeBytes
	aliasTypeBytesTrue AliasTypeBytes = []byte("true")
	aliasTypeBytes8d15 AliasTypeBytes = []byte("8.15")
)

// custom type
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
		cvt.Bool(aliasTypeString0, true)
	}
}

// [function tests]

func TestField_HasDefault(t *testing.T) {
	tests := []struct {
		input  interface{}
		field  interface{}
		def    interface{}
		expect interface{}
	}{
		// supported value, def is not used, def != expect
		{TestStructC{C1: "c1"}, "C1", "C2", "c1"},

		// unsupported value, def == expect
		{TestStructC{C1: "c1"}, "C2", "c3", "c3"},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], def[%+v], expect[%+v]", i, tt.input, tt.def, tt.expect)

		v := cvt.Field(tt.input, tt.field, tt.def)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
	}
}

func TestField_BaseLine(t *testing.T) {
	tests := []struct {
		input  interface{}
		field  interface{}
		expect interface{}
	}{
		{struct {
			*TestStructC
		}{
			&TestStructC{C1: "c1"},
		}, "C1", "c1"},

		{"hello", "NONE", nil},
	}

	for i, tt := range tests {
		msg := fmt.Sprintf("i = %d, input[%+v], expect[%+v]", i, tt.input, tt.expect)

		v := cvt.Field(tt.input, tt.field)
		assertEqual(t, tt.expect, v, "[NonE] "+msg)
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
		{&TestStructB{B1: 1, TestStructC: TestStructC{C1: "c1"}}, "C1", "c1", false},
		{struct {
			*TestStructC
		}{
			&TestStructC{C1: "c1"},
		}, "C1", "c1", false},
		{&struct {
			TestStructC
		}{
			TestStructC{C1: "c1"},
		}, "C1", "c1", false},
		{&struct {
			*TestStructC
		}{
			&TestStructC{C1: "c1"},
		}, "C1", "c1", false},
		{&struct {
			TestStructC
		}{
			TestStructC{C1: "c1"},
		}, "TestStructC", TestStructC{C1: "c1"}, false},
		{struct {
			*TestStructC
		}{
			&TestStructC{C1: "c1"},
		}, "TestStructC", &TestStructC{C1: "c1"}, false},
		{struct {
			AliasTypeInt
		}{8}, "AliasTypeInt", AliasTypeInt(8), false},
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
			assertError(t, err, "[HasErr] "+msg)
			continue
		}

		assertNoError(t, err, "[NoErr] "+msg)
		assertEqual(t, tt.expect, v, "[WithE] "+msg)
	}
}

/* ------------------------------------------------------------------------------ */

// [testing assert functions]

// assert error
func assertError(t *testing.T, err error, msgAndArgs ...interface{}) {
	if err == nil {
		fail(t, "An error is expected but got nil", msgAndArgs...)
		return
	}
}

// assert no error
func assertNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	if err != nil {
		fail(t, fmt.Sprintf("Received unexpected error:\n\t\t\t\t%+v", err), msgAndArgs...)
		return
	}
}

// assert equal
func assertEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if err := validateEqualArgs(expected, actual); err != nil {
		fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
		return
	}

	if !objectsAreEqual(expected, actual) {
		fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual), msgAndArgs...)
		return
	}
}

func fail(t *testing.T, failureMessage string, msgAndArgs ...interface{}) {
	t.Errorf(`
Error: 		%s
Test:  		%s
Messages: 	%s`, failureMessage, t.Name(), messageFromMsgAndArgs(msgAndArgs...))
}

func validateEqualArgs(expected, actual interface{}) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		tpl, ok := msgAndArgs[0].(string)
		if ok {
			return fmt.Sprintf(tpl, msgAndArgs[1:]...)
		}
		for v := range msgAndArgs {
			tpl += fmt.Sprintf("%+v, ", v)
		}
		return strings.TrimRight(tpl, ", ")
	}
	return ""
}
