---
title: Slice
weight: 60
---


{{< toc >}}

## ColumnsE
Return the values from a single column in the input array/slice/map of struct/map, `[]interface{}`.

Like `array_column` in PHP.

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

## KeysE
Return the keys of map, or fields of struct, `[]interface{}`.

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
Convert an interface to a `[]interface{}` type.

## SliceE
Convert an interface to a `[]interface{}` type.

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


## SliceIntE
Convert an interface to a `[]int` type.

```go
cvt.SliceIntE([]string{"1", "2", "3"})              // []int{1, 2, 3}
cvt.SliceIntE(map[int]string{2: "222", 1: "111"})   // []int{111, 222}
```

## SliceInt64E
Convert an interface to a `[]int64` type.

```go
cvt.SliceInt64E([]string{"1", "2", "3"})              // []int64{1, 2, 3}
cvt.SliceInt64E(map[int]string{2: "222", 1: "111"})   // []int64{111, 222}
```


## SliceFloat64E
Convert an interface to a `[]float64` type.

```go
cvt.SliceFloat64E([]string{"1", "2", "3"})              // []float64{1, 2, 3}
cvt.SliceFloat64E(map[int]string{2: "222", 1: "111"})   // []float64{111, 222}
```

## SliceStringE
Convert an interface to a `[]string` type.

```go
cvt.SliceStringE([]float64{1.1, 2.2, 3.0})              // []string{"1.1", "2.2", "3"}
cvt.SliceStringE(map[int]string{2: "222", 1: "11.1"})   // []string{"11.1", "222"}
```

> More case see unit: `slice_test.go`

