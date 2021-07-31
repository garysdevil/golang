// 二分查找
func dichotomy(slice []int, value int) int {
	start, end, mid := 0, len(slice)-1, 0
	for {
		mid = (start + end) / 2
		// fmt.Println(start, end, mid)
		if value > slice[mid] {
			start = mid + 1
		} else if value < slice[mid] {
			end = mid - 1
		} else {
			// fmt.Println(mid)
			// break
			return mid
		}
		if start > end {
			// fmt.Println("There are not target in slice")
			// break
			return -1
		}
	}
}