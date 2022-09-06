---
title: Slice
weight: 60
---



{{< toc >}}

## ColumnsE
返回切片或字典中的某个字段切片，字段可是字典的键值或结构体的字段。

`FieldE` 函数的切片版本，返回 `[]interface{}`。

类似于 `PHP` 中的 `array_column`。

```go
// []interface{}{"D1", "D2", nil}
cvt.ColumnsE([]map[string]interface{}{
	  {"1": 111, "DDD": "D1"},
	  {"2": 222, "DDD": "D2"},
	  {"DDD": nil},
}, "DDD")

// test type
type TestStructD struct {
    D1 int
}
type TestStructE struct {
    D1 int
    DD *TestStructD
}

// []interface{}{11, 22}
cvt.ColumnsE(map[int]TestStructD{1: {11}, 2: {22}}, "D1")

// []interface{}{1, 2}
cvt.ColumnsE([]TestStructE{{D1: 1}, {D1: 2}}, "D1")
```

## KeysE
取 `map` 的键名，或结构体的字段名，返回 `[]interface{}`。

```go
cvt.KeysE()
// key of map
cvt.KeysE(map[float64]float64{0.1: -0.1, -1.2: 1.2}) // []interface{}{-1.2, 0.1}
cvt.KeysE(map[string]interface{}{"A": 1, "2": 2}) // []interface{}{"2", "A"}
cvt.KeysE(map[int]map[string]interface{}{1: {"1": 111, "DDD": 12.3}, -2: {"2": 222, "DDD": "321"}, 3: {"DDD": nil}}) // []interface{}{-2, 1, 3}

// field name of struct
cvt.KeysE(struct{
	  A string
	  B int
	  C float
}{}) // []interface{}{"A", "B", "C"}

type TestStructB {
	  B int
}
cvt.KeysE(struct{
	  A string
    TestStructB
	  C float
}{}) // []interface{}{"A", "B", "C"}
```


## Slice
参考 `SliceE` 方法。

## SliceE
转换成 `[]interface{}`

```go
cvt.SliceE("hello")                             // []interface{}{'h', 'e', 'l', 'l', 'o'}
cvt.SliceE([]byte("hey"))                       // []interface{}{byte('h'), byte('e'), byte('y')}
cvt.SliceE([]int{1, 2, 3})                      // []interface{}{1, 2, 3}
cvt.SliceE([]string{"a", "b", "c"})             // []interface{}{"a", "b", "c"}
cvt.SliceE(map[int]string{1: "111", 2: "222"})  // []interface{}{"111", "222"}

// struct values
type TestStruct struct {
    A int
    B string
}
cvt.SliceE(TestStruct{18,"jhon"}) // []interface{}{18, "jhon"}
```


## SliceInt
参考 `SliceIntE` 方法。

## SliceIntE
转换成 `[]int`

```go
cvt.SliceIntE([]string{"1", "2", "3"})              // []int{1, 2, 3}
cvt.SliceIntE(map[int]string{2: "222", 1: "111"})   // []int{111, 222}
```

## SliceInt64
参考 `SliceInt64E` 方法。

## SliceInt64E
转换成 `[]int64`

```go
cvt.SliceInt64E([]string{"1", "2", "3"})              // []int64{1, 2, 3}
cvt.SliceInt64E(map[int]string{2: "222", 1: "111"})   // []int64{111, 222}
```


## SliceFloat64
参考 `SliceFloat64E` 方法。

## SliceFloat64E
转换成 `[]float64`

```go
cvt.SliceFloat64E([]string{"1", "2", "3"})              // []float64{1, 2, 3}
cvt.SliceFloat64E(map[int]string{2: "222", 1: "111"})   // []float64{111, 222}
```

## SliceString
参考 `SliceStringE` 方法。

## SliceStringE
转换成 `[]string`

```go
cvt.SliceStringE([]float64{1.1, 2.2, 3.0})              // []string{"1.1", "2.2", "3"}
cvt.SliceStringE(map[int]string{2: "222", 1: "11.1"})   // []string{"11.1", "222"}
```

> 更多示例请看单元测试：`slice_test.go`

