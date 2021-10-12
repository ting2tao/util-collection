package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// RandStr 随机字符串
func RandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// RandDigitStr 随机指定长度数字字符串
func RandDigitStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// RandHEXStr 随机指定长度16进制字符串
func RandHEXStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []byte("0123456789abcdef")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// RandUUID 产生随机UUID
func RandUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// RandNumber 在两数之间生产随机数，包含边界值
func RandNumber(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	max = max - min + 1
	data := r.Intn(max)
	data += min
	return data
}

// RandTime 在两个时间之内生成随机时间
func RandTime(min, max time.Time) time.Time {
	minUnix := min.Unix()
	maxUnix := max.Unix()
	delta := maxUnix - minUnix

	sec := rand.Int63n(delta) + minUnix
	return time.Unix(sec, 0)
}
