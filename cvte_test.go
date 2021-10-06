package cvt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/shockerli/cvt"
)

// [test data]

// indirect type
type (
	AliasTypeBool    bool
	AliasTypeInt     int
	AliasTypeInt8    int8
	AliasTypeInt16   int16
	AliasTypeInt32   int32
	AliasTypeInt64   int64
	AliasTypeUint    uint
	AliasTypeUint8   uint8
	AliasTypeUint16  uint16
	AliasTypeUint32  uint32
	AliasTypeUint64  uint64
	AliasTypeFloat32 float32
	AliasTypeFloat64 float64
	AliasTypeString  string
	AliasTypeBytes   []byte
)

var (
	aliasTypeBool_true  AliasTypeBool = true
	aliasTypeBool_false AliasTypeBool = false
)

var (
	aliasTypeInt_0 AliasTypeInt = 0
	aliasTypeInt_1 AliasTypeInt = 1

	aliasTypeUint_0 AliasTypeUint = 0
	aliasTypeUint_1 AliasTypeUint = 1

	aliasTypeString_0          AliasTypeString = "0"
	aliasTypeString_1          AliasTypeString = "1"
	aliasTypeString_8d15       AliasTypeString = "8.15"
	aliasTypeString_8d15_minus AliasTypeString = "-8.15"
	aliasTypeString_on         AliasTypeString = "on"
	aliasTypeString_off        AliasTypeString = "off"
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
		cvt.Bool(aliasTypeString_0, true)
	}
}

// [function tests]

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
