package file

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

// GetFileSize 传入文件路径 返回文件大小
func GetFileSize(localFilePath string) (string, error) {
	var fileSize string
	fd, err := os.Open(localFilePath)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	fInfo, _ := fd.Stat()
	if fInfo.Size() >= 1024 {
		fileSize = fmt.Sprintf("%v%v", math.Ceil(float64(fInfo.Size())/1024), "KB")
	}
	if fInfo.Size() >= 1024*1024 {
		fileSize = fmt.Sprintf("%.2f%v", float64(fInfo.Size())/1048576, "MB")
	}
	return fileSize, nil
}

// RemoveLocalFile 移除当前文件
func RemoveLocalFile(localFilePath string) error {
	return os.Remove(localFilePath)
}

// GetUploadRelativePath 根据os切换相对路径
func GetUploadRelativePath() string {
	if runtime.GOOS == "windows" {
		return "tmp/"
	}
	return "/tmp/"
}
