---
title: Go 类型转换工具包
geekdocNav: false
geekdocAlign: center
geekdocAnchor: false
---

<!-- markdownlint-capture -->
<!-- markdownlint-disable MD033 -->

<span class="badge-placeholder">[![PkgGoDev](https://pkg.go.dev/badge/github.com/shockerli/cvt)](https://pkg.go.dev/github.com/shockerli/cvt)</span>
<span class="badge-placeholder">[![Go Report Card](https://goreportcard.com/badge/github.com/shockerli/cvt)](https://goreportcard.com/report/github.com/shockerli/cvt)</span>
<span class="badge-placeholder">[![Build Status](https://travis-ci.com/shockerli/cvt.svg?branch=master)](https://travis-ci.com/shockerli/cvt)</span>
<span class="badge-placeholder">[![codecov](https://codecov.io/gh/shockerli/cvt/branch/master/graph/badge.svg)](https://codecov.io/gh/shockerli/cvt)</span>
<span class="badge-placeholder">![GitHub](https://img.shields.io/github/license/shockerli/cvt)</span>
<span class="badge-placeholder">[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)</span>

<!-- markdownlint-restore -->

一个简单、安全、高效的转换任意数据类型的 Go 语言工具包，支持自定义类型、提取结构体字段和值

{{< button relref="usage/getting-started" >}}快速上手{{< /button >}}

## 特性

{{< columns >}}

### 安全

支持传入任意类型数据，不会造成恐慌，程序安全运行

<--->

### 轻量

代码少、零依赖，对应用程序几乎无臃肿负担

<--->

### 易用

语义清晰、命名统一、文档详细、直接上手

{{< /columns >}}


## 示例

{{< tabs "index-example" >}}
{{< tab "WITH ERROR" >}}
```go
cvt.IntE("12")          // 12, nil
cvt.Float64E("12.34")   // 12.34, nil
cvt.StringE(12.34)      // "12.34", nil
cvt.BoolE("false")      // false, nil
```
{{< /tab >}}

{{< tab "IGNORE ERROR" >}}
```go
cvt.Int("12")           // 12(success)
cvt.Int(struct{}{})     // 0(failed)
```
{{< /tab >}}

{{< tab "WITH DEFAULT" >}}
```go
cvt.Int(struct{}{}, 12)     // 12
cvt.Float("hello", 12.34)   // 12.34
```
{{< /tab >}}

{{< tab "CUSTOM TYPE" >}}
```go
type Name string

var name Name = "jioby"

cvt.StringE(name)       // jioby, nil
```
{{< /tab >}}

{{< tab "INDIRECT" >}}
```go
var name = "jioby"

cvt.StringE(&name)       // jioby, nil
```
{{< /tab >}}

{{< tab "POINTER" >}}
```go
cvt.BoolP("true")   // (*bool)(0x14000126180)(true)
```
{{< /tab >}}

