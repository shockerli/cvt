# cvt

[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt)
[![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt)
[![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt)
[![codecov](https://codecov.io/gh/shockerli/cvt/branch/master/graph/badge.svg)](https://codecov.io/gh/shockerli/cvt)
![GitHub](https://img.shields.io/github/license/shockerli/cvt)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

> Simple, safe conversion of any type, including indirect/custom types.


## Documents
https://cvt.shockerli.net


## Install
> Go >= 1.13

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


## License

This project is under the terms of the [MIT](LICENSE) license.

## Thanks
- [JetBrains Open Source Support](https://jb.gg/OpenSourceSupport)

![JetBrains Logo (Main) logo](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)
