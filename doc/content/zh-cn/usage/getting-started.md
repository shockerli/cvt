---
title: 快速上手
weight: -20
---

简单、安全、高效的转换任意数据类型的 Go 语言工具包，支持自定义类型、提取结构体字段和值。


{{< hint type=warning >}}
**文档说明**\
`__()`、`__P()` 均是以其对应的 `__E()` 为基础的便捷方法，故其使用方法及支持的类型请参考 `__E()` 的文档，大部分此类便利方法均不再详细阐述。
{{< /hint >}}


<!--more-->

{{< toc >}}

## 环境
- Go >= 1.13


## 安装
```go
go get -u github.com/shockerli/cvt
```

## 引入
```go
import "github.com/shockerli/cvt"
```


## 使用
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

### 数据指针
> 以 `P` 结尾的方法 `__P`，返回转换后数据的指针地址

```go
cvt.BoolP("true")   // (*bool)(0x14000126180)(true)
```


### 更多示例

> 上千个单元测试用例，覆盖率近100%，所有示例可通过单元测试了解：`*_test.go`

