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


> 更多示例请看单元测试：`cvte_test.go`

