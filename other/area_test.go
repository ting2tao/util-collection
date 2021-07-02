package util

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	areaCode2Map := InitAreaInfo()
	fmt.Println(areaCode2Map["11"])
	fmt.Println(areaCode2Map["4301"])
}
