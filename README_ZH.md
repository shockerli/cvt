# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt)
[![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt)
[![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt)
[![codecov](https://codecov.io/gh/shockerli/cvt/branch/master/graph/badge.svg)](https://codecov.io/gh/shockerli/cvt)
![GitHub](https://img.shields.io/github/license/shockerli/cvt)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

> 一个简单、安全、高效的转换任意数据类型的 Go 语言工具包，支持自定义类型、提取结构体字段和值

## 安装

```go
go get -u github.com/shockerli/cvt
```

## 使用

中文 | [English](README.md)

### 支持 `error`

> 以 `E` 结尾的方法 `__E()`: 当转换失败时会返回错误

```go
cvt.IntE("12")          // 12, nil
cvt.Float64E("12.34")   // 12.34, nil
cvt.StringE(12.34)      // "12.34", nil
cvt.BoolE("false")      // false, nil
```

### 自定义类型、指针类型

> 自动解引用，并找到基本类型，完全支持自定义类型的转换

```go
type Name string

var name Name = "jioby"

cvt.StringE(name)       // jioby, nil
cvt.StringE(&name)      // jioby, nil
```

### 忽略 `error`

> 名称不以 `E` 结尾的方法，如果转换失败，不会返回错误，会返回零值

```go
cvt.Int("12")           // 12(success)
cvt.Int(struct{}{})     // 0(failed)
```

### 默认值

> 如果转换失败，返回默认值

```go
cvt.Int(struct{}{}, 12)         // 12
cvt.Float("hello", 12.34)       // 12.34
```

### 更多示例

> 上千个单元测试用例，覆盖率近100%，所有示例可通过单元测试了解：`*_test.go`


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

> 更多示例: [bool_test.go](bool_test.go)


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

> 更多示例: [int_test.go](int_test.go)


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

> 更多示例: [string_test.go](string_test.go)


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

> 更多示例: [float_test.go](float_test.go)

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

> 更多示例: [time_test.go](time_test.go)

### slice
- `ColumnsE`: 类似于 PHP 中的 `array_column`，`FieldE` 函数的切片版本，返回 `[]interface{}`

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

- `FieldE`: 取 `map` 或 `struct` 的字段值，返回 `interface{}`

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

- `KeysE`: 取 `map` 的键名，或结构体的字段名，返回 `[]interface{}`

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

- `Slice` / `SliceE`: 转换成 `[]interface{}`
- `SliceIntE`: 转换成 `[]int`
- `SliceInt64E`: 转换成 `[]int64`
- `SliceFloat64E`: 转换成 `[]float64`
- `SliceStringE`: 转换成 `[]string`

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

> 更多示例: [slice_test.go](slice_test.go)


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

> 更多示例: [map_test.go](map_test.go)



## 开源协议

本项目基于 [MIT](LICENSE) 协议开放源代码。

## 感谢
- [JetBrains Open Source Support](https://jb.gg/OpenSourceSupport)

![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)
