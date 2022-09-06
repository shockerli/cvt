---
title: 'Go type conversion'
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

Simple, safe conversion of any type, including indirect/custom types.

{{< button size="large" relref="usage/getting-started" >}}Getting Started{{< /button >}}

## Feature overview

{{< columns >}}

### Safety

Support to any data type, no panic, safe with application.

<--->

### Lightweight

Less code, zero depend, the application almost no heavy burden.

<--->

### Easily

Semantic clarity, naming, documentation, detailed, directly get started.

{{< /columns >}}

## Examples

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

