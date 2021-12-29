package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(multiply("2", "2"))
	fmt.Println(multiply("3", "5"))
	fmt.Println(multiply("8", "8"))
	fmt.Println(multiply("9", "9"))
	fmt.Println(multiply("11", "11"))
	fmt.Println(multiply("12", "11"))
	fmt.Println(multiply("22", "33"))
	fmt.Println(multiply("22", "33"))
	fmt.Println(multiply("123", "456"))
}

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 示例 1:
//
// 输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
// 输入: num1 = "123", num2 = "456"
//输出: "56088"
//
// 说明：
//
//
// num1 和 num2 的长度小于110。
// num1 和 num2 只包含数字 0-9。
// num1 和 num2 均不以零开头，除非是数字 0 本身。
// 不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//
// Related Topics 数学 字符串 模拟 👍 800 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var (
		n1   = len(num1)
		n2   = len(num2)
		ret  string
		ssum = func(num1 string, num2 string) string {
			i, j := len(num1)-1, len(num2)-1
			add := 0
			ans := ""
			for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
				x, y := 0, 0
				if i >= 0 {
					x = int(num1[i] - '0')
				}
				if j >= 0 {
					y = int(num2[j] - '0')
				}
				result := x + y + add
				ans = strconv.Itoa(result%10) + ans
				add = result / 10
			}
			return ans
		}
	)

	for i := 0; i < n1; i++ {
		var tmp string
		var a, b = num1[n1-i-1] - '0', byte(1)
		var add byte
		for j := 0; j < n2; j++ {
			b = num2[n2-j-1] - '0'
			var prod = a*b + add
			tmp = string(prod%10+'0') + tmp
			add = prod / 10
		}
		for ; add != 0; add /= 10 {
			tmp = string(add%10+'0') + tmp
		}

		if i > 0 {
			tmp += strings.Repeat("0", i)
		}

		ret = ssum(ret, tmp)
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
