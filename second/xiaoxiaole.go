// 消消乐.输入一串字母字符串，剔除相邻相同的字母，输出最后剩余的字母数量
// 例如，输入 aacvv，输出 1
package second

import (
	"fmt"
	"regexp"
	"strings"
)

func Xiaoxiaole() {
	var str string
	fmt.Scan(&str)

	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(str) {
		fmt.Println("Wrong input")
		return
	}

	slice := strings.Split(str, "")
	if len(slice) < 2 {
		fmt.Println(len(slice))
	}
	for {
		slice_result := make([]string, len(slice))
		i, j := 0, 0
		// 剔除相邻相同的字母
		for {
			if slice[i] != slice[i+1] {
				slice_result[j] = slice[i]
				j++
				i++
			} else {
				i = i + 2
			}
			if i == len(slice)-1 {
				slice_result[j] = slice[i]
				j++
				break
			} else if i > len(slice)-1 {
				break
			}
		}

		// 将不相邻的字母再次进行剔除
		slice_result = slice_result[0:j]

		if len(slice) == len(slice_result) {
			fmt.Println(len(slice_result))
			break
		} else if j == 0 || j == 1 {
			fmt.Println(j)
			break
		} else {
			slice = slice_result
		}
	}
}
