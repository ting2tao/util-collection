package util

import (
	"net/url"
	"strings"
)

func EncodeURLChinesePath(urlPath string) string {
	lastIndex := strings.LastIndex(urlPath, "/")
	urlPrefix := urlPath[:lastIndex+1]
	urlLast := urlPath[lastIndex+1:]

	return urlPrefix + url.QueryEscape(urlLast)
}

// FormatMobileStar 手机号中间4位替换为*号
func FormatMobileStar(mobile string) string {
	if len(mobile) < 11 {
		return mobile
	}
	return mobile[:3] + "****" + mobile[7:]
}
