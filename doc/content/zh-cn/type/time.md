---
title: Time
weight: 50
---


{{< toc >}}


## 方法
- Time
- TimeE

## 示例
```go
cvt.Time("2009-11-10 23:00:00 +0000 UTC")
cvt.Time("2018-10-21T23:21:29+0200")
cvt.Time("10 Nov 09 23:00 UTC")
cvt.Time("2009-11-10T23:00:00Z")
cvt.Time("11:00PM")
cvt.Time("2006-01-02")
cvt.Time("2016-03-06 15:28:01")
cvt.Time(1482597504)
cvt.Time(time.Date(2009, 2, 13, 23, 31, 30, 0, time.Local))
```

> 更多示例请看单元测试：`time_test.go`

