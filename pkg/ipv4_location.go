package pkg

import (
	"os"
	"strings"
	"github.com/hou-rong/Ipv4Location/pkg/utils"
	"runtime"
	"path/filepath"
)

type Location struct {
	Country string `json:"country"`
	Area    string `json:"area"`
}

var (
	ipFile *os.File
	ip     uint32
)

func init() {
	_, filePath, _, _ := runtime.Caller(0)
	dataPath := filepath.Join(filepath.Dir(filePath), "../data/ipv4wry.dat")
	file, err := os.Open(dataPath)
	if err != nil {
		panic(err)
	}
	ipFile = file
	ip = 0
}

func readIPArray() []uint32 {
	num := readLongX(4)
	num3 := int64((readLongX(4)-num)/7) + 1
	ipFile.Seek(int64(num), 0)
	var numArray []uint32
	for i := int64(0); i < num3; i++ {
		tmp := make([]byte, 7)
		ipFile.Read(tmp)
		numArray = append(numArray, uint32(bytesToNumber(tmp[:4])))
	}
	ipFile.Seek(int64(num), 0)
	return numArray
}

func readLongX(bytesCount int64) uint32 {
	rawData := make([]byte, bytesCount)
	ipFile.Read(rawData)
	return uint32(bytesToNumber(rawData))
}

func bytesToNumber(rawData []byte) uint32 {
	sum := int64(0)
	for i := len(rawData) - 1; i >= 0; i-- {
		sum += int64(rawData[i])
		sum <<= 8
	}
	sum >>= 8
	return uint32(sum)
}

func searchIp(ipArray []uint32, start, end int64) int64 {
	index := int64((start + end) / 2)
	if index == start {
		return index
	}
	if ip < ipArray[index] {
		return searchIp(ipArray, start, index)
	}
	return searchIp(ipArray, index, end)
}

func readString(flag uint32) string {
	if flag == 1 || flag == 2 {
		ipFile.Seek(int64(readLongX(3)), 0)
	} else {
		ipFile.Seek(-1, 1)
	}
	var rawData []byte
	for {
		tmp := make([]byte, 1)
		ipFile.Read(tmp)
		if tmp[0] == '\x00' {
			break
		}
		rawData = append(rawData, tmp[0])
	}
	rawData, _ = utils.GbkToUtf8(rawData)
	return strings.Trim(string(rawData), " ")
}

func GetIpLocation(strIp string) Location {
	ip = utils.Ip2Long(strIp)
	ipArray := readIPArray()
	num := int64(searchIp(ipArray, 0, int64(len(ipArray)-1))*7) + 4
	ipFile.Seek(num, 1)
	ipFile.Seek(int64(readLongX(3)+4), 0)

	var location Location
	flag := readLongX(1)
	if flag == 1 {
		ipFile.Seek(int64(readLongX(3)), 0)
		flag = readLongX(1)
	}

	position, _ := ipFile.Seek(0, 1)
	location.Country = readString(flag)
	if flag == 2 {
		ipFile.Seek(position+3, 0)
	}

	flag = readLongX(1)
	location.Area = readString(flag)
	ipFile.Seek(0, 0)
	return location
}
