---
title: Bool
weight: 10
---

{{< toc >}}



## Bool
Convert an interface to a bool type, with default value.

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
Convert an interface to a bool type.

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
Convert and store in a new bool value, and returns a pointer to it.

```go
cvt.BoolP("true")   // (*bool)(0x14000126180)(true)
```

## More Examples
More case see unit: `bool_test.go`

