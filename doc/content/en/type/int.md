---
title: Int
weight: 20
---


{{< toc >}}



## Function
- Int
- IntE
- IntP
- Int8
- Int8E
- Int8P
- Int16
- Int16E
- Int16P
- Int32
- Int32E
- Int32P
- Int64
- Int64E
- Int64P
- Uint
- UintE
- UintP
- Uint8
- Uint8E
- Uint8P
- Uint16
- Uint16E
- Uint16P
- Uint32
- Uint32E
- Uint32P
- Uint64
- Uint64E
- Uint64P

## Examples
```go
cvt.Int(int8(8))            // 8
cvt.Int(int32(8))           // 8
cvt.Int("-8.01")            // -8
cvt.Int([]byte("8.00"))     // 8
cvt.Int(nil)                // 0
cvt.IntE("8a")              // 0,err
cvt.IntE([]int{})           // 0,err

// alias type
type OrderType uint8
cvt.Int(OrderType(3))       // 3

var po OrderType = 3
cvt.Int(&po)                // 3

cvt.IntP("12")  // (*int)(0x140000a4180)(12)
```

> More case see unit: `int_test.go`

