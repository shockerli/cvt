---
title: Map
weight: 70
---

{{< toc >}}


## StringMapE

- Map
```go
// expect: map[string]interface{}{"111": "cvt", "222": 3.21}
cvt.StringMapE(map[interface{}]interface{}{111: "cvt", "222": 3.21})
```

- Struct

```go
// expect: map[string]interface{}{"Name": "cvt", "Age": 3}
cvt.StringMapE(struct {
    Name string
    Age  int
}{"cvt", 3})
```

- JSON
```go
// expect: map[string]interface{}{"name": "cvt", "age": 3.21}
cvt.StringMapE(`{"name":"cvt","age":3.21}`)
```


> 更多示例请看单元测试：`map_test.go`

