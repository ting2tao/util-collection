package funcs

import (
	"fmt"
	"github.com/Slovty/util-collection/constant"
	"math"
)

var characterArr = constant.CharacterArr

// 获取excel中的位置
func FindKey(k int, i int) string {
	char := ""
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
	return fmt.Sprintf("%v%v", char, i) //从第二行开始
}
