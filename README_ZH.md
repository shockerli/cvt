# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt) [![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt) [![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt) ![GitHub top language](https://img.shields.io/github/languages/top/shockerli/cvt) ![GitHub](https://img.shields.io/github/license/shockerli/cvt)

> 简单、安全的转换任意类型值，包括自定义类型

## 安装

```go
go get -u github.com/shockerli/cvt
```

## 使用

[English](README.md)

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

### more

> 所有示例，可通过单元测试了解：[cvt_test.go](cvt_test.go)、 [cvte_test.go](cvte_test.go)


## API

- Bool
- BoolE
- ColumnsE
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
- SliceIntE
- SliceInt64E
- SliceFloat64E
- SliceStringE
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


## 开源协议

本项目基于 [MIT](LICENSE) 协议开放源代码。
