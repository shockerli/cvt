# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt)
[![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt)
[![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt)
![GitHub top language](https://img.shields.io/github/languages/top/shockerli/cvt)
[![codecov](https://codecov.io/gh/shockerli/cvt/branch/master/graph/badge.svg)](https://codecov.io/gh/shockerli/cvt)
![GitHub](https://img.shields.io/github/license/shockerli/cvt)

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

### 更多示例

> 上千个单元测试用例，覆盖率近100%，所有示例可通过单元测试了解：`*_test.go`


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
- `ColumnsE`: 类似于 PHP 中的 `array_column`，`FieldE` 函数的切片版本，返回 `[]interface{}`
- `FieldE`: 取 `map` 或 `struct` 的字段值，返回 `interface{}`
- `KeysE`: 取 `map` 的键名，返回 `[]interface{}`
- `Slice`
- `SliceE`: 转换成 `[]interface{}`
- `SliceIntE`: 转换成 `[]int`
- `SliceInt64E`: 转换成 `[]int64`
- `SliceFloat64E`: 转换成 `[]float64`
- `SliceStringE`: 转换成 `[]string`


## 开源协议

本项目基于 [MIT](LICENSE) 协议开放源代码。
