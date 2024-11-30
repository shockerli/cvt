---
title: Others
weight: 80
---


{{< toc >}}


## Field
Reference method `FieldE`.

## FieldE
Return the field value from map/struct, `interface{}`.

```go
// map
cvt.FieldE(map[int]interface{}{123: "112233"}, "123") // "112233"
cvt.FieldE(map[string]interface{}{"123": "112233"}, "123") // "112233"

// struct
cvt.FieldE(struct{
	  A string
	  B int
}{"Hello", 18}, "A") // "Hello"
cvt.FieldE(struct{
	  A string
	  B int
}{"Hello", 18}, "B") // 18
```

Combine with other methods:

```go
cvt.Int(cvt.Field(map[int]interface{}{123: "112233"}, 123)) // 112233
```

## Len
return size of string, slice, array or map.

```go
cvt.Len("Hello") // 5
cvt.Len([]int{1, 2, 3}) // 3
cvt.Len([]interface{}{1, "2", 3.0}}) // 3
cvt.Len(map[int]interface{}{1: 1, 2: 2}) // 2
```

## IsEmpty
checks value for empty state.

```go
cvt.IsEmpty("") // true
cvt.IsEmpty("123") // false
cvt.IsEmpty(nil) // true
cvt.IsEmpty(true) // false
cvt.IsEmpty(false) // true
cvt.IsEmpty(0) // true
cvt.IsEmpty(1) // false
cvt.IsEmpty(180) // false
cvt.IsEmpty(1.23) // false
cvt.IsEmpty([]bool{}) // true
cvt.IsEmpty([]int{1, 2}) // false
cvt.IsEmpty(map[int]interface{}(nil)) // true
cvt.IsEmpty(map[string]string{"1": "1", "2": "2"}) // false
```


> 更多示例请看单元测试：`cvte_test.go`

