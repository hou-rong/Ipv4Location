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
