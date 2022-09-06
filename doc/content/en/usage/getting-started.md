---
title: Getting Started
weight: -20
---

Simple, safe conversion of any type, including indirect/custom types.

<!--more-->

{{< toc >}}

## Requirements
- Go >= 1.13


## Install
```go
go get -u github.com/shockerli/cvt
```

## Import
```go
import "github.com/shockerli/cvt"
```


## Usage

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

### Pointer
> Method `__P`, return the pointer of converted data

```go
cvt.BoolP("true")   // (*bool)(0x14000126180)(true)
```

### more

> 1000+ unit test cases, for more examples, see `*_test.go`

