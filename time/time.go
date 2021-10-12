package util

import (
	"fmt"
	"math"
)

// GetHourMinSecFormat 传入秒 返回时分秒
func GetHourMinSecFormat(seconds uint) (res string) {
	if seconds == 0 {
		return "0秒"
	}

	if seconds >= 3600 {
		res = fmt.Sprintf("%v%v%v", res, math.Floor(float64(seconds)/3600), "小时")
		seconds = seconds % 3600
	}
	if seconds >= 60 {
		res = fmt.Sprintf("%v%v%v", res, math.Floor(float64(seconds)/60), "分钟")
		seconds = seconds % 60
	}
	res = fmt.Sprintf("%v%v%v", res, seconds, "秒")
	return
}
