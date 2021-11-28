# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt)
[![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt)
[![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt)
![GitHub top language](https://img.shields.io/github/languages/top/shockerli/cvt)
[![codecov](https://codecov.io/gh/shockerli/cvt/branch/master/graph/badge.svg)](https://codecov.io/gh/shockerli/cvt)
![GitHub](https://img.shields.io/github/license/shockerli/cvt)

> Simple, safe conversion of any type, including indirect/custom types.

## Install

```go
go get -u github.com/shockerli/cvt
```

## Usage

English | [中文](README_ZH.md)

### with `error`

> Method `__E()`: expect handle error, while unable to convert

```go
cvt.IntE("12")          // 12, nil
cvt.Float64E("12.34")   // 12.34, nil
cvt.StringE(12.34)      // "12.34", nil
cvt.BoolE("false")      // false, nil
```

### custom type and pointers

> dereferencing pointer and reach the original type

```go
type Name string

var name Name = "jioby"

cvt.StringE(name)       // jioby, nil
cvt.StringE(&name)      // jioby, nil
```

### ignore `error`

> Method `__()`: ignore error, while convert failed, will return the zero value of type

```go
cvt.Int("12")           // 12(success)
cvt.Int(struct{}{})     // 0(failed)
```

### with default

> return the default value, while convert failed

```go
cvt.Int(struct{}{}, 12)     // 12
cvt.Float("hello", 12.34)   // 12.34
```

### more

> 1000+ unit test cases, for more examples, see `*_test.go`


## API

### bool
- Bool

```go
cvt.Bool(0)                 // false
cvt.Bool(nil)               // false
cvt.Bool("0")               // false
cvt.Bool("false")           // false
cvt.Bool([]int{})           // false

cvt.Bool(true)              // true
cvt.Bool("true")            // true
cvt.Bool([]int{1, 2})       // true
cvt.Bool([]byte("true"))    // true
```

- BoolE

```go
cvt.BoolE(0)                // false,nil
cvt.BoolE(nil)              // false,nil
cvt.BoolE("0")              // false,nil
cvt.BoolE("false")          // false,nil
cvt.BoolE([]int{})          // false,nil

cvt.BoolE(true)             // true,nil
cvt.BoolE("true")           // true,nil
cvt.BoolE([]int{1, 2})      // true,nil
cvt.BoolE([]byte("true"))   // true,nil
```

> more case see [bool_test.go](bool_test.go)


### int
- Int
- IntE
- Int8
- Int8E
- Int16
- Int16E
- Int32
- Int32E
- Int64
- Int64E
- Uint
- UintE
- Uint8
- Uint8E
- Uint16
- Uint16E
- Uint32
- Uint32E
- Uint64
- Uint64E

```go
cvt.Int(int8(8))            // 8
cvt.Int(int32(8))           // 8
cvt.Int("-8.01")            // -8
cvt.Int([]byte("8.00"))     // 8
cvt.Int(nil)                // 0
cvt.IntE("8a")              // 0,err
cvt.IntE([]int{})           // 0,err

// alias type
type OrderType uint8
cvt.Int(OrderType(3))       // 3

var po OrderType = 3
cvt.Int(&po)                // 3
```

> more case see [int_test.go](int_test.go)


### string
- String
- StringE

```go
cvt.String(uint(8))             // "8"
cvt.String(float32(8.31))       // "8.31"
cvt.String(true)                // "true"
cvt.String([]byte("-8.01"))     // "-8.01"
cvt.String(nil)                 // ""

cvt.String(errors.New("error info"))            // "error info"
cvt.String(time.Friday)                         // "Friday"
cvt.String(big.NewInt(123))                     // "123"
cvt.String(template.URL("https://host.foo"))    // "https://host.foo"
cvt.String(template.HTML("<html></html>"))      // "<html></html>"
cvt.String(json.Number("12.34"))                // "12.34"

// custom type
type TestMarshalJSON struct{}

func (TestMarshalJSON) MarshalJSON() ([]byte, error) {
    return []byte("custom marshal"), nil
}
cvt.String(TestMarshalJSON{})   // "custom marshal"
cvt.String(&TestMarshalJSON{})  // "custom marshal"
```

> more case see [string_test.go](string_test.go)


### float
- Float32
- Float32E
- Float64
- Float64E

```go
cvt.Float64(int32(8))       // 8
cvt.Float64(float32(8.31))  // 8.31
cvt.Float64("-8")           // 8
cvt.Float64("-8.01")        // 8.01
cvt.Float64(nil)            // 0
cvt.Float64(true)           // 1
cvt.Float64(false)          // 0

type AliasTypeInt int
type PointerTypeInt *AliasTypeInt
cvt.Float64(AliasTypeInt(8))            // 8
cvt.Float64((*AliasTypeInt)(nil))       // 0
cvt.FLoat64((*PointerTypeInt)(nil))     // 0
```

> more case see [float_test.go](float_test.go)

### time
- Time
- TimeE

```go
cvt.Time("2009-11-10 23:00:00 +0000 UTC")
cvt.Time("2018-10-21T23:21:29+0200")
cvt.Time("10 Nov 09 23:00 UTC")
cvt.Time("2009-11-10T23:00:00Z")
cvt.Time("11:00PM")
cvt.Time("2006-01-02")
cvt.Time("2016-03-06 15:28:01")
cvt.Time(1482597504)
cvt.Time(time.Date(2009, 2, 13, 23, 31, 30, 0, time.Local))
```

> more case see [time_test.go](time_test.go)

### slice
- `ColumnsE`: the values from a single column in the input array/slice/map of struct/map, `[]interface{}`

```go
// []interface{}{"D1", "D2", nil}
cvt.ColumnsE([]map[string]interface{}{
	  {"1": 111, "DDD": "D1"},
	  {"2": 222, "DDD": "D2"},
	  {"DDD": nil},
}, "DDD")

// test type
type TestStructD struct {
    D1 int
}
type TestStructE struct {
    D1 int
    DD *TestStructD
}

// []interface{}{11, 22}
cvt.ColumnsE(map[int]TestStructD{1: {11}, 2: {22}}, "D1")

// []interface{}{1, 2}
cvt.ColumnsE([]TestStructE{{D1: 1}, {D1: 2}}, "D1")
```

- `FieldE`: the field value from map/struct, `interface{}`

```go
// map
cvt.FieldE(map[int]interface{}{123: "112233"}, "123") // "112233"
cvt.FieldE(map[string]interface{}{"123": "112233"}, "123") // "112233"

// struct
cvt.FieldE(struct{
	  A string
	  B int
}{"Hello", 18}, "A") // "Hello"
cvt.FieldE(struct{
	  A string
	  B int
}{"Hello", 18}, "B") // 18
```

- `KeysE`: the keys of map, or fields of struct, `[]interface{}`

```go
cvt.KeysE()
// key of map
cvt.KeysE(map[float64]float64{0.1: -0.1, -1.2: 1.2}) // []interface{}{-1.2, 0.1}
cvt.KeysE(map[string]interface{}{"A": 1, "2": 2}) // []interface{}{"2", "A"}
cvt.KeysE(map[int]map[string]interface{}{1: {"1": 111, "DDD": 12.3}, -2: {"2": 222, "DDD": "321"}, 3: {"DDD": nil}}) // []interface{}{-2, 1, 3}

// field name of struct
cvt.KeysE(struct{
	  A string
	  B int
	  C float
}{}) // []interface{}{"A", "B", "C"}

type TestStructB {
	  B int
}
cvt.KeysE(struct{
	  A string
    TestStructB
	  C float
}{}) // []interface{}{"A", "B", "C"}
```

- `Slice` / `SliceE`: convert an interface to a `[]interface{}` type
- `SliceIntE`: convert an interface to a `[]int` type
- `SliceInt64E`: convert an interface to a `[]int64` type
- `SliceFloat64E`: convert an interface to a `[]float64` type
- `SliceStringE`: convert an interface to a `[]string` type

```go
cvt.SliceE("hello")                             // []interface{}{'h', 'e', 'l', 'l', 'o'}
cvt.SliceE([]byte("hey"))                       // []interface{}{byte('h'), byte('e'), byte('y')}
cvt.SliceE([]int{1, 2, 3})                      // []interface{}{1, 2, 3}
cvt.SliceE([]string{"a", "b", "c"})             // []interface{}{"a", "b", "c"}
cvt.SliceE(map[int]string{1: "111", 2: "222"})  // []interface{}{"111", "222"}

// struct values
type TestStruct struct {
    A int
    B string
}
cvt.SliceE(TestStruct{18,"jhon"}) // []interface{}{18, "jhon"}

// SliceIntE
cvt.SliceIntE([]string{"1", "2", "3"})              // []int{1, 2, 3}
cvt.SliceIntE(map[int]string{2: "222", 1: "111"})   // []int{111, 222}

// SliceStringE
cvt.SliceStringE([]float64{1.1, 2.2, 3.0})              // []string{"1.1", "2.2", "3"}
cvt.SliceStringE(map[int]string{2: "222", 1: "11.1"})   // []string{"11.1", "222"}
```

> more case see [slice_test.go](slice_test.go)


### map
- StringMapE

```go
// JSON
// expect: map[string]interface{}{"name": "cvt", "age": 3.21}
cvt.StringMapE(`{"name":"cvt","age":3.21}`)

// Map
// expect: map[string]interface{}{"111": "cvt", "222": 3.21}
cvt.StringMapE(map[interface{}]interface{}{111: "cvt", "222": 3.21})

// Struct
// expect: map[string]interface{}{"Name": "cvt", "Age": 3}
cvt.StringMapE(struct {
    Name string
    Age  int
}{"cvt", 3})
```

> more case see [map_test.go](map_test.go)



## License

This project is under the terms of the [MIT](LICENSE) license.
