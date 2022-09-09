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
cvt.Time("2006.01.02")
cvt.Time("2006/01/02")
cvt.Time("2016-03-06 15:28:01")
cvt.Time("2016-03-06 15:28:01.332")
cvt.Time("2016-03-06 15:28:01.332901")
cvt.Time("2016-03-06 15:28:01.332901345")
cvt.Time("2006年01月02日 15时04分05秒")
cvt.Time("Mon Jan 2 15:04:05 2006 -0700")
cvt.Time(1482597504)
cvt.Time(time.Duration(1482597504))
cvt.Time(time.Date(2009, 2, 13, 23, 31, 30, 0, time.Local))
cvt.Time(json.Number("1234567890"))
cvt.Time(json.Number("2016-03-06 15:28:01"))
```

> 更多示例请看单元测试：`time_test.go`

