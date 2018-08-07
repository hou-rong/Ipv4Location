package utils

import (
	"testing"
	"fmt"
)

var Ip2LongTest = []struct {
	a string
	b uint32
}{
	{"127.0.0.1", 2130706433},
	{"192.168.1.1", 3232235777},
	{"202.4.130.95", 3389293151},
	{"114.114.114.114", 1920103026},
}

func TestIp2Long(t *testing.T) {
	for _, tt := range Ip2LongTest {
		ip := Ip2Long(tt.a)
		if ip != tt.b {
			t.Error(fmt.Sprintf("str ip %s convert to int is %d not %d", tt.a, ip, tt.b))
		}
	}
}

func TestLong2Ip(t *testing.T) {
	for _, tt := range Ip2LongTest {
		ip := Long2Ip(tt.b)
		if tt.a != ip {
			t.Error(fmt.Sprintf("int ip %d convert to str is %s not %s", tt.b, ip, tt.a))
		}
	}
}
