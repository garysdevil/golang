// 冒泡排序
func bubbleSort(slice []int) {
	length := len(slice)
	flag := true
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				flag = false
			}
		}
		if flag {
			return 
		}
	}
}
