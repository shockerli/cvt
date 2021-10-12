# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt) [![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt) [![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt) ![GitHub top language](https://img.shields.io/github/languages/top/shockerli/cvt) ![GitHub](https://img.shields.io/github/license/shockerli/cvt)

> Simple, safe conversion of any type, including indirect/custom types.
>
> Inspired by [cast](https://github.com/spf13/cast)

## Install

```go
go get -u github.com/shockerli/cvt
```

## Usage

[中文说明](README_ZH.md)

### with `error`

> Method `__E()`: expect handle error, while unable to convert

```go
cvt.IntE("12")          // 12, nil
cvt.Float64E("12.34")   // 12.34, nil
cvt.StringE(12.34)      // "12.34", nil
cvt.BoolE("false")      // false, nil
```

### custom type and pointers

> dereferencing pointer and reach the base type

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
- BoolE

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

### string
- String
- StringE

### float
- Float32
- Float32E
- Float64
- Float64E

### time
- Time
- TimeE

### slice
- `ColumnsE`: the values from a single column in the input array/slice/map of struct/map, `[]interface{}`
- `FieldE`: the field value from map/struct, `interface{}`
- `KeysE`: the keys of map, `[]interface{}`
- `Slice`
- `SliceE`: convert an interface to a `[]interface{}` type
- `SliceIntE`: convert an interface to a `[]int` type
- `SliceInt64E`: convert an interface to a `[]int64` type
- `SliceFloat64E`: convert an interface to a `[]float64` type
- `SliceStringE`: convert an interface to a `[]string` type


## License

This project is under the terms of the [MIT](LICENSE) license.
