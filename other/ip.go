package util

import (
	"fmt"
	"strconv"
	"strings"
)

// Ip2Int ip转int
func Ip2Int(ip string) (ipNum int, err error) {

	ipSplitArr := strings.Split(ip, ".")
	if len(ipSplitArr) != 4 {
		return 0, fmt.Errorf("ip:%v 非法", ip)
	}

	for i, value := range ipSplitArr {
		intValue, _ := strconv.Atoi(value)
		ipNum += intValue << (24 - i*8)
	}

	return ipNum, nil
}
