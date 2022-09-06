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


## Conversion rule
### nil
Return `false`

### Bool
Return original value

### Number
> Integer, Float, and their derived type, such as `time.Duration` of `int64`

If `val != 0`, then return `true`; otherwise return `false`.

### String
> `string`, `[]byte`, and their derived type

### json.Number
If value can convert to `float64`, then compare `val != 0`; if not a number string, report an error.

### Array/Slice/Map
If the number of elements (len) greater than `0`, the return `true`; Return to `false` otherwise.

### Others
Other types, report an error.


## More Examples
More case see unit: `bool_test.go`

