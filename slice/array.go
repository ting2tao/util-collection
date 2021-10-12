package slice

// InArray 是否在切片里面
func InArray(needle interface{}, arr interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range arr.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range arr.([]int) {
			if key == item {
				return true
			}
		}
	case float64:
		for _, item := range arr.([]float64) {
			if key == item {
				return true
			}
		}
	default:
		return false

	}
	return false
}
