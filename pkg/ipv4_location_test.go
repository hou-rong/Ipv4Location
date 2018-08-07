package pkg

import (
	"testing"
	"fmt"
)

var GetIpLocationTest = []struct {
	ip      string
	country string
	area    string
}{
	{"127.0.0.1", "本机地址", "CZ88.NET"},
	{"192.168.1.1", "局域网", "对方和您在同一内部网"},
	{"202.4.130.95", "北京化工大学", "网络中心"},
	{"114.114.114.114", "江苏省南京市", "信风网络科技有限公司公众DNS服务器"},
}

func TestGetIpLocation(t *testing.T) {
	for _, tt := range GetIpLocationTest {
		location := GetIpLocation(tt.ip)
		if location.Country != tt.country {
			t.Error(fmt.Sprintf("ip %s country should be %s not %s", tt.ip, tt.country, location.Country))
		}
		if location.Area != tt.area {
			t.Error(fmt.Sprintf("ip %s area should be %s not %s", tt.ip, tt.area, location.Area))
		}
	}
}
