package utils

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
)

func Ip2Long(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func Long2Ip(ipInt uint32) string {
	b0 := strconv.FormatUint(uint64(ipInt>>24)&0xff, 10)
	b1 := strconv.FormatUint(uint64(ipInt>>16)&0xff, 10)
	b2 := strconv.FormatUint(uint64(ipInt>>8)&0xff, 10)
	b3 := strconv.FormatUint(uint64(ipInt&0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}
