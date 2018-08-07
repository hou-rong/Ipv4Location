# Ipv4Location

## 简介

用于离线查询 ipv4 地址所在的地理位置（练手项目）。

## 使用方法及结果

```go
package main

import (
	"fmt"
	"github.com/hou-rong/Ipv4Location/pkg"
)

var testIps = []string{
	"127.0.0.1",
	"192.168.1.1",
	"222.199.222.199",
	"114.114.114.114",
	"123.234.132.222",
}

func main() {
	for _, testIp := range testIps {
		var location = pkg.GetIpLocation(testIp)
		fmt.Printf("%s %+v\n", testIp, location)
	}
}
```
**output**

    127.0.0.1 {Country:本机地址 Area:CZ88.NET}
    192.168.1.1 {Country:局域网 Area:对方和您在同一内部网}
    222.199.222.199 {Country:北京市 Area:北京科技大学}
    114.114.114.114 {Country:江苏省南京市 Area:信风网络科技有限公司公众DNS服务器}
    123.234.132.222 {Country:山东省青岛市 Area:联通}