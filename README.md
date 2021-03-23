# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt) [![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt) [![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt) ![GitHub top language](https://img.shields.io/github/languages/top/shockerli/cvt) ![GitHub](https://img.shields.io/github/license/shockerli/cvt)

> Simple, safe conversion of any type, including custom types.
>
> Inspired by [cast](https://github.com/spf13/cast)

## Install

```go
go get -u github.com/shockerli/cvt
```

## Usage

[中文文档](README_ZH.md)

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

> For more examples, see tests [cvt_test.go](cvt_test.go) and [cvte_test.go](cvte_test.go)

## API

- Bool
- BoolE
- ColumnsE
- ColumnsFloat64E
- ColumnsInt64E
- ColumnsIntE
- ColumnsStringE
- FieldE
- Float32
- Float32E
- Float64
- Float64E
- Int
- Int16
- Int16E
- Int32
- Int32E
- Int64
- Int64E
- Int8
- Int8E
- IntE
- Slice
- SliceE
- String
- StringE
- Time
- TimeE
- Uint
- Uint16
- Uint16E
- Uint32
- Uint32E
- Uint64
- Uint64E
- Uint8
- Uint8E
- UintE


## License

This project is licensed under the terms of the [MIT](LICENSE) license.
