// 输入一个正整数X，求连续的正整数相加等于X的数字
// 例如输入5，输出 5； 2，3
package second

import (
	"fmt"
	"math"
)

func AddConsecutiveNum() {
	var T int
	fmt.Scan(&T)

	most := int(math.Ceil(math.Sqrt(float64(T))))
	slice, num := make([]int, most), 1

	// fmt.Println(T)
	for i := 2; i <= int(most); i++ {
		total, start := 0, 0
		mid := T / i // go的特性，将自动向下取整

		if i%2 != 0 {
			start = mid - ((i - 1) / 2)
		} else {
			start = mid - ((i - 2) / 2)
		}
		for j := 0; j < i; j++ {
			target := start + j
			total += target
			slice[j] = target
		}

		if total == T {
			fmt.Println(slice[0:i])
			num++
		}

	}
	fmt.Println("Result:", num)
}
