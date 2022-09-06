---
title: Bool
weight: 10
---

安全的转换数据为布尔类型。

{{< toc >}}


## Bool
转换任意类型数据为布尔类型，忽略错误，支持设置默认值（默认 `false`）。

`Bool` 本质是调用 `BoolE` 方法，如果遇到 `error` 则返回默认值。

该方法永远不会报错。

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

## BoolE
转换任意类型数据为布尔类型，不支持的转换时会返回 `error` 错误。

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

## BoolP
转换任意类型数据为布尔类型的**指针地址**，忽略错误，支持设置默认值（默认 `false`）。

`BoolP` 本质是调用 `Bool` 方法，并返回其指针地址，用法与 `Bool` 完全一致。

```go
cvt.BoolP("true")   // (*bool)(0x14000126180)(true)
```

## 更多示例
更多示例请看单元测试文件：`bool_test.go`

