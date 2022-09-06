---
title: Float
weight: 30
---

安全的转换数据为浮点型。

{{< toc >}}


## 方法
- Float32
- Float32E
- Float32P
- Float64
- Float64E
- Float64P


## 示例
```go
cvt.Float64(int32(8))       // 8
cvt.Float64(float32(8.31))  // 8.31
cvt.Float64("-8")           // 8
cvt.Float64("-8.01")        // 8.01
cvt.Float64(nil)            // 0
cvt.Float64(true)           // 1
cvt.Float64(false)          // 0

type AliasTypeInt int
type PointerTypeInt *AliasTypeInt
cvt.Float64(AliasTypeInt(8))            // 8
cvt.Float64((*AliasTypeInt)(nil))       // 0
cvt.FLoat64((*PointerTypeInt)(nil))     // 0

cvt.Float64P("12.3")    // (*float64)(0x14000126180)(12.3)
```

> 更多示例请看单元测试：`float_test.go`

