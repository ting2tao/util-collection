package funcs

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestFindKey(t *testing.T) {
	fmt.Println(FindKey(23, 1))
	fmt.Println(FindKey(52, 1))
	fmt.Println(FindKey(123, 1))
	fmt.Println(FindKey(627, 1))
	fmt.Println(FindKey(1024, 21))
	fmt.Println(FindKey(2048, 2))
	fmt.Println(FindKey(51305, 2))
}

func TestFindKey3(t *testing.T) {
	tableHeader := []string{}
	for i := 0; i <= 519; i++ {
		tableHeader = append(tableHeader, strconv.Itoa(i))
	}
	tableInfo := []string{}
	for i := 0; i <= 12; i++ {
		tableInfo = append(tableInfo, strconv.Itoa(i))
	}
	dataArr := make([]map[string]string, 0)
	for i, v := range tableInfo {
		data := make(map[string]string)
		for k := range tableHeader {
			char := ""
			key := ""
			if k > 25 {
				rangeTime := int(math.Ceil(float64(k) / 26))
				rangeMod := 0
				for j := 0; j < rangeTime; j++ {
					char = char + characterArr[rangeMod]
					rangeMod = k % 26
				}
				if k%26 == 0 {
					char = char + characterArr[0]
				}
			} else {
				char = characterArr[k]
			}
			key = fmt.Sprintf("%v%v", char, i+2) //从第二行开始
			data[key] = v
		}
		dataArr = append(dataArr, data)
	}
	fmt.Println(dataArr)
}

func TestFindKey2(t *testing.T) {
	fmt.Println(FindKey(0, 2))
}
