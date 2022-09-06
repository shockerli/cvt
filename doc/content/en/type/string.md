---
title: String
weight: 40
---


{{< toc >}}



## Function
- String
- StringE
- StringP

## Examples
```go
cvt.String(uint(8))             // "8"
cvt.String(float32(8.31))       // "8.31"
cvt.String(true)                // "true"
cvt.String([]byte("-8.01"))     // "-8.01"
cvt.String(nil)                 // ""

cvt.String(errors.New("error info"))            // "error info"
cvt.String(time.Friday)                         // "Friday"
cvt.String(big.NewInt(123))                     // "123"
cvt.String(template.URL("https://host.foo"))    // "https://host.foo"
cvt.String(template.HTML("<html></html>"))      // "<html></html>"
cvt.String(json.Number("12.34"))                // "12.34"

// custom type
type TestMarshalJSON struct{}

func (TestMarshalJSON) MarshalJSON() ([]byte, error) {
    return []byte("custom marshal"), nil
}
cvt.String(TestMarshalJSON{})   // "custom marshal"
cvt.String(&TestMarshalJSON{})  // "custom marshal"

cvt.StringP(8.31) // (*string)(0x14000110320)((len=3) "123")
```

> More case see unit: `string_test.go`

